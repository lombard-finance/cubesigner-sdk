package client

import (
	"fmt"
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) CreateKeyRequest(org string, request *v0.CreateKeyRequest) (*v0.DeriveKey200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	response, err := cli.post(fmt.Sprintf("/v0/org/%s/keys", org), encoded)
	if err != nil {
		return nil, errors.Wrap(err, "request CreateKeyRequest")
	}
	decoded, err := decodeJSONResponse[v0.DeriveKey200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}

func (cli *Client) GetKeyInOrg(org string, key string) (*v0.GetKeyInOrg200Response, error) {
	response, err := cli.get(fmt.Sprintf("/v0/org/%s/keys/%s", org, key))
	if err != nil {
		return nil, errors.Wrap(err, "request GetKeyInOrg")
	}
	decoded, err := decodeJSONResponse[v0.GetKeyInOrg200Response](response)
	return &decoded, nil
}

// TODO: CreateKeyImportKeyRequest => CreateKeyImportKey200Response
// TODO: DeleteKeyRequest => SetEmailOtp200Response
// TODO: DeriveKeyRequest => DeriveKey200Response
// TODO: ImportKeyRequest => DeriveKey200Response
// TODO: ListKeyRolesRequest => ListKeyRoles200Response
// TODO: ListKeysInOrgRequest => ListKeysInOrg200Response
// TODO: UpdateKeyRequest => GetKeyInOrg200Response
