package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// ErrorResponseAccepted struct for ErrorResponseAccepted
type ErrorResponseAccepted struct {
	MfaRequired api.AcceptedValueOneOfMfaRequired `json:"MfaRequired"`
}

// NewErrorResponseAccepted instantiates a new ErrorResponseAccepted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseAccepted(mfaRequired api.AcceptedValueOneOfMfaRequired) *ErrorResponseAccepted {
	this := ErrorResponseAccepted{}
	this.MfaRequired = mfaRequired
	return &this
}

// NewErrorResponseAcceptedWithDefaults instantiates a new ErrorResponseAccepted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseAcceptedWithDefaults() *ErrorResponseAccepted {
	this := ErrorResponseAccepted{}
	return &this
}

// GetMfaRequired returns the MfaRequired field value
func (o *ErrorResponseAccepted) GetMfaRequired() api.AcceptedValueOneOfMfaRequired {
	if o == nil {
		var ret api.AcceptedValueOneOfMfaRequired
		return ret
	}

	return o.MfaRequired
}

// GetMfaRequiredOk returns a tuple with the MfaRequired field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseAccepted) GetMfaRequiredOk() (*api.AcceptedValueOneOfMfaRequired, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaRequired, true
}

// SetMfaRequired sets field value
func (o *ErrorResponseAccepted) SetMfaRequired(v api.AcceptedValueOneOfMfaRequired) {
	o.MfaRequired = v
}

func (o ErrorResponseAccepted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MfaRequired"] = o.MfaRequired
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResponseAccepted struct {
	value *ErrorResponseAccepted
	isSet bool
}

func (v NullableErrorResponseAccepted) Get() *ErrorResponseAccepted {
	return v.value
}

func (v *NullableErrorResponseAccepted) Set(val *ErrorResponseAccepted) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseAccepted) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseAccepted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseAccepted(val *ErrorResponseAccepted) *NullableErrorResponseAccepted {
	return &NullableErrorResponseAccepted{value: val, isSet: true}
}

func (v NullableErrorResponseAccepted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseAccepted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
