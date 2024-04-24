package session

import (
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/lombard-finance/cubesigner-sdk/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	DefaultExpirationBuffer = 30 * time.Second
)

type SignerSessionManager struct {
	env        EnvInterface
	orgID      string
	storage    Storage
	refreshing bool
	logger     *logrus.Entry

	tokenExp   time.Time
	sessionExp time.Time

	timeout time.Duration

	stopChan chan bool
}

func NewSignerSessionManager(storage Storage, logger *logrus.Entry, timeout time.Duration) (*SignerSessionManager, error) {
	st, err := storage.retrieve()
	if err != nil {
		return nil, errors.Wrap(err, "retrieve")
	}

	return &SignerSessionManager{
		storage:    storage,
		timeout:    timeout,
		logger:     logger,
		tokenExp:   time.Unix(st.SessionInfo.AuthTokenExp, 0),
		sessionExp: time.Unix(st.Expiration, 0),
		refreshing: false,
		stopChan:   make(chan bool),
	}, nil
}

func (ssm *SignerSessionManager) Client() (*client.Client, error) {
	_ = ssm.refreshIfNeeded()

	if hasTimestampExpired(ssm.tokenExp, 0) || ssm.hasExpired() {
		return nil, SessionExpiredErr
	}

	return ssm.createClient()
}

func (ssm *SignerSessionManager) createClient() (*client.Client, error) {
	storage, err := ssm.storage.retrieve()
	if err != nil {
		return nil, errors.Wrap(err, "retrieve")
	}
	return client.New(storage.Env.DevCubeSignerStack.SignerAPIRoot, storage.OrgID, storage.Token, ssm.logger, ssm.timeout)
}

func (ssm *SignerSessionManager) refreshIfNeeded() error {
	if ssm.isStale() {
		if ssm.refreshing {
			// wait until done refreshing
			for ssm.refreshing {
				time.Sleep(time.Millisecond * 100)
			}
			return nil
		}

		// refresh
		ssm.refreshing = true
		err := ssm.refresh()
		ssm.refreshing = false
		if err != nil {
			return errors.Wrap(err, "refresh")
		}
	}

	return nil
}

// refresh Refreshes the session and **UPDATES/MUTATES** self.
func (ssm *SignerSessionManager) refresh() error {
	if ssm.hasExpired() {
		return SessionExpiredErr
	}

	currSession, err := ssm.storage.retrieve()
	if err != nil {
		return errors.Wrap(err, "retrieve")
	}

	csi := currSession.SessionInfo
	cli, err := ssm.createClient()
	if err != nil {
		return err
	}
	data, err := cli.RefreshToken(v1.NewAuthData(csi.Epoch, csi.EpochToken, csi.RefreshToken))
	if err != nil {
		return errors.Wrap(err, "refresh token")
	}

	newSession := &SignerSessionData{
		OrgID:       currSession.OrgID,
		RoleID:      currSession.RoleID,
		Expiration:  currSession.Expiration,
		Purpose:     currSession.Purpose,
		Env:         currSession.Env,
		Token:       data.Token,
		SessionInfo: data.SessionInfo,
	}

	ssm.storage.save(newSession)
	ssm.tokenExp = time.Unix(newSession.SessionInfo.AuthTokenExp, 0)
	ssm.sessionExp = time.Unix(newSession.Expiration, 0)

	return nil
}

func (ssm *SignerSessionManager) isStale() bool {
	return hasTimestampExpired(ssm.tokenExp, DefaultExpirationBuffer)
}

// hasExpired Return whether this session has expired and cannot be refreshed anymore.
// @return {boolean} Whether this session has expired.
func (ssm *SignerSessionManager) hasExpired() bool {
	return hasTimestampExpired(ssm.sessionExp, 0)
}

func (ssm *SignerSessionManager) AutoRefresh() {
	go func() {
		for {
			select {
			case <-ssm.stopChan:
				return
			default:
				if err := ssm.refreshIfNeeded(); err != nil {
					ssm.logger.WithError(err).Error("failed to auto refresh")
				}
				time.Sleep(DefaultExpirationBuffer)
			}
		}
	}()
}

func (ssm *SignerSessionManager) StopAutoRefresh() {
	// TODO: implement graceful stop with storage export
	ssm.stopChan <- true
}

func hasTimestampExpired(exp time.Time, buffer time.Duration) bool {
	expMsSinceEpoch := exp.UnixMilli()
	nowMsSinceEpoch := time.Now().UnixMilli()
	return expMsSinceEpoch < nowMsSinceEpoch+buffer.Milliseconds()
}
