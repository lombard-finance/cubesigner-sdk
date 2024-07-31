package v0

import (
	"encoding/json"
)

// PsbtSignResponse Response to a PSBT signing request
type PsbtSignResponse struct {
	// The PSBT in standard hex serialization, without leading \"0x\".
	Psbt string `json:"psbt"`
}

// NewPsbtSignResponse instantiates a new PsbtSignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPsbtSignResponse(psbt string) *PsbtSignResponse {
	this := PsbtSignResponse{}
	this.Psbt = psbt
	return &this
}

// NewPsbtSignResponseWithDefaults instantiates a new PsbtSignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPsbtSignResponseWithDefaults() *PsbtSignResponse {
	this := PsbtSignResponse{}
	return &this
}

// GetPsbt returns the Psbt field value
func (o *PsbtSignResponse) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *PsbtSignResponse) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *PsbtSignResponse) SetPsbt(v string) {
	o.Psbt = v
}

func (o PsbtSignResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["psbt"] = o.Psbt
	}
	return json.Marshal(toSerialize)
}

type NullablePsbtSignResponse struct {
	value *PsbtSignResponse
	isSet bool
}

func (v NullablePsbtSignResponse) Get() *PsbtSignResponse {
	return v.value
}

func (v *NullablePsbtSignResponse) Set(val *PsbtSignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePsbtSignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePsbtSignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsbtSignResponse(val *PsbtSignResponse) *NullablePsbtSignResponse {
	return &NullablePsbtSignResponse{value: val, isSet: true}
}

func (v NullablePsbtSignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsbtSignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
