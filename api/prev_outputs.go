package api

import (
	"encoding/json"
	"fmt"
)

// PrevOutputs - Contains outputs of previous transactions.
type PrevOutputs struct {
	PrevOutputsOneOf  *PrevOutputsOneOf
	PrevOutputsOneOf1 *PrevOutputsOneOf1
}

// PrevOutputsOneOfAsPrevOutputs is a convenience function that returns PrevOutputsOneOf wrapped in PrevOutputs
func PrevOutputsOneOfAsPrevOutputs(v *PrevOutputsOneOf) PrevOutputs {
	return PrevOutputs{
		PrevOutputsOneOf: v,
	}
}

// PrevOutputsOneOf1AsPrevOutputs is a convenience function that returns PrevOutputsOneOf1 wrapped in PrevOutputs
func PrevOutputsOneOf1AsPrevOutputs(v *PrevOutputsOneOf1) PrevOutputs {
	return PrevOutputs{
		PrevOutputsOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PrevOutputs) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PrevOutputsOneOf
	err = newStrictDecoder(data).Decode(&dst.PrevOutputsOneOf)
	if err == nil {
		jsonPrevOutputsOneOf, _ := json.Marshal(dst.PrevOutputsOneOf)
		if string(jsonPrevOutputsOneOf) == "{}" { // empty struct
			dst.PrevOutputsOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PrevOutputsOneOf = nil
	}

	// try to unmarshal data into PrevOutputsOneOf1
	err = newStrictDecoder(data).Decode(&dst.PrevOutputsOneOf1)
	if err == nil {
		jsonPrevOutputsOneOf1, _ := json.Marshal(dst.PrevOutputsOneOf1)
		if string(jsonPrevOutputsOneOf1) == "{}" { // empty struct
			dst.PrevOutputsOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.PrevOutputsOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PrevOutputsOneOf = nil
		dst.PrevOutputsOneOf1 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PrevOutputs)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PrevOutputs)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PrevOutputs) MarshalJSON() ([]byte, error) {
	if src.PrevOutputsOneOf != nil {
		return json.Marshal(&src.PrevOutputsOneOf)
	}

	if src.PrevOutputsOneOf1 != nil {
		return json.Marshal(&src.PrevOutputsOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PrevOutputs) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PrevOutputsOneOf != nil {
		return obj.PrevOutputsOneOf
	}

	if obj.PrevOutputsOneOf1 != nil {
		return obj.PrevOutputsOneOf1
	}

	// all schemas are nil
	return nil
}

type NullablePrevOutputs struct {
	value *PrevOutputs
	isSet bool
}

func (v NullablePrevOutputs) Get() *PrevOutputs {
	return v.value
}

func (v *NullablePrevOutputs) Set(val *PrevOutputs) {
	v.value = val
	v.isSet = true
}

func (v NullablePrevOutputs) IsSet() bool {
	return v.isSet
}

func (v *NullablePrevOutputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrevOutputs(val *PrevOutputs) *NullablePrevOutputs {
	return &NullablePrevOutputs{value: val, isSet: true}
}

func (v NullablePrevOutputs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrevOutputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
