package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// AboutMeLegacy200Response struct for AboutMeLegacy200Response
type AboutMeLegacy200Response struct {
	// Optional email
	Email api.NullableString `json:"email,omitempty"`
	// All multi-factor authentication methods configured for this user
	Mfa []ConfiguredMfa `json:"mfa"`
	// MFA policy, applies before logging in and other sensitive operations
	MfaPolicy map[string]interface{} `json:"mfa_policy,omitempty"`
	// Optional name
	Name api.NullableString `json:"name,omitempty"`
	// All organizations the user belongs to. Deprecated in favor of 'orgs'.
	// Deprecated
	OrgIds []string `json:"org_ids"`
	// All organizations the user belongs to, including the membership role in each.
	Orgs []UserInOrgMembership `json:"orgs"`
	// The id of the currently logged in user
	UserId string `json:"user_id"`
}

// NewAboutMeLegacy200Response instantiates a new AboutMeLegacy200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAboutMeLegacy200Response(mfa []ConfiguredMfa, orgIds []string, orgs []UserInOrgMembership, userId string) *AboutMeLegacy200Response {
	this := AboutMeLegacy200Response{}
	this.Mfa = mfa
	this.OrgIds = orgIds
	this.Orgs = orgs
	this.UserId = userId
	return &this
}

// NewAboutMeLegacy200ResponseWithDefaults instantiates a new AboutMeLegacy200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAboutMeLegacy200ResponseWithDefaults() *AboutMeLegacy200Response {
	this := AboutMeLegacy200Response{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AboutMeLegacy200Response) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AboutMeLegacy200Response) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *AboutMeLegacy200Response) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *AboutMeLegacy200Response) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *AboutMeLegacy200Response) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *AboutMeLegacy200Response) UnsetEmail() {
	o.Email.Unset()
}

// GetMfa returns the Mfa field value
func (o *AboutMeLegacy200Response) GetMfa() []ConfiguredMfa {
	if o == nil {
		var ret []ConfiguredMfa
		return ret
	}

	return o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value
// and a boolean to check if the value has been set.
func (o *AboutMeLegacy200Response) GetMfaOk() ([]ConfiguredMfa, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mfa, true
}

// SetMfa sets field value
func (o *AboutMeLegacy200Response) SetMfa(v []ConfiguredMfa) {
	o.Mfa = v
}

// GetMfaPolicy returns the MfaPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AboutMeLegacy200Response) GetMfaPolicy() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MfaPolicy
}

// GetMfaPolicyOk returns a tuple with the MfaPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AboutMeLegacy200Response) GetMfaPolicyOk() (map[string]interface{}, bool) {
	if o == nil || o.MfaPolicy == nil {
		return nil, false
	}
	return o.MfaPolicy, true
}

// HasMfaPolicy returns a boolean if a field has been set.
func (o *AboutMeLegacy200Response) HasMfaPolicy() bool {
	if o != nil && o.MfaPolicy != nil {
		return true
	}

	return false
}

// SetMfaPolicy gets a reference to the given map[string]interface{} and assigns it to the MfaPolicy field.
func (o *AboutMeLegacy200Response) SetMfaPolicy(v map[string]interface{}) {
	o.MfaPolicy = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AboutMeLegacy200Response) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AboutMeLegacy200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AboutMeLegacy200Response) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AboutMeLegacy200Response) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AboutMeLegacy200Response) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AboutMeLegacy200Response) UnsetName() {
	o.Name.Unset()
}

// GetOrgIds returns the OrgIds field value
// Deprecated
func (o *AboutMeLegacy200Response) GetOrgIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrgIds
}

// GetOrgIdsOk returns a tuple with the OrgIds field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *AboutMeLegacy200Response) GetOrgIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgIds, true
}

// SetOrgIds sets field value
// Deprecated
func (o *AboutMeLegacy200Response) SetOrgIds(v []string) {
	o.OrgIds = v
}

// GetOrgs returns the Orgs field value
func (o *AboutMeLegacy200Response) GetOrgs() []UserInOrgMembership {
	if o == nil {
		var ret []UserInOrgMembership
		return ret
	}

	return o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value
// and a boolean to check if the value has been set.
func (o *AboutMeLegacy200Response) GetOrgsOk() ([]UserInOrgMembership, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orgs, true
}

// SetOrgs sets field value
func (o *AboutMeLegacy200Response) SetOrgs(v []UserInOrgMembership) {
	o.Orgs = v
}

// GetUserId returns the UserId field value
func (o *AboutMeLegacy200Response) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AboutMeLegacy200Response) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AboutMeLegacy200Response) SetUserId(v string) {
	o.UserId = v
}

func (o AboutMeLegacy200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if true {
		toSerialize["mfa"] = o.Mfa
	}
	if o.MfaPolicy != nil {
		toSerialize["mfa_policy"] = o.MfaPolicy
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["org_ids"] = o.OrgIds
	}
	if true {
		toSerialize["orgs"] = o.Orgs
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

//
//type NullableAboutMeLegacy200Response struct {
//	value *AboutMeLegacy200Response
//	isSet bool
//}
//
//func (v NullableAboutMeLegacy200Response) Get() *AboutMeLegacy200Response {
//	return v.value
//}
//
//func (v *NullableAboutMeLegacy200Response) Set(val *AboutMeLegacy200Response) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableAboutMeLegacy200Response) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableAboutMeLegacy200Response) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableAboutMeLegacy200Response(val *AboutMeLegacy200Response) *NullableAboutMeLegacy200Response {
//	return &NullableAboutMeLegacy200Response{value: val, isSet: true}
//}
//
//func (v NullableAboutMeLegacy200Response) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableAboutMeLegacy200Response) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
