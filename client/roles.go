package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/pkg/errors"
)

func (cli *Client) CreateRoleToken(request *v0.CreateTokenRequest, roleId string) (*v1.OidcAuth200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	endpoint, err := cli.BuildFullEndpoint(CreateRoleToken, map[string]interface{}{ParamRoleID: roleId}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, _, err := cli.post(endpoint, encoded, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateRoleToken")
	}

	decoded, err := decodeJSONResponse[v1.OidcAuth200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}

func (cli *Client) AddKeysToRole(request *v0.AddKeysToRoleRequest, roleId string) (*v0.AddKeysToRole200Rsponse, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	endpoint, err := cli.BuildFullEndpoint(AddKeysToRole, map[string]interface{}{ParamRoleID: roleId}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, _, err := cli.put(endpoint, encoded, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request AddKeysToRole")
	}

	decoded, err := decodeJSONResponse[v0.AddKeysToRole200Rsponse](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}

func (cli *Client) GetKeysInRole(roleId string, page *Page) (*v0.ListRoleKeys200Response, error) {
	endpoint, err := cli.BuildFullEndpoint(GetKeysInRole, map[string]interface{}{ParamRoleID: roleId}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, nil, page)
	if err != nil {
		return nil, errors.Wrap(err, "request GetKeysInRole")
	}

	decoded, err := decodeJSONResponse[v0.ListRoleKeys200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}
