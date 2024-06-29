package v1

import (
	"encoding/json"
)

// BlobSignResponse struct for BlobSignResponse
type BlobSignResponse struct {
	// The hex-encoded signature.
	Signature string `json:"signature"`
}

// NewBlobSignResponse instantiates a new BlobSignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobSignResponse(signature string) *BlobSignResponse {
	this := BlobSignResponse{}
	this.Signature = signature
	return &this
}

// NewBlobSignResponseWithDefaults instantiates a new BlobSignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobSignResponseWithDefaults() *BlobSignResponse {
	this := BlobSignResponse{}
	return &this
}

// GetSignature returns the Signature field value
func (o *BlobSignResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *BlobSignResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *BlobSignResponse) SetSignature(v string) {
	o.Signature = v
}

func (o BlobSignResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableBlobSignResponse struct {
	value *BlobSignResponse
	isSet bool
}

func (v NullableBlobSignResponse) Get() *BlobSignResponse {
	return v.value
}

func (v *NullableBlobSignResponse) Set(val *BlobSignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobSignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobSignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobSignResponse(val *BlobSignResponse) *NullableBlobSignResponse {
	return &NullableBlobSignResponse{value: val, isSet: true}
}

func (v NullableBlobSignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobSignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
