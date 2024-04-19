package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// EditPolicy struct for EditPolicy
type EditPolicy struct {
	Mfa           NullableEditPolicyMfa `json:"mfa,omitempty"`
	TimeLockUntil api.NullableInt64     `json:"time_lock_until,omitempty"`
}

// NewEditPolicy instantiates a new EditPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditPolicy() *EditPolicy {
	this := EditPolicy{}
	return &this
}

// NewEditPolicyWithDefaults instantiates a new EditPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditPolicyWithDefaults() *EditPolicy {
	this := EditPolicy{}
	return &this
}

// GetMfa returns the Mfa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditPolicy) GetMfa() EditPolicyMfa {
	if o == nil || o.Mfa.Get() == nil {
		var ret EditPolicyMfa
		return ret
	}
	return *o.Mfa.Get()
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditPolicy) GetMfaOk() (*EditPolicyMfa, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mfa.Get(), o.Mfa.IsSet()
}

// HasMfa returns a boolean if a field has been set.
func (o *EditPolicy) HasMfa() bool {
	if o != nil && o.Mfa.IsSet() {
		return true
	}

	return false
}

// SetMfa gets a reference to the given NullableEditPolicyMfa and assigns it to the Mfa field.
func (o *EditPolicy) SetMfa(v EditPolicyMfa) {
	o.Mfa.Set(&v)
}

// SetMfaNil sets the value for Mfa to be an explicit nil
func (o *EditPolicy) SetMfaNil() {
	o.Mfa.Set(nil)
}

// UnsetMfa ensures that no value is present for Mfa, not even an explicit nil
func (o *EditPolicy) UnsetMfa() {
	o.Mfa.Unset()
}

// GetTimeLockUntil returns the TimeLockUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditPolicy) GetTimeLockUntil() int64 {
	if o == nil || o.TimeLockUntil.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TimeLockUntil.Get()
}

// GetTimeLockUntilOk returns a tuple with the TimeLockUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditPolicy) GetTimeLockUntilOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeLockUntil.Get(), o.TimeLockUntil.IsSet()
}

// HasTimeLockUntil returns a boolean if a field has been set.
func (o *EditPolicy) HasTimeLockUntil() bool {
	if o != nil && o.TimeLockUntil.IsSet() {
		return true
	}

	return false
}

// SetTimeLockUntil gets a reference to the given NullableInt64 and assigns it to the TimeLockUntil field.
func (o *EditPolicy) SetTimeLockUntil(v int64) {
	o.TimeLockUntil.Set(&v)
}

// SetTimeLockUntilNil sets the value for TimeLockUntil to be an explicit nil
func (o *EditPolicy) SetTimeLockUntilNil() {
	o.TimeLockUntil.Set(nil)
}

// UnsetTimeLockUntil ensures that no value is present for TimeLockUntil, not even an explicit nil
func (o *EditPolicy) UnsetTimeLockUntil() {
	o.TimeLockUntil.Unset()
}

func (o EditPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mfa.IsSet() {
		toSerialize["mfa"] = o.Mfa.Get()
	}
	if o.TimeLockUntil.IsSet() {
		toSerialize["time_lock_until"] = o.TimeLockUntil.Get()
	}
	return json.Marshal(toSerialize)
}

//
//type NullableEditPolicy struct {
//	value *EditPolicy
//	isSet bool
//}
//
//func (v NullableEditPolicy) Get() *EditPolicy {
//	return v.value
//}
//
//func (v *NullableEditPolicy) Set(val *EditPolicy) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableEditPolicy) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableEditPolicy) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableEditPolicy(val *EditPolicy) *NullableEditPolicy {
//	return &NullableEditPolicy{value: val, isSet: true}
//}
//
//func (v NullableEditPolicy) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableEditPolicy) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
