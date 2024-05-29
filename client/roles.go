package client

import (
	"fmt"
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	v1 "github.com/lombard-finance/cubesigner-sdk/api/v1"
	"github.com/pkg/errors"
	"io/ioutil"
)

func (cli *Client) AddKeysToRole(request *v0.AddKeysToRoleRequest, roleId string) ([]byte, error) {
	roleTokenResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "AddKeysToRole",
	}, roleId)
	if err != nil {
		return nil, errors.Wrap(err, "role token for request")
	}

	overrideHeaders := map[string]string{
		"Authorization": roleTokenResp.Token,
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	response, err := cli.put(fmt.Sprintf("/v0/org/:org_id/roles/%s/add_keys", roleId), encoded, overrideHeaders)
	if err != nil {
		return nil, errors.Wrap(err, "request AddKeysToRole")
	}
	body, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (cli *Client) CreateRoleToken(request *v0.CreateTokenRequest, roleId string) (*v1.OidcAuth200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	response, err := cli.post(fmt.Sprintf("/v0/org/:org_id/roles/%s/tokens", roleId), encoded, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateRoleToken")
	}
	decoded, err := decodeJSONResponse[v1.OidcAuth200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}
