package api

import (
	"encoding/json"
	"fmt"
)

// MemberRole Describes whether a user in an org is an Owner or just a regular member
type MemberRole string

// List of MemberRole
const (
	ALIEN  MemberRole = "Alien"
	MEMBER MemberRole = "Member"
	OWNER  MemberRole = "Owner"
)

// All allowed values of MemberRole enum
var AllowedMemberRoleEnumValues = []MemberRole{
	"Alien",
	"Member",
	"Owner",
}

func (v *MemberRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MemberRole(value)
	for _, existing := range AllowedMemberRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MemberRole", value)
}

// NewMemberRoleFromValue returns a pointer to a valid MemberRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMemberRoleFromValue(v string) (*MemberRole, error) {
	ev := MemberRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MemberRole: valid values are %v", v, AllowedMemberRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MemberRole) IsValid() bool {
	for _, existing := range AllowedMemberRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MemberRole value
func (v MemberRole) Ptr() *MemberRole {
	return &v
}

//
//type NullableMemberRole struct {
//	value *MemberRole
//	isSet bool
//}
//
//func (v NullableMemberRole) Get() *MemberRole {
//	return v.value
//}
//
//func (v *NullableMemberRole) Set(val *MemberRole) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableMemberRole) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableMemberRole) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableMemberRole(val *MemberRole) *NullableMemberRole {
//	return &NullableMemberRole{value: val, isSet: true}
//}
//
//func (v NullableMemberRole) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableMemberRole) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
