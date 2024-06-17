package api

import (
	"encoding/json"
	"fmt"
)

// PolicyErrorCode - struct for PolicyErrorCode
type PolicyErrorCode struct {
	EvmTxDepositErrorCode *EvmTxDepositErrorCode
	PolicyErrorOwnCodes   *PolicyErrorOwnCodes
}

// EvmTxDepositErrorCodeAsPolicyErrorCode is a convenience function that returns EvmTxDepositErrorCode wrapped in PolicyErrorCode
func EvmTxDepositErrorCodeAsPolicyErrorCode(v *EvmTxDepositErrorCode) PolicyErrorCode {
	return PolicyErrorCode{
		EvmTxDepositErrorCode: v,
	}
}

// PolicyErrorOwnCodesAsPolicyErrorCode is a convenience function that returns PolicyErrorOwnCodes wrapped in PolicyErrorCode
func PolicyErrorOwnCodesAsPolicyErrorCode(v *PolicyErrorOwnCodes) PolicyErrorCode {
	return PolicyErrorCode{
		PolicyErrorOwnCodes: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PolicyErrorCode) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EvmTxDepositErrorCode
	err = newStrictDecoder(data).Decode(&dst.EvmTxDepositErrorCode)
	if err == nil {
		jsonEvmTxDepositErrorCode, _ := json.Marshal(dst.EvmTxDepositErrorCode)
		if string(jsonEvmTxDepositErrorCode) == "{}" { // empty struct
			dst.EvmTxDepositErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.EvmTxDepositErrorCode = nil
	}

	// try to unmarshal data into PolicyErrorOwnCodes
	err = newStrictDecoder(data).Decode(&dst.PolicyErrorOwnCodes)
	if err == nil {
		jsonPolicyErrorOwnCodes, _ := json.Marshal(dst.PolicyErrorOwnCodes)
		if string(jsonPolicyErrorOwnCodes) == "{}" { // empty struct
			dst.PolicyErrorOwnCodes = nil
		} else {
			match++
		}
	} else {
		dst.PolicyErrorOwnCodes = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EvmTxDepositErrorCode = nil
		dst.PolicyErrorOwnCodes = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PolicyErrorCode)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PolicyErrorCode)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PolicyErrorCode) MarshalJSON() ([]byte, error) {
	if src.EvmTxDepositErrorCode != nil {
		return json.Marshal(&src.EvmTxDepositErrorCode)
	}

	if src.PolicyErrorOwnCodes != nil {
		return json.Marshal(&src.PolicyErrorOwnCodes)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PolicyErrorCode) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EvmTxDepositErrorCode != nil {
		return obj.EvmTxDepositErrorCode
	}

	if obj.PolicyErrorOwnCodes != nil {
		return obj.PolicyErrorOwnCodes
	}

	// all schemas are nil
	return nil
}

type NullablePolicyErrorCode struct {
	value *PolicyErrorCode
	isSet bool
}

func (v NullablePolicyErrorCode) Get() *PolicyErrorCode {
	return v.value
}

func (v *NullablePolicyErrorCode) Set(val *PolicyErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyErrorCode(val *PolicyErrorCode) *NullablePolicyErrorCode {
	return &NullablePolicyErrorCode{value: val, isSet: true}
}

func (v NullablePolicyErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
