package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) AboutMeLegacy() (*v0.AboutMeLegacy200Response, error) {
	response, err := cli.get("/v0/about_me", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request AboutMeLegacy")
	}
	decoded, err := decodeJSONResponse[v0.AboutMeLegacy200Response](response)
	return &decoded, nil
}

func (cli *Client) AboutMe() (*v0.AboutMeLegacy200Response, error) {
	response, err := cli.get("/v0/org/:org_id/user/me", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request AboutMe")
	}
	decoded, err := decodeJSONResponse[v0.AboutMeLegacy200Response](response)
	return &decoded, nil
}
