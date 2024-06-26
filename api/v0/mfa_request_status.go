package v0

// MfaRequestStatus struct
type MfaRequestStatus struct {
	// Users who are allowed to approve. Must be non-empty.
	AllowedApprovers []string `json:"allowed_approvers"`
	// Allowed approval types. When omitted, defaults to any.
	AllowedMfaTypes []MfaType `json:"allowed_mfa_types,omitempty"`
	// Users who have already approved.
	ApprovedBy map[string]MfaApprovalInfo `json:"approved_by"`
	// How many users must approve (minimum is 0).
	Count int32 `json:"count"`
	// How many auth factors to require per user (minimum is 0).
	NumAuthFactors int32 `json:"num_auth_factors"`
	// TODO: add request_comparer that only has Eq option
}

// NewMfaRequestStatus instantiates a new MfaRequestStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaRequestStatus(allowedApprovers []string, allowedMfaTypes []MfaType, approvedBy map[string]MfaApprovalInfo, count int32, numAuthFactors int32) *MfaRequestStatus {
	this := MfaRequestStatus{}
	this.AllowedApprovers = allowedApprovers
	this.AllowedMfaTypes = allowedMfaTypes
	this.ApprovedBy = approvedBy
	this.Count = count
	this.NumAuthFactors = numAuthFactors
	return &this
}

// NewMfaRequestStatusWithDefaults instantiates a new MfaRequestStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaRequestStatusWithDefaults() *MfaRequestStatus {
	this := MfaRequestStatus{}
	return &this
}

// GetAllowedApprovers returns the AllowedApprovers field value
func (o *MfaRequestStatus) GetAllowedApprovers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedApprovers
}

// GetAllowedApproversOk returns a tuple with the AllowedApprovers field value
// and a boolean to check if the value has been set.
func (o *MfaRequestStatus) GetAllowedApproversOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedApprovers, true
}

// SetAllowedApprovers sets the AllowedApprovers field value
func (o *MfaRequestStatus) SetAllowedApprovers(v []string) {
	o.AllowedApprovers = v
}

// GetAllowedMfaTypes returns the AllowedMfaTypes field value if set, zero value otherwise
func (o *MfaRequestStatus) GetAllowedMfaTypes() []MfaType {
	if o == nil || o.AllowedMfaTypes == nil {
		var ret []MfaType
		return ret
	}
	return o.AllowedMfaTypes
}

// GetAllowedMfaTypesOk returns a tuple with the AllowedMfaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaRequestStatus) GetAllowedMfaTypesOk() (*[]MfaType, bool) {
	if o == nil || o.AllowedMfaTypes == nil {
		return nil, false
	}
	return &o.AllowedMfaTypes, true
}

// SetAllowedMfaTypes sets the AllowedMfaTypes field value
func (o *MfaRequestStatus) SetAllowedMfaTypes(v []MfaType) {
	o.AllowedMfaTypes = v
}

// GetApprovedBy returns the ApprovedBy field value
func (o *MfaRequestStatus) GetApprovedBy() map[string]MfaApprovalInfo {
	if o == nil {
		var ret map[string]MfaApprovalInfo
		return ret
	}
	return o.ApprovedBy
}

// GetApprovedByOk returns a tuple with the ApprovedBy field value
// and a boolean to check if the value has been set.
func (o *MfaRequestStatus) GetApprovedByOk() (*map[string]MfaApprovalInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovedBy, true
}

// SetApprovedBy sets the ApprovedBy field value
func (o *MfaRequestStatus) SetApprovedBy(v map[string]MfaApprovalInfo) {
	o.ApprovedBy = v
}

// GetCount returns the Count field value
func (o *MfaRequestStatus) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MfaRequestStatus) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets the Count field value
func (o *MfaRequestStatus) SetCount(v int32) {
	o.Count = v
}

// GetNumAuthFactors returns the NumAuthFactors field value
func (o *MfaRequestStatus) GetNumAuthFactors() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.NumAuthFactors
}

// GetNumAuthFactorsOk returns a tuple with the NumAuthFactors field value
// and a boolean to check if the value has been set.
func (o *MfaRequestStatus) GetNumAuthFactorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAuthFactors, true
}

// SetNumAuthFactors sets the NumAuthFactors field value
func (o *MfaRequestStatus) SetNumAuthFactors(v int32) {
	o.NumAuthFactors = v
}
