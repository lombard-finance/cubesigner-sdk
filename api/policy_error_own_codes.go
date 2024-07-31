package api

import (
	"encoding/json"
	"fmt"
)

// PolicyErrorOwnCodes the model 'PolicyErrorOwnCodes'
type PolicyErrorOwnCodes string

// List of PolicyErrorOwnCodes
const (
	BTC_TX_DISALLOWED_OUTPUTS            PolicyErrorOwnCodes = "BtcTxDisallowedOutputs"
	EVM_TX_RECEIVER_MISMATCH             PolicyErrorOwnCodes = "EvmTxReceiverMismatch"
	EVM_TX_SENDER_MISMATCH               PolicyErrorOwnCodes = "EvmTxSenderMismatch"
	EVM_TX_EXCEEDED_VALUE                PolicyErrorOwnCodes = "EvmTxExceededValue"
	EVM_TX_VALUE_UNDEFINED               PolicyErrorOwnCodes = "EvmTxValueUndefined"
	EVM_TX_EXCEEDED_GAS_COST             PolicyErrorOwnCodes = "EvmTxExceededGasCost"
	EVM_TX_GAS_COST_UNDEFINED            PolicyErrorOwnCodes = "EvmTxGasCostUndefined"
	ERC20_DATA_INVALID                   PolicyErrorOwnCodes = "Erc20DataInvalid"
	ERC20_CONTRACT_ADDRESS_UNDEFINED     PolicyErrorOwnCodes = "Erc20ContractAddressUndefined"
	ERC20_CONTRACT_CHAIN_ID_UNDEFINED    PolicyErrorOwnCodes = "Erc20ContractChainIdUndefined"
	ERC20_NOT_IN_CONTRACT_ALLOWLIST      PolicyErrorOwnCodes = "Erc20NotInContractAllowlist"
	ERC20_EXCEEDED_TRANSFER_LIMIT        PolicyErrorOwnCodes = "Erc20ExceededTransferLimit"
	ERC20_RECEIVER_MISMATCH              PolicyErrorOwnCodes = "Erc20ReceiverMismatch"
	ERC20_EXCEEDED_APPROVE_LIMIT         PolicyErrorOwnCodes = "Erc20ExceededApproveLimit"
	ERC20_SPENDER_MISMATCH               PolicyErrorOwnCodes = "Erc20SpenderMismatch"
	POLICY_DISJUNCTION_ERROR             PolicyErrorOwnCodes = "PolicyDisjunctionError"
	POLICY_NEGATION_ERROR                PolicyErrorOwnCodes = "PolicyNegationError"
	ETH2_EXCEEDED_MAX_UNSTAKE            PolicyErrorOwnCodes = "Eth2ExceededMaxUnstake"
	ETH2_CONCURRENT_UNSTAKING            PolicyErrorOwnCodes = "Eth2ConcurrentUnstaking"
	NOT_IN_IPV4_ALLOWLIST                PolicyErrorOwnCodes = "NotInIpv4Allowlist"
	NOT_IN_ORIGIN_ALLOWLIST              PolicyErrorOwnCodes = "NotInOriginAllowlist"
	INVALID_SOURCE_IP                    PolicyErrorOwnCodes = "InvalidSourceIp"
	RAW_SIGNING_NOT_ALLOWED              PolicyErrorOwnCodes = "RawSigningNotAllowed"
	EIP712_SIGNING_NOT_ALLOWED           PolicyErrorOwnCodes = "Eip712SigningNotAllowed"
	OIDC_SOURCE_NOT_ALLOWED              PolicyErrorOwnCodes = "OidcSourceNotAllowed"
	NO_OIDC_AUTH_SOURCES_DEFINED         PolicyErrorOwnCodes = "NoOidcAuthSourcesDefined"
	ADD_KEY_TO_ROLE_DISALLOWED           PolicyErrorOwnCodes = "AddKeyToRoleDisallowed"
	KEYS_ALREADY_IN_ROLE                 PolicyErrorOwnCodes = "KeysAlreadyInRole"
	KEY_IN_MULTIPLE_ROLES                PolicyErrorOwnCodes = "KeyInMultipleRoles"
	KEY_ACCESS_ERROR                     PolicyErrorOwnCodes = "KeyAccessError"
	EIP191_SIGNING_NOT_ALLOWED           PolicyErrorOwnCodes = "Eip191SigningNotAllowed"
	TAPROOT_SIGNING                      PolicyErrorOwnCodes = "TaprootSigning"
	TIME_LOCKED                          PolicyErrorOwnCodes = "TimeLocked"
	BABYLON_STAKING_NETWORK              PolicyErrorOwnCodes = "BabylonStakingNetwork"
	BABYLON_STAKING_PARAMS_VERSION       PolicyErrorOwnCodes = "BabylonStakingParamsVersion"
	BABYLON_STAKING_EXPLICIT_PARAMS      PolicyErrorOwnCodes = "BabylonStakingExplicitParams"
	BABYLON_STAKING_STAKER_PK            PolicyErrorOwnCodes = "BabylonStakingStakerPk"
	BABYLON_STAKING_FINALITY_PROVIDER_PK PolicyErrorOwnCodes = "BabylonStakingFinalityProviderPk"
	BABYLON_STAKING_LOCK_TIME            PolicyErrorOwnCodes = "BabylonStakingLockTime"
	BABYLON_STAKING_VALUE                PolicyErrorOwnCodes = "BabylonStakingValue"
	BABYLON_STAKING_CHANGE_ADDRESS       PolicyErrorOwnCodes = "BabylonStakingChangeAddress"
	BABYLON_STAKING_FEE                  PolicyErrorOwnCodes = "BabylonStakingFee"
	BABYLON_STAKING_WITHDRAWAL_ADDRESS   PolicyErrorOwnCodes = "BabylonStakingWithdrawalAddress"
)

// All allowed values of PolicyErrorOwnCodes enum
var AllowedPolicyErrorOwnCodesEnumValues = []PolicyErrorOwnCodes{
	"BtcTxDisallowedOutputs",
	"EvmTxReceiverMismatch",
	"EvmTxSenderMismatch",
	"EvmTxExceededValue",
	"EvmTxValueUndefined",
	"EvmTxExceededGasCost",
	"EvmTxGasCostUndefined",
	"Erc20DataInvalid",
	"Erc20ContractAddressUndefined",
	"Erc20ContractChainIdUndefined",
	"Erc20NotInContractAllowlist",
	"Erc20ExceededTransferLimit",
	"Erc20ReceiverMismatch",
	"Erc20ExceededApproveLimit",
	"Erc20SpenderMismatch",
	"PolicyDisjunctionError",
	"PolicyNegationError",
	"Eth2ExceededMaxUnstake",
	"Eth2ConcurrentUnstaking",
	"NotInIpv4Allowlist",
	"NotInOriginAllowlist",
	"InvalidSourceIp",
	"RawSigningNotAllowed",
	"Eip712SigningNotAllowed",
	"OidcSourceNotAllowed",
	"NoOidcAuthSourcesDefined",
	"AddKeyToRoleDisallowed",
	"KeysAlreadyInRole",
	"KeyInMultipleRoles",
	"KeyAccessError",
	"Eip191SigningNotAllowed",
	"TaprootSigning",
	"TimeLocked",
	"BabylonStakingNetwork",
	"BabylonStakingParamsVersion",
	"BabylonStakingExplicitParams",
	"BabylonStakingStakerPk",
	"BabylonStakingFinalityProviderPk",
	"BabylonStakingLockTime",
	"BabylonStakingValue",
	"BabylonStakingChangeAddress",
	"BabylonStakingFee",
	"BabylonStakingWithdrawalAddress",
}

func (v *PolicyErrorOwnCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyErrorOwnCodes(value)
	for _, existing := range AllowedPolicyErrorOwnCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyErrorOwnCodes", value)
}

// NewPolicyErrorOwnCodesFromValue returns a pointer to a valid PolicyErrorOwnCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyErrorOwnCodesFromValue(v string) (*PolicyErrorOwnCodes, error) {
	ev := PolicyErrorOwnCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyErrorOwnCodes: valid values are %v", v, AllowedPolicyErrorOwnCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyErrorOwnCodes) IsValid() bool {
	for _, existing := range AllowedPolicyErrorOwnCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyErrorOwnCodes value
func (v PolicyErrorOwnCodes) Ptr() *PolicyErrorOwnCodes {
	return &v
}

type NullablePolicyErrorOwnCodes struct {
	value *PolicyErrorOwnCodes
	isSet bool
}

func (v NullablePolicyErrorOwnCodes) Get() *PolicyErrorOwnCodes {
	return v.value
}

func (v *NullablePolicyErrorOwnCodes) Set(val *PolicyErrorOwnCodes) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyErrorOwnCodes) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyErrorOwnCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyErrorOwnCodes(val *PolicyErrorOwnCodes) *NullablePolicyErrorOwnCodes {
	return &NullablePolicyErrorOwnCodes{value: val, isSet: true}
}

func (v NullablePolicyErrorOwnCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyErrorOwnCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
