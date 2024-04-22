package session

import (
	"encoding/base64"
	"encoding/json"
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/pkg/errors"
)

type SignerSessionData struct {
	OrgID      string `json:"org_id"`
	RoleID     string `json:"role_id,omitempty"`
	Expiration int64  `json:"expiration"`
	Purpose    string `json:"purpose,omitempty"`
	Token      string `json:"token"`
	Env        struct {
		DevCubeSignerStack EnvInterface `json:"Dev-CubeSignerStack"`
	} `json:"env"`
	SessionInfo v1.ClientSessionInfo `json:"session_info"`
}

type SignerSessionLifetime struct {
	// Session lifetime (in seconds). Defaults to one week (604800)
	Session *int
	// Auth token lifetime (in seconds). Defaults to five minutes (300)
	Auth *int
	//* Refresh token lifetime (in seconds). Defaults to one day (86400)
	Refresh *int
	//* Grace lifetime (in seconds). Defaults to 30 seconds (30)
	Grace *int
}

type EnvInterface struct {
	ClientID                string `json:"ClientId"`
	Region                  string `json:"Region"`
	UserPoolID              string `json:"UserPoolId"`
	SignerAPIRoot           string `json:"SignerApiRoot"`
	DefaultCredentialRpID   string `json:"DefaultCredentialRpId"`
	EncExportS3BucketName   any    `json:"EncExportS3BucketName"`
	DeletedKeysS3BucketName any    `json:"DeletedKeysS3BucketName"`
}

func NewFromBase64(data string) (*SignerSessionData, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, errors.Wrap(err, "decode string")
	}

	var sessionData SignerSessionData
	err = json.Unmarshal(bytes, &sessionData)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return &sessionData, err
}
