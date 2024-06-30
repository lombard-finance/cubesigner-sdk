package v0

// BabylonStakingDeposit struct for BabylonStakingDeposit
type BabylonStakingDeposit struct {
	BabylonScriptData
	// The change address, specified as a Bitcoin spend script.
	Change string `json:"change"`
	// The transaction fee value.
	Fee int64 `json:"fee"`
	// Determines whether the fee is a fixed fee in sats or a rate in sats per (estimated) virtual byte of transaction weight (i.e., sats per vb).
	FeeType FeeType `json:"fee_type"`
	// When set, indicates that existing outputs in the PSBT will be ignored (i.e., deleted from the PSBT) when creating the deposit transaction.
	IgnorePSBTOutputs *bool `json:"ignore_psbt_outputs,omitempty"`
	// A hex-serialized PSBT (version 0) containing the transaction inputs and all necessary information for signing (e.g., taproot path and leaf information). This PSBT must not have any transaction outputs; they will be added to the returned PSBT.
	PSBT string `json:"psbt"`
	// An optional lock height (in blocks) for this transaction. The resulting transaction cannot be mined before the specified block height.
	TxnLockHeight *int32 `json:"txn_lock_height,omitempty"`
	// The value to be staked in sats.
	Value int64 `json:"value"`
}

func NewBabylonStakingDeposit(scriptData BabylonScriptData, change string, fee int64, feeType FeeType, psbt string, value int64) *BabylonStakingDeposit {
	this := BabylonStakingDeposit{}
	this.BabylonScriptData = scriptData
	this.Change = change
	this.Fee = fee
	this.FeeType = feeType
	this.PSBT = psbt
	this.Value = value
	return &this
}

func NewBabylonStakingDepositWithDefaults() *BabylonStakingDeposit {
	this := BabylonStakingDeposit{}
	return &this
}

func (o *BabylonStakingDeposit) GetChange() string {
	if o == nil {
		return ""
	}
	return o.Change
}

func (o *BabylonStakingDeposit) GetChangeOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.Change, true
}

func (o *BabylonStakingDeposit) SetChange(v string) {
	o.Change = v
}

func (o *BabylonStakingDeposit) GetFee() int64 {
	if o == nil {
		return 0
	}
	return o.Fee
}

func (o *BabylonStakingDeposit) GetFeeOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.Fee, true
}

func (o *BabylonStakingDeposit) SetFee(v int64) {
	o.Fee = v
}

func (o *BabylonStakingDeposit) GetFeeType() FeeType {
	if o == nil {
		return ""
	}
	return o.FeeType
}

func (o *BabylonStakingDeposit) GetFeeTypeOk() (FeeType, bool) {
	if o == nil {
		return "", false
	}
	return o.FeeType, true
}

func (o *BabylonStakingDeposit) SetFeeType(v FeeType) {
	o.FeeType = v
}

func (o *BabylonStakingDeposit) GetIgnorePSBTOutputs() *bool {
	if o == nil {
		return nil
	}
	return o.IgnorePSBTOutputs
}

func (o *BabylonStakingDeposit) GetIgnorePSBTOutputsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IgnorePSBTOutputs, true
}

func (o *BabylonStakingDeposit) SetIgnorePSBTOutputs(v *bool) {
	o.IgnorePSBTOutputs = v
}

func (o *BabylonStakingDeposit) GetPSBT() string {
	if o == nil {
		return ""
	}
	return o.PSBT
}

func (o *BabylonStakingDeposit) GetPSBTOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.PSBT, true
}

func (o *BabylonStakingDeposit) SetPSBT(v string) {
	o.PSBT = v
}

func (o *BabylonStakingDeposit) GetTxnLockHeight() *int32 {
	if o == nil {
		return nil
	}
	return o.TxnLockHeight
}

func (o *BabylonStakingDeposit) GetTxnLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxnLockHeight, true
}

func (o *BabylonStakingDeposit) SetTxnLockHeight(v *int32) {
	o.TxnLockHeight = v
}

func (o *BabylonStakingDeposit) GetValue() int64 {
	if o == nil {
		return 0
	}
	return o.Value
}

func (o *BabylonStakingDeposit) GetValueOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.Value, true
}

func (o *BabylonStakingDeposit) SetValue(v int64) {
	o.Value = v
}