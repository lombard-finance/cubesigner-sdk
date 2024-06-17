package api

import (
	"encoding/json"
	"fmt"
)

// PreconditionErrorCode - struct for PreconditionErrorCode
type PreconditionErrorCode struct {
	PolicyErrorCode           *PolicyErrorCode
	PreconditionErrorOwnCodes *PreconditionErrorOwnCodes
}

// PolicyErrorCodeAsPreconditionErrorCode is a convenience function that returns PolicyErrorCode wrapped in PreconditionErrorCode
func PolicyErrorCodeAsPreconditionErrorCode(v *PolicyErrorCode) PreconditionErrorCode {
	return PreconditionErrorCode{
		PolicyErrorCode: v,
	}
}

// PreconditionErrorOwnCodesAsPreconditionErrorCode is a convenience function that returns PreconditionErrorOwnCodes wrapped in PreconditionErrorCode
func PreconditionErrorOwnCodesAsPreconditionErrorCode(v *PreconditionErrorOwnCodes) PreconditionErrorCode {
	return PreconditionErrorCode{
		PreconditionErrorOwnCodes: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PreconditionErrorCode) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PolicyErrorCode
	err = newStrictDecoder(data).Decode(&dst.PolicyErrorCode)
	if err == nil {
		jsonPolicyErrorCode, _ := json.Marshal(dst.PolicyErrorCode)
		if string(jsonPolicyErrorCode) == "{}" { // empty struct
			dst.PolicyErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.PolicyErrorCode = nil
	}

	// try to unmarshal data into PreconditionErrorOwnCodes
	err = newStrictDecoder(data).Decode(&dst.PreconditionErrorOwnCodes)
	if err == nil {
		jsonPreconditionErrorOwnCodes, _ := json.Marshal(dst.PreconditionErrorOwnCodes)
		if string(jsonPreconditionErrorOwnCodes) == "{}" { // empty struct
			dst.PreconditionErrorOwnCodes = nil
		} else {
			match++
		}
	} else {
		dst.PreconditionErrorOwnCodes = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PolicyErrorCode = nil
		dst.PreconditionErrorOwnCodes = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PreconditionErrorCode)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PreconditionErrorCode)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PreconditionErrorCode) MarshalJSON() ([]byte, error) {
	if src.PolicyErrorCode != nil {
		return json.Marshal(&src.PolicyErrorCode)
	}

	if src.PreconditionErrorOwnCodes != nil {
		return json.Marshal(&src.PreconditionErrorOwnCodes)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PreconditionErrorCode) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PolicyErrorCode != nil {
		return obj.PolicyErrorCode
	}

	if obj.PreconditionErrorOwnCodes != nil {
		return obj.PreconditionErrorOwnCodes
	}

	// all schemas are nil
	return nil
}

type NullablePreconditionErrorCode struct {
	value *PreconditionErrorCode
	isSet bool
}

func (v NullablePreconditionErrorCode) Get() *PreconditionErrorCode {
	return v.value
}

func (v *NullablePreconditionErrorCode) Set(val *PreconditionErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullablePreconditionErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullablePreconditionErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreconditionErrorCode(val *PreconditionErrorCode) *NullablePreconditionErrorCode {
	return &NullablePreconditionErrorCode{value: val, isSet: true}
}

func (v NullablePreconditionErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreconditionErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
