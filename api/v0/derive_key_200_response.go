package v0

import (
	"encoding/json"
)

// DeriveKey200Response struct for DeriveKey200Response
type DeriveKey200Response struct {
	// The info about the created keys
	Keys []KeyInfo `json:"keys"`
}

// NewDeriveKey200Response instantiates a new DeriveKey200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeriveKey200Response(keys []KeyInfo) *DeriveKey200Response {
	this := DeriveKey200Response{}
	this.Keys = keys
	return &this
}

// NewDeriveKey200ResponseWithDefaults instantiates a new DeriveKey200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeriveKey200ResponseWithDefaults() *DeriveKey200Response {
	this := DeriveKey200Response{}
	return &this
}

// GetKeys returns the Keys field value
func (o *DeriveKey200Response) GetKeys() []KeyInfo {
	if o == nil {
		var ret []KeyInfo
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *DeriveKey200Response) GetKeysOk() ([]KeyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *DeriveKey200Response) SetKeys(v []KeyInfo) {
	o.Keys = v
}

func (o DeriveKey200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableDeriveKey200Response struct {
	value *DeriveKey200Response
	isSet bool
}

func (v NullableDeriveKey200Response) Get() *DeriveKey200Response {
	return v.value
}

func (v *NullableDeriveKey200Response) Set(val *DeriveKey200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeriveKey200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeriveKey200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeriveKey200Response(val *DeriveKey200Response) *NullableDeriveKey200Response {
	return &NullableDeriveKey200Response{value: val, isSet: true}
}

func (v NullableDeriveKey200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeriveKey200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
