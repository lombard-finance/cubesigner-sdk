package api

import (
	"encoding/json"
	"fmt"
)

// BtcSighashType the model 'BtcSighashType'
type BtcSighashType string

// List of BtcSighashType
const (
	ALL                        BtcSighashType = "All"
	NONE                       BtcSighashType = "None"
	SINGLE                     BtcSighashType = "Single"
	ALL_PLUS_ANYONE_CAN_PAY    BtcSighashType = "AllPlusAnyoneCanPay"
	NONE_PLUS_ANYONE_CAN_PAY   BtcSighashType = "NonePlusAnyoneCanPay"
	SINGLE_PLUS_ANYONE_CAN_PAY BtcSighashType = "SinglePlusAnyoneCanPay"
)

// All allowed values of BtcSighashType enum
var AllowedBtcSighashTypeEnumValues = []BtcSighashType{
	"All",
	"None",
	"Single",
	"AllPlusAnyoneCanPay",
	"NonePlusAnyoneCanPay",
	"SinglePlusAnyoneCanPay",
}

func (v *BtcSighashType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BtcSighashType(value)
	for _, existing := range AllowedBtcSighashTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BtcSighashType", value)
}

// NewBtcSighashTypeFromValue returns a pointer to a valid BtcSighashType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBtcSighashTypeFromValue(v string) (*BtcSighashType, error) {
	ev := BtcSighashType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BtcSighashType: valid values are %v", v, AllowedBtcSighashTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BtcSighashType) IsValid() bool {
	for _, existing := range AllowedBtcSighashTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BtcSighashType value
func (v BtcSighashType) Ptr() *BtcSighashType {
	return &v
}

type NullableBtcSighashType struct {
	value *BtcSighashType
	isSet bool
}

func (v NullableBtcSighashType) Get() *BtcSighashType {
	return v.value
}

func (v *NullableBtcSighashType) Set(val *BtcSighashType) {
	v.value = val
	v.isSet = true
}

func (v NullableBtcSighashType) IsSet() bool {
	return v.isSet
}

func (v *NullableBtcSighashType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtcSighashType(val *BtcSighashType) *NullableBtcSighashType {
	return &NullableBtcSighashType{value: val, isSet: true}
}

func (v NullableBtcSighashType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtcSighashType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
