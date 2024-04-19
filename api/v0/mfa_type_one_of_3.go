package v0

import (
	"encoding/json"
)

// MfaTypeOneOf3 struct for MfaTypeOneOf3
type MfaTypeOneOf3 struct {
	FidoKey MfaTypeOneOf3FidoKey `json:"FidoKey"`
}

// NewMfaTypeOneOf3 instantiates a new MfaTypeOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaTypeOneOf3(fidoKey MfaTypeOneOf3FidoKey) *MfaTypeOneOf3 {
	this := MfaTypeOneOf3{}
	this.FidoKey = fidoKey
	return &this
}

// NewMfaTypeOneOf3WithDefaults instantiates a new MfaTypeOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaTypeOneOf3WithDefaults() *MfaTypeOneOf3 {
	this := MfaTypeOneOf3{}
	return &this
}

// GetFidoKey returns the FidoKey field value
func (o *MfaTypeOneOf3) GetFidoKey() MfaTypeOneOf3FidoKey {
	if o == nil {
		var ret MfaTypeOneOf3FidoKey
		return ret
	}

	return o.FidoKey
}

// GetFidoKeyOk returns a tuple with the FidoKey field value
// and a boolean to check if the value has been set.
func (o *MfaTypeOneOf3) GetFidoKeyOk() (*MfaTypeOneOf3FidoKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FidoKey, true
}

// SetFidoKey sets field value
func (o *MfaTypeOneOf3) SetFidoKey(v MfaTypeOneOf3FidoKey) {
	o.FidoKey = v
}

func (o MfaTypeOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["FidoKey"] = o.FidoKey
	}
	return json.Marshal(toSerialize)
}

//
//type NullableMfaTypeOneOf3 struct {
//	value *MfaTypeOneOf3
//	isSet bool
//}
//
//func (v NullableMfaTypeOneOf3) Get() *MfaTypeOneOf3 {
//	return v.value
//}
//
//func (v *NullableMfaTypeOneOf3) Set(val *MfaTypeOneOf3) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMfaTypeOneOf3) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMfaTypeOneOf3) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMfaTypeOneOf3(val *MfaTypeOneOf3) *NullableMfaTypeOneOf3 {
//	return &NullableMfaTypeOneOf3{value: val, isSet: true}
//}
//
//func (v NullableMfaTypeOneOf3) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMfaTypeOneOf3) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
