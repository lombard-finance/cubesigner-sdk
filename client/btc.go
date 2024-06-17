package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
)

func (cli *Client) SignTaproot(roleId, pubkey string, request *v0.TaprootSignRequest) (*v0.TaprootSignResponse, string, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign taproot",
		Scopes:  []string{"sign:btc:taproot"},
	}, roleId)
	if err != nil {
		return nil, "", errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, "", errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/btc/taproot/sign/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "request SignTaproot")
	}

	if statusCode == http.StatusAccepted {
		mfaId, err := decodeAcceptedResponse(response)
		if err != nil {
			return nil, "", errors.Wrap(err, "decode accepted response")
		}
		return nil, mfaId, nil
	}

	decoded, err := decodeJSONResponse[v0.TaprootSignResponse](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode")
	}
	return &decoded, "", nil
}

func (cli *Client) SignSegWit(roleId, pubkey string, request *v0.BtcSignRequest) (*v0.BtcSign200Response, string, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign segwit",
		Scopes:  []string{"sign:btc:segwit"},
	}, roleId)
	if err != nil {
		return nil, "", errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, "", errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/btc/sign/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "request SignSegwit")
	}

	if statusCode == http.StatusAccepted {
		mfaId, err := decodeAcceptedResponse(response)
		if err != nil {
			return nil, "", errors.Wrap(err, "decode accepted response")
		}
		return nil, mfaId, nil
	}

	decoded, err := decodeJSONResponse[v0.BtcSign200Response](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode response")
	}
	return &decoded, "", nil
}
