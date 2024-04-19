package v0

import (
	"encoding/json"
)

// KeyDerivationInfo Derivation-related metadata for keys derived from a long-lived mnemonic
type KeyDerivationInfo struct {
	// The derivation path used to derive this key
	DerivationPath string `json:"derivation_path"`
	// The mnemonic-id of the key's parent mnemonic
	MnemonicId string `json:"mnemonic_id"`
}

// NewKeyDerivationInfo instantiates a new KeyDerivationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyDerivationInfo(derivationPath string, mnemonicId string) *KeyDerivationInfo {
	this := KeyDerivationInfo{}
	this.DerivationPath = derivationPath
	this.MnemonicId = mnemonicId
	return &this
}

// NewKeyDerivationInfoWithDefaults instantiates a new KeyDerivationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyDerivationInfoWithDefaults() *KeyDerivationInfo {
	this := KeyDerivationInfo{}
	return &this
}

// GetDerivationPath returns the DerivationPath field value
func (o *KeyDerivationInfo) GetDerivationPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DerivationPath
}

// GetDerivationPathOk returns a tuple with the DerivationPath field value
// and a boolean to check if the value has been set.
func (o *KeyDerivationInfo) GetDerivationPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DerivationPath, true
}

// SetDerivationPath sets field value
func (o *KeyDerivationInfo) SetDerivationPath(v string) {
	o.DerivationPath = v
}

// GetMnemonicId returns the MnemonicId field value
func (o *KeyDerivationInfo) GetMnemonicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MnemonicId
}

// GetMnemonicIdOk returns a tuple with the MnemonicId field value
// and a boolean to check if the value has been set.
func (o *KeyDerivationInfo) GetMnemonicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MnemonicId, true
}

// SetMnemonicId sets field value
func (o *KeyDerivationInfo) SetMnemonicId(v string) {
	o.MnemonicId = v
}

func (o KeyDerivationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["derivation_path"] = o.DerivationPath
	}
	if true {
		toSerialize["mnemonic_id"] = o.MnemonicId
	}
	return json.Marshal(toSerialize)
}

type NullableKeyDerivationInfo struct {
	value *KeyDerivationInfo
	isSet bool
}

func (v NullableKeyDerivationInfo) Get() *KeyDerivationInfo {
	return v.value
}

func (v *NullableKeyDerivationInfo) Set(val *KeyDerivationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyDerivationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyDerivationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyDerivationInfo(val *KeyDerivationInfo) *NullableKeyDerivationInfo {
	return &NullableKeyDerivationInfo{value: val, isSet: true}
}

func (v NullableKeyDerivationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyDerivationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
