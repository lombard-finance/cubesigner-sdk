package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// ListRoleKeys200Response Response type that wraps another type and adds base64url-encoded encrypted `last_evaluated_key` value (which can the user pass back to use as a url query parameter to continue pagination).
type ListRoleKeys200Response struct {
	// All keys in a role
	Keys []KeyInRoleInfo `json:"keys"`
	// If set, the content of `response` does not contain the entire result set. To fetch the next page of the result set, call the same endpoint but specify this value as the 'page.start' query parameter.
	LastEvaluatedKey api.NullableString `json:"last_evaluated_key,omitempty"`
}

// NewListRoleKeys200Response instantiates a new ListRoleKeys200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRoleKeys200Response(keys []KeyInRoleInfo) *ListRoleKeys200Response {
	this := ListRoleKeys200Response{}
	this.Keys = keys
	return &this
}

// NewListRoleKeys200ResponseWithDefaults instantiates a new ListRoleKeys200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRoleKeys200ResponseWithDefaults() *ListRoleKeys200Response {
	this := ListRoleKeys200Response{}
	return &this
}

// GetKeys returns the Keys field value
func (o *ListRoleKeys200Response) GetKeys() []KeyInRoleInfo {
	if o == nil {
		var ret []KeyInRoleInfo
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *ListRoleKeys200Response) GetKeysOk() ([]KeyInRoleInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *ListRoleKeys200Response) SetKeys(v []KeyInRoleInfo) {
	o.Keys = v
}

// GetLastEvaluatedKey returns the LastEvaluatedKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListRoleKeys200Response) GetLastEvaluatedKey() string {
	if o == nil || o.LastEvaluatedKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastEvaluatedKey.Get()
}

// GetLastEvaluatedKeyOk returns a tuple with the LastEvaluatedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListRoleKeys200Response) GetLastEvaluatedKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastEvaluatedKey.Get(), o.LastEvaluatedKey.IsSet()
}

// HasLastEvaluatedKey returns a boolean if a field has been set.
func (o *ListRoleKeys200Response) HasLastEvaluatedKey() bool {
	if o != nil && o.LastEvaluatedKey.IsSet() {
		return true
	}

	return false
}

// SetLastEvaluatedKey gets a reference to the given NullableString and assigns it to the LastEvaluatedKey field.
func (o *ListRoleKeys200Response) SetLastEvaluatedKey(v string) {
	o.LastEvaluatedKey.Set(&v)
}

// SetLastEvaluatedKeyNil sets the value for LastEvaluatedKey to be an explicit nil
func (o *ListRoleKeys200Response) SetLastEvaluatedKeyNil() {
	o.LastEvaluatedKey.Set(nil)
}

// UnsetLastEvaluatedKey ensures that no value is present for LastEvaluatedKey, not even an explicit nil
func (o *ListRoleKeys200Response) UnsetLastEvaluatedKey() {
	o.LastEvaluatedKey.Unset()
}

func (o ListRoleKeys200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["keys"] = o.Keys
	}
	if o.LastEvaluatedKey.IsSet() {
		toSerialize["last_evaluated_key"] = o.LastEvaluatedKey.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListRoleKeys200Response struct {
	value *ListRoleKeys200Response
	isSet bool
}

func (v NullableListRoleKeys200Response) Get() *ListRoleKeys200Response {
	return v.value
}

func (v *NullableListRoleKeys200Response) Set(val *ListRoleKeys200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListRoleKeys200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListRoleKeys200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRoleKeys200Response(val *ListRoleKeys200Response) *NullableListRoleKeys200Response {
	return &NullableListRoleKeys200Response{value: val, isSet: true}
}

func (v NullableListRoleKeys200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRoleKeys200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
