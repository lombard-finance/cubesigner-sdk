package v0

// ListMfaResponse struct represents all pending MFA requests
type ListMfaResponse struct {
	MfaRequests []MfaRequestInfo `json:"mfa_requests"`
}

// NewListMfaResponse instantiates a new ListMfaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMfaResponse(mfaRequests []MfaRequestInfo) *ListMfaResponse {
	this := ListMfaResponse{}
	this.MfaRequests = mfaRequests
	return &this
}

// NewListMfaResponseWithDefaults instantiates a new ListMfaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMfaResponseWithDefaults() *ListMfaResponse {
	this := ListMfaResponse{}
	return &this
}

// GetMfaRequests returns the MfaRequests field value
func (o *ListMfaResponse) GetMfaRequests() []MfaRequestInfo {
	if o == nil {
		var ret []MfaRequestInfo
		return ret
	}

	return o.MfaRequests
}

// GetMfaRequestsOk returns a tuple with the MfaRequests field value
// and a boolean to check if the value has been set.
func (o *ListMfaResponse) GetMfaRequestsOk() (*[]MfaRequestInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaRequests, true
}

// SetMfaRequests sets the MfaRequests field value
func (o *ListMfaResponse) SetMfaRequests(v []MfaRequestInfo) {
	o.MfaRequests = v
}
