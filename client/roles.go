package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/lombard-finance/cubesigner-sdk/client/pagination"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

func (cli *Client) CreateRoleToken(request *v0.CreateTokenRequest, roleId string) (*v1.OidcAuth200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	endpoint := strings.Replace("/v0/org/:org_id/roles/:role_id/tokens", ":role_id", url.PathEscape(roleId), -1)
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
	endpoint := strings.Replace("/v0/org/:org_id/roles/:role_id/add_keys", ":role_id", url.PathEscape(roleId), -1)
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

func (cli *Client) GetKeysInRole(roleId string, page *pagination.Page) (*v0.ListRoleKeys200Response, error) {
	endpoint := strings.Replace("/v0/org/:org_id/roles/:role_id/keys", ":role_id", url.PathEscape(roleId), -1)

	response, err := cli.get(endpoint, nil, page)
	if err != nil {
		return nil, errors.Wrap(err, "request ListRoleKeys")
	}
	decoded, err := decodeJSONResponse[v0.ListRoleKeys200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}
