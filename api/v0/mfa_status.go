package v0

// MfaStatus struct
type MfaStatus struct {
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

// NewMfaStatus instantiates a new MfaStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaStatus(allowedApprovers []string, allowedMfaTypes []MfaType, approvedBy map[string]MfaApprovalInfo, count int32, numAuthFactors int32) *MfaStatus {
	this := MfaStatus{}
	this.AllowedApprovers = allowedApprovers
	this.AllowedMfaTypes = allowedMfaTypes
	this.ApprovedBy = approvedBy
	this.Count = count
	this.NumAuthFactors = numAuthFactors
	return &this
}

// NewMfaStatusWithDefaults instantiates a new MfaStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaStatusWithDefaults() *MfaStatus {
	this := MfaStatus{}
	return &this
}

// GetAllowedApprovers returns the AllowedApprovers field value
func (o *MfaStatus) GetAllowedApprovers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedApprovers
}

// GetAllowedApproversOk returns a tuple with the AllowedApprovers field value
// and a boolean to check if the value has been set.
func (o *MfaStatus) GetAllowedApproversOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedApprovers, true
}

// SetAllowedApprovers sets the AllowedApprovers field value
func (o *MfaStatus) SetAllowedApprovers(v []string) {
	o.AllowedApprovers = v
}

// GetAllowedMfaTypes returns the AllowedMfaTypes field value if set, zero value otherwise
func (o *MfaStatus) GetAllowedMfaTypes() []MfaType {
	if o == nil || o.AllowedMfaTypes == nil {
		var ret []MfaType
		return ret
	}
	return o.AllowedMfaTypes
}

// GetAllowedMfaTypesOk returns a tuple with the AllowedMfaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaStatus) GetAllowedMfaTypesOk() (*[]MfaType, bool) {
	if o == nil || o.AllowedMfaTypes == nil {
		return nil, false
	}
	return &o.AllowedMfaTypes, true
}

// SetAllowedMfaTypes sets the AllowedMfaTypes field value
func (o *MfaStatus) SetAllowedMfaTypes(v []MfaType) {
	o.AllowedMfaTypes = v
}

// GetApprovedBy returns the ApprovedBy field value
func (o *MfaStatus) GetApprovedBy() map[string]MfaApprovalInfo {
	if o == nil {
		var ret map[string]MfaApprovalInfo
		return ret
	}
	return o.ApprovedBy
}

// GetApprovedByOk returns a tuple with the ApprovedBy field value
// and a boolean to check if the value has been set.
func (o *MfaStatus) GetApprovedByOk() (*map[string]MfaApprovalInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovedBy, true
}

// SetApprovedBy sets the ApprovedBy field value
func (o *MfaStatus) SetApprovedBy(v map[string]MfaApprovalInfo) {
	o.ApprovedBy = v
}

// GetCount returns the Count field value
func (o *MfaStatus) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MfaStatus) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets the Count field value
func (o *MfaStatus) SetCount(v int32) {
	o.Count = v
}

// GetNumAuthFactors returns the NumAuthFactors field value
func (o *MfaStatus) GetNumAuthFactors() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.NumAuthFactors
}

// GetNumAuthFactorsOk returns a tuple with the NumAuthFactors field value
// and a boolean to check if the value has been set.
func (o *MfaStatus) GetNumAuthFactorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAuthFactors, true
}

// SetNumAuthFactors sets the NumAuthFactors field value
func (o *MfaStatus) SetNumAuthFactors(v int32) {
	o.NumAuthFactors = v
}
