package client

import (
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/pkg/errors"
)

func (cli *Client) RefreshToken(request *v1.AuthData) (*v1.OidcAuth200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	response, _, err := cli.patch("/v1/org/:org_id/token/refresh", encoded, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateKeyRequest")
	}
	decoded, err := decodeJSONResponse[v1.OidcAuth200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}
