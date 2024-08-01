package v0

import (
	"encoding/json"
)

// PsbtSignRequest A request to sign a PSBT
type PsbtSignRequest struct {
	// A hex-serialized PSBT (version 0) containing the transaction inputs and all necessary information for signing (e.g., taproot path and leaf information).
	Psbt string `json:"psbt"`
	// When true, unconditionally sign every input to the PSBT controlled by a script spend. Otherwise (false, the default), this endpoint uses a heuristic to decide whether the script controlling a given UTXO requires a signature from this key.
	SignAllScripts *bool `json:"sign_all_scripts,omitempty"`
}

// NewPsbtSignRequest instantiates a new PsbtSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPsbtSignRequest(psbt string) *PsbtSignRequest {
	this := PsbtSignRequest{}
	this.Psbt = psbt
	return &this
}

// NewPsbtSignRequestWithDefaults instantiates a new PsbtSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPsbtSignRequestWithDefaults() *PsbtSignRequest {
	this := PsbtSignRequest{}
	return &this
}

// GetPsbt returns the Psbt field value
func (o *PsbtSignRequest) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *PsbtSignRequest) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *PsbtSignRequest) SetPsbt(v string) {
	o.Psbt = v
}

// GetSignAllScripts returns the SignAllScripts field value if set, zero value otherwise.
func (o *PsbtSignRequest) GetSignAllScripts() bool {
	if o == nil || o.SignAllScripts == nil {
		var ret bool
		return ret
	}
	return *o.SignAllScripts
}

// GetSignAllScriptsOk returns a tuple with the SignAllScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsbtSignRequest) GetSignAllScriptsOk() (*bool, bool) {
	if o == nil || o.SignAllScripts == nil {
		return nil, false
	}
	return o.SignAllScripts, true
}

// HasSignAllScripts returns a boolean if a field has been set.
func (o *PsbtSignRequest) HasSignAllScripts() bool {
	if o != nil && o.SignAllScripts != nil {
		return true
	}

	return false
}

// SetSignAllScripts gets a reference to the given bool and assigns it to the SignAllScripts field.
func (o *PsbtSignRequest) SetSignAllScripts(v bool) {
	o.SignAllScripts = &v
}

func (o PsbtSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["psbt"] = o.Psbt
	}
	if o.SignAllScripts != nil {
		toSerialize["sign_all_scripts"] = o.SignAllScripts
	}
	return json.Marshal(toSerialize)
}

type NullablePsbtSignRequest struct {
	value *PsbtSignRequest
	isSet bool
}

func (v NullablePsbtSignRequest) Get() *PsbtSignRequest {
	return v.value
}

func (v *NullablePsbtSignRequest) Set(val *PsbtSignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePsbtSignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePsbtSignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsbtSignRequest(val *PsbtSignRequest) *NullablePsbtSignRequest {
	return &NullablePsbtSignRequest{value: val, isSet: true}
}

func (v NullablePsbtSignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsbtSignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
