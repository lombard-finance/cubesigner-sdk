package api

import (
	"encoding/json"
)

// BtcSignatureKindOneOfSegwit Segregated Witness
type BtcSignatureKindOneOfSegwit struct {
	// Transaction input index
	InputIndex int32 `json:"input_index"`
	// Script
	ScriptCode  string         `json:"script_code"`
	SighashType BtcSighashType `json:"sighash_type"`
	// Amount in satoshis
	Value int64 `json:"value"`
}

// NewBtcSignatureKindOneOfSegwit instantiates a new BtcSignatureKindOneOfSegwit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcSignatureKindOneOfSegwit(inputIndex int32, scriptCode string, sighashType BtcSighashType, value int64) *BtcSignatureKindOneOfSegwit {
	this := BtcSignatureKindOneOfSegwit{}
	this.InputIndex = inputIndex
	this.ScriptCode = scriptCode
	this.SighashType = sighashType
	this.Value = value
	return &this
}

// NewBtcSignatureKindOneOfSegwitWithDefaults instantiates a new BtcSignatureKindOneOfSegwit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcSignatureKindOneOfSegwitWithDefaults() *BtcSignatureKindOneOfSegwit {
	this := BtcSignatureKindOneOfSegwit{}
	return &this
}

// GetInputIndex returns the InputIndex field value
func (o *BtcSignatureKindOneOfSegwit) GetInputIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InputIndex
}

// GetInputIndexOk returns a tuple with the InputIndex field value
// and a boolean to check if the value has been set.
func (o *BtcSignatureKindOneOfSegwit) GetInputIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputIndex, true
}

// SetInputIndex sets field value
func (o *BtcSignatureKindOneOfSegwit) SetInputIndex(v int32) {
	o.InputIndex = v
}

// GetScriptCode returns the ScriptCode field value
func (o *BtcSignatureKindOneOfSegwit) GetScriptCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptCode
}

// GetScriptCodeOk returns a tuple with the ScriptCode field value
// and a boolean to check if the value has been set.
func (o *BtcSignatureKindOneOfSegwit) GetScriptCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptCode, true
}

// SetScriptCode sets field value
func (o *BtcSignatureKindOneOfSegwit) SetScriptCode(v string) {
	o.ScriptCode = v
}

// GetSighashType returns the SighashType field value
func (o *BtcSignatureKindOneOfSegwit) GetSighashType() BtcSighashType {
	if o == nil {
		var ret BtcSighashType
		return ret
	}

	return o.SighashType
}

// GetSighashTypeOk returns a tuple with the SighashType field value
// and a boolean to check if the value has been set.
func (o *BtcSignatureKindOneOfSegwit) GetSighashTypeOk() (*BtcSighashType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SighashType, true
}

// SetSighashType sets field value
func (o *BtcSignatureKindOneOfSegwit) SetSighashType(v BtcSighashType) {
	o.SighashType = v
}

// GetValue returns the Value field value
func (o *BtcSignatureKindOneOfSegwit) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BtcSignatureKindOneOfSegwit) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BtcSignatureKindOneOfSegwit) SetValue(v int64) {
	o.Value = v
}

func (o BtcSignatureKindOneOfSegwit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["input_index"] = o.InputIndex
	}
	if true {
		toSerialize["script_code"] = o.ScriptCode
	}
	if true {
		toSerialize["sighash_type"] = o.SighashType
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBtcSignatureKindOneOfSegwit struct {
	value *BtcSignatureKindOneOfSegwit
	isSet bool
}

func (v NullableBtcSignatureKindOneOfSegwit) Get() *BtcSignatureKindOneOfSegwit {
	return v.value
}

func (v *NullableBtcSignatureKindOneOfSegwit) Set(val *BtcSignatureKindOneOfSegwit) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcSignatureKindOneOfSegwit) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcSignatureKindOneOfSegwit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcSignatureKindOneOfSegwit(val *BtcSignatureKindOneOfSegwit) *NullableBtcSignatureKindOneOfSegwit {
	return &NullableBtcSignatureKindOneOfSegwit{value: val, isSet: true}
}

func (v NullableBtcSignatureKindOneOfSegwit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcSignatureKindOneOfSegwit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
