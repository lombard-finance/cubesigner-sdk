package api

import (
	"encoding/json"
	"fmt"
)

// MfaTypeOneOf Log in with CubeSigner user credentials
type MfaTypeOneOf string

// List of MfaType_oneOf
const (
	CUBE_SIGNER MfaTypeOneOf = "CubeSigner"
)

// All allowed values of MfaTypeOneOf enum
var AllowedMfaTypeOneOfEnumValues = []MfaTypeOneOf{
	"CubeSigner",
}

func (v *MfaTypeOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MfaTypeOneOf(value)
	for _, existing := range AllowedMfaTypeOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MfaTypeOneOf", value)
}

// NewMfaTypeOneOfFromValue returns a pointer to a valid MfaTypeOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMfaTypeOneOfFromValue(v string) (*MfaTypeOneOf, error) {
	ev := MfaTypeOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MfaTypeOneOf: valid values are %v", v, AllowedMfaTypeOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MfaTypeOneOf) IsValid() bool {
	for _, existing := range AllowedMfaTypeOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MfaType_oneOf value
func (v MfaTypeOneOf) Ptr() *MfaTypeOneOf {
	return &v
}

//
//type NullableMfaTypeOneOf struct {
//	value *MfaTypeOneOf
//	isSet bool
//}
//
//func (v NullableMfaTypeOneOf) Get() *MfaTypeOneOf {
//	return v.value
//}
//
//func (v *NullableMfaTypeOneOf) Set(val *MfaTypeOneOf) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMfaTypeOneOf) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMfaTypeOneOf) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMfaTypeOneOf(val *MfaTypeOneOf) *NullableMfaTypeOneOf {
//	return &NullableMfaTypeOneOf{value: val, isSet: true}
//}
//
//func (v NullableMfaTypeOneOf) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMfaTypeOneOf) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
