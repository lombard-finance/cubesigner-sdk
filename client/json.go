package client

import (
	"bytes"
	"encoding/json"
	"io"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func decodeJSONResponse[T any](body io.Reader) (T, error) {
	var res T

	if body == nil {
		return res, errors.New("no body to read")
	}

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return res, errors.Wrap(err, "parse response")
	}

	return res, nil
}

func encodeJSONRequest(request any) (io.Reader, error) {
	encoded, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshal request")
	}
	return bytes.NewBuffer(encoded), nil
}

func decodeAcceptedResponse(response io.Reader) (string, error) {
	decoded, err := decodeJSONResponse[v0.AcceptedResponse](response)
	if err != nil {
		return "", err
	}

	errorCode, ok := decoded.GetErrorCodeOk()
	if ok {

		errCode := errorCode.GetActualInstance()
		if errCode != errorCode.AcceptedValueCode {
			// TODO: errCode is pointer, need to parse to string somehow
			return "", errors.Errorf("accepted response with error (%v): %s", errCode, decoded.GetMessage())
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
