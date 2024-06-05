package api

import (
	"encoding/json"
)

// PrevOutputsOneOfOne `One` variant allows provision of the single previous output needed. It's useful, for example, when modifier `SIGHASH_ANYONECANPAY` is provided, only previous output of the current input is needed. The first `index` argument is the input index this output is referring to.
type PrevOutputsOneOfOne struct {
	Index int32    `json:"index"`
	TxOut BtcTxOut `json:"tx_out"`
}

// NewPrevOutputsOneOfOne instantiates a new PrevOutputsOneOfOne object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrevOutputsOneOfOne(index int32, txOut BtcTxOut) *PrevOutputsOneOfOne {
	this := PrevOutputsOneOfOne{}
	this.Index = index
	this.TxOut = txOut
	return &this
}

// NewPrevOutputsOneOfOneWithDefaults instantiates a new PrevOutputsOneOfOne object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrevOutputsOneOfOneWithDefaults() *PrevOutputsOneOfOne {
	this := PrevOutputsOneOfOne{}
	return &this
}

// GetIndex returns the Index field value
func (o *PrevOutputsOneOfOne) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *PrevOutputsOneOfOne) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *PrevOutputsOneOfOne) SetIndex(v int32) {
	o.Index = v
}

// GetTxOut returns the TxOut field value
func (o *PrevOutputsOneOfOne) GetTxOut() BtcTxOut {
	if o == nil {
		var ret BtcTxOut
		return ret
	}

	return o.TxOut
}

// GetTxOutOk returns a tuple with the TxOut field value
// and a boolean to check if the value has been set.
func (o *PrevOutputsOneOfOne) GetTxOutOk() (*BtcTxOut, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxOut, true
}

// SetTxOut sets field value
func (o *PrevOutputsOneOfOne) SetTxOut(v BtcTxOut) {
	o.TxOut = v
}

func (o PrevOutputsOneOfOne) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["tx_out"] = o.TxOut
	}
	return json.Marshal(toSerialize)
}

type NullablePrevOutputsOneOfOne struct {
	value *PrevOutputsOneOfOne
	isSet bool
}

func (v NullablePrevOutputsOneOfOne) Get() *PrevOutputsOneOfOne {
	return v.value
}

func (v *NullablePrevOutputsOneOfOne) Set(val *PrevOutputsOneOfOne) {
	v.value = val
	v.isSet = true
}

func (v NullablePrevOutputsOneOfOne) IsSet() bool {
	return v.isSet
}

func (v *NullablePrevOutputsOneOfOne) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrevOutputsOneOfOne(val *PrevOutputsOneOfOne) *NullablePrevOutputsOneOfOne {
	return &NullablePrevOutputsOneOfOne{value: val, isSet: true}
}

func (v NullablePrevOutputsOneOfOne) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrevOutputsOneOfOne) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
