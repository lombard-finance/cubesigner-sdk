package client

import (
	"github.com/lombard-finance/cubesigner-sdk/api"
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) CreateKeyRequest(request *v0.CreateKeyRequest) (*v0.DeriveKey200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	endpoint, err := cli.BuildFullEndpoint(CreateKey, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, _, err := cli.post(endpoint, encoded, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateKeyRequest")
	}
	decoded, err := decodeJSONResponse[v0.DeriveKey200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}

func (cli *Client) GetKeyInOrg(key string) (*v0.GetKeyInOrg200Response, error) {
	endpoint, err := cli.BuildFullEndpoint(GetKeyInOrg, map[string]interface{}{ParamKeyID: key}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, nil, nil)

	if err != nil {
		return nil, errors.Wrap(err, "request GetKeyInOrg")
	}
	decoded, err := decodeJSONResponse[v0.GetKeyInOrg200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

func (cli *Client) GetKeyInOrgForRole(key, role string) (*v0.GetKeyInOrg200Response, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "get key",
		Scopes:  []api.Scope{api.SIGNBTC}, // NOTE: cubesigner requires at least one scope to pass, but it anyway allows to get key
	}, role)
	if err != nil {
		return nil, errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

	endpoint, err := cli.BuildFullEndpoint(GetKeyInOrg, map[string]interface{}{ParamKeyID: key}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, headers, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request GetKeyInOrg")
	}
	decoded, err := decodeJSONResponse[v0.GetKeyInOrg200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

// TODO: CreateKeyImportKeyRequest => CreateKeyImportKey200Response
// TODO: DeleteKeyRequest => SetEmailOtp200Response
// TODO: DeriveKeyRequest => DeriveKey200Response
// TODO: ImportKeyRequest => DeriveKey200Response
// TODO: ListKeyRolesRequest => ListKeyRoles200Response
// TODO: ListKeysInOrgRequest => ListKeysInOrg200Response
// TODO: UpdateKeyRequest => GetKeyInOrg200Response
