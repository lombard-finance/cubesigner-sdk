package api

import (
	"encoding/json"
)

// BtcTxOut struct for BtcTxOut
type BtcTxOut struct {
	// The script which must be satisfied for the output to be spent.
	ScriptPubkey string `json:"script_pubkey"`
	// The value of the output, in satoshis.
	Value int64 `json:"value"`
}

// NewBtcTxOut instantiates a new BtcTxOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcTxOut(scriptPubkey string, value int64) *BtcTxOut {
	this := BtcTxOut{}
	this.ScriptPubkey = scriptPubkey
	this.Value = value
	return &this
}

// NewBtcTxOutWithDefaults instantiates a new BtcTxOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcTxOutWithDefaults() *BtcTxOut {
	this := BtcTxOut{}
	return &this
}

// GetScriptPubkey returns the ScriptPubkey field value
func (o *BtcTxOut) GetScriptPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptPubkey
}

// GetScriptPubkeyOk returns a tuple with the ScriptPubkey field value
// and a boolean to check if the value has been set.
func (o *BtcTxOut) GetScriptPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubkey, true
}

// SetScriptPubkey sets field value
func (o *BtcTxOut) SetScriptPubkey(v string) {
	o.ScriptPubkey = v
}

// GetValue returns the Value field value
func (o *BtcTxOut) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BtcTxOut) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BtcTxOut) SetValue(v int64) {
	o.Value = v
}

func (o BtcTxOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["script_pubkey"] = o.ScriptPubkey
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBtcTxOut struct {
	value *BtcTxOut
	isSet bool
}

func (v NullableBtcTxOut) Get() *BtcTxOut {
	return v.value
}

func (v *NullableBtcTxOut) Set(val *BtcTxOut) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcTxOut) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcTxOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcTxOut(val *BtcTxOut) *NullableBtcTxOut {
	return &NullableBtcTxOut{value: val, isSet: true}
}

func (v NullableBtcTxOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcTxOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
