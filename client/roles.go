package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/url"
	"strings"
)

func (cli *Client) CreateRoleToken(request *v0.CreateTokenRequest, roleId string) (*v1.OidcAuth200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	endpoint := strings.Replace("/v0/org/:org_id/roles/:role_id/tokens", ":role_id", url.PathEscape(roleId), -1)
	response, err := cli.post(endpoint, encoded, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateRoleToken")
	}
	decoded, err := decodeJSONResponse[v1.OidcAuth200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}

func (cli *Client) AddKeysToRole(request *v0.AddKeysToRoleRequest, roleId string) ([]byte, error) {
	//roleTokenResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
	//	Purpose: "AddKeysToRole",
	//}, roleId)
	//if err != nil {
	//	return nil, errors.Wrap(err, "role token for request")
	//}
	//
	//overrideHeaders := map[string]string{
	//	"Authorization": roleTokenResp.Token,
	//	"SignerAuth":    roleTokenResp.Token,
	//}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	endpoint := strings.Replace("/v0/org/:org_id/roles/:role_id/add_keys", ":role_id", url.PathEscape(roleId), -1)
	response, err := cli.put(endpoint, encoded, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request AddKeysToRole")
	}
	body, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (cli *Client) GetKeysInRole(roleId string) (*v0.ListRoleKeys200Response, error) {
	//roleTokenResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
	//	Purpose: "ListRoleKeys",
	//}, roleId)
	//if err != nil {
	//	return nil, errors.Wrap(err, "role token for request")
	//}

	//overrideHeaders := map[string]string{
	//	"Authorization": roleTokenResp.Token,
	//	"SignerAuth":    roleTokenResp.Token,
	//}

	endpoint := strings.Replace("/v0/org/:org_id/roles/:role_id/keys", ":role_id", url.PathEscape(roleId), -1)

	// TODO: implement pagination
	response, err := cli.get(endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request ListRoleKeys")
	}
	decoded, err := decodeJSONResponse[v0.ListRoleKeys200Response](response)
	return &decoded, nil
}
