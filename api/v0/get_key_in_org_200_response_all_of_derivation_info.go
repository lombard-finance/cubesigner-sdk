package v0

import (
	"encoding/json"
)

// GetKeyInOrg200ResponseAllOfDerivationInfo struct for GetKeyInOrg200ResponseAllOfDerivationInfo
type GetKeyInOrg200ResponseAllOfDerivationInfo struct {
	// The derivation path used to derive this key
	DerivationPath string `json:"derivation_path"`
	// The mnemonic-id of the key's parent mnemonic
	MnemonicId string `json:"mnemonic_id"`
}

// NewGetKeyInOrg200ResponseAllOfDerivationInfo instantiates a new GetKeyInOrg200ResponseAllOfDerivationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKeyInOrg200ResponseAllOfDerivationInfo(derivationPath string, mnemonicId string) *GetKeyInOrg200ResponseAllOfDerivationInfo {
	this := GetKeyInOrg200ResponseAllOfDerivationInfo{}
	this.DerivationPath = derivationPath
	this.MnemonicId = mnemonicId
	return &this
}

// NewGetKeyInOrg200ResponseAllOfDerivationInfoWithDefaults instantiates a new GetKeyInOrg200ResponseAllOfDerivationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKeyInOrg200ResponseAllOfDerivationInfoWithDefaults() *GetKeyInOrg200ResponseAllOfDerivationInfo {
	this := GetKeyInOrg200ResponseAllOfDerivationInfo{}
	return &this
}

// GetDerivationPath returns the DerivationPath field value
func (o *GetKeyInOrg200ResponseAllOfDerivationInfo) GetDerivationPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DerivationPath
}

// GetDerivationPathOk returns a tuple with the DerivationPath field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200ResponseAllOfDerivationInfo) GetDerivationPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DerivationPath, true
}

// SetDerivationPath sets field value
func (o *GetKeyInOrg200ResponseAllOfDerivationInfo) SetDerivationPath(v string) {
	o.DerivationPath = v
}

// GetMnemonicId returns the MnemonicId field value
func (o *GetKeyInOrg200ResponseAllOfDerivationInfo) GetMnemonicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MnemonicId
}

// GetMnemonicIdOk returns a tuple with the MnemonicId field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200ResponseAllOfDerivationInfo) GetMnemonicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MnemonicId, true
}

// SetMnemonicId sets field value
func (o *GetKeyInOrg200ResponseAllOfDerivationInfo) SetMnemonicId(v string) {
	o.MnemonicId = v
}

func (o GetKeyInOrg200ResponseAllOfDerivationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["derivation_path"] = o.DerivationPath
	}
	if true {
		toSerialize["mnemonic_id"] = o.MnemonicId
	}
	return json.Marshal(toSerialize)
}

type NullableGetKeyInOrg200ResponseAllOfDerivationInfo struct {
	value *GetKeyInOrg200ResponseAllOfDerivationInfo
	isSet bool
}

func (v NullableGetKeyInOrg200ResponseAllOfDerivationInfo) Get() *GetKeyInOrg200ResponseAllOfDerivationInfo {
	return v.value
}

func (v *NullableGetKeyInOrg200ResponseAllOfDerivationInfo) Set(val *GetKeyInOrg200ResponseAllOfDerivationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKeyInOrg200ResponseAllOfDerivationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKeyInOrg200ResponseAllOfDerivationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKeyInOrg200ResponseAllOfDerivationInfo(val *GetKeyInOrg200ResponseAllOfDerivationInfo) *NullableGetKeyInOrg200ResponseAllOfDerivationInfo {
	return &NullableGetKeyInOrg200ResponseAllOfDerivationInfo{value: val, isSet: true}
}

func (v NullableGetKeyInOrg200ResponseAllOfDerivationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKeyInOrg200ResponseAllOfDerivationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
