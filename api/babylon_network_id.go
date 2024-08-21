package api

import (
	"encoding/json"
	"fmt"
)

// BabylonNetworkId The network-id for Babylon staking
type BabylonNetworkId string

// List of BabylonNetworkId
const (
	BBT4 BabylonNetworkId = "bbt4"
	BBN1 BabylonNetworkId = "bbn1"
)

// All allowed values of BabylonNetworkId enum
var AllowedBabylonNetworkIdEnumValues = []BabylonNetworkId{
	"bbt4",
	"bbn1",
}

func (v *BabylonNetworkId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BabylonNetworkId(value)
	for _, existing := range AllowedBabylonNetworkIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BabylonNetworkId", value)
}

// NewBabylonNetworkIdFromValue returns a pointer to a valid BabylonNetworkId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBabylonNetworkIdFromValue(v string) (*BabylonNetworkId, error) {
	ev := BabylonNetworkId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BabylonNetworkId: valid values are %v", v, AllowedBabylonNetworkIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BabylonNetworkId) IsValid() bool {
	for _, existing := range AllowedBabylonNetworkIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BabylonNetworkId value
func (v BabylonNetworkId) Ptr() *BabylonNetworkId {
	return &v
}

type NullableBabylonNetworkId struct {
	value *BabylonNetworkId
	isSet bool
}

func (v NullableBabylonNetworkId) Get() *BabylonNetworkId {
	return v.value
}

func (v *NullableBabylonNetworkId) Set(val *BabylonNetworkId) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonNetworkId) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonNetworkId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonNetworkId(val *BabylonNetworkId) *NullableBabylonNetworkId {
	return &NullableBabylonNetworkId{value: val, isSet: true}
}

func (v NullableBabylonNetworkId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonNetworkId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
