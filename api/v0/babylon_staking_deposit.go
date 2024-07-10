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
	// When set, indicates that existing outputs in the Psbt will be ignored (i.e., deleted from the Psbt) when creating the deposit transaction.
	IgnorePsbtOutputs *bool `json:"ignore_psbt_outputs,omitempty"`
	// A hex-serialized Psbt (version 0) containing the transaction inputs and all necessary information for signing (e.g., taproot path and leaf information). This Psbt must not have any transaction outputs; they will be added to the returned Psbt.
	Psbt string `json:"psbt"`
	// An optional lock height (in blocks) for this transaction. The resulting transaction cannot be mined before the specified block height.
	TxnLockHeight *int32 `json:"txn_lock_height,omitempty"`
	// The value to be staked in sats.
	Value int64 `json:"value"`
	// The action type.
	Action BabylonStakingAction `json:"action"`
}

func NewBabylonStakingDeposit(scriptData BabylonScriptData, change string, fee int64, feeType FeeType, psbt string, value int64, action BabylonStakingAction) *BabylonStakingDeposit {
	this := BabylonStakingDeposit{}
	this.Action = action
	this.BabylonScriptData = scriptData
	this.Change = change
	this.Fee = fee
	this.FeeType = feeType
	this.Psbt = psbt
	this.Value = value
	return &this
}

func NewBabylonStakingDepositWithDefaults() *BabylonStakingDeposit {
	this := BabylonStakingDeposit{}
	return &this
}

func (o *BabylonStakingDeposit) GetAction() BabylonStakingAction {
	if o == nil {
		return ""
	}
	return o.Action
}

func (o *BabylonStakingDeposit) GetActionOk() (BabylonStakingAction, bool) {
	if o == nil {
		return "", false
	}
	return o.Action, true
}

func (o *BabylonStakingDeposit) SetAction(v BabylonStakingAction) {
	o.Action = v
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

func (o *BabylonStakingDeposit) GetIgnorePsbtOutputs() *bool {
	if o == nil {
		return nil
	}
	return o.IgnorePsbtOutputs
}

func (o *BabylonStakingDeposit) GetIgnorePsbtOutputsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IgnorePsbtOutputs, true
}

func (o *BabylonStakingDeposit) SetIgnorePsbtOutputs(v *bool) {
	o.IgnorePsbtOutputs = v
}

func (o *BabylonStakingDeposit) GetPsbt() string {
	if o == nil {
		return ""
	}
	return o.Psbt
}

func (o *BabylonStakingDeposit) GetPsbtOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.Psbt, true
}

func (o *BabylonStakingDeposit) SetPsbt(v string) {
	o.Psbt = v
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
