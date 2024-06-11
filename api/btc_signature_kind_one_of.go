package api

import (
	"encoding/json"
)

// BtcSignatureKindOneOf struct for BtcSignatureKindOneOf
type BtcSignatureKindOneOf struct {
	Segwit BtcSignatureKindOneOfSegwit `json:"Segwit"`
}

// NewBtcSignatureKindOneOf instantiates a new BtcSignatureKindOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcSignatureKindOneOf(segwit BtcSignatureKindOneOfSegwit) *BtcSignatureKindOneOf {
	this := BtcSignatureKindOneOf{}
	this.Segwit = segwit
	return &this
}

// NewBtcSignatureKindOneOfWithDefaults instantiates a new BtcSignatureKindOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcSignatureKindOneOfWithDefaults() *BtcSignatureKindOneOf {
	this := BtcSignatureKindOneOf{}
	return &this
}

// GetSegwit returns the Segwit field value
func (o *BtcSignatureKindOneOf) GetSegwit() BtcSignatureKindOneOfSegwit {
	if o == nil {
		var ret BtcSignatureKindOneOfSegwit
		return ret
	}

	return o.Segwit
}

// GetSegwitOk returns a tuple with the Segwit field value
// and a boolean to check if the value has been set.
func (o *BtcSignatureKindOneOf) GetSegwitOk() (*BtcSignatureKindOneOfSegwit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Segwit, true
}

// SetSegwit sets field value
func (o *BtcSignatureKindOneOf) SetSegwit(v BtcSignatureKindOneOfSegwit) {
	o.Segwit = v
}

func (o BtcSignatureKindOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Segwit"] = o.Segwit
	}
	return json.Marshal(toSerialize)
}

type NullableBtcSignatureKindOneOf struct {
	value *BtcSignatureKindOneOf
	isSet bool
}

func (v NullableBtcSignatureKindOneOf) Get() *BtcSignatureKindOneOf {
	return v.value
}

func (v *NullableBtcSignatureKindOneOf) Set(val *BtcSignatureKindOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcSignatureKindOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcSignatureKindOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcSignatureKindOneOf(val *BtcSignatureKindOneOf) *NullableBtcSignatureKindOneOf {
	return &NullableBtcSignatureKindOneOf{value: val, isSet: true}
}

func (v NullableBtcSignatureKindOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcSignatureKindOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
