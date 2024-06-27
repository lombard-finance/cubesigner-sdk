package client

import (
	"io"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func decodeAcceptedResponse(response io.Reader) (string, error) {
	decoded, err := decodeJSONResponse[v0.AcceptedResponse](response)
	if err != nil {
		return "", err
	}

	errorCode, ok := decoded.GetErrorCodeOk()
	if ok {

		errCode := errorCode.GetActualInstance()
		if errCode == errorCode.AcceptedValueCode {
			return "", errors.Errorf("accepted response with error (%v): %s", &errCode, decoded.GetMessage())
		}
	}

	if accepted, ok := decoded.GetAcceptedOk(); ok {
		if mfaRequired, ok := accepted.GetMfaRequiredOk(); ok {
			if id, ok := mfaRequired.GetIdOk(); ok {
				return *id, nil
			} else {
				return "", errors.New("mfa id is missed")
			}
		} else {
			return "", errors.New("mfa required is nil")
		}
	} else {
		return "", errors.New("accepted is nil")
	}
}

func getMfaHeaders(mfaId string, mfaConfirmation string, orgId string) map[string]string {
	return map[string]string{
		"x-cubist-mfa-id":           mfaId,
		"x-cubist-mfa-org-id":       orgId,
		"x-cubist-mfa-confirmation": mfaConfirmation,
	}
}
