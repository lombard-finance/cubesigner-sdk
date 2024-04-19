package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) RefreshToken(request *v0.AuthData) (*v0.OidcAuth200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	response, err := cli.patch("/v1/org/:org_id/token/refresh", encoded)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateKeyRequest")
	}
	decoded, err := decodeJSONResponse[v0.OidcAuth200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}
