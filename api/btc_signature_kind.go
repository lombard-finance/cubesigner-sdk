package api

import (
	"encoding/json"
	"fmt"
)

// BtcSignatureKind - struct for BtcSignatureKind
type BtcSignatureKind struct {
	BtcSignatureKindOneOf *BtcSignatureKindOneOf
}

// BtcSignatureKindOneOfAsBtcSignatureKind is a convenience function that returns BtcSignatureKindOneOf wrapped in BtcSignatureKind
func BtcSignatureKindOneOfAsBtcSignatureKind(v *BtcSignatureKindOneOf) BtcSignatureKind {
	return BtcSignatureKind{
		BtcSignatureKindOneOf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BtcSignatureKind) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BtcSignatureKindOneOf
	err = newStrictDecoder(data).Decode(&dst.BtcSignatureKindOneOf)
	if err == nil {
		jsonBtcSignatureKindOneOf, _ := json.Marshal(dst.BtcSignatureKindOneOf)
		if string(jsonBtcSignatureKindOneOf) == "{}" { // empty struct
			dst.BtcSignatureKindOneOf = nil
		} else {
			match++
		}
	} else {
		dst.BtcSignatureKindOneOf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BtcSignatureKindOneOf = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(BtcSignatureKind)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(BtcSignatureKind)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BtcSignatureKind) MarshalJSON() ([]byte, error) {
	if src.BtcSignatureKindOneOf != nil {
		return json.Marshal(&src.BtcSignatureKindOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BtcSignatureKind) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BtcSignatureKindOneOf != nil {
		return obj.BtcSignatureKindOneOf
	}

	// all schemas are nil
	return nil
}

type NullableBtcSignatureKind struct {
	value *BtcSignatureKind
	isSet bool
}

func (v NullableBtcSignatureKind) Get() *BtcSignatureKind {
	return v.value
}

func (v *NullableBtcSignatureKind) Set(val *BtcSignatureKind) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcSignatureKind) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcSignatureKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcSignatureKind(val *BtcSignatureKind) *NullableBtcSignatureKind {
	return &NullableBtcSignatureKind{value: val, isSet: true}
}

func (v NullableBtcSignatureKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcSignatureKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
