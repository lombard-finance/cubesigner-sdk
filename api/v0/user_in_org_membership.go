package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// UserInOrgMembership Information about a user's membership in an organization (without including any info about the user)
type UserInOrgMembership struct {
	Membership api.MemberRole `json:"membership"`
	// Organization id
	OrgId  string               `json:"org_id"`
	Status api.MembershipStatus `json:"status"`
}

// NewUserInOrgMembership instantiates a new UserInOrgMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInOrgMembership(membership api.MemberRole, orgId string, status api.MembershipStatus) *UserInOrgMembership {
	this := UserInOrgMembership{}
	this.Membership = membership
	this.OrgId = orgId
	this.Status = status
	return &this
}

// NewUserInOrgMembershipWithDefaults instantiates a new UserInOrgMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInOrgMembershipWithDefaults() *UserInOrgMembership {
	this := UserInOrgMembership{}
	return &this
}

// GetMembership returns the Membership field value
func (o *UserInOrgMembership) GetMembership() api.MemberRole {
	if o == nil {
		var ret api.MemberRole
		return ret
	}

	return o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value
// and a boolean to check if the value has been set.
func (o *UserInOrgMembership) GetMembershipOk() (*api.MemberRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Membership, true
}

// SetMembership sets field value
func (o *UserInOrgMembership) SetMembership(v api.MemberRole) {
	o.Membership = v
}

// GetOrgId returns the OrgId field value
func (o *UserInOrgMembership) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UserInOrgMembership) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *UserInOrgMembership) SetOrgId(v string) {
	o.OrgId = v
}

// GetStatus returns the Status field value
func (o *UserInOrgMembership) GetStatus() api.MembershipStatus {
	if o == nil {
		var ret api.MembershipStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserInOrgMembership) GetStatusOk() (*api.MembershipStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserInOrgMembership) SetStatus(v api.MembershipStatus) {
	o.Status = v
}

func (o UserInOrgMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["membership"] = o.Membership
	}
	if true {
		toSerialize["org_id"] = o.OrgId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

//
//type NullableUserInOrgMembership struct {
//	value *UserInOrgMembership
//	isSet bool
//}
//
//func (v NullableUserInOrgMembership) Get() *UserInOrgMembership {
//	return v.value
//}
//
//func (v *NullableUserInOrgMembership) Set(val *UserInOrgMembership) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableUserInOrgMembership) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableUserInOrgMembership) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableUserInOrgMembership(val *UserInOrgMembership) *NullableUserInOrgMembership {
//	return &NullableUserInOrgMembership{value: val, isSet: true}
//}
//
//func (v NullableUserInOrgMembership) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableUserInOrgMembership) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
