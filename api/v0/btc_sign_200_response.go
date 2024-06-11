package v0

import (
	"encoding/json"
)

// BtcSign200Response struct for BtcSign200Response
type BtcSign200Response struct {
	// The hex-encoded signature in compact format.
	Signature string `json:"signature"`
}

// NewBtcSign200Response instantiates a new BtcSign200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcSign200Response(signature string) *BtcSign200Response {
	this := BtcSign200Response{}
	this.Signature = signature
	return &this
}

// NewBtcSign200ResponseWithDefaults instantiates a new BtcSign200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcSign200ResponseWithDefaults() *BtcSign200Response {
	this := BtcSign200Response{}
	return &this
}

// GetSignature returns the Signature field value
func (o *BtcSign200Response) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *BtcSign200Response) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *BtcSign200Response) SetSignature(v string) {
	o.Signature = v
}

func (o BtcSign200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableBtcSign200Response struct {
	value *BtcSign200Response
	isSet bool
}

func (v NullableBtcSign200Response) Get() *BtcSign200Response {
	return v.value
}

func (v *NullableBtcSign200Response) Set(val *BtcSign200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcSign200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcSign200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcSign200Response(val *BtcSign200Response) *NullableBtcSign200Response {
	return &NullableBtcSign200Response{value: val, isSet: true}
}

func (v NullableBtcSign200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcSign200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
