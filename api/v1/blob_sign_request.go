package v1

import (
	"encoding/json"
)

// BlobSignRequest struct for BlobSignRequest
type BlobSignRequest struct {
	// The blob to sign, encoded as a base64 string.  Note that certain signing keys impose additional requirements on the contents of the message. For example, Secp256k1 keys require that the message is 32 bytes long.
	MessageBase64 string `json:"message_base64"`
}

// NewBlobSignRequest instantiates a new BlobSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobSignRequest(messageBase64 string) *BlobSignRequest {
	this := BlobSignRequest{}
	this.MessageBase64 = messageBase64
	return &this
}

// NewBlobSignRequestWithDefaults instantiates a new BlobSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobSignRequestWithDefaults() *BlobSignRequest {
	this := BlobSignRequest{}
	return &this
}

// GetMessageBase64 returns the MessageBase64 field value
func (o *BlobSignRequest) GetMessageBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageBase64
}

// GetMessageBase64Ok returns a tuple with the MessageBase64 field value
// and a boolean to check if the value has been set.
func (o *BlobSignRequest) GetMessageBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageBase64, true
}

// SetMessageBase64 sets field value
func (o *BlobSignRequest) SetMessageBase64(v string) {
	o.MessageBase64 = v
}

func (o BlobSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message_base64"] = o.MessageBase64
	}
	return json.Marshal(toSerialize)
}

type NullableBlobSignRequest struct {
	value *BlobSignRequest
	isSet bool
}

func (v NullableBlobSignRequest) Get() *BlobSignRequest {
	return v.value
}

func (v *NullableBlobSignRequest) Set(val *BlobSignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobSignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobSignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobSignRequest(val *BlobSignRequest) *NullableBlobSignRequest {
	return &NullableBlobSignRequest{value: val, isSet: true}
}

func (v NullableBlobSignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobSignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
