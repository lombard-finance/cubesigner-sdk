package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonStakingDeposit struct for BabylonStakingDeposit
type BabylonStakingDeposit struct {
	ExplicitParams api.NullableBabylonStakingParams `json:"explicit_params,omitempty"`
	// The Schnorr public key (i.e., 32-byte X-coordinate) of the finality provider to which the stake is delegated.
	FinalityProviderPk string `json:"finality_provider_pk"`
	// The lock time used for the withdrawal output in the staking deposit transaction
	LockTime int32            `json:"lock_time"`
	Network  api.BabylonNetworkId `json:"network"`
	// The Schnorr public key (i.e., 32-byte X-coordinate) of the staker. This is the key that signs the slashing, withdrawal, and unbonding scripts.
	StakerPk string `json:"staker_pk"`
	// The parameter version to use. If `None`, uses the latest version.
	Version api.NullableInt32 `json:"version,omitempty"`
	// The change address, specified as a Bitcoin spend script
	Change string `json:"change"`
	// The transaction fee value. The `fee_type` field determines whether this is a fixed fee in sats or a rate in sats per (estimated) virtual byte of transaction weight (i.e., sats per vb).
	Fee     int64       `json:"fee"`
	FeeType api.FeeType `json:"fee_type"`
	// By default, the PSBT in this request can only specify transaction inputs: PSBTs that specify outputs will result in an error. When this flag is set, existing outputs in the PSBT will instead be ignored (i.e., deleted from the PSBT) when creating the deposit transaction.
	IgnorePsbtOutputs *bool `json:"ignore_psbt_outputs,omitempty"`
	// A hex-serialized PSBT (version 0) containing the transaction inputs and all necessary information for signing (e.g., taproot path and leaf information). This PSBT must not have any transaction outputs; they will be added to the returned PSBT.
	Psbt string `json:"psbt"`
	// An optional lock height (in blocks) for this transaction. The resulting transaction cannot be mined before the specified block height.
	TxnLockHeight api.NullableInt32 `json:"txn_lock_height,omitempty"`
	// The value to be staked in sats
	Value int64 `json:"value"`
}

// NewBabylonStakingDeposit instantiates a new BabylonStakingDeposit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingDeposit(finalityProviderPk string, lockTime int32, network api.BabylonNetworkId, stakerPk string, change string, fee int64, feeType api.FeeType, psbt string, value int64) *BabylonStakingDeposit {
	this := BabylonStakingDeposit{}
	this.FinalityProviderPk = finalityProviderPk
	this.LockTime = lockTime
	this.Network = network
	this.StakerPk = stakerPk
	this.Change = change
	this.Fee = fee
	this.FeeType = feeType
	this.Psbt = psbt
	this.Value = value
	return &this
}

// NewBabylonStakingDepositWithDefaults instantiates a new BabylonStakingDeposit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingDepositWithDefaults() *BabylonStakingDeposit {
	this := BabylonStakingDeposit{}
	return &this
}

// GetExplicitParams returns the ExplicitParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingDeposit) GetExplicitParams() api.BabylonStakingParams {
	if o == nil || o.ExplicitParams.Get() == nil {
		var ret api.BabylonStakingParams
		return ret
	}
	return *o.ExplicitParams.Get()
}

// GetExplicitParamsOk returns a tuple with the ExplicitParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingDeposit) GetExplicitParamsOk() (*api.BabylonStakingParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExplicitParams.Get(), o.ExplicitParams.IsSet()
}

// HasExplicitParams returns a boolean if a field has been set.
func (o *BabylonStakingDeposit) HasExplicitParams() bool {
	if o != nil && o.ExplicitParams.IsSet() {
		return true
	}

	return false
}

// SetExplicitParams gets a reference to the given NullableBabylonStakingParams and assigns it to the ExplicitParams field.
func (o *BabylonStakingDeposit) SetExplicitParams(v api.BabylonStakingParams) {
	o.ExplicitParams.Set(&v)
}

// SetExplicitParamsNil sets the value for ExplicitParams to be an explicit nil
func (o *BabylonStakingDeposit) SetExplicitParamsNil() {
	o.ExplicitParams.Set(nil)
}

// UnsetExplicitParams ensures that no value is present for ExplicitParams, not even an explicit nil
func (o *BabylonStakingDeposit) UnsetExplicitParams() {
	o.ExplicitParams.Unset()
}

// GetFinalityProviderPk returns the FinalityProviderPk field value
func (o *BabylonStakingDeposit) GetFinalityProviderPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinalityProviderPk
}

// GetFinalityProviderPkOk returns a tuple with the FinalityProviderPk field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetFinalityProviderPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalityProviderPk, true
}

// SetFinalityProviderPk sets field value
func (o *BabylonStakingDeposit) SetFinalityProviderPk(v string) {
	o.FinalityProviderPk = v
}

// GetLockTime returns the LockTime field value
func (o *BabylonStakingDeposit) GetLockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetLockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTime, true
}

// SetLockTime sets field value
func (o *BabylonStakingDeposit) SetLockTime(v int32) {
	o.LockTime = v
}

// GetNetwork returns the Network field value
func (o *BabylonStakingDeposit) GetNetwork() api.BabylonNetworkId {
	if o == nil {
		var ret api.BabylonNetworkId
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetNetworkOk() (*api.BabylonNetworkId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *BabylonStakingDeposit) SetNetwork(v api.BabylonNetworkId) {
	o.Network = v
}

// GetStakerPk returns the StakerPk field value
func (o *BabylonStakingDeposit) GetStakerPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakerPk
}

// GetStakerPkOk returns a tuple with the StakerPk field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetStakerPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakerPk, true
}

// SetStakerPk sets field value
func (o *BabylonStakingDeposit) SetStakerPk(v string) {
	o.StakerPk = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingDeposit) GetVersion() int32 {
	if o == nil || o.Version.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingDeposit) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *BabylonStakingDeposit) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableInt32 and assigns it to the Version field.
func (o *BabylonStakingDeposit) SetVersion(v int32) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *BabylonStakingDeposit) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *BabylonStakingDeposit) UnsetVersion() {
	o.Version.Unset()
}

// GetChange returns the Change field value
func (o *BabylonStakingDeposit) GetChange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Change
}

// GetChangeOk returns a tuple with the Change field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Change, true
}

// SetChange sets field value
func (o *BabylonStakingDeposit) SetChange(v string) {
	o.Change = v
}

// GetFee returns the Fee field value
func (o *BabylonStakingDeposit) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *BabylonStakingDeposit) SetFee(v int64) {
	o.Fee = v
}

// GetFeeType returns the FeeType field value
func (o *BabylonStakingDeposit) GetFeeType() api.FeeType {
	if o == nil {
		var ret api.FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetFeeTypeOk() (*api.FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *BabylonStakingDeposit) SetFeeType(v api.FeeType) {
	o.FeeType = v
}

// GetIgnorePsbtOutputs returns the IgnorePsbtOutputs field value if set, zero value otherwise.
func (o *BabylonStakingDeposit) GetIgnorePsbtOutputs() bool {
	if o == nil || o.IgnorePsbtOutputs == nil {
		var ret bool
		return ret
	}
	return *o.IgnorePsbtOutputs
}

// GetIgnorePsbtOutputsOk returns a tuple with the IgnorePsbtOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetIgnorePsbtOutputsOk() (*bool, bool) {
	if o == nil || o.IgnorePsbtOutputs == nil {
		return nil, false
	}
	return o.IgnorePsbtOutputs, true
}

// HasIgnorePsbtOutputs returns a boolean if a field has been set.
func (o *BabylonStakingDeposit) HasIgnorePsbtOutputs() bool {
	if o != nil && o.IgnorePsbtOutputs != nil {
		return true
	}

	return false
}

// SetIgnorePsbtOutputs gets a reference to the given bool and assigns it to the IgnorePsbtOutputs field.
func (o *BabylonStakingDeposit) SetIgnorePsbtOutputs(v bool) {
	o.IgnorePsbtOutputs = &v
}

// GetPsbt returns the Psbt field value
func (o *BabylonStakingDeposit) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *BabylonStakingDeposit) SetPsbt(v string) {
	o.Psbt = v
}

// GetTxnLockHeight returns the TxnLockHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingDeposit) GetTxnLockHeight() int32 {
	if o == nil || o.TxnLockHeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TxnLockHeight.Get()
}

// GetTxnLockHeightOk returns a tuple with the TxnLockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingDeposit) GetTxnLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxnLockHeight.Get(), o.TxnLockHeight.IsSet()
}

// HasTxnLockHeight returns a boolean if a field has been set.
func (o *BabylonStakingDeposit) HasTxnLockHeight() bool {
	if o != nil && o.TxnLockHeight.IsSet() {
		return true
	}

	return false
}

// SetTxnLockHeight gets a reference to the given NullableInt32 and assigns it to the TxnLockHeight field.
func (o *BabylonStakingDeposit) SetTxnLockHeight(v int32) {
	o.TxnLockHeight.Set(&v)
}

// SetTxnLockHeightNil sets the value for TxnLockHeight to be an explicit nil
func (o *BabylonStakingDeposit) SetTxnLockHeightNil() {
	o.TxnLockHeight.Set(nil)
}

// UnsetTxnLockHeight ensures that no value is present for TxnLockHeight, not even an explicit nil
func (o *BabylonStakingDeposit) UnsetTxnLockHeight() {
	o.TxnLockHeight.Unset()
}

// GetValue returns the Value field value
func (o *BabylonStakingDeposit) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDeposit) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BabylonStakingDeposit) SetValue(v int64) {
	o.Value = v
}

func (o BabylonStakingDeposit) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["change"] = o.Change
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["fee_type"] = o.FeeType
	}
	if o.IgnorePsbtOutputs != nil {
		toSerialize["ignore_psbt_outputs"] = o.IgnorePsbtOutputs
	}
	if true {
		toSerialize["psbt"] = o.Psbt
	}
	if o.TxnLockHeight.IsSet() {
		toSerialize["txn_lock_height"] = o.TxnLockHeight.Get()
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBabylonStakingDeposit struct {
	value *BabylonStakingDeposit
	isSet bool
}

func (v NullableBabylonStakingDeposit) Get() *BabylonStakingDeposit {
	return v.value
}

func (v *NullableBabylonStakingDeposit) Set(val *BabylonStakingDeposit) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingDeposit) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingDeposit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingDeposit(val *BabylonStakingDeposit) *NullableBabylonStakingDeposit {
	return &NullableBabylonStakingDeposit{value: val, isSet: true}
}

func (v NullableBabylonStakingDeposit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingDeposit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
