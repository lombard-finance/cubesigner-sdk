package v0

import (
	"encoding/json"
)

// MfaTypeOneOf3FidoKey Answer a FIDO challenge with a specific FIDO key
type MfaTypeOneOf3FidoKey struct {
	// The ID of the FIDO key that must be use to approve the request
	KeyId string `json:"key_id"`
}

// NewMfaTypeOneOf3FidoKey instantiates a new MfaTypeOneOf3FidoKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaTypeOneOf3FidoKey(keyId string) *MfaTypeOneOf3FidoKey {
	this := MfaTypeOneOf3FidoKey{}
	this.KeyId = keyId
	return &this
}

// NewMfaTypeOneOf3FidoKeyWithDefaults instantiates a new MfaTypeOneOf3FidoKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaTypeOneOf3FidoKeyWithDefaults() *MfaTypeOneOf3FidoKey {
	this := MfaTypeOneOf3FidoKey{}
	return &this
}

// GetKeyId returns the KeyId field value
func (o *MfaTypeOneOf3FidoKey) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *MfaTypeOneOf3FidoKey) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *MfaTypeOneOf3FidoKey) SetKeyId(v string) {
	o.KeyId = v
}

func (o MfaTypeOneOf3FidoKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	return json.Marshal(toSerialize)
}

//
//type NullableMfaTypeOneOf3FidoKey struct {
//	value *MfaTypeOneOf3FidoKey
//	isSet bool
//}
//
//func (v NullableMfaTypeOneOf3FidoKey) Get() *MfaTypeOneOf3FidoKey {
//	return v.value
//}
//
//func (v *NullableMfaTypeOneOf3FidoKey) Set(val *MfaTypeOneOf3FidoKey) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMfaTypeOneOf3FidoKey) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMfaTypeOneOf3FidoKey) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMfaTypeOneOf3FidoKey(val *MfaTypeOneOf3FidoKey) *NullableMfaTypeOneOf3FidoKey {
//	return &NullableMfaTypeOneOf3FidoKey{value: val, isSet: true}
//}
//
//func (v NullableMfaTypeOneOf3FidoKey) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMfaTypeOneOf3FidoKey) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
