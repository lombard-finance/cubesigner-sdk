package api

// BabylonStakingAction represents the type of action when signing a Babylon staking transaction.
type BabylonStakingAction string

const (
	DepositAction             BabylonStakingAction = "deposit"
	EarlyUnbondAction         BabylonStakingAction = "early_unbond"
	WithdrawTimelockAction    BabylonStakingAction = "withdraw_timelock"
	WithdrawEarlyUnbondAction BabylonStakingAction = "withdraw_early_unbond"
	WithdrawSlashing          BabylonStakingAction = "withdraw_slashing"
	SlashDepositAction        BabylonStakingAction = "slash_deposit"
	SlashEarlyUnbondAction    BabylonStakingAction = "slash_early_unbond"
)
