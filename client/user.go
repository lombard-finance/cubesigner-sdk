package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) AboutMeLegacy() (*v0.AboutMeLegacy200Response, error) {
	endpoint, err := cli.BuildFullEndpoint(AboutMeLegacy, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request AboutMeLegacy")
	}

	decoded, err := decodeJSONResponse[v0.AboutMeLegacy200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

func (cli *Client) AboutMe() (*v0.AboutMeLegacy200Response, error) {
	endpoint, err := cli.BuildFullEndpoint(AboutMe, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request AboutMe")
	}

	decoded, err := decodeJSONResponse[v0.AboutMeLegacy200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}
