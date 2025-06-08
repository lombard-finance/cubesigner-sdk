package client

import (
	"net/http"

	"github.com/lombard-finance/cubesigner-sdk/api"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) SignBabylonStaking(
	roleId, pubkey string,
	request *v0.BabylonStakingRequest,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonStaking200Response, string, error) {
	var scope api.Scope
	switch request.Action {
	case api.DepositAction:
		scope = api.SIGNBABYLONSTAKINGDEPOSIT
	case api.EarlyUnbondAction:
		scope = api.SIGNBABYLONSTAKINGUNBOND
	case api.WithdrawEarlyUnbondAction, api.WithdrawTimelockAction:
		scope = api.SIGNBABYLONSTAKINGWITHDRAW
	case api.SlashDepositAction, api.SlashEarlyUnbondAction, api.WithdrawSlashing:
		scope = api.SIGNBABYLONSTAKINGSLASH
	default:
		return nil, "", errors.New("not implemented")
	}

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

	endpoint, err := cli.BuildFullEndpoint(SignBabylonStaking, map[string]interface{}{ParamPubkey: pubkey}, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "build endpoint")
	}

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

	decoded, err := decodeJSONResponse[v0.BabylonStaking200Response](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode")
	}
	return &decoded, "", nil
}

func (cli *Client) SignBabylonRegistration(
	roleId, pubkey string,
	request *v0.BabylonRegistrationRequest,
	mfaId, mfaConfirmation *string,
) (*v0.BabylonRegistration200Response, string, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign babylon registration",
		Scopes:  []api.Scope{api.SIGNBABYLONREGISTRATION},
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

	endpoint, err := cli.BuildFullEndpoint(SignBabylonRegistration, map[string]interface{}{ParamPubkey: pubkey}, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "build endpoint")
	}

	response, statusCode, err := cli.post(endpoint, encoded, headers, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "request SignBabylonRegistration")
	}

	if statusCode == http.StatusAccepted {
		mfaId, err := decodeAcceptedResponse(response)
		if err != nil {
			return nil, "", errors.Wrap(err, "decode accepted response")
		}
		return nil, mfaId, nil
	}

	decoded, err := decodeJSONResponse[v0.BabylonRegistration200Response](response)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode")
	}
	return &decoded, "", nil
}
