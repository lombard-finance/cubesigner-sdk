package v0

import (
	"encoding/json"
)

// TaprootSignResponse struct for TaprootSignResponse
type TaprootSignResponse struct {
	// The 64-byte signature, encoded as defined in BIP0340.
	Signature string `json:"signature"`
}

// NewTaprootSignResponse instantiates a new TaprootSignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaprootSignResponse(signature string) *TaprootSignResponse {
	this := TaprootSignResponse{}
	this.Signature = signature
	return &this
}

// NewTaprootSignResponseWithDefaults instantiates a new TaprootSignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaprootSignResponseWithDefaults() *TaprootSignResponse {
	this := TaprootSignResponse{}
	return &this
}

// GetSignature returns the Signature field value
func (o *TaprootSignResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *TaprootSignResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *TaprootSignResponse) SetSignature(v string) {
	o.Signature = v
}

func (o TaprootSignResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableTaprootSignResponse struct {
	value *TaprootSignResponse
	isSet bool
}

func (v NullableTaprootSignResponse) Get() *TaprootSignResponse {
	return v.value
}

func (v *NullableTaprootSignResponse) Set(val *TaprootSignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaprootSignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaprootSignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaprootSignResponse(val *TaprootSignResponse) *NullableTaprootSignResponse {
	return &NullableTaprootSignResponse{value: val, isSet: true}
}

func (v NullableTaprootSignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaprootSignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
