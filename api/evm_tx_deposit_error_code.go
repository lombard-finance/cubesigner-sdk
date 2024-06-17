package api

import (
	"encoding/json"
	"fmt"
)

// EvmTxDepositErrorCode the model 'EvmTxDepositErrorCode'
type EvmTxDepositErrorCode string

// List of EvmTxDepositErrorCode
const (
	EVM_TX_DEPOSIT_RECEIVER_MISMATCH                 EvmTxDepositErrorCode = "EvmTxDepositReceiverMismatch"
	EVM_TX_DEPOSIT_EMPTY_DATA                        EvmTxDepositErrorCode = "EvmTxDepositEmptyData"
	EVM_TX_DEPOSIT_EMPTY_CHAIN_ID                    EvmTxDepositErrorCode = "EvmTxDepositEmptyChainId"
	EVM_TX_DEPOSIT_EMPTY_RECEIVER                    EvmTxDepositErrorCode = "EvmTxDepositEmptyReceiver"
	EVM_TX_DEPOSIT_UNEXPECTED_VALUE                  EvmTxDepositErrorCode = "EvmTxDepositUnexpectedValue"
	EVM_TX_DEPOSIT_UNEXPECTED_DATA_LENGTH            EvmTxDepositErrorCode = "EvmTxDepositUnexpectedDataLength"
	EVM_TX_DEPOSIT_NO_ABI                            EvmTxDepositErrorCode = "EvmTxDepositNoAbi"
	EVM_TX_DEPOSIT_NO_DEPOSIT_FUNCTION               EvmTxDepositErrorCode = "EvmTxDepositNoDepositFunction"
	EVM_TX_DEPOSIT_UNEXPECTED_FUNCTION_NAME          EvmTxDepositErrorCode = "EvmTxDepositUnexpectedFunctionName"
	EVM_TX_DEPOSIT_UNEXPECTED_VALIDATOR_KEY          EvmTxDepositErrorCode = "EvmTxDepositUnexpectedValidatorKey"
	EVM_TX_DEPOSIT_INVALID_VALIDATOR_KEY             EvmTxDepositErrorCode = "EvmTxDepositInvalidValidatorKey"
	EVM_TX_DEPOSIT_MISSING_DEPOSIT_ARG               EvmTxDepositErrorCode = "EvmTxDepositMissingDepositArg"
	EVM_TX_DEPOSIT_WRONG_DEPOSIT_ARG_TYPE            EvmTxDepositErrorCode = "EvmTxDepositWrongDepositArgType"
	EVM_TX_DEPOSIT_VALIDATOR_KEY_NOT_IN_ROLE         EvmTxDepositErrorCode = "EvmTxDepositValidatorKeyNotInRole"
	EVM_TX_DEPOSIT_UNEXPECTED_WITHDRAWAL_CREDENTIALS EvmTxDepositErrorCode = "EvmTxDepositUnexpectedWithdrawalCredentials"
	EVM_TX_DEPOSIT_UNRESOLVED_ROLE                   EvmTxDepositErrorCode = "EvmTxDepositUnresolvedRole"
	EVM_TX_DEPOSIT_INVALID_DEPOSIT_ENCODING          EvmTxDepositErrorCode = "EvmTxDepositInvalidDepositEncoding"
)

// All allowed values of EvmTxDepositErrorCode enum
var AllowedEvmTxDepositErrorCodeEnumValues = []EvmTxDepositErrorCode{
	"EvmTxDepositReceiverMismatch",
	"EvmTxDepositEmptyData",
	"EvmTxDepositEmptyChainId",
	"EvmTxDepositEmptyReceiver",
	"EvmTxDepositUnexpectedValue",
	"EvmTxDepositUnexpectedDataLength",
	"EvmTxDepositNoAbi",
	"EvmTxDepositNoDepositFunction",
	"EvmTxDepositUnexpectedFunctionName",
	"EvmTxDepositUnexpectedValidatorKey",
	"EvmTxDepositInvalidValidatorKey",
	"EvmTxDepositMissingDepositArg",
	"EvmTxDepositWrongDepositArgType",
	"EvmTxDepositValidatorKeyNotInRole",
	"EvmTxDepositUnexpectedWithdrawalCredentials",
	"EvmTxDepositUnresolvedRole",
	"EvmTxDepositInvalidDepositEncoding",
}

func (v *EvmTxDepositErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EvmTxDepositErrorCode(value)
	for _, existing := range AllowedEvmTxDepositErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EvmTxDepositErrorCode", value)
}

// NewEvmTxDepositErrorCodeFromValue returns a pointer to a valid EvmTxDepositErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEvmTxDepositErrorCodeFromValue(v string) (*EvmTxDepositErrorCode, error) {
	ev := EvmTxDepositErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EvmTxDepositErrorCode: valid values are %v", v, AllowedEvmTxDepositErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EvmTxDepositErrorCode) IsValid() bool {
	for _, existing := range AllowedEvmTxDepositErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EvmTxDepositErrorCode value
func (v EvmTxDepositErrorCode) Ptr() *EvmTxDepositErrorCode {
	return &v
}

type NullableEvmTxDepositErrorCode struct {
	value *EvmTxDepositErrorCode
	isSet bool
}

func (v NullableEvmTxDepositErrorCode) Get() *EvmTxDepositErrorCode {
	return v.value
}

func (v *NullableEvmTxDepositErrorCode) Set(val *EvmTxDepositErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmTxDepositErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmTxDepositErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmTxDepositErrorCode(val *EvmTxDepositErrorCode) *NullableEvmTxDepositErrorCode {
	return &NullableEvmTxDepositErrorCode{value: val, isSet: true}
}

func (v NullableEvmTxDepositErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmTxDepositErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
