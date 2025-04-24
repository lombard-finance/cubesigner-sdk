package api

import (
	"encoding/json"
)

type IBabylonStakingParams interface {
	// GetCovenantPks returns the covenant public keys
	GetCovenantPks() []string

	// GetCovenantQuorum returns the covenant quorum threshold
	GetCovenantQuorum() uint32

	// GetMinStakingValueSat returns the minimum staking value in satoshis
	GetMinStakingValueSat() int64

	// GetMaxStakingValueSat returns the maximum staking value in satoshis
	GetMaxStakingValueSat() int64

	// GetMinStakingTimeBlocks returns the minimum staking time in blocks
	GetMinStakingTimeBlocks() int64

	// GetMaxStakingTimeBlocks returns the maximum staking time in blocks
	GetMaxStakingTimeBlocks() int64

	// GetSlashingPkScript returns the Base64-encoded slashing PK script
	GetSlashingPkScript() string

	// GetMinSlashingTxFeeSat returns the minimum slashing transaction fee in satoshis
	GetMinSlashingTxFeeSat() int64

	// GetSlashingRate returns the slashing rate as a string
	GetSlashingRate() string

	// GetUnbondingTimeBlocks returns the unbonding time in blocks
	GetUnbondingTimeBlocks() int64

	// GetUnbondingFeeSat returns the unbonding fee in satoshis
	GetUnbondingFeeSat() int64

	// GetMinCommissionRate returns the minimum commission rate as a string
	GetMinCommissionRate() string

	// GetDelegationCreationBaseGasFee returns the delegation creation base gas fee
	GetDelegationCreationBaseGasFee() int64

	// GetAllowListExpirationHeight returns the allow list expiration height
	GetAllowListExpirationHeight() string

	// GetBTCActivationHeight returns the BTC activation height
	GetBTCActivationHeight() int64

	// GetVersion returns the Version field value
	GetVersion() int32
}

// BabylonStakingParams Parameter set for Babylon staking. The latest parameter sets are available from <https://github.com/babylonchain/networks>
type BabylonStakingParams struct {
	CovenantPks                  []string `json:"covenant_pks"`
	CovenantQuorum               uint32   `json:"covenant_quorum"`
	MinStakingValueSat           int64    `json:"min_staking_value_sat"`
	MaxStakingValueSat           int64    `json:"max_staking_value_sat"`
	MinStakingTimeBlocks         uint32   `json:"min_staking_time_blocks"`
	MaxStakingTimeBlocks         uint32   `json:"max_staking_time_blocks"`
	SlashingPkScript             string   `json:"slashing_pk_script"`
	MinSlashingTxFeeSat          int64    `json:"min_slashing_tx_fee_sat"`
	SlashingRate                 string   `json:"slashing_rate"`
	UnbondingTimeBlocks          uint32   `json:"unbonding_time_blocks"`
	UnbondingFeeSat              int64    `json:"unbonding_fee_sat"`
	MinCommissionRate            string   `json:"min_commission_rate"`
	DelegationCreationBaseGasFee uint64   `json:"delegation_creation_base_gas_fee"`
	AllowListExpirationHeight    uint64   `json:"allow_list_expiration_height"`
	BTCActivationHeight          uint32   `json:"btc_activation_height"`
	Version                      int32    `json:"version"`
}

// NewBabylonStakingParams instantiates a new BabylonStakingParams object
func NewBabylonStakingParams(
	covenantPks []string,
	covenantQuorum uint32,
	minStakingValueSat int64,
	maxStakingValueSat int64,
	minStakingTimeBlocks uint32,
	maxStakingTimeBlocks uint32,
	slashingPKScript string,
	minSlashingTxFeeSat int64,
	slashingRate string,
	unbondingTimeBlocks uint32,
	unbondingFeeSat int64,
	minCommissionRate string,
	delegationCreationBaseGasFee uint64,
	allowListExpirationHeight uint64,
	btcActivationHeight uint32,
	version int32,
) *BabylonStakingParams {
	this := BabylonStakingParams{}
	this.CovenantPks = covenantPks
	this.CovenantQuorum = covenantQuorum
	this.MinStakingValueSat = minStakingValueSat
	this.MaxStakingValueSat = maxStakingValueSat
	this.MinStakingTimeBlocks = minStakingTimeBlocks
	this.MaxStakingTimeBlocks = maxStakingTimeBlocks
	this.SlashingPkScript = slashingPKScript
	this.MinSlashingTxFeeSat = minSlashingTxFeeSat
	this.SlashingRate = slashingRate
	this.UnbondingTimeBlocks = unbondingTimeBlocks
	this.UnbondingFeeSat = unbondingFeeSat
	this.MinCommissionRate = minCommissionRate
	this.DelegationCreationBaseGasFee = delegationCreationBaseGasFee
	this.AllowListExpirationHeight = allowListExpirationHeight
	this.BTCActivationHeight = btcActivationHeight
	this.Version = version
	return &this
}

// NewBabylonStakingParamsWithDefaults instantiates a new BabylonStakingParams object with default values
func NewBabylonStakingParamsWithDefaults() *BabylonStakingParams {
	this := BabylonStakingParams{}
	return &this
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
func (o *BabylonStakingParams) GetCovenantQuorum() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}
	return o.CovenantQuorum
}

// GetCovenantQuorumOk returns a tuple with the CovenantQuorum field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetCovenantQuorumOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CovenantQuorum, true
}

// SetCovenantQuorum sets field value
func (o *BabylonStakingParams) SetCovenantQuorum(v uint32) {
	o.CovenantQuorum = v
}

// GetMinStakingValueSat returns the MinStakingValueSat field value
func (o *BabylonStakingParams) GetMinStakingValueSat() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MinStakingValueSat
}

// GetMinStakingValueSatOk returns a tuple with the MinStakingValueSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinStakingValueSatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinStakingValueSat, true
}

// SetMinStakingValueSat sets field value
func (o *BabylonStakingParams) SetMinStakingValueSat(v int64) {
	o.MinStakingValueSat = v
}

// GetMaxStakingValueSat returns the MaxStakingValueSat field value
func (o *BabylonStakingParams) GetMaxStakingValueSat() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxStakingValueSat
}

// GetMaxStakingValueSatOk returns a tuple with the MaxStakingValueSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMaxStakingValueSatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingValueSat, true
}

// SetMaxStakingValueSat sets field value
func (o *BabylonStakingParams) SetMaxStakingValueSat(v int64) {
	o.MaxStakingValueSat = v
}

// GetMinStakingTimeBlocks returns the MinStakingTimeBlocks field value
func (o *BabylonStakingParams) GetMinStakingTimeBlocks() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}
	return o.MinStakingTimeBlocks
}

// GetMinStakingTimeBlocksOk returns a tuple with the MinStakingTimeBlocks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinStakingTimeBlocksOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinStakingTimeBlocks, true
}

// SetMinStakingTimeBlocks sets field value
func (o *BabylonStakingParams) SetMinStakingTimeBlocks(v uint32) {
	o.MinStakingTimeBlocks = v
}

// GetMaxStakingTimeBlocks returns the MaxStakingTimeBlocks field value
func (o *BabylonStakingParams) GetMaxStakingTimeBlocks() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}
	return o.MaxStakingTimeBlocks
}

// GetMaxStakingTimeBlocksOk returns a tuple with the MaxStakingTimeBlocks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMaxStakingTimeBlocksOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingTimeBlocks, true
}

// SetMaxStakingTimeBlocks sets field value
func (o *BabylonStakingParams) SetMaxStakingTimeBlocks(v uint32) {
	o.MaxStakingTimeBlocks = v
}

// GetSlashingPkScript returns the SlashingPkScript field value
func (o *BabylonStakingParams) GetSlashingPkScript() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SlashingPkScript
}

// GetSlashingPkScriptOk returns a tuple with the SlashingPkScript field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetSlashingPkScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashingPkScript, true
}

// SetSlashingPkScript sets field value
func (o *BabylonStakingParams) SetSlashingPkScript(v string) {
	o.SlashingPkScript = v
}

// GetMinSlashingTxFeeSat returns the MinSlashingTxFeeSat field value
func (o *BabylonStakingParams) GetMinSlashingTxFeeSat() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MinSlashingTxFeeSat
}

// GetMinSlashingTxFeeSatOk returns a tuple with the MinSlashingTxFeeSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinSlashingTxFeeSatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinSlashingTxFeeSat, true
}

// SetMinSlashingTxFeeSat sets field value
func (o *BabylonStakingParams) SetMinSlashingTxFeeSat(v int64) {
	o.MinSlashingTxFeeSat = v
}

// GetSlashingRate returns the SlashingRate field value
func (o *BabylonStakingParams) GetSlashingRate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SlashingRate
}

// GetSlashingRateOk returns a tuple with the SlashingRate field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetSlashingRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashingRate, true
}

// SetSlashingRate sets field value
func (o *BabylonStakingParams) SetSlashingRate(v string) {
	o.SlashingRate = v
}

// GetUnbondingTimeBlocks returns the UnbondingTimeBlocks field value
func (o *BabylonStakingParams) GetUnbondingTimeBlocks() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}
	return o.UnbondingTimeBlocks
}

// GetUnbondingTimeBlocksOk returns a tuple with the UnbondingTimeBlocks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetUnbondingTimeBlocksOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondingTimeBlocks, true
}

// SetUnbondingTimeBlocks sets field value
func (o *BabylonStakingParams) SetUnbondingTimeBlocks(v uint32) {
	o.UnbondingTimeBlocks = v
}

// GetUnbondingFeeSat returns the UnbondingFeeSat field value
func (o *BabylonStakingParams) GetUnbondingFeeSat() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UnbondingFeeSat
}

// GetUnbondingFeeSatOk returns a tuple with the UnbondingFeeSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetUnbondingFeeSatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondingFeeSat, true
}

// SetUnbondingFeeSat sets field value
func (o *BabylonStakingParams) SetUnbondingFeeSat(v int64) {
	o.UnbondingFeeSat = v
}

// GetMinCommissionRate returns the MinCommissionRate field value
func (o *BabylonStakingParams) GetMinCommissionRate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MinCommissionRate
}

// GetMinCommissionRateOk returns a tuple with the MinCommissionRate field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinCommissionRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinCommissionRate, true
}

// SetMinCommissionRate sets field value
func (o *BabylonStakingParams) SetMinCommissionRate(v string) {
	o.MinCommissionRate = v
}

// GetDelegationCreationBaseGasFee returns the DelegationCreationBaseGasFee field value
func (o *BabylonStakingParams) GetDelegationCreationBaseGasFee() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}
	return o.DelegationCreationBaseGasFee
}

// GetDelegationCreationBaseGasFeeOk returns a tuple with the DelegationCreationBaseGasFee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetDelegationCreationBaseGasFeeOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelegationCreationBaseGasFee, true
}

// SetDelegationCreationBaseGasFee sets field value
func (o *BabylonStakingParams) SetDelegationCreationBaseGasFee(v uint64) {
	o.DelegationCreationBaseGasFee = v
}

// GetAllowListExpirationHeight returns the AllowListExpirationHeight field value
func (o *BabylonStakingParams) GetAllowListExpirationHeight() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}
	return o.AllowListExpirationHeight
}

// GetAllowListExpirationHeightOk returns a tuple with the AllowListExpirationHeight field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetAllowListExpirationHeightOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowListExpirationHeight, true
}

// SetAllowListExpirationHeight sets field value
func (o *BabylonStakingParams) SetAllowListExpirationHeight(v uint64) {
	o.AllowListExpirationHeight = v
}

// GetBTCActivationHeight returns the BTCActivationHeight field value
func (o *BabylonStakingParams) GetBTCActivationHeight() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}
	return o.BTCActivationHeight
}

// GetBTCActivationHeightOk returns a tuple with the BTCActivationHeight field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetBTCActivationHeightOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BTCActivationHeight, true
}

// SetBTCActivationHeight sets field value
func (o *BabylonStakingParams) SetBTCActivationHeight(v uint32) {
	o.BTCActivationHeight = v
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

// MarshalJSON marshals the BabylonStakingParams to JSON
func (o BabylonStakingParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	toSerialize["covenant_pks"] = o.CovenantPks
	toSerialize["covenant_quorum"] = o.CovenantQuorum
	toSerialize["min_staking_value_sat"] = o.MinStakingValueSat
	toSerialize["max_staking_value_sat"] = o.MaxStakingValueSat
	toSerialize["min_staking_time_blocks"] = o.MinStakingTimeBlocks
	toSerialize["max_staking_time_blocks"] = o.MaxStakingTimeBlocks
	toSerialize["slashing_pk_script"] = o.SlashingPkScript
	toSerialize["min_slashing_tx_fee_sat"] = o.MinSlashingTxFeeSat
	toSerialize["slashing_rate"] = o.SlashingRate
	toSerialize["unbonding_time_blocks"] = o.UnbondingTimeBlocks
	toSerialize["unbonding_fee_sat"] = o.UnbondingFeeSat
	toSerialize["min_commission_rate"] = o.MinCommissionRate
	toSerialize["delegation_creation_base_gas_fee"] = o.DelegationCreationBaseGasFee
	toSerialize["allow_list_expiration_height"] = o.AllowListExpirationHeight
	toSerialize["btc_activation_height"] = o.BTCActivationHeight
	toSerialize["version"] = o.Version

	return json.Marshal(toSerialize)
}

// UnmarshalJSON unmarshals BabylonStakingParams from JSON
func (o *BabylonStakingParams) UnmarshalJSON(bytes []byte) error {
	var temp map[string]interface{}
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	if covenantPks, ok := temp["covenant_pks"].([]interface{}); ok {
		o.CovenantPks = make([]string, len(covenantPks))
		for i, v := range covenantPks {
			o.CovenantPks[i] = v.(string)
		}
	}

	if covenantQuorum, ok := temp["covenant_quorum"].(float64); ok {
		o.CovenantQuorum = uint32(covenantQuorum)
	}

	if minStakingValueSat, ok := temp["min_staking_value_sat"].(float64); ok {
		o.MinStakingValueSat = int64(minStakingValueSat)
	}

	if maxStakingValueSat, ok := temp["max_staking_value_sat"].(float64); ok {
		o.MaxStakingValueSat = int64(maxStakingValueSat)
	}

	if minStakingTimeBlocks, ok := temp["min_staking_time_blocks"].(float64); ok {
		o.MinStakingTimeBlocks = uint32(minStakingTimeBlocks)
	}

	if maxStakingTimeBlocks, ok := temp["max_staking_time_blocks"].(float64); ok {
		o.MaxStakingTimeBlocks = uint32(maxStakingTimeBlocks)
	}

	if slashingPKScript, ok := temp["slashing_pk_script"].(string); ok {
		o.SlashingPkScript = slashingPKScript
	}

	if minSlashingTxFeeSat, ok := temp["min_slashing_tx_fee_sat"].(float64); ok {
		o.MinSlashingTxFeeSat = int64(minSlashingTxFeeSat)
	}

	if slashingRate, ok := temp["slashing_rate"].(string); ok {
		o.SlashingRate = slashingRate
	}

	if unbondingTimeBlocks, ok := temp["unbonding_time_blocks"].(float64); ok {
		o.UnbondingTimeBlocks = uint32(unbondingTimeBlocks)
	}

	if unbondingFeeSat, ok := temp["unbonding_fee_sat"].(float64); ok {
		o.UnbondingFeeSat = int64(unbondingFeeSat)
	}

	if minCommissionRate, ok := temp["min_commission_rate"].(string); ok {
		o.MinCommissionRate = minCommissionRate
	}

	if delegationCreationBaseGasFee, ok := temp["delegation_creation_base_gas_fee"].(float64); ok {
		o.DelegationCreationBaseGasFee = uint64(delegationCreationBaseGasFee)
	}

	if allowListExpirationHeight, ok := temp["allow_list_expiration_height"].(float64); ok {
		o.AllowListExpirationHeight = uint64(allowListExpirationHeight)
	}

	if btcActivationHeight, ok := temp["btc_activation_height"].(float64); ok {
		o.BTCActivationHeight = uint32(btcActivationHeight)
	}

	if version, ok := temp["version"].(float64); ok {
		o.Version = int32(version)
	}

	return nil
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
