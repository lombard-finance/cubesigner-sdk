package client

import (
	"net/http"
	"net/url"
	"strings"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) SignEip712(pubkey string, request *v0.Eip712SignRequest, mfaId, mfaConfirmation *string) (*v0.EvmSignResponse, string, error) {
	headers := map[string]string{}

	// add mfa headers
	if mfaConfirmation != nil && *mfaConfirmation != "" {
		mfaHeaders := getMfaHeaders(*mfaId, *mfaConfirmation, cli.orgID)
		for k, v := range mfaHeaders {
			headers[k] = v
		}
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, "", errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/evm/eip712/sign/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

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
