package v0

// BabylonStakingParams struct for BabylonStakingParams
type BabylonStakingParams struct {
	// The block height at which these params will enter use.
	ActivationHeight int64 `json:"activation_height"`
	// The block height above which staking is disabled.
	CapHeight *int64 `json:"cap_height,omitempty"`
	// The number of confirmations before a staking txn is considered finalized by the covenant committee.
	ConfirmationDepth int32 `json:"confirmation_depth"`
	// The public keys of the covenant signers.
	CovenantPks []string `json:"covenant_pks"`
	// The quorum for covenant signer.
	CovenantQuorum int `json:"covenant_quorum"`
	// The max amount that can be staked in a single txn.
	MaxStakingAmount int64 `json:"max_staking_amount"`
	// The max timelock for staking.
	MaxStakingTime int32 `json:"max_staking_time"`
	// The min amount that must be staked.
	MinStakingAmount int64 `json:"min_staking_amount"`
	// The min timelock for staking.
	MinStakingTime int32 `json:"min_staking_time"`
	// The max total amount staked.
	StakingCap *int64 `json:"staking_cap,omitempty"`
	// The "Magic bytes" tag for staking metadata.
	Tag string `json:"tag"`
	// The fee that must be spent as part of the unbonding txn.
	UnbondingFee int64 `json:"unbonding_fee"`
	// The min timelock for an unbonding script.
	UnbondingTime int32 `json:"unbonding_time"`
	// The parameter version.
	Version int32 `json:"version"`
}

func NewBabylonStakingParams(activationHeight int64, confirmationDepth int32, covenantPks []string, covenantQuorum int, maxStakingAmount int64, maxStakingTime int32, minStakingAmount int64, minStakingTime int32, tag string, unbondingFee int64, unbondingTime int32, version int32) *BabylonStakingParams {
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

func NewBabylonStakingParamsWithDefaults() *BabylonStakingParams {
	this := BabylonStakingParams{}
	return &this
}

func (o *BabylonStakingParams) GetActivationHeight() int64 {
	if o == nil {
		return 0
	}
	return o.ActivationHeight
}

func (o *BabylonStakingParams) GetActivationHeightOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.ActivationHeight, true
}

func (o *BabylonStakingParams) SetActivationHeight(v int64) {
	o.ActivationHeight = v
}

func (o *BabylonStakingParams) GetCapHeight() *int64 {
	if o == nil {
		return nil
	}
	return o.CapHeight
}

func (o *BabylonStakingParams) GetCapHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CapHeight, true
}

func (o *BabylonStakingParams) SetCapHeight(v *int64) {
	o.CapHeight = v
}

func (o *BabylonStakingParams) GetConfirmationDepth() int32 {
	if o == nil {
		return 0
	}
	return o.ConfirmationDepth
}

func (o *BabylonStakingParams) GetConfirmationDepthOk() (int32, bool) {
	if o == nil {
		return 0, false
	}
	return o.ConfirmationDepth, true
}

func (o *BabylonStakingParams) SetConfirmationDepth(v int32) {
	o.ConfirmationDepth = v
}

func (o *BabylonStakingParams) GetCovenantPks() []string {
	if o == nil {
		return nil
	}
	return o.CovenantPks
}

func (o *BabylonStakingParams) GetCovenantPksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CovenantPks, true
}

func (o *BabylonStakingParams) SetCovenantPks(v []string) {
	o.CovenantPks = v
}

func (o *BabylonStakingParams) GetCovenantQuorum() int {
	if o == nil {
		return 0
	}
	return o.CovenantQuorum
}

func (o *BabylonStakingParams) GetCovenantQuorumOk() (int, bool) {
	if o == nil {
		return 0, false
	}
	return o.CovenantQuorum, true
}

func (o *BabylonStakingParams) SetCovenantQuorum(v int) {
	o.CovenantQuorum = v
}

func (o *BabylonStakingParams) GetMaxStakingAmount() int64 {
	if o == nil {
		return 0
	}
	return o.MaxStakingAmount
}

func (o *BabylonStakingParams) GetMaxStakingAmountOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.MaxStakingAmount, true
}

func (o *BabylonStakingParams) SetMaxStakingAmount(v int64) {
	o.MaxStakingAmount = v
}

func (o *BabylonStakingParams) GetMaxStakingTime() int32 {
	if o == nil {
		return 0
	}
	return o.MaxStakingTime
}

func (o *BabylonStakingParams) GetMaxStakingTimeOk() (int32, bool) {
	if o == nil {
		return 0, false
	}
	return o.MaxStakingTime, true
}

func (o *BabylonStakingParams) SetMaxStakingTime(v int32) {
	o.MaxStakingTime = v
}

func (o *BabylonStakingParams) GetMinStakingAmount() int64 {
	if o == nil {
		return 0
	}
	return o.MinStakingAmount
}

func (o *BabylonStakingParams) GetMinStakingAmountOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.MinStakingAmount, true
}

func (o *BabylonStakingParams) SetMinStakingAmount(v int64) {
	o.MinStakingAmount = v
}

func (o *BabylonStakingParams) GetMinStakingTime() int32 {
	if o == nil {
		return 0
	}
	return o.MinStakingTime
}

func (o *BabylonStakingParams) GetMinStakingTimeOk() (int32, bool) {
	if o == nil {
		return 0, false
	}
	return o.MinStakingTime, true
}

func (o *BabylonStakingParams) SetMinStakingTime(v int32) {
	o.MinStakingTime = v
}

func (o *BabylonStakingParams) GetStakingCap() *int64 {
	if o == nil {
		return nil
	}
	return o.StakingCap
}

func (o *BabylonStakingParams) GetStakingCapOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakingCap, true
}

func (o *BabylonStakingParams) SetStakingCap(v *int64) {
	o.StakingCap = v
}

func (o *BabylonStakingParams) GetTag() string {
	if o == nil {
		return ""
	}
	return o.Tag
}

func (o *BabylonStakingParams) GetTagOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.Tag, true
}

func (o *BabylonStakingParams) SetTag(v string) {
	o.Tag = v
}

func (o *BabylonStakingParams) GetUnbondingFee() int64 {
	if o == nil {
		return 0
	}
	return o.UnbondingFee
}

func (o *BabylonStakingParams) GetUnbondingFeeOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.UnbondingFee, true
}

func (o *BabylonStakingParams) SetUnbondingFee(v int64) {
	o.UnbondingFee = v
}

func (o *BabylonStakingParams) GetUnbondingTime() int32 {
	if o == nil {
		return 0
	}
	return o.UnbondingTime
}

func (o *BabylonStakingParams) GetUnbondingTimeOk() (int32, bool) {
	if o == nil {
		return 0, false
	}
	return o.UnbondingTime, true
}

func (o *BabylonStakingParams) SetUnbondingTime(v int32) {
	o.UnbondingTime = v
}

func (o *BabylonStakingParams) GetVersion() int32 {
	if o == nil {
		return 0
	}
	return o.Version
}

func (o *BabylonStakingParams) GetVersionOk() (int32, bool) {
	if o == nil {
		return 0, false
	}
	return o.Version, true
}

func (o *BabylonStakingParams) SetVersion(v int32) {
	o.Version = v
}
