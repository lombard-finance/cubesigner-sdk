package api

import (
	"encoding/json"
	"fmt"
)

// FeeType This type specifies the interpretation of the `fee` field in Babylon staking requests. If `sats`, the field is intpreted as a fixed value in sats. If `sats_pr_vb`, the field is interpreted as a rate and the fee is computed by estimating the weight of the output transaction.
type FeeType string

// List of FeeType
const (
	SATS        FeeType = "sats"
	SATS_PER_VB FeeType = "sats_per_vb"
)

// All allowed values of FeeType enum
var AllowedFeeTypeEnumValues = []FeeType{
	"sats",
	"sats_per_vb",
}

func (v *FeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeeType(value)
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeeType", value)
}

// NewFeeTypeFromValue returns a pointer to a valid FeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeeTypeFromValue(v string) (*FeeType, error) {
	ev := FeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeeType: valid values are %v", v, AllowedFeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeeType) IsValid() bool {
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeeType value
func (v FeeType) Ptr() *FeeType {
	return &v
}

type NullableFeeType struct {
	value *FeeType
	isSet bool
}

func (v NullableFeeType) Get() *FeeType {
	return v.value
}

func (v *NullableFeeType) Set(val *FeeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeType(val *FeeType) *NullableFeeType {
	return &NullableFeeType{value: val, isSet: true}
}

func (v NullableFeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
