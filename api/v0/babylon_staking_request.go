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

	// Unmarshal into a temporary map to analyze the fields
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return fmt.Errorf("failed to unmarshal into map: %v", err)
	}

	// Check if the data seems to match BabylonStakingWithdrawal by looking for its unique fields
	hasRecipient := false
	if _, hasRecipient := tempMap["recipient"]; hasRecipient {
		err = api.NewStrictDecoder(data).Decode(&dst.BabylonStakingWithdrawal)
		if err == nil && !isEmptyStruct(dst.BabylonStakingWithdrawal) {
			match++
		} else {
			dst.BabylonStakingWithdrawal = nil
		}
	}

	// Check if the data seems to match BabylonStakingEarlyUnbond by looking for its unique fields
	hasTxid := false
	if _, hasTxid := tempMap["txid"]; hasTxid && !hasRecipient {
		err = api.NewStrictDecoder(data).Decode(&dst.BabylonStakingEarlyUnbond)
		if err == nil && !isEmptyStruct(dst.BabylonStakingEarlyUnbond) {
			match++
		} else {
			dst.BabylonStakingEarlyUnbond = nil
		}
	}

	// Check if the data seems to match BabylonStakingDeposit (you can add unique checks for deposit as well)
	if !hasTxid && !hasRecipient {
		err = api.NewStrictDecoder(data).Decode(&dst.BabylonStakingDeposit)
		if err == nil && !isEmptyStruct(dst.BabylonStakingDeposit) {
			match++
		} else {
			dst.BabylonStakingDeposit = nil
		}
	}

	if match > 1 {
		// more than one match, reset all
		dst.BabylonStakingDeposit = nil
		dst.BabylonStakingEarlyUnbond = nil
		dst.BabylonStakingWithdrawal = nil
		return fmt.Errorf("data matches more than one schema in oneOf(BabylonStakingRequest)")
	} else if match == 1 {
		return nil // exactly one match
	}

	// No match
	return fmt.Errorf("data failed to match schemas in oneOf(BabylonStakingRequest)")
}

func isEmptyStruct(v interface{}) bool {
	jsonData, _ := json.Marshal(v)
	return string(jsonData) == "{}"
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
