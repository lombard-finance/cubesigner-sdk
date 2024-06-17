package api

import (
	"encoding/json"
	"fmt"
)

// AcceptedValueCode the model 'AcceptedValueCode'
type AcceptedValueCode string

// List of AcceptedValueCode
const (
	MFA_REQUIRED AcceptedValueCode = "MfaRequired"
)

// All allowed values of AcceptedValueCode enum
var AllowedAcceptedValueCodeEnumValues = []AcceptedValueCode{
	"MfaRequired",
}

func (v *AcceptedValueCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AcceptedValueCode(value)
	for _, existing := range AllowedAcceptedValueCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AcceptedValueCode", value)
}

// NewAcceptedValueCodeFromValue returns a pointer to a valid AcceptedValueCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcceptedValueCodeFromValue(v string) (*AcceptedValueCode, error) {
	ev := AcceptedValueCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcceptedValueCode: valid values are %v", v, AllowedAcceptedValueCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcceptedValueCode) IsValid() bool {
	for _, existing := range AllowedAcceptedValueCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcceptedValueCode value
func (v AcceptedValueCode) Ptr() *AcceptedValueCode {
	return &v
}

type NullableAcceptedValueCode struct {
	value *AcceptedValueCode
	isSet bool
}

func (v NullableAcceptedValueCode) Get() *AcceptedValueCode {
	return v.value
}

func (v *NullableAcceptedValueCode) Set(val *AcceptedValueCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedValueCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedValueCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedValueCode(val *AcceptedValueCode) *NullableAcceptedValueCode {
	return &NullableAcceptedValueCode{value: val, isSet: true}
}

func (v NullableAcceptedValueCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedValueCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
