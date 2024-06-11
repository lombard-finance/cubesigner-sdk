package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

func (cli *Client) SignTaproot(roleId, pubkey string, request *v0.TaprootSignRequest) (*v0.TaprootSignResponse, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign taproot",
		Scopes:  []string{"sign:btc:taproot"},
	}, roleId)
	if err != nil {
		return nil, errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/btc/taproot/sign/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

	response, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request SignTaproot")
	}
	decoded, err := decodeJSONResponse[v0.TaprootSignResponse](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}

func (cli *Client) SignSegWit(roleId, pubkey string, request *v0.BtcSignRequest) (*v0.BtcSign200Response, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign segwit",
		Scopes:  []string{"sign:btc:segwit"},
	}, roleId)
	if err != nil {
		return nil, errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/btc/sign/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

	response, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request SignSegwit")
	}
	decoded, err := decodeJSONResponse[v0.BtcSign200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}
