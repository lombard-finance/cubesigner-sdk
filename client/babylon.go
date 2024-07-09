package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/lombard-finance/cubesigner-sdk/api"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) signBabylonStakingRequest(
	roleId, pubkey string,
	request interface{},
	scope api.Scope,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonStakingResponse, string, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign babylon staking",
		Scopes:  []api.Scope{scope},
	}, roleId)
	if err != nil {
		return nil, "", errors.Wrap(err, "create role token")
	}

	headers := map[string]string{
		"Authorization": authResp.GetToken(),
	}

	// add mfa headers
	if mfaConfirmation != nil && *mfaConfirmation != "" {
		mfaHeaders := getMfaHeaders(*mfaId, *mfaConfirmation, cli.orgID)
		for k, v := range mfaHeaders {
			headers[k] = v
		}
	}

	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, "", errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/babylon/staking/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		bodyBytes, err := io.ReadAll(encoded)
		if err != nil {
			return nil, "", err
		}
		return nil, "", errors.Wrap(err, fmt.Sprintf("request SignBabylonStaking: %s", string(bodyBytes)))
	}

	if statusCode == http.StatusAccepted {
		mfaId, err := decodeAcceptedResponse(response)
		if err != nil {
			return nil, "", errors.Wrap(err, "decode accepted response")
		}
		return nil, mfaId, nil
	}

	decoded, err := decodeJSONResponse[v0.BabylonStakingResponse](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode")
	}
	return &decoded, "", nil
}

func (cli *Client) SignBabylonStakingDeposit(
	roleId, pubkey string,
	request *v0.BabylonStakingDeposit,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonStakingResponse, string, error) {
	return cli.signBabylonStakingRequest(roleId, pubkey, request, api.SIGNBABYLONSTAKINGDEPOSIT, mfaId, mfaConfirmation)
}

func (cli *Client) SignBabylonStakingEarlyUnbond(
	roleId, pubkey string,
	request *v0.BabylonStakingEarlyUnbond,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonStakingResponse, string, error) {
	return cli.signBabylonStakingRequest(roleId, pubkey, request, api.SIGNBABYLONSTAKINGUNBOND, mfaId, mfaConfirmation)
}

func (cli *Client) SignBabylonStakingWithdrawal(
	roleId, pubkey string,
	request *v0.BabylonStakingWithdrawal,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonStakingResponse, string, error) {
	return cli.signBabylonStakingRequest(roleId, pubkey, request, api.SIGNBABYLONSTAKINGWITHDRAW, mfaId, mfaConfirmation)
}
