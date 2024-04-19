package api

import (
	"encoding/json"
	"fmt"
)

// MfaTypeOneOf2 Answer a FIDO challenge using any registered FIDO key
type MfaTypeOneOf2 string

// List of MfaType_oneOf_2
const (
	FIDO MfaTypeOneOf2 = "Fido"
)

// All allowed values of MfaTypeOneOf2 enum
var AllowedMfaTypeOneOf2EnumValues = []MfaTypeOneOf2{
	"Fido",
}

func (v *MfaTypeOneOf2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MfaTypeOneOf2(value)
	for _, existing := range AllowedMfaTypeOneOf2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MfaTypeOneOf2", value)
}

// NewMfaTypeOneOf2FromValue returns a pointer to a valid MfaTypeOneOf2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMfaTypeOneOf2FromValue(v string) (*MfaTypeOneOf2, error) {
	ev := MfaTypeOneOf2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MfaTypeOneOf2: valid values are %v", v, AllowedMfaTypeOneOf2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MfaTypeOneOf2) IsValid() bool {
	for _, existing := range AllowedMfaTypeOneOf2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MfaType_oneOf_2 value
func (v MfaTypeOneOf2) Ptr() *MfaTypeOneOf2 {
	return &v
}

//
//type NullableMfaTypeOneOf2 struct {
//	value *MfaTypeOneOf2
//	isSet bool
//}
//
//func (v NullableMfaTypeOneOf2) Get() *MfaTypeOneOf2 {
//	return v.value
//}
//
//func (v *NullableMfaTypeOneOf2) Set(val *MfaTypeOneOf2) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMfaTypeOneOf2) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMfaTypeOneOf2) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMfaTypeOneOf2(val *MfaTypeOneOf2) *NullableMfaTypeOneOf2 {
//	return &NullableMfaTypeOneOf2{value: val, isSet: true}
//}
//
//func (v NullableMfaTypeOneOf2) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMfaTypeOneOf2) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
