package api

import (
	"encoding/json"
)

type IBabylonStakingParams interface {
	// GetCovenantPKs returns the covenant public keys
	GetCovenantPKs() []string

	// GetCovenantQuorum returns the covenant quorum threshold
	GetCovenantQuorum() int

	// GetMinStakingValueSat returns the minimum staking value in satoshis
	GetMinStakingValueSat() int64

	// GetMaxStakingValueSat returns the maximum staking value in satoshis
	GetMaxStakingValueSat() int64

	// GetMinStakingTimeBlocks returns the minimum staking time in blocks
	GetMinStakingTimeBlocks() int64

	// GetMaxStakingTimeBlocks returns the maximum staking time in blocks
	GetMaxStakingTimeBlocks() int64

	// GetSlashingPKScript returns the Base64-encoded slashing PK script
	GetSlashingPKScript() string

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
	CovenantPKs                  []string `json:"covenant_pks"`
	CovenantQuorum               int      `json:"covenant_quorum"`
	MinStakingValueSat           string   `json:"min_staking_value_sat"`
	MaxStakingValueSat           string   `json:"max_staking_value_sat"`
	MinStakingTimeBlocks         int64    `json:"min_staking_time_blocks"`
	MaxStakingTimeBlocks         int64    `json:"max_staking_time_blocks"`
	SlashingPKScript             string   `json:"slashing_pk_script"`
	MinSlashingTxFeeSat          string   `json:"min_slashing_tx_fee_sat"`
	SlashingRate                 string   `json:"slashing_rate"`
	UnbondingTimeBlocks          int64    `json:"unbonding_time_blocks"`
	UnbondingFeeSat              string   `json:"unbonding_fee_sat"`
	MinCommissionRate            string   `json:"min_commission_rate"`
	DelegationCreationBaseGasFee string   `json:"delegation_creation_base_gas_fee"`
	AllowListExpirationHeight    string   `json:"allow_list_expiration_height"`
	BTCActivationHeight          int64    `json:"btc_activation_height"`
	Version                      int32    `json:"version"`
}

// NewBabylonStakingParams instantiates a new BabylonStakingParams object
func NewBabylonStakingParams(
	covenantPKs []string,
	covenantQuorum int,
	minStakingValueSat string,
	maxStakingValueSat string,
	minStakingTimeBlocks int64,
	maxStakingTimeBlocks int64,
	slashingPKScript string,
	minSlashingTxFeeSat string,
	slashingRate string,
	unbondingTimeBlocks int64,
	unbondingFeeSat string,
	minCommissionRate string,
	delegationCreationBaseGasFee string,
	allowListExpirationHeight string,
	btcActivationHeight int64,
	version int32,
) *BabylonStakingParams {
	this := BabylonStakingParams{}
	this.CovenantPKs = covenantPKs
	this.CovenantQuorum = covenantQuorum
	this.MinStakingValueSat = minStakingValueSat
	this.MaxStakingValueSat = maxStakingValueSat
	this.MinStakingTimeBlocks = minStakingTimeBlocks
	this.MaxStakingTimeBlocks = maxStakingTimeBlocks
	this.SlashingPKScript = slashingPKScript
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

// GetCovenantPKs returns the CovenantPKs field value
func (o *BabylonStakingParams) GetCovenantPKs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CovenantPKs
}

// GetCovenantPKsOk returns a tuple with the CovenantPKs field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetCovenantPKsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CovenantPKs, true
}

// SetCovenantPKs sets field value
func (o *BabylonStakingParams) SetCovenantPKs(v []string) {
	o.CovenantPKs = v
}

// GetCovenantQuorum returns the CovenantQuorum field value
func (o *BabylonStakingParams) GetCovenantQuorum() int {
	if o == nil {
		var ret int
		return ret
	}
	return o.CovenantQuorum
}

// GetCovenantQuorumOk returns a tuple with the CovenantQuorum field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetCovenantQuorumOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CovenantQuorum, true
}

// SetCovenantQuorum sets field value
func (o *BabylonStakingParams) SetCovenantQuorum(v int) {
	o.CovenantQuorum = v
}

// GetMinStakingValueSat returns the MinStakingValueSat field value
func (o *BabylonStakingParams) GetMinStakingValueSat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MinStakingValueSat
}

// GetMinStakingValueSatOk returns a tuple with the MinStakingValueSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinStakingValueSatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinStakingValueSat, true
}

// SetMinStakingValueSat sets field value
func (o *BabylonStakingParams) SetMinStakingValueSat(v string) {
	o.MinStakingValueSat = v
}

// GetMaxStakingValueSat returns the MaxStakingValueSat field value
func (o *BabylonStakingParams) GetMaxStakingValueSat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MaxStakingValueSat
}

// GetMaxStakingValueSatOk returns a tuple with the MaxStakingValueSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMaxStakingValueSatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingValueSat, true
}

// SetMaxStakingValueSat sets field value
func (o *BabylonStakingParams) SetMaxStakingValueSat(v string) {
	o.MaxStakingValueSat = v
}

// GetMinStakingTimeBlocks returns the MinStakingTimeBlocks field value
func (o *BabylonStakingParams) GetMinStakingTimeBlocks() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MinStakingTimeBlocks
}

// GetMinStakingTimeBlocksOk returns a tuple with the MinStakingTimeBlocks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinStakingTimeBlocksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinStakingTimeBlocks, true
}

// SetMinStakingTimeBlocks sets field value
func (o *BabylonStakingParams) SetMinStakingTimeBlocks(v int64) {
	o.MinStakingTimeBlocks = v
}

// GetMaxStakingTimeBlocks returns the MaxStakingTimeBlocks field value
func (o *BabylonStakingParams) GetMaxStakingTimeBlocks() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxStakingTimeBlocks
}

// GetMaxStakingTimeBlocksOk returns a tuple with the MaxStakingTimeBlocks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMaxStakingTimeBlocksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingTimeBlocks, true
}

// SetMaxStakingTimeBlocks sets field value
func (o *BabylonStakingParams) SetMaxStakingTimeBlocks(v int64) {
	o.MaxStakingTimeBlocks = v
}

// GetSlashingPKScript returns the SlashingPKScript field value
func (o *BabylonStakingParams) GetSlashingPKScript() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SlashingPKScript
}

// GetSlashingPKScriptOk returns a tuple with the SlashingPKScript field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetSlashingPKScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashingPKScript, true
}

// SetSlashingPKScript sets field value
func (o *BabylonStakingParams) SetSlashingPKScript(v string) {
	o.SlashingPKScript = v
}

// GetMinSlashingTxFeeSat returns the MinSlashingTxFeeSat field value
func (o *BabylonStakingParams) GetMinSlashingTxFeeSat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MinSlashingTxFeeSat
}

// GetMinSlashingTxFeeSatOk returns a tuple with the MinSlashingTxFeeSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetMinSlashingTxFeeSatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinSlashingTxFeeSat, true
}

// SetMinSlashingTxFeeSat sets field value
func (o *BabylonStakingParams) SetMinSlashingTxFeeSat(v string) {
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
func (o *BabylonStakingParams) GetUnbondingTimeBlocks() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UnbondingTimeBlocks
}

// GetUnbondingTimeBlocksOk returns a tuple with the UnbondingTimeBlocks field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetUnbondingTimeBlocksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondingTimeBlocks, true
}

// SetUnbondingTimeBlocks sets field value
func (o *BabylonStakingParams) SetUnbondingTimeBlocks(v int64) {
	o.UnbondingTimeBlocks = v
}

// GetUnbondingFeeSat returns the UnbondingFeeSat field value
func (o *BabylonStakingParams) GetUnbondingFeeSat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnbondingFeeSat
}

// GetUnbondingFeeSatOk returns a tuple with the UnbondingFeeSat field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetUnbondingFeeSatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondingFeeSat, true
}

// SetUnbondingFeeSat sets field value
func (o *BabylonStakingParams) SetUnbondingFeeSat(v string) {
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
func (o *BabylonStakingParams) GetDelegationCreationBaseGasFee() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DelegationCreationBaseGasFee
}

// GetDelegationCreationBaseGasFeeOk returns a tuple with the DelegationCreationBaseGasFee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetDelegationCreationBaseGasFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelegationCreationBaseGasFee, true
}

// SetDelegationCreationBaseGasFee sets field value
func (o *BabylonStakingParams) SetDelegationCreationBaseGasFee(v string) {
	o.DelegationCreationBaseGasFee = v
}

// GetAllowListExpirationHeight returns the AllowListExpirationHeight field value
func (o *BabylonStakingParams) GetAllowListExpirationHeight() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AllowListExpirationHeight
}

// GetAllowListExpirationHeightOk returns a tuple with the AllowListExpirationHeight field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetAllowListExpirationHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowListExpirationHeight, true
}

// SetAllowListExpirationHeight sets field value
func (o *BabylonStakingParams) SetAllowListExpirationHeight(v string) {
	o.AllowListExpirationHeight = v
}

// GetBTCActivationHeight returns the BTCActivationHeight field value
func (o *BabylonStakingParams) GetBTCActivationHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BTCActivationHeight
}

// GetBTCActivationHeightOk returns a tuple with the BTCActivationHeight field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingParams) GetBTCActivationHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BTCActivationHeight, true
}

// SetBTCActivationHeight sets field value
func (o *BabylonStakingParams) SetBTCActivationHeight(v int64) {
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

	toSerialize["covenant_pks"] = o.CovenantPKs
	toSerialize["covenant_quorum"] = o.CovenantQuorum
	toSerialize["min_staking_value_sat"] = o.MinStakingValueSat
	toSerialize["max_staking_value_sat"] = o.MaxStakingValueSat
	toSerialize["min_staking_time_blocks"] = o.MinStakingTimeBlocks
	toSerialize["max_staking_time_blocks"] = o.MaxStakingTimeBlocks
	toSerialize["slashing_pk_script"] = o.SlashingPKScript
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

	if covenantPKs, ok := temp["covenant_pks"].([]interface{}); ok {
		o.CovenantPKs = make([]string, len(covenantPKs))
		for i, v := range covenantPKs {
			o.CovenantPKs[i] = v.(string)
		}
	}

	if covenantQuorum, ok := temp["covenant_quorum"].(float64); ok {
		o.CovenantQuorum = int(covenantQuorum)
	}

	if minStakingValueSat, ok := temp["min_staking_value_sat"].(string); ok {
		o.MinStakingValueSat = minStakingValueSat
	}

	if maxStakingValueSat, ok := temp["max_staking_value_sat"].(string); ok {
		o.MaxStakingValueSat = maxStakingValueSat
	}

	if minStakingTimeBlocks, ok := temp["min_staking_time_blocks"].(float64); ok {
		o.MinStakingTimeBlocks = int64(minStakingTimeBlocks)
	}

	if maxStakingTimeBlocks, ok := temp["max_staking_time_blocks"].(float64); ok {
		o.MaxStakingTimeBlocks = int64(maxStakingTimeBlocks)
	}

	if slashingPKScript, ok := temp["slashing_pk_script"].(string); ok {
		o.SlashingPKScript = slashingPKScript
	}

	if minSlashingTxFeeSat, ok := temp["min_slashing_tx_fee_sat"].(string); ok {
		o.MinSlashingTxFeeSat = minSlashingTxFeeSat
	}

	if slashingRate, ok := temp["slashing_rate"].(string); ok {
		o.SlashingRate = slashingRate
	}

	if unbondingTimeBlocks, ok := temp["unbonding_time_blocks"].(float64); ok {
		o.UnbondingTimeBlocks = int64(unbondingTimeBlocks)
	}

	if unbondingFeeSat, ok := temp["unbonding_fee_sat"].(string); ok {
		o.UnbondingFeeSat = unbondingFeeSat
	}

	if minCommissionRate, ok := temp["min_commission_rate"].(string); ok {
		o.MinCommissionRate = minCommissionRate
	}

	if delegationCreationBaseGasFee, ok := temp["delegation_creation_base_gas_fee"].(string); ok {
		o.DelegationCreationBaseGasFee = delegationCreationBaseGasFee
	}

	if allowListExpirationHeight, ok := temp["allow_list_expiration_height"].(string); ok {
		o.AllowListExpirationHeight = allowListExpirationHeight
	}

	if btcActivationHeight, ok := temp["btc_activation_height"].(float64); ok {
		o.BTCActivationHeight = int64(btcActivationHeight)
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
