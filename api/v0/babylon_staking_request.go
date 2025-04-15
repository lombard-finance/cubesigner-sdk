package v0

import (
	"encoding/json"
	"fmt"

	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonStakingRequest - The actions possible via the Babylon Staking endpoint
type BabylonStakingRequest struct {
	BabylonStakingDeposit     *BabylonStakingDeposit
	BabylonStakingEarlyUnbond *BabylonStakingEarlyUnbond
	BabylonStakingWithdrawal  *BabylonStakingWithdrawal

	Action api.BabylonStakingAction `json:"action"`
}

// BabylonStakingDepositAsBabylonStakingRequest is a convenience function that returns BabylonStakingDeposit wrapped in BabylonStakingRequest
func BabylonStakingDepositAsBabylonStakingRequest(v *BabylonStakingDeposit) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingDeposit: v,
		Action:                api.DepositAction,
	}
}

// BabylonStakingEarlyUnbondAsBabylonStakingRequest is a convenience function that returns BabylonStakingEarlyUnbond wrapped in BabylonStakingRequest
func BabylonStakingEarlyUnbondAsBabylonStakingRequest(v *BabylonStakingEarlyUnbond) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingEarlyUnbond: v,
		Action:                    api.EarlyUnbondAction,
	}
}

// BabylonStakingWithdrawTimelockAsBabylonStakingRequest is a convenience function that returns BabylonStakingWithdrawal wrapped in BabylonStakingRequest
func BabylonStakingWithdrawTimelockAsBabylonStakingRequest(v *BabylonStakingWithdrawal) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingWithdrawal: v,
		Action:                   api.WithdrawTimelockAction,
	}
}

// BabylonStakingWithdrawEarlyUnbondActionAsBabylonStakingRequest is a convenience function that returns BabylonStakingWithdrawal wrapped in BabylonStakingRequest
func BabylonStakingWithdrawEarlyUnbondActionAsBabylonStakingRequest(v *BabylonStakingWithdrawal) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingWithdrawal: v,
		Action:                   api.WithdrawEarlyUnbondAction,
	}
}

// BabylonStakingSlashDepositAsBabylonStakingRequest is a convenience function that returns BabylonStakingSlashDeposit wrapped in BabylonStakingRequest
func BabylonStakingSlashDepositAsBabylonStakingRequest(v *BabylonStakingEarlyUnbond) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingEarlyUnbond: v,
		Action:                    api.SlashDepositAction,
	}
}

// BabylonStakingSlashEarlyUnbondAsBabylonStakingRequest is a convenience function that returns BabylonStakingSlashEarlyUnbond wrapped in BabylonStakingRequest
func BabylonStakingSlashEarlyUnbondAsBabylonStakingRequest(v *BabylonStakingEarlyUnbond) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingEarlyUnbond: v,
		Action:                    api.SlashEarlyUnbondAction,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BabylonStakingRequest) UnmarshalJSON(data []byte) error {
	// Create a temporary struct to parse the action field
	type ActionOnly struct {
		Action api.BabylonStakingAction `json:"action"`
	}

	var actionObj ActionOnly
	if err := json.Unmarshal(data, &actionObj); err != nil {
		return fmt.Errorf("failed to unmarshal BabylonStakingRequest action: %w", err)
	}

	// Set the action field in the destination
	dst.Action = actionObj.Action

	// Based on the action type, unmarshal into the appropriate struct
	switch actionObj.Action {
	case api.DepositAction:
		var deposit BabylonStakingDeposit
		if err := json.Unmarshal(data, &deposit); err != nil {
			return fmt.Errorf("failed to unmarshal BabylonStakingDeposit for BabylonStakingAction %s: %w", actionObj.Action, err)
		}
		dst.BabylonStakingDeposit = &deposit

	case api.EarlyUnbondAction, api.SlashDepositAction, api.SlashEarlyUnbondAction:
		var earlyUnbond BabylonStakingEarlyUnbond
		if err := json.Unmarshal(data, &earlyUnbond); err != nil {
			return fmt.Errorf("failed to unmarshal BabylonStakingEarlyUnbond for BabylonStakingAction %s: %w", actionObj.Action, err)
		}
		dst.BabylonStakingEarlyUnbond = &earlyUnbond

	case api.WithdrawTimelockAction, api.WithdrawEarlyUnbondAction:
		var withdrawal BabylonStakingWithdrawal
		if err := json.Unmarshal(data, &withdrawal); err != nil {
			return fmt.Errorf("failed to unmarshal BabylonStakingWithdrawal for BabylonStakingAction %s: %w", actionObj.Action, err)
		}
		dst.BabylonStakingWithdrawal = &withdrawal

	default:
		return fmt.Errorf("unknown BabylonStakingAction: %s", actionObj.Action)
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BabylonStakingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if src.BabylonStakingDeposit != nil {
		toSerialize = src.BabylonStakingDeposit.Serialize()
	}

	if src.BabylonStakingEarlyUnbond != nil {
		toSerialize = src.BabylonStakingEarlyUnbond.Serialize()
	}

	if src.BabylonStakingWithdrawal != nil {
		toSerialize = src.BabylonStakingWithdrawal.Serialize()
	}

	toSerialize["action"] = src.Action

	return json.Marshal(toSerialize)
}

// Get the actual instance
func (obj *BabylonStakingRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BabylonStakingDeposit != nil {
		return obj.BabylonStakingDeposit
	}

	if obj.BabylonStakingEarlyUnbond != nil {
		return obj.BabylonStakingEarlyUnbond
	}

	if obj.BabylonStakingWithdrawal != nil {
		return obj.BabylonStakingWithdrawal
	}

	// all schemas are nil
	return nil
}
