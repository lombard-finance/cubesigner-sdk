package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// GetKeyInOrg200Response struct for GetKeyInOrg200Response
type GetKeyInOrg200Response struct {
	Created      api.NullableInt64 `json:"created,omitempty"`
	EditPolicy   *EditPolicy       `json:"edit_policy,omitempty"`
	LastModified api.NullableInt64 `json:"last_modified,omitempty"`
	// User-defined metadata. When rendering (e.g., in the browser) you should treat it as untrusted user data (and avoid injecting metadata into HTML directly) if untrusted users can create/update keys (or their metadata).
	Metadata interface{} `json:"metadata,omitempty"`
	// Version of this object
	Version        *int64                                            `json:"version,omitempty"`
	DerivationInfo NullableGetKeyInOrg200ResponseAllOfDerivationInfo `json:"derivation_info,omitempty"`
	// Whether the key is enabled (only enabled keys may be used for signing)
	Enabled bool `json:"enabled"`
	// The id of the key: \"Key#\" followed by a unique identifier specific to the type of key (such as a public key for BLS or an ethereum address for Secp)
	KeyId   string      `json:"key_id"`
	KeyType api.KeyType `json:"key_type"`
	// A unique identifier specific to the type of key, such as a public key or an ethereum address
	MaterialId string `json:"material_id"`
	// Owner of the key
	Owner string `json:"owner"`
	// Key policy
	Policy []map[string]interface{} `json:"policy"`
	// Hex-encoded, serialized public key. The format used depends on the key type: - Secp256k1 keys use 65-byte uncompressed SECG format; - Stark keys use 33-byte compressed SECG format; - BLS keys use 48-byte compressed BLS12-381 (ZCash) format; - Ed25519 keys use the canonical 32-byte encoding specified in RFC 8032.
	PublicKey string `json:"public_key"`
	// The purpose for which the key can be used (e.g., chain id for which the key is allowed to sign messages)
	Purpose string `json:"purpose"`
}

// NewGetKeyInOrg200Response instantiates a new GetKeyInOrg200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKeyInOrg200Response(enabled bool, keyId string, keyType api.KeyType, materialId string, owner string, policy []map[string]interface{}, publicKey string, purpose string) *GetKeyInOrg200Response {
	this := GetKeyInOrg200Response{}
	this.Enabled = enabled
	this.KeyId = keyId
	this.KeyType = keyType
	this.MaterialId = materialId
	this.Owner = owner
	this.Policy = policy
	this.PublicKey = publicKey
	this.Purpose = purpose
	return &this
}

// NewGetKeyInOrg200ResponseWithDefaults instantiates a new GetKeyInOrg200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKeyInOrg200ResponseWithDefaults() *GetKeyInOrg200Response {
	this := GetKeyInOrg200Response{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetKeyInOrg200Response) GetCreated() int64 {
	if o == nil || o.Created.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetKeyInOrg200Response) GetCreatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *GetKeyInOrg200Response) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableInt64 and assigns it to the Created field.
func (o *GetKeyInOrg200Response) SetCreated(v int64) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *GetKeyInOrg200Response) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *GetKeyInOrg200Response) UnsetCreated() {
	o.Created.Unset()
}

// GetEditPolicy returns the EditPolicy field value if set, zero value otherwise.
func (o *GetKeyInOrg200Response) GetEditPolicy() EditPolicy {
	if o == nil || o.EditPolicy == nil {
		var ret EditPolicy
		return ret
	}
	return *o.EditPolicy
}

// GetEditPolicyOk returns a tuple with the EditPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetEditPolicyOk() (*EditPolicy, bool) {
	if o == nil || o.EditPolicy == nil {
		return nil, false
	}
	return o.EditPolicy, true
}

// HasEditPolicy returns a boolean if a field has been set.
func (o *GetKeyInOrg200Response) HasEditPolicy() bool {
	if o != nil && o.EditPolicy != nil {
		return true
	}

	return false
}

// SetEditPolicy gets a reference to the given EditPolicy and assigns it to the EditPolicy field.
func (o *GetKeyInOrg200Response) SetEditPolicy(v EditPolicy) {
	o.EditPolicy = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetKeyInOrg200Response) GetLastModified() int64 {
	if o == nil || o.LastModified.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetKeyInOrg200Response) GetLastModifiedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *GetKeyInOrg200Response) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableInt64 and assigns it to the LastModified field.
func (o *GetKeyInOrg200Response) SetLastModified(v int64) {
	o.LastModified.Set(&v)
}

// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *GetKeyInOrg200Response) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *GetKeyInOrg200Response) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetKeyInOrg200Response) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetKeyInOrg200Response) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetKeyInOrg200Response) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GetKeyInOrg200Response) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetKeyInOrg200Response) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetKeyInOrg200Response) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *GetKeyInOrg200Response) SetVersion(v int64) {
	o.Version = &v
}

// GetDerivationInfo returns the DerivationInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetKeyInOrg200Response) GetDerivationInfo() GetKeyInOrg200ResponseAllOfDerivationInfo {
	if o == nil || o.DerivationInfo.Get() == nil {
		var ret GetKeyInOrg200ResponseAllOfDerivationInfo
		return ret
	}
	return *o.DerivationInfo.Get()
}

// GetDerivationInfoOk returns a tuple with the DerivationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetKeyInOrg200Response) GetDerivationInfoOk() (*GetKeyInOrg200ResponseAllOfDerivationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DerivationInfo.Get(), o.DerivationInfo.IsSet()
}

// HasDerivationInfo returns a boolean if a field has been set.
func (o *GetKeyInOrg200Response) HasDerivationInfo() bool {
	if o != nil && o.DerivationInfo.IsSet() {
		return true
	}

	return false
}

// SetDerivationInfo gets a reference to the given NullableGetKeyInOrg200ResponseAllOfDerivationInfo and assigns it to the DerivationInfo field.
func (o *GetKeyInOrg200Response) SetDerivationInfo(v GetKeyInOrg200ResponseAllOfDerivationInfo) {
	o.DerivationInfo.Set(&v)
}

// SetDerivationInfoNil sets the value for DerivationInfo to be an explicit nil
func (o *GetKeyInOrg200Response) SetDerivationInfoNil() {
	o.DerivationInfo.Set(nil)
}

// UnsetDerivationInfo ensures that no value is present for DerivationInfo, not even an explicit nil
func (o *GetKeyInOrg200Response) UnsetDerivationInfo() {
	o.DerivationInfo.Unset()
}

// GetEnabled returns the Enabled field value
func (o *GetKeyInOrg200Response) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetKeyInOrg200Response) SetEnabled(v bool) {
	o.Enabled = v
}

// GetKeyId returns the KeyId field value
func (o *GetKeyInOrg200Response) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *GetKeyInOrg200Response) SetKeyId(v string) {
	o.KeyId = v
}

// GetKeyType returns the KeyType field value
func (o *GetKeyInOrg200Response) GetKeyType() api.KeyType {
	if o == nil {
		var ret api.KeyType
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetKeyTypeOk() (*api.KeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *GetKeyInOrg200Response) SetKeyType(v api.KeyType) {
	o.KeyType = v
}

// GetMaterialId returns the MaterialId field value
func (o *GetKeyInOrg200Response) GetMaterialId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaterialId
}

// GetMaterialIdOk returns a tuple with the MaterialId field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetMaterialIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaterialId, true
}

// SetMaterialId sets field value
func (o *GetKeyInOrg200Response) SetMaterialId(v string) {
	o.MaterialId = v
}

// GetOwner returns the Owner field value
func (o *GetKeyInOrg200Response) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *GetKeyInOrg200Response) SetOwner(v string) {
	o.Owner = v
}

// GetPolicy returns the Policy field value
func (o *GetKeyInOrg200Response) GetPolicy() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetPolicyOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy, true
}

// SetPolicy sets field value
func (o *GetKeyInOrg200Response) SetPolicy(v []map[string]interface{}) {
	o.Policy = v
}

// GetPublicKey returns the PublicKey field value
func (o *GetKeyInOrg200Response) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *GetKeyInOrg200Response) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPurpose returns the Purpose field value
func (o *GetKeyInOrg200Response) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *GetKeyInOrg200Response) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *GetKeyInOrg200Response) SetPurpose(v string) {
	o.Purpose = v
}

func (o GetKeyInOrg200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.EditPolicy != nil {
		toSerialize["edit_policy"] = o.EditPolicy
	}
	if o.LastModified.IsSet() {
		toSerialize["last_modified"] = o.LastModified.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.DerivationInfo.IsSet() {
		toSerialize["derivation_info"] = o.DerivationInfo.Get()
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	if true {
		toSerialize["key_type"] = o.KeyType
	}
	if true {
		toSerialize["material_id"] = o.MaterialId
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
	}
	if true {
		toSerialize["purpose"] = o.Purpose
	}
	return json.Marshal(toSerialize)
}

//
//type NullableGetKeyInOrg200Response struct {
//	value *GetKeyInOrg200Response
//	isSet bool
//}
//
//func (v NullableGetKeyInOrg200Response) Get() *GetKeyInOrg200Response {
//	return v.value
//}
//
//func (v *NullableGetKeyInOrg200Response) Set(val *GetKeyInOrg200Response) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableGetKeyInOrg200Response) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableGetKeyInOrg200Response) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableGetKeyInOrg200Response(val *GetKeyInOrg200Response) *NullableGetKeyInOrg200Response {
//	return &NullableGetKeyInOrg200Response{value: val, isSet: true}
//}
//
//func (v NullableGetKeyInOrg200Response) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableGetKeyInOrg200Response) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
