package v0

import "encoding/json"

// AddKeysToRoleRequest struct for AddKeysToRoleRequest
type AddKeysToRoleRequest struct {
	// A list of keys to add to a role
	KeyIds []string `json:"key_ids"`
	// Optional policies to apply for each key
	Policy []map[string]interface{} `json:"policy,omitempty"`
}

// NewAddKeysToRoleRequest instantiates a new AddKeysToRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddKeysToRoleRequest(keyIds []string) *AddKeysToRoleRequest {
	this := AddKeysToRoleRequest{}
	this.KeyIds = keyIds
	return &this
}

// NewAddKeysToRoleRequestWithDefaults instantiates a new AddKeysToRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddKeysToRoleRequestWithDefaults() *AddKeysToRoleRequest {
	this := AddKeysToRoleRequest{}
	return &this
}

// GetKeyIds returns the KeyIds field value
func (o *AddKeysToRoleRequest) GetKeyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KeyIds
}

// GetKeyIdsOk returns a tuple with the KeyIds field value
// and a boolean to check if the value has been set.
func (o *AddKeysToRoleRequest) GetKeyIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyIds, true
}

// SetKeyIds sets field value
func (o *AddKeysToRoleRequest) SetKeyIds(v []string) {
	o.KeyIds = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddKeysToRoleRequest) GetPolicy() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddKeysToRoleRequest) GetPolicyOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *AddKeysToRoleRequest) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given []map[string]interface{} and assigns it to the Policy field.
func (o *AddKeysToRoleRequest) SetPolicy(v []map[string]interface{}) {
	o.Policy = v
}

func (o AddKeysToRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_ids"] = o.KeyIds
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableAddKeysToRoleRequest struct {
	value *AddKeysToRoleRequest
	isSet bool
}

func (v NullableAddKeysToRoleRequest) Get() *AddKeysToRoleRequest {
	return v.value
}

func (v *NullableAddKeysToRoleRequest) Set(val *AddKeysToRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddKeysToRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddKeysToRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddKeysToRoleRequest(val *AddKeysToRoleRequest) *NullableAddKeysToRoleRequest {
	return &NullableAddKeysToRoleRequest{value: val, isSet: true}
}

func (v NullableAddKeysToRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddKeysToRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
