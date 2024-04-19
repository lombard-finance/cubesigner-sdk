package v0

import (
	"encoding/json"
)

// ConfiguredMfaOneOf struct for ConfiguredMfaOneOf
type ConfiguredMfaOneOf struct {
	Type string `json:"type"`
}

// NewConfiguredMfaOneOf instantiates a new ConfiguredMfaOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguredMfaOneOf(type_ string) *ConfiguredMfaOneOf {
	this := ConfiguredMfaOneOf{}
	this.Type = type_
	return &this
}

// NewConfiguredMfaOneOfWithDefaults instantiates a new ConfiguredMfaOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguredMfaOneOfWithDefaults() *ConfiguredMfaOneOf {
	this := ConfiguredMfaOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *ConfiguredMfaOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfiguredMfaOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfiguredMfaOneOf) SetType(v string) {
	o.Type = v
}

func (o ConfiguredMfaOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

//
//type NullableConfiguredMfaOneOf struct {
//	value *ConfiguredMfaOneOf
//	isSet bool
//}
//
//func (v NullableConfiguredMfaOneOf) Get() *ConfiguredMfaOneOf {
//	return v.value
//}
//
//func (v *NullableConfiguredMfaOneOf) Set(val *ConfiguredMfaOneOf) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableConfiguredMfaOneOf) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableConfiguredMfaOneOf) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableConfiguredMfaOneOf(val *ConfiguredMfaOneOf) *NullableConfiguredMfaOneOf {
//	return &NullableConfiguredMfaOneOf{value: val, isSet: true}
//}
//
//func (v NullableConfiguredMfaOneOf) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableConfiguredMfaOneOf) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
