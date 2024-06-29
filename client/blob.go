package client

import (
	"github.com/lombard-finance/cubesigner-sdk/api"
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"net/http"
	"net/url"
	"strings"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) SignBlob(roleId, key string, request *v1.BlobSignRequest, mfaId *string, mfaConfirmation *string) (*v1.BlobSignResponse, string, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign blob",
		Scopes:  []api.Scope{api.SIGNBLOB},
	}, roleId)
	if err != nil {
		return nil, "", errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

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
	endpoint := strings.Replace("/v1/org/:org_id/blob/sign/:key_id", ":key_id", url.PathEscape(key), -1)

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "request /blob/sign")
	}

	if statusCode == http.StatusAccepted {
		mfaId, err := decodeAcceptedResponse(response)
		if err != nil {
			return nil, "", errors.Wrap(err, "decode accepted response")
		}
		return nil, mfaId, nil
	}

	decoded, err := decodeJSONResponse[v1.BlobSignResponse](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode response")
	}
	return &decoded, "", nil
}
