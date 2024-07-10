package v0

import (
	"encoding/json"
	"fmt"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonStakingRequest - The actions possible via the Babylon Staking endpoint
type BabylonStakingRequest struct {
	*BabylonStakingDeposit
	*BabylonStakingEarlyUnbond
	*BabylonStakingWithdrawal

	Action api.BabylonStakingAction `json:"action"`
}

// BabylonStakingDepositAsBabylonStakingRequest is a convenience function that returns BabylonStakingDeposit wrapped in BabylonStakingRequest
func BabylonStakingDepositAsBabylonStakingRequest(v *BabylonStakingDeposit) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingDeposit: v,
	}
}

// BabylonStakingEarlyUnbondAsBabylonStakingRequest is a convenience function that returns BabylonStakingEarlyUnbond wrapped in BabylonStakingRequest
func BabylonStakingEarlyUnbondAsBabylonStakingRequest(v *BabylonStakingEarlyUnbond) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingEarlyUnbond: v,
	}
}

// BabylonStakingWithdrawalAsBabylonStakingRequest is a convenience function that returns BabylonStakingWithdrawal wrapped in BabylonStakingRequest
func BabylonStakingWithdrawalAsBabylonStakingRequest(v *BabylonStakingWithdrawal) BabylonStakingRequest {
	return BabylonStakingRequest{
		BabylonStakingWithdrawal: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BabylonStakingRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BabylonStakingDeposit
	err = api.NewStrictDecoder(data).Decode(&dst.BabylonStakingDeposit)
	if err == nil {
		jsonBabylonStakingDeposit, _ := json.Marshal(dst.BabylonStakingDeposit)
		if string(jsonBabylonStakingDeposit) == "{}" { // empty struct
			dst.BabylonStakingDeposit = nil
		} else {
			match++
		}
	} else {
		dst.BabylonStakingDeposit = nil
	}

	// try to unmarshal data into BabylonStakingEarlyUnbond
	err = api.NewStrictDecoder(data).Decode(&dst.BabylonStakingEarlyUnbond)
	if err == nil {
		jsonBabylonStakingEarlyUnbond, _ := json.Marshal(dst.BabylonStakingEarlyUnbond)
		if string(jsonBabylonStakingEarlyUnbond) == "{}" { // empty struct
			dst.BabylonStakingEarlyUnbond = nil
		} else {
			match++
		}
	} else {
		dst.BabylonStakingEarlyUnbond = nil
	}

	// try to unmarshal data into BabylonStakingWithdrawal
	err = api.NewStrictDecoder(data).Decode(&dst.BabylonStakingWithdrawal)
	if err == nil {
		jsonBabylonStakingWithdrawal, _ := json.Marshal(dst.BabylonStakingWithdrawal)
		if string(jsonBabylonStakingWithdrawal) == "{}" { // empty struct
			dst.BabylonStakingWithdrawal = nil
		} else {
			match++
		}
	} else {
		dst.BabylonStakingWithdrawal = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BabylonStakingDeposit = nil
		dst.BabylonStakingEarlyUnbond = nil
		dst.BabylonStakingWithdrawal = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(BabylonStakingRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(BabylonStakingRequest)")
	}
}

//// Marshal data from the first non-nil pointers in the struct to JSON
//func (src BabylonStakingRequest) MarshalJSON() ([]byte, error) {
//	if src.BabylonStakingDeposit != nil {
//		return json.Marshal(&src.BabylonStakingDeposit)
//	}
//
//	if src.BabylonStakingEarlyUnbond != nil {
//		return json.Marshal(&src.BabylonStakingEarlyUnbond)
//	}
//
//	if src.BabylonStakingWithdrawal != nil {
//		return json.Marshal(&src.BabylonStakingWithdrawal)
//	}
//
//	return nil, nil // no data in oneOf schemas
//}

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
