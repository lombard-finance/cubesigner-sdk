package api

import (
	"encoding/json"
	"fmt"
)

// BadRequestErrorCode the model 'BadRequestErrorCode'
type BadRequestErrorCode string

// List of BadRequestErrorCode
const (
	GENERIC_BAD_REQUEST                                    BadRequestErrorCode = "GenericBadRequest"
	EMAIL_PASSWORD_NOT_FOUND                               BadRequestErrorCode = "EmailPasswordNotFound"
	ONE_TIME_CODE_EXPIRED                                  BadRequestErrorCode = "OneTimeCodeExpired"
	INVALID_BODY                                           BadRequestErrorCode = "InvalidBody"
	TOKEN_REQUEST_ERROR                                    BadRequestErrorCode = "TokenRequestError"
	INVALID_MFA_RECEIPT                                    BadRequestErrorCode = "InvalidMfaReceipt"
	INVALID_MFA_POLICY_COUNT                               BadRequestErrorCode = "InvalidMfaPolicyCount"
	INVALID_MFA_POLICY_NUM_AUTH_FACTORS                    BadRequestErrorCode = "InvalidMfaPolicyNumAuthFactors"
	INVALID_MFA_POLICY_NUM_ALLOWED_APPROVERS               BadRequestErrorCode = "InvalidMfaPolicyNumAllowedApprovers"
	INVALID_MFA_POLICY_GRACE_PERIOD_TOO_LONG               BadRequestErrorCode = "InvalidMfaPolicyGracePeriodTooLong"
	INVALID_MFA_POLICY_REDUNDANT_RULE                      BadRequestErrorCode = "InvalidMfaPolicyRedundantRule"
	INVALID_BABYLON_STAKING_POLICY_PARAMS                  BadRequestErrorCode = "InvalidBabylonStakingPolicyParams"
	INVALID_BTC_TX_RECEIVERS_EMPTY_ALLOWLIST               BadRequestErrorCode = "InvalidBtcTxReceiversEmptyAllowlist"
	INVALID_CREATE_KEY_COUNT                               BadRequestErrorCode = "InvalidCreateKeyCount"
	ORG_INVITE_EXISTING_USER                               BadRequestErrorCode = "OrgInviteExistingUser"
	ORG_USER_ALREADY_EXISTS                                BadRequestErrorCode = "OrgUserAlreadyExists"
	ORG_NAME_TAKEN                                         BadRequestErrorCode = "OrgNameTaken"
	ROLE_NAME_TAKEN                                        BadRequestErrorCode = "RoleNameTaken"
	ADD_KEY_TO_ROLE_COUNT_TOO_HIGH                         BadRequestErrorCode = "AddKeyToRoleCountTooHigh"
	INVALID_KEY_ID                                         BadRequestErrorCode = "InvalidKeyId"
	INVALID_TIME_LOCK_ALREADY_IN_THE_PAST                  BadRequestErrorCode = "InvalidTimeLockAlreadyInThePast"
	INVALID_RESTRICTED_SCOPES                              BadRequestErrorCode = "InvalidRestrictedScopes"
	INVALID_UPDATE                                         BadRequestErrorCode = "InvalidUpdate"
	INVALID_METADATA_LENGTH                                BadRequestErrorCode = "InvalidMetadataLength"
	INVALID_KEY_MATERIAL_ID                                BadRequestErrorCode = "InvalidKeyMaterialId"
	KEY_NOT_FOUND                                          BadRequestErrorCode = "KeyNotFound"
	USER_EXPORT_DERIVED_KEY                                BadRequestErrorCode = "UserExportDerivedKey"
	USER_EXPORT_PUBLIC_KEY_INVALID                         BadRequestErrorCode = "UserExportPublicKeyInvalid"
	UNABLE_TO_ACCESS_SMTP_RELAY                            BadRequestErrorCode = "UnableToAccessSmtpRelay"
	USER_EXPORT_IN_PROGRESS                                BadRequestErrorCode = "UserExportInProgress"
	ROLE_NOT_FOUND                                         BadRequestErrorCode = "RoleNotFound"
	INVALID_ROLE_NAME_OR_ID                                BadRequestErrorCode = "InvalidRoleNameOrId"
	INVALID_MFA_RECEIPT_ORG_ID_MISSING                     BadRequestErrorCode = "InvalidMfaReceiptOrgIdMissing"
	INVALID_MFA_RECEIPT_INVALID_ORG_ID                     BadRequestErrorCode = "InvalidMfaReceiptInvalidOrgId"
	MFA_REQUEST_NOT_FOUND                                  BadRequestErrorCode = "MfaRequestNotFound"
	INVALID_KEY_TYPE                                       BadRequestErrorCode = "InvalidKeyType"
	INVALID_KEY_MATERIAL                                   BadRequestErrorCode = "InvalidKeyMaterial"
	INVALID_HEX_VALUE                                      BadRequestErrorCode = "InvalidHexValue"
	INVALID_BASE32_VALUE                                   BadRequestErrorCode = "InvalidBase32Value"
	INVALID_BASE58_VALUE                                   BadRequestErrorCode = "InvalidBase58Value"
	INVALID_SS58_VALUE                                     BadRequestErrorCode = "InvalidSs58Value"
	INVALID_FORK_VERSION_LENGTH                            BadRequestErrorCode = "InvalidForkVersionLength"
	INVALID_ETH_ADDRESS                                    BadRequestErrorCode = "InvalidEthAddress"
	INVALID_STELLAR_ADDRESS                                BadRequestErrorCode = "InvalidStellarAddress"
	INVALID_ORG_NAME_OR_ID                                 BadRequestErrorCode = "InvalidOrgNameOrId"
	INVALID_STAKE_DEPOSIT                                  BadRequestErrorCode = "InvalidStakeDeposit"
	INVALID_BLOB_SIGN_REQUEST                              BadRequestErrorCode = "InvalidBlobSignRequest"
	INVALID_SOLANA_SIGN_REQUEST                            BadRequestErrorCode = "InvalidSolanaSignRequest"
	INVALID_EIP712_SIGN_REQUEST                            BadRequestErrorCode = "InvalidEip712SignRequest"
	INVALID_EVM_SIGN_REQUEST                               BadRequestErrorCode = "InvalidEvmSignRequest"
	INVALID_ETH2_SIGN_REQUEST                              BadRequestErrorCode = "InvalidEth2SignRequest"
	INVALID_DERIVE_KEY_REQUEST                             BadRequestErrorCode = "InvalidDeriveKeyRequest"
	INVALID_STAKING_AMOUNT                                 BadRequestErrorCode = "InvalidStakingAmount"
	CUSTOM_STAKING_AMOUNT_NOT_ALLOWED_FOR_WRAPPER_CONTRACT BadRequestErrorCode = "CustomStakingAmountNotAllowedForWrapperContract"
	INVALID_UNSTAKE_REQUEST                                BadRequestErrorCode = "InvalidUnstakeRequest"
	INVALID_CREATE_USER_REQUEST                            BadRequestErrorCode = "InvalidCreateUserRequest"
	USER_ALREADY_EXISTS                                    BadRequestErrorCode = "UserAlreadyExists"
	IDP_USER_ALREADY_EXISTS                                BadRequestErrorCode = "IdpUserAlreadyExists"
	COGNITO_USER_ALREADY_ORG_MEMBER                        BadRequestErrorCode = "CognitoUserAlreadyOrgMember"
	USER_NOT_FOUND                                         BadRequestErrorCode = "UserNotFound"
	POLICY_RULE_KEY_MISMATCH                               BadRequestErrorCode = "PolicyRuleKeyMismatch"
	EMPTY_SCOPES                                           BadRequestErrorCode = "EmptyScopes"
	INVALID_SCOPES_FOR_ROLE_SESSION                        BadRequestErrorCode = "InvalidScopesForRoleSession"
	INVALID_LIFETIME                                       BadRequestErrorCode = "InvalidLifetime"
	NO_SINGLE_KEY_FOR_USER                                 BadRequestErrorCode = "NoSingleKeyForUser"
	INVALID_ORG_POLICY_RULE                                BadRequestErrorCode = "InvalidOrgPolicyRule"
	SOURCE_IP_ALLOWLIST_EMPTY                              BadRequestErrorCode = "SourceIpAllowlistEmpty"
	LIMIT_WINDOW_TOO_LONG                                  BadRequestErrorCode = "LimitWindowTooLong"
	ERC20_CONTRACT_DISALLOWED                              BadRequestErrorCode = "Erc20ContractDisallowed"
	EMPTY_RULE_ERROR                                       BadRequestErrorCode = "EmptyRuleError"
	OPTIONAL_LIST_EMPTY                                    BadRequestErrorCode = "OptionalListEmpty"
	INVALID_ORG_POLICY_REPEATED_RULE                       BadRequestErrorCode = "InvalidOrgPolicyRepeatedRule"
	AVA_SIGN_HASH_ERROR                                    BadRequestErrorCode = "AvaSignHashError"
	AVA_SIGN_ERROR                                         BadRequestErrorCode = "AvaSignError"
	BTC_SEGWIT_HASH_ERROR                                  BadRequestErrorCode = "BtcSegwitHashError"
	BTC_TAPROOT_HASH_ERROR                                 BadRequestErrorCode = "BtcTaprootHashError"
	BTC_SIGN_ERROR                                         BadRequestErrorCode = "BtcSignError"
	TAPROOT_SIGN_ERROR                                     BadRequestErrorCode = "TaprootSignError"
	EIP712_SIGN_ERROR                                      BadRequestErrorCode = "Eip712SignError"
	INVALID_MEMBER_ROLE_IN_USER_ADD                        BadRequestErrorCode = "InvalidMemberRoleInUserAdd"
	THIRD_PARTY_USER_ALREADY_EXISTS                        BadRequestErrorCode = "ThirdPartyUserAlreadyExists"
	OIDC_IDENTITY_ALREADY_EXISTS                           BadRequestErrorCode = "OidcIdentityAlreadyExists"
	THIRD_PARTY_USER_NOT_FOUND                             BadRequestErrorCode = "ThirdPartyUserNotFound"
	DELETE_OIDC_USER_ERROR                                 BadRequestErrorCode = "DeleteOidcUserError"
	DELETE_USER_ERROR                                      BadRequestErrorCode = "DeleteUserError"
	SESSION_ROLE_MISMATCH                                  BadRequestErrorCode = "SessionRoleMismatch"
	INVALID_OIDC_TOKEN                                     BadRequestErrorCode = "InvalidOidcToken"
	INVALID_OIDC_IDENTITY                                  BadRequestErrorCode = "InvalidOidcIdentity"
	OIDC_ISSUER_UNSUPPORTED                                BadRequestErrorCode = "OidcIssuerUnsupported"
	OIDC_ISSUER_NOT_ALLOWED                                BadRequestErrorCode = "OidcIssuerNotAllowed"
	OIDC_ISSUER_NO_APPLICABLE_JWK                          BadRequestErrorCode = "OidcIssuerNoApplicableJwk"
	FIDO_KEY_ALREADY_REGISTERED                            BadRequestErrorCode = "FidoKeyAlreadyRegistered"
	FIDO_KEY_SIGN_COUNT_TOO_LOW                            BadRequestErrorCode = "FidoKeySignCountTooLow"
	FIDO_VERIFICATION_FAILED                               BadRequestErrorCode = "FidoVerificationFailed"
	FIDO_CHALLENGE_MFA_MISMATCH                            BadRequestErrorCode = "FidoChallengeMfaMismatch"
	UNSUPPORTED_LEGACY_COGNITO_SESSION                     BadRequestErrorCode = "UnsupportedLegacyCognitoSession"
	INVALID_IDENTITY_PROOF                                 BadRequestErrorCode = "InvalidIdentityProof"
	PAGINATION_DATA_EXPIRED                                BadRequestErrorCode = "PaginationDataExpired"
	EXISTING_KEYS_VIOLATE_EXCLUSIVE_KEY_ACCESS             BadRequestErrorCode = "ExistingKeysViolateExclusiveKeyAccess"
	EXPORT_DELAY_TOO_SHORT                                 BadRequestErrorCode = "ExportDelayTooShort"
	EXPORT_WINDOW_TOO_LONG                                 BadRequestErrorCode = "ExportWindowTooLong"
	INVALID_TOTP_FAILURE_LIMIT                             BadRequestErrorCode = "InvalidTotpFailureLimit"
	INVALID_EIP191_SIGN_REQUEST                            BadRequestErrorCode = "InvalidEip191SignRequest"
	CANNOT_RESEND_USER_INVITATION                          BadRequestErrorCode = "CannotResendUserInvitation"
	INVALID_NOTIFICATION_ENDPOINT_COUNT                    BadRequestErrorCode = "InvalidNotificationEndpointCount"
	CANNOT_DELETE_PENDING_SUBSCRIPTION                     BadRequestErrorCode = "CannotDeletePendingSubscription"
	INVALID_NOTIFICATION_URL_PROTOCOL                      BadRequestErrorCode = "InvalidNotificationUrlProtocol"
	EMPTY_ONE_OF_ORG_EVENT_FILTER                          BadRequestErrorCode = "EmptyOneOfOrgEventFilter"
	EMPTY_ALL_EXCEPT_ORG_EVENT_FILTER                      BadRequestErrorCode = "EmptyAllExceptOrgEventFilter"
	INVALID_TAP_NODE_HASH                                  BadRequestErrorCode = "InvalidTapNodeHash"
	INVALID_ONE_TIME_CODE                                  BadRequestErrorCode = "InvalidOneTimeCode"
	MESSAGE_NOT_FOUND                                      BadRequestErrorCode = "MessageNotFound"
	MESSAGE_ALREADY_SIGNED                                 BadRequestErrorCode = "MessageAlreadySigned"
	MESSAGE_REJECTED                                       BadRequestErrorCode = "MessageRejected"
	MESSAGE_REPLACED                                       BadRequestErrorCode = "MessageReplaced"
	INVALID_MESSAGE_TYPE                                   BadRequestErrorCode = "InvalidMessageType"
	EMPTY_ADDRESS                                          BadRequestErrorCode = "EmptyAddress"
	INVALID_ETH2_SIGNING_POLICY_SLOT_RANGE                 BadRequestErrorCode = "InvalidEth2SigningPolicySlotRange"
	INVALID_ETH2_SIGNING_POLICY_EPOCH_RANGE                BadRequestErrorCode = "InvalidEth2SigningPolicyEpochRange"
	INVALID_ETH2_SIGNING_POLICY_TIMESTAMP_RANGE            BadRequestErrorCode = "InvalidEth2SigningPolicyTimestampRange"
	INVALID_ETH2_SIGNING_POLICY_OVERLAPPING_RULE           BadRequestErrorCode = "InvalidEth2SigningPolicyOverlappingRule"
	MMI_RPC_URL_MISSING                                    BadRequestErrorCode = "MmiRpcUrlMissing"
	MMI_CHAIN_ID_MISSING                                   BadRequestErrorCode = "MmiChainIdMissing"
	ETHERS_INVALID_RPC_URL                                 BadRequestErrorCode = "EthersInvalidRpcUrl"
	ETHERS_GET_TRANSACTION_COUNT_ERROR                     BadRequestErrorCode = "EthersGetTransactionCountError"
	INVALID_PASSWORD                                       BadRequestErrorCode = "InvalidPassword"
	BABYLON_STAKING_FEE_PLUS_DUST_OVERFLOW                 BadRequestErrorCode = "BabylonStakingFeePlusDustOverflow"
	BABYLON_STAKING                                        BadRequestErrorCode = "BabylonStaking"
	BABYLON_STAKING_INCORRECT_KEY                          BadRequestErrorCode = "BabylonStakingIncorrectKey"
	PSBT_SIGNING                                           BadRequestErrorCode = "PsbtSigning"
)

// All allowed values of BadRequestErrorCode enum
var AllowedBadRequestErrorCodeEnumValues = []BadRequestErrorCode{
	"GenericBadRequest",
	"EmailPasswordNotFound",
	"OneTimeCodeExpired",
	"InvalidBody",
	"TokenRequestError",
	"InvalidMfaReceipt",
	"InvalidMfaPolicyCount",
	"InvalidMfaPolicyNumAuthFactors",
	"InvalidMfaPolicyNumAllowedApprovers",
	"InvalidMfaPolicyGracePeriodTooLong",
	"InvalidMfaPolicyRedundantRule",
	"InvalidBabylonStakingPolicyParams",
	"InvalidBtcTxReceiversEmptyAllowlist",
	"InvalidCreateKeyCount",
	"OrgInviteExistingUser",
	"OrgUserAlreadyExists",
	"OrgNameTaken",
	"RoleNameTaken",
	"AddKeyToRoleCountTooHigh",
	"InvalidKeyId",
	"InvalidTimeLockAlreadyInThePast",
	"InvalidRestrictedScopes",
	"InvalidUpdate",
	"InvalidMetadataLength",
	"InvalidKeyMaterialId",
	"KeyNotFound",
	"UserExportDerivedKey",
	"UserExportPublicKeyInvalid",
	"UnableToAccessSmtpRelay",
	"UserExportInProgress",
	"RoleNotFound",
	"InvalidRoleNameOrId",
	"InvalidMfaReceiptOrgIdMissing",
	"InvalidMfaReceiptInvalidOrgId",
	"MfaRequestNotFound",
	"InvalidKeyType",
	"InvalidKeyMaterial",
	"InvalidHexValue",
	"InvalidBase32Value",
	"InvalidBase58Value",
	"InvalidSs58Value",
	"InvalidForkVersionLength",
	"InvalidEthAddress",
	"InvalidStellarAddress",
	"InvalidOrgNameOrId",
	"InvalidStakeDeposit",
	"InvalidBlobSignRequest",
	"InvalidSolanaSignRequest",
	"InvalidEip712SignRequest",
	"InvalidEvmSignRequest",
	"InvalidEth2SignRequest",
	"InvalidDeriveKeyRequest",
	"InvalidStakingAmount",
	"CustomStakingAmountNotAllowedForWrapperContract",
	"InvalidUnstakeRequest",
	"InvalidCreateUserRequest",
	"UserAlreadyExists",
	"IdpUserAlreadyExists",
	"CognitoUserAlreadyOrgMember",
	"UserNotFound",
	"PolicyRuleKeyMismatch",
	"EmptyScopes",
	"InvalidScopesForRoleSession",
	"InvalidLifetime",
	"NoSingleKeyForUser",
	"InvalidOrgPolicyRule",
	"SourceIpAllowlistEmpty",
	"LimitWindowTooLong",
	"Erc20ContractDisallowed",
	"EmptyRuleError",
	"OptionalListEmpty",
	"InvalidOrgPolicyRepeatedRule",
	"AvaSignHashError",
	"AvaSignError",
	"BtcSegwitHashError",
	"BtcTaprootHashError",
	"BtcSignError",
	"TaprootSignError",
	"Eip712SignError",
	"InvalidMemberRoleInUserAdd",
	"ThirdPartyUserAlreadyExists",
	"OidcIdentityAlreadyExists",
	"ThirdPartyUserNotFound",
	"DeleteOidcUserError",
	"DeleteUserError",
	"SessionRoleMismatch",
	"InvalidOidcToken",
	"InvalidOidcIdentity",
	"OidcIssuerUnsupported",
	"OidcIssuerNotAllowed",
	"OidcIssuerNoApplicableJwk",
	"FidoKeyAlreadyRegistered",
	"FidoKeySignCountTooLow",
	"FidoVerificationFailed",
	"FidoChallengeMfaMismatch",
	"UnsupportedLegacyCognitoSession",
	"InvalidIdentityProof",
	"PaginationDataExpired",
	"ExistingKeysViolateExclusiveKeyAccess",
	"ExportDelayTooShort",
	"ExportWindowTooLong",
	"InvalidTotpFailureLimit",
	"InvalidEip191SignRequest",
	"CannotResendUserInvitation",
	"InvalidNotificationEndpointCount",
	"CannotDeletePendingSubscription",
	"InvalidNotificationUrlProtocol",
	"EmptyOneOfOrgEventFilter",
	"EmptyAllExceptOrgEventFilter",
	"InvalidTapNodeHash",
	"InvalidOneTimeCode",
	"MessageNotFound",
	"MessageAlreadySigned",
	"MessageRejected",
	"MessageReplaced",
	"InvalidMessageType",
	"EmptyAddress",
	"InvalidEth2SigningPolicySlotRange",
	"InvalidEth2SigningPolicyEpochRange",
	"InvalidEth2SigningPolicyTimestampRange",
	"InvalidEth2SigningPolicyOverlappingRule",
	"MmiRpcUrlMissing",
	"MmiChainIdMissing",
	"EthersInvalidRpcUrl",
	"EthersGetTransactionCountError",
	"InvalidPassword",
	"BabylonStakingFeePlusDustOverflow",
	"BabylonStaking",
	"BabylonStakingIncorrectKey",
	"PsbtSigning",
}

func (v *BadRequestErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BadRequestErrorCode(value)
	for _, existing := range AllowedBadRequestErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BadRequestErrorCode", value)
}

// NewBadRequestErrorCodeFromValue returns a pointer to a valid BadRequestErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBadRequestErrorCodeFromValue(v string) (*BadRequestErrorCode, error) {
	ev := BadRequestErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BadRequestErrorCode: valid values are %v", v, AllowedBadRequestErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BadRequestErrorCode) IsValid() bool {
	for _, existing := range AllowedBadRequestErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BadRequestErrorCode value
func (v BadRequestErrorCode) Ptr() *BadRequestErrorCode {
	return &v
}

type NullableBadRequestErrorCode struct {
	value *BadRequestErrorCode
	isSet bool
}

func (v NullableBadRequestErrorCode) Get() *BadRequestErrorCode {
	return v.value
}

func (v *NullableBadRequestErrorCode) Set(val *BadRequestErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestErrorCode(val *BadRequestErrorCode) *NullableBadRequestErrorCode {
	return &NullableBadRequestErrorCode{value: val, isSet: true}
}

func (v NullableBadRequestErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
