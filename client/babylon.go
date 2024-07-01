package client

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/lombard-finance/cubesigner-sdk/api"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

func (cli *Client) SignBabylonStaking(roleId, pubkey string, request *v0.BabylonStakingRequest, mfaId *string, mfaConfirmation *string) (*v0.BabylonStakingResponse, string, error) {
	authResp, err := cli.CreateRoleToken(&v0.CreateTokenRequest{
		Purpose: "sign babylon staking",
		Scopes:  []api.Scope{api.SIGNBABYLONSTAKING},
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
