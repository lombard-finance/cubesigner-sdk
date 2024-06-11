package v0

import (
	"encoding/json"
)

// BtcLeafHashCodeSeparator Leaf hash and code, as per BIP341 and https://github.com/rust-bitcoin/rust-bitcoin/blob/464202109d2b2c96e9b4867461bffe420dbd8177/bitcoin/src/crypto/sighash.rs#L691
type BtcLeafHashCodeSeparator struct {
	// Code separator
	CodeSeparator int32 `json:"code_separator"`
	// Taproot-tagged hash with tag \"TapLeaf\".
	LeafHash string `json:"leaf_hash"`
}

// NewBtcLeafHashCodeSeparator instantiates a new BtcLeafHashCodeSeparator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcLeafHashCodeSeparator(codeSeparator int32, leafHash string) *BtcLeafHashCodeSeparator {
	this := BtcLeafHashCodeSeparator{}
	this.CodeSeparator = codeSeparator
	this.LeafHash = leafHash
	return &this
}

// NewBtcLeafHashCodeSeparatorWithDefaults instantiates a new BtcLeafHashCodeSeparator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcLeafHashCodeSeparatorWithDefaults() *BtcLeafHashCodeSeparator {
	this := BtcLeafHashCodeSeparator{}
	return &this
}

// GetCodeSeparator returns the CodeSeparator field value
func (o *BtcLeafHashCodeSeparator) GetCodeSeparator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CodeSeparator
}

// GetCodeSeparatorOk returns a tuple with the CodeSeparator field value
// and a boolean to check if the value has been set.
func (o *BtcLeafHashCodeSeparator) GetCodeSeparatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeSeparator, true
}

// SetCodeSeparator sets field value
func (o *BtcLeafHashCodeSeparator) SetCodeSeparator(v int32) {
	o.CodeSeparator = v
}

// GetLeafHash returns the LeafHash field value
func (o *BtcLeafHashCodeSeparator) GetLeafHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeafHash
}

// GetLeafHashOk returns a tuple with the LeafHash field value
// and a boolean to check if the value has been set.
func (o *BtcLeafHashCodeSeparator) GetLeafHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeafHash, true
}

// SetLeafHash sets field value
func (o *BtcLeafHashCodeSeparator) SetLeafHash(v string) {
	o.LeafHash = v
}

func (o BtcLeafHashCodeSeparator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code_separator"] = o.CodeSeparator
	}
	if true {
		toSerialize["leaf_hash"] = o.LeafHash
	}
	return json.Marshal(toSerialize)
}

type NullableBtcLeafHashCodeSeparator struct {
	value *BtcLeafHashCodeSeparator
	isSet bool
}

func (v NullableBtcLeafHashCodeSeparator) Get() *BtcLeafHashCodeSeparator {
	return v.value
}

func (v *NullableBtcLeafHashCodeSeparator) Set(val *BtcLeafHashCodeSeparator) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcLeafHashCodeSeparator) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcLeafHashCodeSeparator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcLeafHashCodeSeparator(val *BtcLeafHashCodeSeparator) *NullableBtcLeafHashCodeSeparator {
	return &NullableBtcLeafHashCodeSeparator{value: val, isSet: true}
}

func (v NullableBtcLeafHashCodeSeparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcLeafHashCodeSeparator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
