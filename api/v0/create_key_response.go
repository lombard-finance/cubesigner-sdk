package v0

import (
	"encoding/json"
)

// CreateKeyResponse struct for CreateKeyResponse
type CreateKeyResponse struct {
	// The info about the created keys
	Keys []KeyInfo `json:"keys"`
}

// NewCreateKeyResponse instantiates a new CreateKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyResponse(keys []KeyInfo) *CreateKeyResponse {
	this := CreateKeyResponse{}
	this.Keys = keys
	return &this
}

// NewCreateKeyResponseWithDefaults instantiates a new CreateKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyResponseWithDefaults() *CreateKeyResponse {
	this := CreateKeyResponse{}
	return &this
}

// GetKeys returns the Keys field value
func (o *CreateKeyResponse) GetKeys() []KeyInfo {
	if o == nil {
		var ret []KeyInfo
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *CreateKeyResponse) GetKeysOk() ([]KeyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *CreateKeyResponse) SetKeys(v []KeyInfo) {
	o.Keys = v
}

func (o CreateKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

//
//type NullableCreateKeyResponse struct {
//	value *CreateKeyResponse
//	isSet bool
//}
//
//func (v NullableCreateKeyResponse) Get() *CreateKeyResponse {
//	return v.value
//}
//
//func (v *NullableCreateKeyResponse) Set(val *CreateKeyResponse) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableCreateKeyResponse) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableCreateKeyResponse) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableCreateKeyResponse(val *CreateKeyResponse) *NullableCreateKeyResponse {
//	return &NullableCreateKeyResponse{value: val, isSet: true}
//}
//
//func (v NullableCreateKeyResponse) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableCreateKeyResponse) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
