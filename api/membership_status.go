package api

import (
	"encoding/json"
	"fmt"
)

// MembershipStatus the model 'MembershipStatus'
type MembershipStatus string

// List of MembershipStatus
const (
	ENABLED  MembershipStatus = "enabled"
	DISABLED MembershipStatus = "disabled"
)

// All allowed values of MembershipStatus enum
var AllowedMembershipStatusEnumValues = []MembershipStatus{
	"enabled",
	"disabled",
}

func (v *MembershipStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MembershipStatus(value)
	for _, existing := range AllowedMembershipStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MembershipStatus", value)
}

// NewMembershipStatusFromValue returns a pointer to a valid MembershipStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMembershipStatusFromValue(v string) (*MembershipStatus, error) {
	ev := MembershipStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MembershipStatus: valid values are %v", v, AllowedMembershipStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MembershipStatus) IsValid() bool {
	for _, existing := range AllowedMembershipStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MembershipStatus value
func (v MembershipStatus) Ptr() *MembershipStatus {
	return &v
}

//
//type NullableMembershipStatus struct {
//	value *MembershipStatus
//	isSet bool
//}
//
//func (v NullableMembershipStatus) Get() *MembershipStatus {
//	return v.value
//}
//
//func (v *NullableMembershipStatus) Set(val *MembershipStatus) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMembershipStatus) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMembershipStatus) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMembershipStatus(val *MembershipStatus) *NullableMembershipStatus {
//	return &NullableMembershipStatus{value: val, isSet: true}
//}
//
//func (v NullableMembershipStatus) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMembershipStatus) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
