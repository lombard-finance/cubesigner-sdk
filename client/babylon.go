package client

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/lombard-finance/cubesigner-sdk/api"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) StakingBabylon(
	roleId, pubkey string,
	request *v0.BabylonStakingRequest,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonStakingResponse, string, error) {
	switch request.GetActualInstance().(type) {
	case v0.BabylonStakingDeposit:
		return cli.stakingBabylon(roleId, pubkey, request, api.SIGNBABYLONSTAKINGDEPOSIT, mfaId, mfaConfirmation)
	case v0.BabylonStakingEarlyUnbond:
		return cli.stakingBabylon(roleId, pubkey, request, api.SIGNBABYLONSTAKINGUNBOND, mfaId, mfaConfirmation)
	case v0.BabylonStakingWithdrawal:
		return cli.stakingBabylon(roleId, pubkey, request, api.SIGNBABYLONSTAKINGWITHDRAW, mfaId, mfaConfirmation)
	}

	return nil, "", errors.New("not implemented")
}

func (cli *Client) stakingBabylon(
	roleId, pubkey string,
	request *v0.BabylonStakingRequest,
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
	endpoint := strings.Replace("/v0/org/:org_id/babylon/staking/:pubkey", ":pubkey", url.PathEscape(parameterToString(pubkey, "")), -1)

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "request SignBabylonStaking")
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
