package v1

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BlobSignRequest struct for BlobSignRequest
type BlobSignRequest struct {
	// The blob to sign, encoded as a base64 string.  Note that certain signing keys impose additional requirements on the contents of the message. For example, Secp256k1 keys require that the message is 32 bytes long.
	MessageBase64 string `json:"message_base64"`
	// An optional tweak value for use *only* with Taproot keys. This field is ignored for all other key types.  If this field is not present or null, no tweak is applied. If the field is an empty string, the key is tweaked with an unspendable script path per BIP0341. Otherwise, this field must contain a 32-byte, base-64 encoded hex string representing the Merkle root with which to tweak the key before signing.
	TaprootTweak api.NullableString `json:"taproot_tweak,omitempty"`
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

// GetTaprootTweak returns the TaprootTweak field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlobSignRequest) GetTaprootTweak() string {
	if o == nil || o.TaprootTweak.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaprootTweak.Get()
}

// GetTaprootTweakOk returns a tuple with the TaprootTweak field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlobSignRequest) GetTaprootTweakOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaprootTweak.Get(), o.TaprootTweak.IsSet()
}

// HasTaprootTweak returns a boolean if a field has been set.
func (o *BlobSignRequest) HasTaprootTweak() bool {
	if o != nil && o.TaprootTweak.IsSet() {
		return true
	}

	return false
}

// SetTaprootTweak gets a reference to the given NullableString and assigns it to the TaprootTweak field.
func (o *BlobSignRequest) SetTaprootTweak(v string) {
	o.TaprootTweak.Set(&v)
}

// SetTaprootTweakNil sets the value for TaprootTweak to be an explicit nil
func (o *BlobSignRequest) SetTaprootTweakNil() {
	o.TaprootTweak.Set(nil)
}

// UnsetTaprootTweak ensures that no value is present for TaprootTweak, not even an explicit nil
func (o *BlobSignRequest) UnsetTaprootTweak() {
	o.TaprootTweak.Unset()
}

func (o BlobSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message_base64"] = o.MessageBase64
	}
	if o.TaprootTweak.IsSet() {
		toSerialize["taproot_tweak"] = o.TaprootTweak.Get()
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
