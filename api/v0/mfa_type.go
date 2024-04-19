package v0

import (
	"encoding/json"
	"fmt"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// MfaType - struct for MfaType
type MfaType struct {
	MfaTypeOneOf  *api.MfaTypeOneOf
	MfaTypeOneOf1 *api.MfaTypeOneOf1
	MfaTypeOneOf2 *api.MfaTypeOneOf2
	MfaTypeOneOf3 *MfaTypeOneOf3
}

// MfaTypeOneOfAsMfaType is a convenience function that returns MfaTypeOneOf wrapped in MfaType
func MfaTypeOneOfAsMfaType(v *api.MfaTypeOneOf) MfaType {
	return MfaType{
		MfaTypeOneOf: v,
	}
}

// MfaTypeOneOf1AsMfaType is a convenience function that returns MfaTypeOneOf1 wrapped in MfaType
func MfaTypeOneOf1AsMfaType(v *api.MfaTypeOneOf1) MfaType {
	return MfaType{
		MfaTypeOneOf1: v,
	}
}

// MfaTypeOneOf2AsMfaType is a convenience function that returns MfaTypeOneOf2 wrapped in MfaType
func MfaTypeOneOf2AsMfaType(v *api.MfaTypeOneOf2) MfaType {
	return MfaType{
		MfaTypeOneOf2: v,
	}
}

// MfaTypeOneOf3AsMfaType is a convenience function that returns MfaTypeOneOf3 wrapped in MfaType
func MfaTypeOneOf3AsMfaType(v *MfaTypeOneOf3) MfaType {
	return MfaType{
		MfaTypeOneOf3: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MfaType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MfaTypeOneOf
	err = api.NewStrictDecoder(data).Decode(&dst.MfaTypeOneOf)
	if err == nil {
		jsonMfaTypeOneOf, _ := json.Marshal(dst.MfaTypeOneOf)
		if string(jsonMfaTypeOneOf) == "{}" { // empty struct
			dst.MfaTypeOneOf = nil
		} else {
			match++
		}
	} else {
		dst.MfaTypeOneOf = nil
	}

	// try to unmarshal data into MfaTypeOneOf1
	err = api.NewStrictDecoder(data).Decode(&dst.MfaTypeOneOf1)
	if err == nil {
		jsonMfaTypeOneOf1, _ := json.Marshal(dst.MfaTypeOneOf1)
		if string(jsonMfaTypeOneOf1) == "{}" { // empty struct
			dst.MfaTypeOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.MfaTypeOneOf1 = nil
	}

	// try to unmarshal data into MfaTypeOneOf2
	err = api.NewStrictDecoder(data).Decode(&dst.MfaTypeOneOf2)
	if err == nil {
		jsonMfaTypeOneOf2, _ := json.Marshal(dst.MfaTypeOneOf2)
		if string(jsonMfaTypeOneOf2) == "{}" { // empty struct
			dst.MfaTypeOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.MfaTypeOneOf2 = nil
	}

	// try to unmarshal data into MfaTypeOneOf3
	err = api.NewStrictDecoder(data).Decode(&dst.MfaTypeOneOf3)
	if err == nil {
		jsonMfaTypeOneOf3, _ := json.Marshal(dst.MfaTypeOneOf3)
		if string(jsonMfaTypeOneOf3) == "{}" { // empty struct
			dst.MfaTypeOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.MfaTypeOneOf3 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MfaTypeOneOf = nil
		dst.MfaTypeOneOf1 = nil
		dst.MfaTypeOneOf2 = nil
		dst.MfaTypeOneOf3 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(MfaType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(MfaType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MfaType) MarshalJSON() ([]byte, error) {
	if src.MfaTypeOneOf != nil {
		return json.Marshal(&src.MfaTypeOneOf)
	}

	if src.MfaTypeOneOf1 != nil {
		return json.Marshal(&src.MfaTypeOneOf1)
	}

	if src.MfaTypeOneOf2 != nil {
		return json.Marshal(&src.MfaTypeOneOf2)
	}

	if src.MfaTypeOneOf3 != nil {
		return json.Marshal(&src.MfaTypeOneOf3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MfaType) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MfaTypeOneOf != nil {
		return obj.MfaTypeOneOf
	}

	if obj.MfaTypeOneOf1 != nil {
		return obj.MfaTypeOneOf1
	}

	if obj.MfaTypeOneOf2 != nil {
		return obj.MfaTypeOneOf2
	}

	if obj.MfaTypeOneOf3 != nil {
		return obj.MfaTypeOneOf3
	}

	// all schemas are nil
	return nil
}

//
//type NullableMfaType struct {
//	value *MfaType
//	isSet bool
//}
//
//func (v NullableMfaType) Get() *MfaType {
//	return v.value
//}
//
//func (v *NullableMfaType) Set(val *MfaType) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMfaType) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMfaType) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMfaType(val *MfaType) *NullableMfaType {
//	return &NullableMfaType{value: val, isSet: true}
//}
//
//func (v NullableMfaType) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMfaType) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
