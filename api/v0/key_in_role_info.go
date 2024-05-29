package v0

import (
	"encoding/json"
)

// KeyInRoleInfo struct for KeyInRoleInfo
type KeyInRoleInfo struct {
	// Key ID
	KeyId string `json:"key_id"`
	// Policies that are checked before this key is used on behalf of this role
	Policy []map[string]interface{} `json:"policy,omitempty"`
	// Role ID
	RoleId string `json:"role_id"`
}

// NewKeyInRoleInfo instantiates a new KeyInRoleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyInRoleInfo(keyId string, roleId string) *KeyInRoleInfo {
	this := KeyInRoleInfo{}
	this.KeyId = keyId
	this.RoleId = roleId
	return &this
}

// NewKeyInRoleInfoWithDefaults instantiates a new KeyInRoleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyInRoleInfoWithDefaults() *KeyInRoleInfo {
	this := KeyInRoleInfo{}
	return &this
}

// GetKeyId returns the KeyId field value
func (o *KeyInRoleInfo) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *KeyInRoleInfo) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *KeyInRoleInfo) SetKeyId(v string) {
	o.KeyId = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *KeyInRoleInfo) GetPolicy() []map[string]interface{} {
	if o == nil || o.Policy == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInRoleInfo) GetPolicyOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *KeyInRoleInfo) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given []map[string]interface{} and assigns it to the Policy field.
func (o *KeyInRoleInfo) SetPolicy(v []map[string]interface{}) {
	o.Policy = v
}

// GetRoleId returns the RoleId field value
func (o *KeyInRoleInfo) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *KeyInRoleInfo) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *KeyInRoleInfo) SetRoleId(v string) {
	o.RoleId = v
}

func (o KeyInRoleInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["role_id"] = o.RoleId
	}
	return json.Marshal(toSerialize)
}
