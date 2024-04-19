package api

import (
	"encoding/json"
	"fmt"
)

// MfaTypeOneOf1 Provide TOTP code
type MfaTypeOneOf1 string

// List of MfaType_oneOf_1
const (
	TOTP MfaTypeOneOf1 = "Totp"
)

// All allowed values of MfaTypeOneOf1 enum
var AllowedMfaTypeOneOf1EnumValues = []MfaTypeOneOf1{
	"Totp",
}

func (v *MfaTypeOneOf1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MfaTypeOneOf1(value)
	for _, existing := range AllowedMfaTypeOneOf1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MfaTypeOneOf1", value)
}

// NewMfaTypeOneOf1FromValue returns a pointer to a valid MfaTypeOneOf1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMfaTypeOneOf1FromValue(v string) (*MfaTypeOneOf1, error) {
	ev := MfaTypeOneOf1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MfaTypeOneOf1: valid values are %v", v, AllowedMfaTypeOneOf1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MfaTypeOneOf1) IsValid() bool {
	for _, existing := range AllowedMfaTypeOneOf1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MfaType_oneOf_1 value
func (v MfaTypeOneOf1) Ptr() *MfaTypeOneOf1 {
	return &v
}

//
//type NullableMfaTypeOneOf1 struct {
//	value *MfaTypeOneOf1
//	isSet bool
//}
//
//func (v NullableMfaTypeOneOf1) Get() *MfaTypeOneOf1 {
//	return v.value
//}
//
//func (v *NullableMfaTypeOneOf1) Set(val *MfaTypeOneOf1) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMfaTypeOneOf1) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMfaTypeOneOf1) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMfaTypeOneOf1(val *MfaTypeOneOf1) *NullableMfaTypeOneOf1 {
//	return &NullableMfaTypeOneOf1{value: val, isSet: true}
//}
//
//func (v NullableMfaTypeOneOf1) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMfaTypeOneOf1) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
