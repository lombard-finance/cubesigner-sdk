package v0

import (
	"encoding/json"
)

// EvmSignResponse struct for EvmSignResponse
type EvmSignResponse struct {
	// The hex-encoded resulting signature.
	Signature string `json:"signature"`
}

// NewEvmSignResponse instantiates a new EvmSignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmSignResponse(signature string) *EvmSignResponse {
	this := EvmSignResponse{}
	this.Signature = signature
	return &this
}

// NewEvmSignResponseWithDefaults instantiates a new EvmSignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmSignResponseWithDefaults() *EvmSignResponse {
	this := EvmSignResponse{}
	return &this
}

// GetSignature returns the Signature field value
func (o *EvmSignResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *EvmSignResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *EvmSignResponse) SetSignature(v string) {
	o.Signature = v
}

func (o EvmSignResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableEvmSignResponse struct {
	value *EvmSignResponse
	isSet bool
}

func (v NullableEvmSignResponse) Get() *EvmSignResponse {
	return v.value
}

func (v *NullableEvmSignResponse) Set(val *EvmSignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmSignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmSignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmSignResponse(val *EvmSignResponse) *NullableEvmSignResponse {
	return &NullableEvmSignResponse{value: val, isSet: true}
}

func (v NullableEvmSignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmSignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
