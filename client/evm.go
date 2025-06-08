package client

import (
	"net/http"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) SignEip712(
	pubkey string,
	request *v0.Eip712SignRequest,
	mfaHeaders *MfaHeaders,
) (*v0.EvmSignResponse, string, error) {
	headers := map[string]string{}
	if mfaHeaders != nil {
		for k, v := range cli.buildMfaHeaders(*mfaHeaders) {
			headers[k] = v
		}
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, "", errors.Wrap(err, "encode")
	}

	endpoint, err := cli.BuildFullEndpoint(SignEvmEip712, map[string]interface{}{ParamPubkey: pubkey}, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "build endpoint")
	}

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "request SignTaproot")
	}

	if statusCode == http.StatusAccepted {
		mfaId, err := decodeAcceptedResponse(response)
		if err != nil {
			return nil, "", errors.Wrap(err, "decode accepted response")
		}
		return nil, mfaId, nil
	}

	decoded, err := decodeJSONResponse[v0.EvmSignResponse](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode")
	}
	return &decoded, "", nil
}
