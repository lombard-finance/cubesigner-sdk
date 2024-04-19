package session

import "github.com/pkg/errors"

var (
	NoDataErr         = errors.New("no data")
	SessionExpiredErr = errors.New("session expired")
)
