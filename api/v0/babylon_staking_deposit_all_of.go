package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonStakingDepositAllOf Data needed to create the Babylon deposit scripts
type BabylonStakingDepositAllOf struct {
	ExplicitParams api.NullableBabylonStakingParams `json:"explicit_params,omitempty"`
	// The Schnorr public key (i.e., 32-byte X-coordinate) of the finality provider to which the stake is delegated.
	FinalityProviderPk string `json:"finality_provider_pk"`
	// The lock time used for the withdrawal output in the staking deposit transaction
	LockTime int32 `json:"lock_time"`
	Network api.BabylonNetworkId `json:"network"`
	// The Schnorr public key (i.e., 32-byte X-coordinate) of the staker. This is the key that signs the slashing, withdrawal, and unbonding scripts.
	StakerPk string `json:"staker_pk"`
	// The parameter version to use. If `None`, uses the latest version.
	Version api.NullableInt32 `json:"version,omitempty"`
}

// NewBabylonStakingDepositAllOf instantiates a new BabylonStakingDepositAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingDepositAllOf(finalityProviderPk string, lockTime int32, network api.BabylonNetworkId, stakerPk string) *BabylonStakingDepositAllOf {
	this := BabylonStakingDepositAllOf{}
	this.FinalityProviderPk = finalityProviderPk
	this.LockTime = lockTime
	this.Network = network
	this.StakerPk = stakerPk
	return &this
}

// NewBabylonStakingDepositAllOfWithDefaults instantiates a new BabylonStakingDepositAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingDepositAllOfWithDefaults() *BabylonStakingDepositAllOf {
	this := BabylonStakingDepositAllOf{}
	return &this
}

// GetExplicitParams returns the ExplicitParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingDepositAllOf) GetExplicitParams() api.BabylonStakingParams {
	if o == nil || o.ExplicitParams.Get() == nil {
		var ret api.BabylonStakingParams
		return ret
	}
	return *o.ExplicitParams.Get()
}

// GetExplicitParamsOk returns a tuple with the ExplicitParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingDepositAllOf) GetExplicitParamsOk() (*api.BabylonStakingParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExplicitParams.Get(), o.ExplicitParams.IsSet()
}

// HasExplicitParams returns a boolean if a field has been set.
func (o *BabylonStakingDepositAllOf) HasExplicitParams() bool {
	if o != nil && o.ExplicitParams.IsSet() {
		return true
	}

	return false
}

// SetExplicitParams gets a reference to the given NullableBabylonStakingParams and assigns it to the ExplicitParams field.
func (o *BabylonStakingDepositAllOf) SetExplicitParams(v api.BabylonStakingParams) {
	o.ExplicitParams.Set(&v)
}
// SetExplicitParamsNil sets the value for ExplicitParams to be an explicit nil
func (o *BabylonStakingDepositAllOf) SetExplicitParamsNil() {
	o.ExplicitParams.Set(nil)
}

// UnsetExplicitParams ensures that no value is present for ExplicitParams, not even an explicit nil
func (o *BabylonStakingDepositAllOf) UnsetExplicitParams() {
	o.ExplicitParams.Unset()
}

// GetFinalityProviderPk returns the FinalityProviderPk field value
func (o *BabylonStakingDepositAllOf) GetFinalityProviderPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinalityProviderPk
}

// GetFinalityProviderPkOk returns a tuple with the FinalityProviderPk field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf) GetFinalityProviderPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalityProviderPk, true
}

// SetFinalityProviderPk sets field value
func (o *BabylonStakingDepositAllOf) SetFinalityProviderPk(v string) {
	o.FinalityProviderPk = v
}

// GetLockTime returns the LockTime field value
func (o *BabylonStakingDepositAllOf) GetLockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf) GetLockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTime, true
}

// SetLockTime sets field value
func (o *BabylonStakingDepositAllOf) SetLockTime(v int32) {
	o.LockTime = v
}

// GetNetwork returns the Network field value
func (o *BabylonStakingDepositAllOf) GetNetwork() api.BabylonNetworkId {
	if o == nil {
		var ret api.BabylonNetworkId
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf) GetNetworkOk() (*api.BabylonNetworkId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *BabylonStakingDepositAllOf) SetNetwork(v api.BabylonNetworkId) {
	o.Network = v
}

// GetStakerPk returns the StakerPk field value
func (o *BabylonStakingDepositAllOf) GetStakerPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakerPk
}

// GetStakerPkOk returns a tuple with the StakerPk field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf) GetStakerPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakerPk, true
}

// SetStakerPk sets field value
func (o *BabylonStakingDepositAllOf) SetStakerPk(v string) {
	o.StakerPk = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingDepositAllOf) GetVersion() int32 {
	if o == nil || o.Version.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingDepositAllOf) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *BabylonStakingDepositAllOf) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableInt32 and assigns it to the Version field.
func (o *BabylonStakingDepositAllOf) SetVersion(v int32) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *BabylonStakingDepositAllOf) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *BabylonStakingDepositAllOf) UnsetVersion() {
	o.Version.Unset()
}

func (o BabylonStakingDepositAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExplicitParams.IsSet() {
		toSerialize["explicit_params"] = o.ExplicitParams.Get()
	}
	if true {
		toSerialize["finality_provider_pk"] = o.FinalityProviderPk
	}
	if true {
		toSerialize["lock_time"] = o.LockTime
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["staker_pk"] = o.StakerPk
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBabylonStakingDepositAllOf struct {
	value *BabylonStakingDepositAllOf
	isSet bool
}

func (v NullableBabylonStakingDepositAllOf) Get() *BabylonStakingDepositAllOf {
	return v.value
}

func (v *NullableBabylonStakingDepositAllOf) Set(val *BabylonStakingDepositAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingDepositAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingDepositAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingDepositAllOf(val *BabylonStakingDepositAllOf) *NullableBabylonStakingDepositAllOf {
	return &NullableBabylonStakingDepositAllOf{value: val, isSet: true}
}

func (v NullableBabylonStakingDepositAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingDepositAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


