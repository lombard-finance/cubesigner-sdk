package api

import (
	"encoding/json"
)

// PrevOutputsOneOf1 struct for PrevOutputsOneOf1
type PrevOutputsOneOf1 struct {
	// When `SIGHASH_ANYONECANPAY` is not provided, or when the caller is giving all previous outputs so the same variable can be used for multiple inputs.
	All []BtcTxOut `json:"All"`
}

// NewPrevOutputsOneOf1 instantiates a new PrevOutputsOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrevOutputsOneOf1(all []BtcTxOut) *PrevOutputsOneOf1 {
	this := PrevOutputsOneOf1{}
	this.All = all
	return &this
}

// NewPrevOutputsOneOf1WithDefaults instantiates a new PrevOutputsOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrevOutputsOneOf1WithDefaults() *PrevOutputsOneOf1 {
	this := PrevOutputsOneOf1{}
	return &this
}

// GetAll returns the All field value
func (o *PrevOutputsOneOf1) GetAll() []BtcTxOut {
	if o == nil {
		var ret []BtcTxOut
		return ret
	}

	return o.All
}

// GetAllOk returns a tuple with the All field value
// and a boolean to check if the value has been set.
func (o *PrevOutputsOneOf1) GetAllOk() ([]BtcTxOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.All, true
}

// SetAll sets field value
func (o *PrevOutputsOneOf1) SetAll(v []BtcTxOut) {
	o.All = v
}

func (o PrevOutputsOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["All"] = o.All
	}
	return json.Marshal(toSerialize)
}

type NullablePrevOutputsOneOf1 struct {
	value *PrevOutputsOneOf1
	isSet bool
}

func (v NullablePrevOutputsOneOf1) Get() *PrevOutputsOneOf1 {
	return v.value
}

func (v *NullablePrevOutputsOneOf1) Set(val *PrevOutputsOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePrevOutputsOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePrevOutputsOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrevOutputsOneOf1(val *PrevOutputsOneOf1) *NullablePrevOutputsOneOf1 {
	return &NullablePrevOutputsOneOf1{value: val, isSet: true}
}

func (v NullablePrevOutputsOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrevOutputsOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
