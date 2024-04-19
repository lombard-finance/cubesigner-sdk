package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// CreateKeyRequest struct for CreateKeyRequest
type CreateKeyRequest struct {
	EditPolicy NullableCreateAndUpdateKeyPropertiesEditPolicy `json:"edit_policy,omitempty"`
	// Set this key's metadata. If this value is `null`, the metadata is erased. If the field is missing, the metadata remains unchanged.
	Metadata interface{} `json:"metadata,omitempty"`
	// Specify a user other than themselves to be the (potentially new) owner of the key. The specified owner must be an existing user who is a member of the same org.
	Owner api.NullableString `json:"owner,omitempty"`
	// Set this key's policies. For an existing key, this overwrites all its policies.
	Policy []map[string]interface{} `json:"policy,omitempty"`
	// Chain id for which the key is allowed to sign messages
	ChainId api.NullableInt64 `json:"chain_id,omitempty"`
	// Number of keys to create
	Count   int32       `json:"count"`
	KeyType api.KeyType `json:"key_type"`
}

// NewCreateKeyRequest instantiates a new CreateKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyRequest(count int32, keyType api.KeyType) *CreateKeyRequest {
	this := CreateKeyRequest{}
	this.Count = count
	this.KeyType = keyType
	return &this
}

// NewCreateKeyRequestWithDefaults instantiates a new CreateKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyRequestWithDefaults() *CreateKeyRequest {
	this := CreateKeyRequest{}
	return &this
}

// GetEditPolicy returns the EditPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKeyRequest) GetEditPolicy() CreateAndUpdateKeyPropertiesEditPolicy {
	if o == nil || o.EditPolicy.Get() == nil {
		var ret CreateAndUpdateKeyPropertiesEditPolicy
		return ret
	}
	return *o.EditPolicy.Get()
}

// GetEditPolicyOk returns a tuple with the EditPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeyRequest) GetEditPolicyOk() (*CreateAndUpdateKeyPropertiesEditPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditPolicy.Get(), o.EditPolicy.IsSet()
}

// HasEditPolicy returns a boolean if a field has been set.
func (o *CreateKeyRequest) HasEditPolicy() bool {
	if o != nil && o.EditPolicy.IsSet() {
		return true
	}

	return false
}

// SetEditPolicy gets a reference to the given NullableCreateAndUpdateKeyPropertiesEditPolicy and assigns it to the EditPolicy field.
func (o *CreateKeyRequest) SetEditPolicy(v CreateAndUpdateKeyPropertiesEditPolicy) {
	o.EditPolicy.Set(&v)
}

// SetEditPolicyNil sets the value for EditPolicy to be an explicit nil
func (o *CreateKeyRequest) SetEditPolicyNil() {
	o.EditPolicy.Set(nil)
}

// UnsetEditPolicy ensures that no value is present for EditPolicy, not even an explicit nil
func (o *CreateKeyRequest) UnsetEditPolicy() {
	o.EditPolicy.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKeyRequest) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeyRequest) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateKeyRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *CreateKeyRequest) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKeyRequest) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeyRequest) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateKeyRequest) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *CreateKeyRequest) SetOwner(v string) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *CreateKeyRequest) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *CreateKeyRequest) UnsetOwner() {
	o.Owner.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKeyRequest) GetPolicy() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeyRequest) GetPolicyOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *CreateKeyRequest) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given []map[string]interface{} and assigns it to the Policy field.
func (o *CreateKeyRequest) SetPolicy(v []map[string]interface{}) {
	o.Policy = v
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKeyRequest) GetChainId() int64 {
	if o == nil || o.ChainId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeyRequest) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *CreateKeyRequest) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *CreateKeyRequest) SetChainId(v int64) {
	o.ChainId.Set(&v)
}

// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *CreateKeyRequest) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *CreateKeyRequest) UnsetChainId() {
	o.ChainId.Unset()
}

// GetCount returns the Count field value
func (o *CreateKeyRequest) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CreateKeyRequest) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *CreateKeyRequest) SetCount(v int32) {
	o.Count = v
}

// GetKeyType returns the KeyType field value
func (o *CreateKeyRequest) GetKeyType() api.KeyType {
	if o == nil {
		var ret api.KeyType
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *CreateKeyRequest) GetKeyTypeOk() (*api.KeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *CreateKeyRequest) SetKeyType(v api.KeyType) {
	o.KeyType = v
}

func (o CreateKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EditPolicy.IsSet() {
		toSerialize["edit_policy"] = o.EditPolicy.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["key_type"] = o.KeyType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKeyRequest struct {
	value *CreateKeyRequest
	isSet bool
}

func (v NullableCreateKeyRequest) Get() *CreateKeyRequest {
	return v.value
}

func (v *NullableCreateKeyRequest) Set(val *CreateKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyRequest(val *CreateKeyRequest) *NullableCreateKeyRequest {
	return &NullableCreateKeyRequest{value: val, isSet: true}
}

func (v NullableCreateKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
