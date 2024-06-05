package api

import (
	"encoding/json"
)

// PrevOutputsOneOf struct for PrevOutputsOneOf
type PrevOutputsOneOf struct {
	One PrevOutputsOneOfOne `json:"One"`
}

// NewPrevOutputsOneOf instantiates a new PrevOutputsOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrevOutputsOneOf(one PrevOutputsOneOfOne) *PrevOutputsOneOf {
	this := PrevOutputsOneOf{}
	this.One = one
	return &this
}

// NewPrevOutputsOneOfWithDefaults instantiates a new PrevOutputsOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrevOutputsOneOfWithDefaults() *PrevOutputsOneOf {
	this := PrevOutputsOneOf{}
	return &this
}

// GetOne returns the One field value
func (o *PrevOutputsOneOf) GetOne() PrevOutputsOneOfOne {
	if o == nil {
		var ret PrevOutputsOneOfOne
		return ret
	}

	return o.One
}

// GetOneOk returns a tuple with the One field value
// and a boolean to check if the value has been set.
func (o *PrevOutputsOneOf) GetOneOk() (*PrevOutputsOneOfOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.One, true
}

// SetOne sets field value
func (o *PrevOutputsOneOf) SetOne(v PrevOutputsOneOfOne) {
	o.One = v
}

func (o PrevOutputsOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["One"] = o.One
	}
	return json.Marshal(toSerialize)
}

type NullablePrevOutputsOneOf struct {
	value *PrevOutputsOneOf
	isSet bool
}

func (v NullablePrevOutputsOneOf) Get() *PrevOutputsOneOf {
	return v.value
}

func (v *NullablePrevOutputsOneOf) Set(val *PrevOutputsOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePrevOutputsOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePrevOutputsOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrevOutputsOneOf(val *PrevOutputsOneOf) *NullablePrevOutputsOneOf {
	return &NullablePrevOutputsOneOf{value: val, isSet: true}
}

func (v NullablePrevOutputsOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrevOutputsOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
