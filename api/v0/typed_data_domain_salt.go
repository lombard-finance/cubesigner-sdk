package v0

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TypedDataDomainSalt - A disambiguating salt for the protocol. This can be used as a domain separator of last resort. Can be either a hex-encoded string or byte array
type TypedDataDomainSalt struct {
	ArrayOfFloat32 *[]float32
	String         *string
}

// []float32AsTypedDataDomainSalt is a convenience function that returns []float32 wrapped in TypedDataDomainSalt
func ArrayOfFloat32AsTypedDataDomainSalt(v *[]float32) TypedDataDomainSalt {
	return TypedDataDomainSalt{
		ArrayOfFloat32: v,
	}
}

// stringAsTypedDataDomainSalt is a convenience function that returns string wrapped in TypedDataDomainSalt
func StringAsTypedDataDomainSalt(v *string) TypedDataDomainSalt {
	return TypedDataDomainSalt{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TypedDataDomainSalt) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfFloat32
	err = newStrictDecoder(data).Decode(&dst.ArrayOfFloat32)
	if err == nil {
		jsonArrayOfFloat32, _ := json.Marshal(dst.ArrayOfFloat32)
		if string(jsonArrayOfFloat32) == "{}" { // empty struct
			dst.ArrayOfFloat32 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfFloat32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfFloat32 = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TypedDataDomainSalt)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TypedDataDomainSalt)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TypedDataDomainSalt) MarshalJSON() ([]byte, error) {
	if src.ArrayOfFloat32 != nil {
		return json.Marshal(&src.ArrayOfFloat32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TypedDataDomainSalt) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfFloat32 != nil {
		return obj.ArrayOfFloat32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableTypedDataDomainSalt struct {
	value *TypedDataDomainSalt
	isSet bool
}

func (v NullableTypedDataDomainSalt) Get() *TypedDataDomainSalt {
	return v.value
}

func (v *NullableTypedDataDomainSalt) Set(val *TypedDataDomainSalt) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedDataDomainSalt) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedDataDomainSalt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedDataDomainSalt(val *TypedDataDomainSalt) *NullableTypedDataDomainSalt {
	return &NullableTypedDataDomainSalt{value: val, isSet: true}
}

func (v NullableTypedDataDomainSalt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedDataDomainSalt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}
