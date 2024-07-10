package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonStakingDepositAllOf1 struct for BabylonStakingDepositAllOf1
type BabylonStakingDepositAllOf1 struct {
	// The change address, specified as a Bitcoin spend script
	Change string `json:"change"`
	// The transaction fee value. The `fee_type` field determines whether this is a fixed fee in sats or a rate in sats per (estimated) virtual byte of transaction weight (i.e., sats per vb).
	Fee int64 `json:"fee"`
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

// NewBabylonStakingDepositAllOf1 instantiates a new BabylonStakingDepositAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingDepositAllOf1(change string, fee int64, feeType api.FeeType, psbt string, value int64) *BabylonStakingDepositAllOf1 {
	this := BabylonStakingDepositAllOf1{}
	this.Change = change
	this.Fee = fee
	this.FeeType = feeType
	this.Psbt = psbt
	this.Value = value
	return &this
}

// NewBabylonStakingDepositAllOf1WithDefaults instantiates a new BabylonStakingDepositAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingDepositAllOf1WithDefaults() *BabylonStakingDepositAllOf1 {
	this := BabylonStakingDepositAllOf1{}
	return &this
}

// GetChange returns the Change field value
func (o *BabylonStakingDepositAllOf1) GetChange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Change
}

// GetChangeOk returns a tuple with the Change field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf1) GetChangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Change, true
}

// SetChange sets field value
func (o *BabylonStakingDepositAllOf1) SetChange(v string) {
	o.Change = v
}

// GetFee returns the Fee field value
func (o *BabylonStakingDepositAllOf1) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf1) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *BabylonStakingDepositAllOf1) SetFee(v int64) {
	o.Fee = v
}

// GetFeeType returns the FeeType field value
func (o *BabylonStakingDepositAllOf1) GetFeeType() api.FeeType {
	if o == nil {
		var ret api.FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf1) GetFeeTypeOk() (*api.FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *BabylonStakingDepositAllOf1) SetFeeType(v api.FeeType) {
	o.FeeType = v
}

// GetIgnorePsbtOutputs returns the IgnorePsbtOutputs field value if set, zero value otherwise.
func (o *BabylonStakingDepositAllOf1) GetIgnorePsbtOutputs() bool {
	if o == nil || o.IgnorePsbtOutputs == nil {
		var ret bool
		return ret
	}
	return *o.IgnorePsbtOutputs
}

// GetIgnorePsbtOutputsOk returns a tuple with the IgnorePsbtOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf1) GetIgnorePsbtOutputsOk() (*bool, bool) {
	if o == nil || o.IgnorePsbtOutputs == nil {
		return nil, false
	}
	return o.IgnorePsbtOutputs, true
}

// HasIgnorePsbtOutputs returns a boolean if a field has been set.
func (o *BabylonStakingDepositAllOf1) HasIgnorePsbtOutputs() bool {
	if o != nil && o.IgnorePsbtOutputs != nil {
		return true
	}

	return false
}

// SetIgnorePsbtOutputs gets a reference to the given bool and assigns it to the IgnorePsbtOutputs field.
func (o *BabylonStakingDepositAllOf1) SetIgnorePsbtOutputs(v bool) {
	o.IgnorePsbtOutputs = &v
}

// GetPsbt returns the Psbt field value
func (o *BabylonStakingDepositAllOf1) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf1) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *BabylonStakingDepositAllOf1) SetPsbt(v string) {
	o.Psbt = v
}

// GetTxnLockHeight returns the TxnLockHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingDepositAllOf1) GetTxnLockHeight() int32 {
	if o == nil || o.TxnLockHeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TxnLockHeight.Get()
}

// GetTxnLockHeightOk returns a tuple with the TxnLockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingDepositAllOf1) GetTxnLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxnLockHeight.Get(), o.TxnLockHeight.IsSet()
}

// HasTxnLockHeight returns a boolean if a field has been set.
func (o *BabylonStakingDepositAllOf1) HasTxnLockHeight() bool {
	if o != nil && o.TxnLockHeight.IsSet() {
		return true
	}

	return false
}

// SetTxnLockHeight gets a reference to the given NullableInt32 and assigns it to the TxnLockHeight field.
func (o *BabylonStakingDepositAllOf1) SetTxnLockHeight(v int32) {
	o.TxnLockHeight.Set(&v)
}
// SetTxnLockHeightNil sets the value for TxnLockHeight to be an explicit nil
func (o *BabylonStakingDepositAllOf1) SetTxnLockHeightNil() {
	o.TxnLockHeight.Set(nil)
}

// UnsetTxnLockHeight ensures that no value is present for TxnLockHeight, not even an explicit nil
func (o *BabylonStakingDepositAllOf1) UnsetTxnLockHeight() {
	o.TxnLockHeight.Unset()
}

// GetValue returns the Value field value
func (o *BabylonStakingDepositAllOf1) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingDepositAllOf1) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BabylonStakingDepositAllOf1) SetValue(v int64) {
	o.Value = v
}

func (o BabylonStakingDepositAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBabylonStakingDepositAllOf1 struct {
	value *BabylonStakingDepositAllOf1
	isSet bool
}

func (v NullableBabylonStakingDepositAllOf1) Get() *BabylonStakingDepositAllOf1 {
	return v.value
}

func (v *NullableBabylonStakingDepositAllOf1) Set(val *BabylonStakingDepositAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingDepositAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingDepositAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingDepositAllOf1(val *BabylonStakingDepositAllOf1) *NullableBabylonStakingDepositAllOf1 {
	return &NullableBabylonStakingDepositAllOf1{value: val, isSet: true}
}

func (v NullableBabylonStakingDepositAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingDepositAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


