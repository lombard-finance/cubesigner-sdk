package v0

import (
	"encoding/json"
)

// PsbtSign200Response Response to a PSBT signing request
type PsbtSign200Response struct {
	// The PSBT in standard hex serialization, without leading \"0x\".
	Psbt string `json:"psbt"`
}

// NewPsbtSign200Response instantiates a new PsbtSign200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPsbtSign200Response(psbt string) *PsbtSign200Response {
	this := PsbtSign200Response{}
	this.Psbt = psbt
	return &this
}

// NewPsbtSign200ResponseWithDefaults instantiates a new PsbtSign200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPsbtSign200ResponseWithDefaults() *PsbtSign200Response {
	this := PsbtSign200Response{}
	return &this
}

// GetPsbt returns the Psbt field value
func (o *PsbtSign200Response) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *PsbtSign200Response) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *PsbtSign200Response) SetPsbt(v string) {
	o.Psbt = v
}

func (o PsbtSign200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["psbt"] = o.Psbt
	}
	return json.Marshal(toSerialize)
}

type NullablePsbtSign200Response struct {
	value *PsbtSign200Response
	isSet bool
}

func (v NullablePsbtSign200Response) Get() *PsbtSign200Response {
	return v.value
}

func (v *NullablePsbtSign200Response) Set(val *PsbtSign200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePsbtSign200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePsbtSign200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsbtSign200Response(val *PsbtSign200Response) *NullablePsbtSign200Response {
	return &NullablePsbtSign200Response{value: val, isSet: true}
}

func (v NullablePsbtSign200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsbtSign200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
