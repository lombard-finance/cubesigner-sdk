package client

import (
	"fmt"
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) CreateKeyRequest(request *v0.CreateKeyRequest) (*v0.DeriveKey200Response, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	response, err := cli.post("/v0/org/:org_id/keys", encoded)
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
	response, err := cli.get(fmt.Sprintf("/v0/org/:org_id/keys/%s", key))
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
