package v0

import (
	"encoding/json"
)

// AcceptedValueOneOfMfaRequired struct for AcceptedValueOneOfMfaRequired
type AcceptedValueOneOfMfaRequired struct {
	// MFA request id
	Id string `json:"id"`
	// Organization id
	OrgId   string                     `json:"org_id"`
	Session NullableNewSessionResponse `json:"session,omitempty"`
}

// NewAcceptedValueOneOfMfaRequired instantiates a new AcceptedValueOneOfMfaRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptedValueOneOfMfaRequired(id string, orgId string) *AcceptedValueOneOfMfaRequired {
	this := AcceptedValueOneOfMfaRequired{}
	this.Id = id
	this.OrgId = orgId
	return &this
}

// NewAcceptedValueOneOfMfaRequiredWithDefaults instantiates a new AcceptedValueOneOfMfaRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedValueOneOfMfaRequiredWithDefaults() *AcceptedValueOneOfMfaRequired {
	this := AcceptedValueOneOfMfaRequired{}
	return &this
}

// GetId returns the Id field value
func (o *AcceptedValueOneOfMfaRequired) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AcceptedValueOneOfMfaRequired) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AcceptedValueOneOfMfaRequired) SetId(v string) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *AcceptedValueOneOfMfaRequired) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *AcceptedValueOneOfMfaRequired) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *AcceptedValueOneOfMfaRequired) SetOrgId(v string) {
	o.OrgId = v
}

// GetSession returns the Session field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceptedValueOneOfMfaRequired) GetSession() NewSessionResponse {
	if o == nil || o.Session.Get() == nil {
		var ret NewSessionResponse
		return ret
	}
	return *o.Session.Get()
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceptedValueOneOfMfaRequired) GetSessionOk() (*NewSessionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Session.Get(), o.Session.IsSet()
}

// HasSession returns a boolean if a field has been set.
func (o *AcceptedValueOneOfMfaRequired) HasSession() bool {
	if o != nil && o.Session.IsSet() {
		return true
	}

	return false
}

// SetSession gets a reference to the given NullableNewSessionResponse and assigns it to the Session field.
func (o *AcceptedValueOneOfMfaRequired) SetSession(v NewSessionResponse) {
	o.Session.Set(&v)
}

// SetSessionNil sets the value for Session to be an explicit nil
func (o *AcceptedValueOneOfMfaRequired) SetSessionNil() {
	o.Session.Set(nil)
}

// UnsetSession ensures that no value is present for Session, not even an explicit nil
func (o *AcceptedValueOneOfMfaRequired) UnsetSession() {
	o.Session.Unset()
}

func (o AcceptedValueOneOfMfaRequired) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Session.IsSet() {
		toSerialize["session"] = o.Session.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptedValueOneOfMfaRequired struct {
	value *AcceptedValueOneOfMfaRequired
	isSet bool
}

func (v NullableAcceptedValueOneOfMfaRequired) Get() *AcceptedValueOneOfMfaRequired {
	return v.value
}

func (v *NullableAcceptedValueOneOfMfaRequired) Set(val *AcceptedValueOneOfMfaRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedValueOneOfMfaRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedValueOneOfMfaRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedValueOneOfMfaRequired(val *AcceptedValueOneOfMfaRequired) *NullableAcceptedValueOneOfMfaRequired {
	return &NullableAcceptedValueOneOfMfaRequired{value: val, isSet: true}
}

func (v NullableAcceptedValueOneOfMfaRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedValueOneOfMfaRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
