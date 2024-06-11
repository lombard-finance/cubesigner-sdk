package v0

import (
	"encoding/json"
)

// BtcSignResponse struct for BtcSignResponse
type BtcSignResponse struct {
	// The hex-encoded signature in compact format.
	Signature string `json:"signature"`
}

// NewBtcSignResponse instantiates a new BtcSignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcSignResponse(signature string) *BtcSignResponse {
	this := BtcSignResponse{}
	this.Signature = signature
	return &this
}

// NewBtcSignResponseWithDefaults instantiates a new BtcSignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcSignResponseWithDefaults() *BtcSignResponse {
	this := BtcSignResponse{}
	return &this
}

// GetSignature returns the Signature field value
func (o *BtcSignResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *BtcSignResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *BtcSignResponse) SetSignature(v string) {
	o.Signature = v
}

func (o BtcSignResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableBtcSignResponse struct {
	value *BtcSignResponse
	isSet bool
}

func (v NullableBtcSignResponse) Get() *BtcSignResponse {
	return v.value
}

func (v *NullableBtcSignResponse) Set(val *BtcSignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcSignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcSignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcSignResponse(val *BtcSignResponse) *NullableBtcSignResponse {
	return &NullableBtcSignResponse{value: val, isSet: true}
}

func (v NullableBtcSignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcSignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
