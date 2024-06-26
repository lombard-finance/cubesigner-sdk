package client

import (
	"io"

	"github.com/lombard-finance/cubesigner-sdk/api"
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
		return "", errors.Errorf("accepted response with error (%v): %s", errorCode.GetActualInstance(), decoded.GetMessage())
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

func getMfaHeaders(mfaReceipt api.MfaReceiptHeader) map[string]string {
	return map[string]string{
		"x-cubist-mfa-id":           mfaReceipt.Id,
		"x-cubist-mfa-org-id":       mfaReceipt.OrganizationId,
		"x-cubist-mfa-confirmation": mfaReceipt.Confirmation,
	}
}
