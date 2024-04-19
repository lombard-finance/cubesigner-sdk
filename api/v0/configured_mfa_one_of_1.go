package v0

import (
	"encoding/json"
)

// ConfiguredMfaOneOf1 Named FIDO device (multiple can be configured per user)
type ConfiguredMfaOneOf1 struct {
	// A unique credential id
	Id string `json:"id"`
	// A human-readable name given to the key
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewConfiguredMfaOneOf1 instantiates a new ConfiguredMfaOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguredMfaOneOf1(id string, name string, type_ string) *ConfiguredMfaOneOf1 {
	this := ConfiguredMfaOneOf1{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewConfiguredMfaOneOf1WithDefaults instantiates a new ConfiguredMfaOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguredMfaOneOf1WithDefaults() *ConfiguredMfaOneOf1 {
	this := ConfiguredMfaOneOf1{}
	return &this
}

// GetId returns the Id field value
func (o *ConfiguredMfaOneOf1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfiguredMfaOneOf1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfiguredMfaOneOf1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConfiguredMfaOneOf1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfiguredMfaOneOf1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfiguredMfaOneOf1) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ConfiguredMfaOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfiguredMfaOneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfiguredMfaOneOf1) SetType(v string) {
	o.Type = v
}

func (o ConfiguredMfaOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

//
//type NullableConfiguredMfaOneOf1 struct {
//	value *ConfiguredMfaOneOf1
//	isSet bool
//}
//
//func (v NullableConfiguredMfaOneOf1) Get() *ConfiguredMfaOneOf1 {
//	return v.value
//}
//
//func (v *NullableConfiguredMfaOneOf1) Set(val *ConfiguredMfaOneOf1) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableConfiguredMfaOneOf1) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableConfiguredMfaOneOf1) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableConfiguredMfaOneOf1(val *ConfiguredMfaOneOf1) *NullableConfiguredMfaOneOf1 {
//	return &NullableConfiguredMfaOneOf1{value: val, isSet: true}
//}
//
//func (v NullableConfiguredMfaOneOf1) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableConfiguredMfaOneOf1) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
