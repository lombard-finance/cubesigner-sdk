package v0

// BabylonStakingRequest struct for BabylonStakingRequest
type BabylonStakingRequest struct {
	// The action type.
	Action string `json:"action"`
	// The deposit data.
	Deposit *BabylonStakingDeposit `json:"deposit,omitempty"`
	// The early unbond data.
	EarlyUnbond *BabylonStakingEarlyUnbond `json:"early_unbond,omitempty"`
	// The withdrawal data.
	Withdrawal *BabylonStakingWithdrawal `json:"withdrawal,omitempty"`
}

func NewBabylonStakingDepositRequest(deposit *BabylonStakingDeposit) *BabylonStakingRequest {
	this := BabylonStakingRequest{}
	this.Action = "deposit"
	this.Deposit = deposit
	return &this
}

func NewBabylonStakingEarlyUnbondRequest(earlyUnbond *BabylonStakingEarlyUnbond) *BabylonStakingRequest {
	this := BabylonStakingRequest{}
	this.Action = "early_unbond"
	this.EarlyUnbond = earlyUnbond
	return &this
}

func NewBabylonStakingWithdrawalTimelockRequest(withdrawal *BabylonStakingWithdrawal) *BabylonStakingRequest {
	this := BabylonStakingRequest{}
	this.Action = "withdraw_timelock"
	this.Withdrawal = withdrawal
	return &this
}

func NewBabylonStakingWithdrawalEarlyUnbondRequest(withdrawal *BabylonStakingWithdrawal) *BabylonStakingRequest {
	this := BabylonStakingRequest{}
	this.Action = "withdraw_early_unbond"
	this.Withdrawal = withdrawal
	return &this
}

func NewBabylonStakingRequestWithDefaults() *BabylonStakingRequest {
	this := BabylonStakingRequest{}
	return &this
}

func (o *BabylonStakingRequest) GetAction() string {
	if o == nil {
		return ""
	}
	return o.Action
}

func (o *BabylonStakingRequest) GetActionOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.Action, true
}

func (o *BabylonStakingRequest) SetAction(v string) {
	o.Action = v
}

func (o *BabylonStakingRequest) GetDeposit() *BabylonStakingDeposit {
	if o == nil {
		return nil
	}
	return o.Deposit
}

func (o *BabylonStakingRequest) GetDepositOk() (*BabylonStakingDeposit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deposit, true
}

func (o *BabylonStakingRequest) SetDeposit(v *BabylonStakingDeposit) {
	o.Deposit = v
}

func (o *BabylonStakingRequest) GetEarlyUnbond() *BabylonStakingEarlyUnbond {
	if o == nil {
		return nil
	}
	return o.EarlyUnbond
}

func (o *BabylonStakingRequest) GetEarlyUnbondOk() (*BabylonStakingEarlyUnbond, bool) {
	if o == nil {
		return nil, false
	}
	return o.EarlyUnbond, true
}

func (o *BabylonStakingRequest) SetEarlyUnbond(v *BabylonStakingEarlyUnbond) {
	o.EarlyUnbond = v
}

func (o *BabylonStakingRequest) GetWithdrawal() *BabylonStakingWithdrawal {
	if o == nil {
		return nil
	}
	return o.Withdrawal
}

func (o *BabylonStakingRequest) GetWithdrawalOk() (*BabylonStakingWithdrawal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Withdrawal, true
}

func (o *BabylonStakingRequest) SetWithdrawal(v *BabylonStakingWithdrawal) {
	o.Withdrawal = v
}
