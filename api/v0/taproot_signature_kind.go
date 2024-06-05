package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// TaprootSignatureKind struct for TaprootSignatureKind
type TaprootSignatureKind struct {
	// Optional annex, as per BIP341
	Annex api.NullableString `json:"annex,omitempty"`
	// Transaction input index
	InputIndex            int32                                             `json:"input_index"`
	LeafHashCodeSeparator NullableTaprootSignatureKindLeafHashCodeSeparator `json:"leaf_hash_code_separator,omitempty"`
	// If this field is not present or null, no tweak is applied. If the field is an empty string, the key is tweaked with an unspendable script path per BIP0341. Otherwise, this field must contain a 32-byte, base-64 encoded hex string representing the Merkle root with which to tweak the key before signing.
	MerkleRoot api.NullableString `json:"merkle_root,omitempty"`
	Prevouts   api.PrevOutputs    `json:"prevouts"`
	// Hash type of an input's signature, encoded in the last byte of the signature. Possible values: - SIGHASH_ALL - SIGHASH_ALL|SIGHASH_ANYONECANPAY - SIGHASH_DEFAULT - SIGHASH_NONE - SIGHASH_NONE|SIGHASH_ANYONECANPAY - SIGHASH_SINGLE - SIGHASH_SINGLE|SIGHASH_ANYONECANPAY
	SighashType string `json:"sighash_type"`
}

// NewTaprootSignatureKind instantiates a new TaprootSignatureKind object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaprootSignatureKind(inputIndex int32, prevouts api.PrevOutputs, sighashType string) *TaprootSignatureKind {
	this := TaprootSignatureKind{}
	this.InputIndex = inputIndex
	this.Prevouts = prevouts
	this.SighashType = sighashType
	return &this
}

// NewTaprootSignatureKindWithDefaults instantiates a new TaprootSignatureKind object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaprootSignatureKindWithDefaults() *TaprootSignatureKind {
	this := TaprootSignatureKind{}
	return &this
}

// GetAnnex returns the Annex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaprootSignatureKind) GetAnnex() string {
	if o == nil || o.Annex.Get() == nil {
		var ret string
		return ret
	}
	return *o.Annex.Get()
}

// GetAnnexOk returns a tuple with the Annex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaprootSignatureKind) GetAnnexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annex.Get(), o.Annex.IsSet()
}

// HasAnnex returns a boolean if a field has been set.
func (o *TaprootSignatureKind) HasAnnex() bool {
	if o != nil && o.Annex.IsSet() {
		return true
	}

	return false
}

// SetAnnex gets a reference to the given NullableString and assigns it to the Annex field.
func (o *TaprootSignatureKind) SetAnnex(v string) {
	o.Annex.Set(&v)
}

// SetAnnexNil sets the value for Annex to be an explicit nil
func (o *TaprootSignatureKind) SetAnnexNil() {
	o.Annex.Set(nil)
}

// UnsetAnnex ensures that no value is present for Annex, not even an explicit nil
func (o *TaprootSignatureKind) UnsetAnnex() {
	o.Annex.Unset()
}

// GetInputIndex returns the InputIndex field value
func (o *TaprootSignatureKind) GetInputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InputIndex
}

// GetInputIndexOk returns a tuple with the InputIndex field value
// and a boolean to check if the value has been set.
func (o *TaprootSignatureKind) GetInputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputIndex, true
}

// SetInputIndex sets field value
func (o *TaprootSignatureKind) SetInputIndex(v int32) {
	o.InputIndex = v
}

// GetLeafHashCodeSeparator returns the LeafHashCodeSeparator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaprootSignatureKind) GetLeafHashCodeSeparator() TaprootSignatureKindLeafHashCodeSeparator {
	if o == nil || o.LeafHashCodeSeparator.Get() == nil {
		var ret TaprootSignatureKindLeafHashCodeSeparator
		return ret
	}
	return *o.LeafHashCodeSeparator.Get()
}

// GetLeafHashCodeSeparatorOk returns a tuple with the LeafHashCodeSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaprootSignatureKind) GetLeafHashCodeSeparatorOk() (*TaprootSignatureKindLeafHashCodeSeparator, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeafHashCodeSeparator.Get(), o.LeafHashCodeSeparator.IsSet()
}

// HasLeafHashCodeSeparator returns a boolean if a field has been set.
func (o *TaprootSignatureKind) HasLeafHashCodeSeparator() bool {
	if o != nil && o.LeafHashCodeSeparator.IsSet() {
		return true
	}

	return false
}

// SetLeafHashCodeSeparator gets a reference to the given NullableTaprootSignatureKindLeafHashCodeSeparator and assigns it to the LeafHashCodeSeparator field.
func (o *TaprootSignatureKind) SetLeafHashCodeSeparator(v TaprootSignatureKindLeafHashCodeSeparator) {
	o.LeafHashCodeSeparator.Set(&v)
}

// SetLeafHashCodeSeparatorNil sets the value for LeafHashCodeSeparator to be an explicit nil
func (o *TaprootSignatureKind) SetLeafHashCodeSeparatorNil() {
	o.LeafHashCodeSeparator.Set(nil)
}

// UnsetLeafHashCodeSeparator ensures that no value is present for LeafHashCodeSeparator, not even an explicit nil
func (o *TaprootSignatureKind) UnsetLeafHashCodeSeparator() {
	o.LeafHashCodeSeparator.Unset()
}

// GetMerkleRoot returns the MerkleRoot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaprootSignatureKind) GetMerkleRoot() string {
	if o == nil || o.MerkleRoot.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerkleRoot.Get()
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaprootSignatureKind) GetMerkleRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerkleRoot.Get(), o.MerkleRoot.IsSet()
}

// HasMerkleRoot returns a boolean if a field has been set.
func (o *TaprootSignatureKind) HasMerkleRoot() bool {
	if o != nil && o.MerkleRoot.IsSet() {
		return true
	}

	return false
}

// SetMerkleRoot gets a reference to the given NullableString and assigns it to the MerkleRoot field.
func (o *TaprootSignatureKind) SetMerkleRoot(v string) {
	o.MerkleRoot.Set(&v)
}

// SetMerkleRootNil sets the value for MerkleRoot to be an explicit nil
func (o *TaprootSignatureKind) SetMerkleRootNil() {
	o.MerkleRoot.Set(nil)
}

// UnsetMerkleRoot ensures that no value is present for MerkleRoot, not even an explicit nil
func (o *TaprootSignatureKind) UnsetMerkleRoot() {
	o.MerkleRoot.Unset()
}

// GetPrevouts returns the Prevouts field value
func (o *TaprootSignatureKind) GetPrevouts() api.PrevOutputs {
	if o == nil {
		var ret api.PrevOutputs
		return ret
	}

	return o.Prevouts
}

// GetPrevoutsOk returns a tuple with the Prevouts field value
// and a boolean to check if the value has been set.
func (o *TaprootSignatureKind) GetPrevoutsOk() (*api.PrevOutputs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prevouts, true
}

// SetPrevouts sets field value
func (o *TaprootSignatureKind) SetPrevouts(v api.PrevOutputs) {
	o.Prevouts = v
}

// GetSighashType returns the SighashType field value
func (o *TaprootSignatureKind) GetSighashType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SighashType
}

// GetSighashTypeOk returns a tuple with the SighashType field value
// and a boolean to check if the value has been set.
func (o *TaprootSignatureKind) GetSighashTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SighashType, true
}

// SetSighashType sets field value
func (o *TaprootSignatureKind) SetSighashType(v string) {
	o.SighashType = v
}

func (o TaprootSignatureKind) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annex.IsSet() {
		toSerialize["annex"] = o.Annex.Get()
	}
	if true {
		toSerialize["input_index"] = o.InputIndex
	}
	if o.LeafHashCodeSeparator.IsSet() {
		toSerialize["leaf_hash_code_separator"] = o.LeafHashCodeSeparator.Get()
	}
	if o.MerkleRoot.IsSet() {
		toSerialize["merkle_root"] = o.MerkleRoot.Get()
	}
	if true {
		toSerialize["prevouts"] = o.Prevouts
	}
	if true {
		toSerialize["sighash_type"] = o.SighashType
	}
	return json.Marshal(toSerialize)
}

type NullableTaprootSignatureKind struct {
	value *TaprootSignatureKind
	isSet bool
}

func (v NullableTaprootSignatureKind) Get() *TaprootSignatureKind {
	return v.value
}

func (v *NullableTaprootSignatureKind) Set(val *TaprootSignatureKind) {
	v.value = val
	v.isSet = true
}

func (v NullableTaprootSignatureKind) IsSet() bool {
	return v.isSet
}

func (v *NullableTaprootSignatureKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaprootSignatureKind(val *TaprootSignatureKind) *NullableTaprootSignatureKind {
	return &NullableTaprootSignatureKind{value: val, isSet: true}
}

func (v NullableTaprootSignatureKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaprootSignatureKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
