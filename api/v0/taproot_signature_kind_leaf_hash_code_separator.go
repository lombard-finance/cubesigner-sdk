package v0

import (
	"encoding/json"
)

// TaprootSignatureKindLeafHashCodeSeparator struct for TaprootSignatureKindLeafHashCodeSeparator
type TaprootSignatureKindLeafHashCodeSeparator struct {
	// Code separator
	CodeSeparator int32 `json:"code_separator"`
	// Taproot-tagged hash with tag \"TapLeaf\".
	LeafHash string `json:"leaf_hash"`
}

// NewTaprootSignatureKindLeafHashCodeSeparator instantiates a new TaprootSignatureKindLeafHashCodeSeparator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaprootSignatureKindLeafHashCodeSeparator(codeSeparator int32, leafHash string) *TaprootSignatureKindLeafHashCodeSeparator {
	this := TaprootSignatureKindLeafHashCodeSeparator{}
	this.CodeSeparator = codeSeparator
	this.LeafHash = leafHash
	return &this
}

// NewTaprootSignatureKindLeafHashCodeSeparatorWithDefaults instantiates a new TaprootSignatureKindLeafHashCodeSeparator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaprootSignatureKindLeafHashCodeSeparatorWithDefaults() *TaprootSignatureKindLeafHashCodeSeparator {
	this := TaprootSignatureKindLeafHashCodeSeparator{}
	return &this
}

// GetCodeSeparator returns the CodeSeparator field value
func (o *TaprootSignatureKindLeafHashCodeSeparator) GetCodeSeparator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CodeSeparator
}

// GetCodeSeparatorOk returns a tuple with the CodeSeparator field value
// and a boolean to check if the value has been set.
func (o *TaprootSignatureKindLeafHashCodeSeparator) GetCodeSeparatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeSeparator, true
}

// SetCodeSeparator sets field value
func (o *TaprootSignatureKindLeafHashCodeSeparator) SetCodeSeparator(v int32) {
	o.CodeSeparator = v
}

// GetLeafHash returns the LeafHash field value
func (o *TaprootSignatureKindLeafHashCodeSeparator) GetLeafHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeafHash
}

// GetLeafHashOk returns a tuple with the LeafHash field value
// and a boolean to check if the value has been set.
func (o *TaprootSignatureKindLeafHashCodeSeparator) GetLeafHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeafHash, true
}

// SetLeafHash sets field value
func (o *TaprootSignatureKindLeafHashCodeSeparator) SetLeafHash(v string) {
	o.LeafHash = v
}

func (o TaprootSignatureKindLeafHashCodeSeparator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code_separator"] = o.CodeSeparator
	}
	if true {
		toSerialize["leaf_hash"] = o.LeafHash
	}
	return json.Marshal(toSerialize)
}

type NullableTaprootSignatureKindLeafHashCodeSeparator struct {
	value *TaprootSignatureKindLeafHashCodeSeparator
	isSet bool
}

func (v NullableTaprootSignatureKindLeafHashCodeSeparator) Get() *TaprootSignatureKindLeafHashCodeSeparator {
	return v.value
}

func (v *NullableTaprootSignatureKindLeafHashCodeSeparator) Set(val *TaprootSignatureKindLeafHashCodeSeparator) {
	v.value = val
	v.isSet = true
}

func (v NullableTaprootSignatureKindLeafHashCodeSeparator) IsSet() bool {
	return v.isSet
}

func (v *NullableTaprootSignatureKindLeafHashCodeSeparator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaprootSignatureKindLeafHashCodeSeparator(val *TaprootSignatureKindLeafHashCodeSeparator) *NullableTaprootSignatureKindLeafHashCodeSeparator {
	return &NullableTaprootSignatureKindLeafHashCodeSeparator{value: val, isSet: true}
}

func (v NullableTaprootSignatureKindLeafHashCodeSeparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaprootSignatureKindLeafHashCodeSeparator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
