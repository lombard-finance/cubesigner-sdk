package api

import (
	"encoding/json"
)

// BabylonStakingParams Parameter set for Babylon staking. The latest parameter sets are available from <https://github.com/babylonchain/networks>
type BabylonStakingParams struct {
	// Block height at which these params will enter use
	ActivationHeight int64 `json:"activation_height"`
	// Block height above which staking is disabled
	CapHeight *int64 `json:"cap_height,omitempty"`
	// Number of confirmations before a staking txn is considered finalized by the covenant committee
	ConfirmationDepth int32 `json:"confirmation_depth"`
	// Public keys of the covenant signers
	CovenantPks []string `json:"covenant_pks"`
	// Quorum for covenant signer
	CovenantQuorum int32 `json:"covenant_quorum"`
	// Max amount that can be staked in a single txn
	MaxStakingAmount int64 `json:"max_staking_amount"`
	// Max timelock for staking
	MaxStakingTime int32 `json:"max_staking_time"`
	// Min amount that must be staked
	MinStakingAmount int64 `json:"min_staking_amount"`
	// Min timelock for staking
	MinStakingTime int32 `json:"min_staking_time"`
	// Max total amount staked
	StakingCap *int64 `json:"staking_cap,omitempty"`
	// \"Magic bytes\" tag for staking metadata
	Tag string `json:"tag"`
	// Fee that must be spent as part of the unbonding txn
	UnbondingFee int64 `json:"unbonding_fee"`
	// Min timelock for an unbonding script
	UnbondingTime int32 `json:"unbonding_time"`
	// Parameter version
	Version int32 `json:"version"`
}

// NewBabylonStakingParams instantiates a new BabylonStakingParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingParams(activationHeight int64, confirmationDepth int32, covenantPks []string, covenantQuorum int32, maxStakingAmount int64, maxStakingTime int32, minStakingAmount int64, minStakingTime int32, tag string, unbondingFee int64, unbondingTime int32, version int32) *BabylonStakingParams {
	this := BabylonStakingParams{}
	this.ActivationHeight = activationHeight
	this.ConfirmationDepth = confirmationDepth
	this.CovenantPks = covenantPks
	this.CovenantQuorum = covenantQuorum
	this.MaxStakingAmount = maxStakingAmount
	this.MaxStakingTime = maxStakingTime
	this.MinStakingAmount = minStakingAmount
	this.MinStakingTime = minStakingTime
	this.Tag = tag
	this.UnbondingFee = unbondingFee
	this.UnbondingTime = unbondingTime
	this.Version = version
	return &this
}

// NewBabylonStakingParamsWithDefaults instantiates a new BabylonStakingParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingParamsWithDefaults() *BabylonStakingParams {
	this := BabylonStakingParams{}
	return &this
}

// GetActivationHeight returns the ActivationHeight field value
func (o *BabylonStakingParams) GetActivationHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ActivationHeight
}

// GetActivationHeightOk returns a tuple with the ActivationHeight field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetActivationHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationHeight, true
}

// SetActivationHeight sets field value
func (o *BabylonStakingParams) SetActivationHeight(v int64) {
	o.ActivationHeight = v
}

// GetCapHeight returns the CapHeight field value if set, zero value otherwise.
func (o *BabylonStakingParams) GetCapHeight() int64 {
	if o == nil || o.CapHeight == nil {
		var ret int64
		return ret
	}
	return *o.CapHeight
}

// GetCapHeightOk returns a tuple with the CapHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetCapHeightOk() (*int64, bool) {
	if o == nil || o.CapHeight == nil {
		return nil, false
	}
	return o.CapHeight, true
}

// HasCapHeight returns a boolean if a field has been set.
func (o *BabylonStakingParams) HasCapHeight() bool {
	if o != nil && o.CapHeight != nil {
		return true
	}

	return false
}

// SetCapHeight gets a reference to the given int64 and assigns it to the CapHeight field.
func (o *BabylonStakingParams) SetCapHeight(v int64) {
	o.CapHeight = &v
}

// GetConfirmationDepth returns the ConfirmationDepth field value
func (o *BabylonStakingParams) GetConfirmationDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationDepth
}

// GetConfirmationDepthOk returns a tuple with the ConfirmationDepth field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetConfirmationDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationDepth, true
}

// SetConfirmationDepth sets field value
func (o *BabylonStakingParams) SetConfirmationDepth(v int32) {
	o.ConfirmationDepth = v
}

// GetCovenantPks returns the CovenantPks field value
func (o *BabylonStakingParams) GetCovenantPks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CovenantPks
}

// GetCovenantPksOk returns a tuple with the CovenantPks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetCovenantPksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CovenantPks, true
}

// SetCovenantPks sets field value
func (o *BabylonStakingParams) SetCovenantPks(v []string) {
	o.CovenantPks = v
}

// GetCovenantQuorum returns the CovenantQuorum field value
func (o *BabylonStakingParams) GetCovenantQuorum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CovenantQuorum
}

// GetCovenantQuorumOk returns a tuple with the CovenantQuorum field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetCovenantQuorumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CovenantQuorum, true
}

// SetCovenantQuorum sets field value
func (o *BabylonStakingParams) SetCovenantQuorum(v int32) {
	o.CovenantQuorum = v
}

// GetMaxStakingAmount returns the MaxStakingAmount field value
func (o *BabylonStakingParams) GetMaxStakingAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxStakingAmount
}

// GetMaxStakingAmountOk returns a tuple with the MaxStakingAmount field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMaxStakingAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingAmount, true
}

// SetMaxStakingAmount sets field value
func (o *BabylonStakingParams) SetMaxStakingAmount(v int64) {
	o.MaxStakingAmount = v
}

// GetMaxStakingTime returns the MaxStakingTime field value
func (o *BabylonStakingParams) GetMaxStakingTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxStakingTime
}

// GetMaxStakingTimeOk returns a tuple with the MaxStakingTime field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMaxStakingTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingTime, true
}

// SetMaxStakingTime sets field value
func (o *BabylonStakingParams) SetMaxStakingTime(v int32) {
	o.MaxStakingTime = v
}

// GetMinStakingAmount returns the MinStakingAmount field value
func (o *BabylonStakingParams) GetMinStakingAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinStakingAmount
}

// GetMinStakingAmountOk returns a tuple with the MinStakingAmount field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinStakingAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinStakingAmount, true
}

// SetMinStakingAmount sets field value
func (o *BabylonStakingParams) SetMinStakingAmount(v int64) {
	o.MinStakingAmount = v
}

// GetMinStakingTime returns the MinStakingTime field value
func (o *BabylonStakingParams) GetMinStakingTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinStakingTime
}

// GetMinStakingTimeOk returns a tuple with the MinStakingTime field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinStakingTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinStakingTime, true
}

// SetMinStakingTime sets field value
func (o *BabylonStakingParams) SetMinStakingTime(v int32) {
	o.MinStakingTime = v
}

// GetStakingCap returns the StakingCap field value if set, zero value otherwise.
func (o *BabylonStakingParams) GetStakingCap() int64 {
	if o == nil || o.StakingCap == nil {
		var ret int64
		return ret
	}
	return *o.StakingCap
}

// GetStakingCapOk returns a tuple with the StakingCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetStakingCapOk() (*int64, bool) {
	if o == nil || o.StakingCap == nil {
		return nil, false
	}
	return o.StakingCap, true
}

// HasStakingCap returns a boolean if a field has been set.
func (o *BabylonStakingParams) HasStakingCap() bool {
	if o != nil && o.StakingCap != nil {
		return true
	}

	return false
}

// SetStakingCap gets a reference to the given int64 and assigns it to the StakingCap field.
func (o *BabylonStakingParams) SetStakingCap(v int64) {
	o.StakingCap = &v
}

// GetTag returns the Tag field value
func (o *BabylonStakingParams) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *BabylonStakingParams) SetTag(v string) {
	o.Tag = v
}

// GetUnbondingFee returns the UnbondingFee field value
func (o *BabylonStakingParams) GetUnbondingFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnbondingFee
}

// GetUnbondingFeeOk returns a tuple with the UnbondingFee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetUnbondingFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondingFee, true
}

// SetUnbondingFee sets field value
func (o *BabylonStakingParams) SetUnbondingFee(v int64) {
	o.UnbondingFee = v
}

// GetUnbondingTime returns the UnbondingTime field value
func (o *BabylonStakingParams) GetUnbondingTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnbondingTime
}

// GetUnbondingTimeOk returns a tuple with the UnbondingTime field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetUnbondingTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondingTime, true
}

// SetUnbondingTime sets field value
func (o *BabylonStakingParams) SetUnbondingTime(v int32) {
	o.UnbondingTime = v
}

// GetVersion returns the Version field value
func (o *BabylonStakingParams) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *BabylonStakingParams) SetVersion(v int32) {
	o.Version = v
}

func (o BabylonStakingParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["activation_height"] = o.ActivationHeight
	}
	if o.CapHeight != nil {
		toSerialize["cap_height"] = o.CapHeight
	}
	if true {
		toSerialize["confirmation_depth"] = o.ConfirmationDepth
	}
	if true {
		toSerialize["covenant_pks"] = o.CovenantPks
	}
	if true {
		toSerialize["covenant_quorum"] = o.CovenantQuorum
	}
	if true {
		toSerialize["max_staking_amount"] = o.MaxStakingAmount
	}
	if true {
		toSerialize["max_staking_time"] = o.MaxStakingTime
	}
	if true {
		toSerialize["min_staking_amount"] = o.MinStakingAmount
	}
	if true {
		toSerialize["min_staking_time"] = o.MinStakingTime
	}
	if o.StakingCap != nil {
		toSerialize["staking_cap"] = o.StakingCap
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["unbonding_fee"] = o.UnbondingFee
	}
	if true {
		toSerialize["unbonding_time"] = o.UnbondingTime
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBabylonStakingParams struct {
	value *BabylonStakingParams
	isSet bool
}

func (v NullableBabylonStakingParams) Get() *BabylonStakingParams {
	return v.value
}

func (v *NullableBabylonStakingParams) Set(val *BabylonStakingParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingParams(val *BabylonStakingParams) *NullableBabylonStakingParams {
	return &NullableBabylonStakingParams{value: val, isSet: true}
}

func (v NullableBabylonStakingParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
