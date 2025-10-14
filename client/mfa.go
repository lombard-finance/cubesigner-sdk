package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

// Retrieves and returns all pending MFA requests that are accessible to the
// current user, i.e., those in which the current user is listed as an approver.
// No pagination.
func (cli *Client) ListMfaRequests() (*v0.ListMfaResponse, error) {
	endpoint, err := cli.BuildFullEndpoint(ListMfaRequests, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request ListMfaRequests")
	}
	decoded, err := decodeJSONResponse[v0.ListMfaResponse](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

// Retrieves and returns a pending MFA request by its id.
func (cli *Client) GetMfaRequest(mfaId string) (*v0.MfaRequestInfo, error) {
	endpoint, err := cli.BuildFullEndpoint(MfaRequest, map[string]interface{}{ParamMfaID: mfaId}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, err := cli.get(endpoint, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request GetMfaRequest")
	}
	decoded, err := decodeJSONResponse[v0.MfaRequestInfo](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

// Approve or Reject MFA Request after logging in with CubeSigner.
// If approving, adds the currently-logged user as an approver of a pending MFA
// request of the [Status::RequiredApprovers] kind. If the required number of
// approvers is reached, the MFA request is approved; the confirmation receipt
// can be used to resume the original HTTP request.
// If rejecting, immediately deletes the pending MFA request.
func (cli *Client) ApproveOrRejectMfaRequest(mfaId string, mfaVote v0.MfaVote) (*v0.MfaRequestInfo, error) {
	if !mfaVote.IsValid() {
		return nil, errors.New("invalid MfaVote value")
	}

	endpoint, err := cli.BuildFullEndpoint(MfaRequest, map[string]interface{}{ParamMfaID: mfaId}, map[string]interface{}{QueryParamVote: mfaVote})
	if err != nil {
		return nil, errors.Wrap(err, "build endpoint")
	}

	response, _, err := cli.patch(endpoint, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request ApproveOrRejectMfaRequest")
	}
	decoded, err := decodeJSONResponse[v0.MfaRequestInfo](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}
