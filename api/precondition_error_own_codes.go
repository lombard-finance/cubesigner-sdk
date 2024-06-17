package api

import (
	"encoding/json"
	"fmt"
)

// PreconditionErrorOwnCodes the model 'PreconditionErrorOwnCodes'
type PreconditionErrorOwnCodes string

// List of PreconditionErrorOwnCodes
const (
	ETH2_PROPOSER_SLOT_TOO_LOW                               PreconditionErrorOwnCodes = "Eth2ProposerSlotTooLow"
	ETH2_ATTESTATION_SOURCE_EPOCH_TOO_LOW                    PreconditionErrorOwnCodes = "Eth2AttestationSourceEpochTooLow"
	ETH2_ATTESTATION_TARGET_EPOCH_TOO_LOW                    PreconditionErrorOwnCodes = "Eth2AttestationTargetEpochTooLow"
	ETH2_CONCURRENT_BLOCK_SIGNING                            PreconditionErrorOwnCodes = "Eth2ConcurrentBlockSigning"
	ETH2_CONCURRENT_ATTESTATION_SIGNING                      PreconditionErrorOwnCodes = "Eth2ConcurrentAttestationSigning"
	ETH2_MULTI_DEPOSIT_TO_NON_GENERATED_KEY                  PreconditionErrorOwnCodes = "Eth2MultiDepositToNonGeneratedKey"
	ETH2_MULTI_DEPOSIT_UNKNOWN_INITIAL_DEPOSIT               PreconditionErrorOwnCodes = "Eth2MultiDepositUnknownInitialDeposit"
	ETH2_MULTI_DEPOSIT_WITHDRAWAL_ADDRESS_MISMATCH           PreconditionErrorOwnCodes = "Eth2MultiDepositWithdrawalAddressMismatch"
	EVM_CONCURRENT_SIGNING_WHEN_TIME_LIMIT_POLICY_IS_DEFINED PreconditionErrorOwnCodes = "EvmConcurrentSigningWhenTimeLimitPolicyIsDefined"
	BABYLON_EOTS_CONCURRENT_SIGNING                          PreconditionErrorOwnCodes = "BabylonEotsConcurrentSigning"
)

// All allowed values of PreconditionErrorOwnCodes enum
var AllowedPreconditionErrorOwnCodesEnumValues = []PreconditionErrorOwnCodes{
	"Eth2ProposerSlotTooLow",
	"Eth2AttestationSourceEpochTooLow",
	"Eth2AttestationTargetEpochTooLow",
	"Eth2ConcurrentBlockSigning",
	"Eth2ConcurrentAttestationSigning",
	"Eth2MultiDepositToNonGeneratedKey",
	"Eth2MultiDepositUnknownInitialDeposit",
	"Eth2MultiDepositWithdrawalAddressMismatch",
	"EvmConcurrentSigningWhenTimeLimitPolicyIsDefined",
	"BabylonEotsConcurrentSigning",
}

func (v *PreconditionErrorOwnCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PreconditionErrorOwnCodes(value)
	for _, existing := range AllowedPreconditionErrorOwnCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PreconditionErrorOwnCodes", value)
}

// NewPreconditionErrorOwnCodesFromValue returns a pointer to a valid PreconditionErrorOwnCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPreconditionErrorOwnCodesFromValue(v string) (*PreconditionErrorOwnCodes, error) {
	ev := PreconditionErrorOwnCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PreconditionErrorOwnCodes: valid values are %v", v, AllowedPreconditionErrorOwnCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PreconditionErrorOwnCodes) IsValid() bool {
	for _, existing := range AllowedPreconditionErrorOwnCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PreconditionErrorOwnCodes value
func (v PreconditionErrorOwnCodes) Ptr() *PreconditionErrorOwnCodes {
	return &v
}

type NullablePreconditionErrorOwnCodes struct {
	value *PreconditionErrorOwnCodes
	isSet bool
}

func (v NullablePreconditionErrorOwnCodes) Get() *PreconditionErrorOwnCodes {
	return v.value
}

func (v *NullablePreconditionErrorOwnCodes) Set(val *PreconditionErrorOwnCodes) {
	v.value = val
	v.isSet = true
}

func (v NullablePreconditionErrorOwnCodes) IsSet() bool {
	return v.isSet
}

func (v *NullablePreconditionErrorOwnCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreconditionErrorOwnCodes(val *PreconditionErrorOwnCodes) *NullablePreconditionErrorOwnCodes {
	return &NullablePreconditionErrorOwnCodes{value: val, isSet: true}
}

func (v NullablePreconditionErrorOwnCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreconditionErrorOwnCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
