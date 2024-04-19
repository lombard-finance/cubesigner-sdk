package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// EditPolicyMfa struct for EditPolicyMfa
type EditPolicyMfa struct {
	// Users who are allowed to approve. If empty at creation time, default to the current user.
	AllowedApprovers []string `json:"allowed_approvers,omitempty"`
	// Allowed approval types. When omitted, defaults to any.
	AllowedMfaTypes []MfaType `json:"allowed_mfa_types,omitempty"`
	// How many users to require to approve (defaults to 1).
	Count *int32 `json:"count,omitempty"`
	// Duration measured in seconds A wrapper type for serialization that encodes a `Duration` as a `u64` representing the number of seconds.
	Lifetime *int64 `json:"lifetime,omitempty"`
	// How many auth factors to require per user (defaults to 1).
	NumAuthFactors *int32 `json:"num_auth_factors,omitempty"`
	// CubeSigner operations to which this policy should apply. When omitted, applies to all operations.
	RestrictedOperations []api.OperationKind `json:"restricted_operations,omitempty"`
}

// NewEditPolicyMfa instantiates a new EditPolicyMfa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditPolicyMfa() *EditPolicyMfa {
	this := EditPolicyMfa{}
	return &this
}

// NewEditPolicyMfaWithDefaults instantiates a new EditPolicyMfa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditPolicyMfaWithDefaults() *EditPolicyMfa {
	this := EditPolicyMfa{}
	return &this
}

// GetAllowedApprovers returns the AllowedApprovers field value if set, zero value otherwise.
func (o *EditPolicyMfa) GetAllowedApprovers() []string {
	if o == nil || o.AllowedApprovers == nil {
		var ret []string
		return ret
	}
	return o.AllowedApprovers
}

// GetAllowedApproversOk returns a tuple with the AllowedApprovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPolicyMfa) GetAllowedApproversOk() ([]string, bool) {
	if o == nil || o.AllowedApprovers == nil {
		return nil, false
	}
	return o.AllowedApprovers, true
}

// HasAllowedApprovers returns a boolean if a field has been set.
func (o *EditPolicyMfa) HasAllowedApprovers() bool {
	if o != nil && o.AllowedApprovers != nil {
		return true
	}

	return false
}

// SetAllowedApprovers gets a reference to the given []string and assigns it to the AllowedApprovers field.
func (o *EditPolicyMfa) SetAllowedApprovers(v []string) {
	o.AllowedApprovers = v
}

// GetAllowedMfaTypes returns the AllowedMfaTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditPolicyMfa) GetAllowedMfaTypes() []MfaType {
	if o == nil {
		var ret []MfaType
		return ret
	}
	return o.AllowedMfaTypes
}

// GetAllowedMfaTypesOk returns a tuple with the AllowedMfaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditPolicyMfa) GetAllowedMfaTypesOk() ([]MfaType, bool) {
	if o == nil || o.AllowedMfaTypes == nil {
		return nil, false
	}
	return o.AllowedMfaTypes, true
}

// HasAllowedMfaTypes returns a boolean if a field has been set.
func (o *EditPolicyMfa) HasAllowedMfaTypes() bool {
	if o != nil && o.AllowedMfaTypes != nil {
		return true
	}

	return false
}

// SetAllowedMfaTypes gets a reference to the given []MfaType and assigns it to the AllowedMfaTypes field.
func (o *EditPolicyMfa) SetAllowedMfaTypes(v []MfaType) {
	o.AllowedMfaTypes = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EditPolicyMfa) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPolicyMfa) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EditPolicyMfa) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *EditPolicyMfa) SetCount(v int32) {
	o.Count = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *EditPolicyMfa) GetLifetime() int64 {
	if o == nil || o.Lifetime == nil {
		var ret int64
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPolicyMfa) GetLifetimeOk() (*int64, bool) {
	if o == nil || o.Lifetime == nil {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *EditPolicyMfa) HasLifetime() bool {
	if o != nil && o.Lifetime != nil {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given int64 and assigns it to the Lifetime field.
func (o *EditPolicyMfa) SetLifetime(v int64) {
	o.Lifetime = &v
}

// GetNumAuthFactors returns the NumAuthFactors field value if set, zero value otherwise.
func (o *EditPolicyMfa) GetNumAuthFactors() int32 {
	if o == nil || o.NumAuthFactors == nil {
		var ret int32
		return ret
	}
	return *o.NumAuthFactors
}

// GetNumAuthFactorsOk returns a tuple with the NumAuthFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPolicyMfa) GetNumAuthFactorsOk() (*int32, bool) {
	if o == nil || o.NumAuthFactors == nil {
		return nil, false
	}
	return o.NumAuthFactors, true
}

// HasNumAuthFactors returns a boolean if a field has been set.
func (o *EditPolicyMfa) HasNumAuthFactors() bool {
	if o != nil && o.NumAuthFactors != nil {
		return true
	}

	return false
}

// SetNumAuthFactors gets a reference to the given int32 and assigns it to the NumAuthFactors field.
func (o *EditPolicyMfa) SetNumAuthFactors(v int32) {
	o.NumAuthFactors = &v
}

// GetRestrictedOperations returns the RestrictedOperations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditPolicyMfa) GetRestrictedOperations() []api.OperationKind {
	if o == nil {
		var ret []api.OperationKind
		return ret
	}
	return o.RestrictedOperations
}

// GetRestrictedOperationsOk returns a tuple with the RestrictedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditPolicyMfa) GetRestrictedOperationsOk() ([]api.OperationKind, bool) {
	if o == nil || o.RestrictedOperations == nil {
		return nil, false
	}
	return o.RestrictedOperations, true
}

// HasRestrictedOperations returns a boolean if a field has been set.
func (o *EditPolicyMfa) HasRestrictedOperations() bool {
	if o != nil && o.RestrictedOperations != nil {
		return true
	}

	return false
}

// SetRestrictedOperations gets a reference to the given []OperationKind and assigns it to the RestrictedOperations field.
func (o *EditPolicyMfa) SetRestrictedOperations(v []api.OperationKind) {
	o.RestrictedOperations = v
}

func (o EditPolicyMfa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedApprovers != nil {
		toSerialize["allowed_approvers"] = o.AllowedApprovers
	}
	if o.AllowedMfaTypes != nil {
		toSerialize["allowed_mfa_types"] = o.AllowedMfaTypes
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Lifetime != nil {
		toSerialize["lifetime"] = o.Lifetime
	}
	if o.NumAuthFactors != nil {
		toSerialize["num_auth_factors"] = o.NumAuthFactors
	}
	if o.RestrictedOperations != nil {
		toSerialize["restricted_operations"] = o.RestrictedOperations
	}
	return json.Marshal(toSerialize)
}

type NullableEditPolicyMfa struct {
	value *EditPolicyMfa
	isSet bool
}

func (v NullableEditPolicyMfa) Get() *EditPolicyMfa {
	return v.value
}

func (v *NullableEditPolicyMfa) Set(val *EditPolicyMfa) {
	v.value = val
	v.isSet = true
}

func (v NullableEditPolicyMfa) IsSet() bool {
	return v.isSet
}

func (v *NullableEditPolicyMfa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditPolicyMfa(val *EditPolicyMfa) *NullableEditPolicyMfa {
	return &NullableEditPolicyMfa{value: val, isSet: true}
}

func (v NullableEditPolicyMfa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditPolicyMfa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
