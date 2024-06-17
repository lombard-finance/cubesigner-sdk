package api

import (
	"encoding/json"
	"fmt"
)

// ForbiddenErrorCode the model 'ForbiddenErrorCode'
type ForbiddenErrorCode string

// List of ForbiddenErrorCode
const (
	FIDO_REQUIRED_TO_REMOVE_TOTP    ForbiddenErrorCode = "FidoRequiredToRemoveTotp"
	EMAIL_OTP_NOT_CONFIGURED        ForbiddenErrorCode = "EmailOtpNotConfigured"
	MFA_CHALLENGE_EXPIRED           ForbiddenErrorCode = "MfaChallengeExpired"
	CHAIN_ID_NOT_ALLOWED            ForbiddenErrorCode = "ChainIdNotAllowed"
	INVALID_ORG                     ForbiddenErrorCode = "InvalidOrg"
	ORG_ID_MISMATCH                 ForbiddenErrorCode = "OrgIdMismatch"
	SESSION_FOR_WRONG_ORG           ForbiddenErrorCode = "SessionForWrongOrg"
	SELF_DELETE                     ForbiddenErrorCode = "SelfDelete"
	SELF_DISABLE                    ForbiddenErrorCode = "SelfDisable"
	USER_MFA_NOT_CONFIGURED         ForbiddenErrorCode = "UserMfaNotConfigured"
	USER_DISABLED                   ForbiddenErrorCode = "UserDisabled"
	ORG_DISABLED                    ForbiddenErrorCode = "OrgDisabled"
	ORG_NOT_FOUND                   ForbiddenErrorCode = "OrgNotFound"
	ORG_WITHOUT_OWNER               ForbiddenErrorCode = "OrgWithoutOwner"
	ORPHANED_USER                   ForbiddenErrorCode = "OrphanedUser"
	OIDC_USER_NOT_FOUND             ForbiddenErrorCode = "OidcUserNotFound"
	USER_NOT_IN_ORG                 ForbiddenErrorCode = "UserNotInOrg"
	USER_NOT_ORG_OWNER              ForbiddenErrorCode = "UserNotOrgOwner"
	USER_NOT_KEY_OWNER              ForbiddenErrorCode = "UserNotKeyOwner"
	INVALID_ROLE                    ForbiddenErrorCode = "InvalidRole"
	DISABLED_ROLE                   ForbiddenErrorCode = "DisabledRole"
	KEY_DISABLED                    ForbiddenErrorCode = "KeyDisabled"
	ROLE_NOT_IN_ORG                 ForbiddenErrorCode = "RoleNotInOrg"
	KEY_NOT_IN_ROLE                 ForbiddenErrorCode = "KeyNotInRole"
	KEY_NOT_IN_ORG                  ForbiddenErrorCode = "KeyNotInOrg"
	USER_EXPORT_REQUEST_NOT_IN_ORG  ForbiddenErrorCode = "UserExportRequestNotInOrg"
	USER_EXPORT_REQUEST_INVALID     ForbiddenErrorCode = "UserExportRequestInvalid"
	USER_NOT_ORIGINAL_KEY_OWNER     ForbiddenErrorCode = "UserNotOriginalKeyOwner"
	USER_NOT_IN_ROLE                ForbiddenErrorCode = "UserNotInRole"
	MUST_BE_FULL_MEMBER             ForbiddenErrorCode = "MustBeFullMember"
	SESSION_EXPIRED                 ForbiddenErrorCode = "SessionExpired"
	SESSION_CHANGED                 ForbiddenErrorCode = "SessionChanged"
	SESSION_REVOKED                 ForbiddenErrorCode = "SessionRevoked"
	EXPECTED_USER_SESSION           ForbiddenErrorCode = "ExpectedUserSession"
	SESSION_ROLE_CHANGED            ForbiddenErrorCode = "SessionRoleChanged"
	SCOPED_NAME_NOT_FOUND           ForbiddenErrorCode = "ScopedNameNotFound"
	SESSION_INVALID_EPOCH_TOKEN     ForbiddenErrorCode = "SessionInvalidEpochToken"
	SESSION_INVALID_REFRESH_TOKEN   ForbiddenErrorCode = "SessionInvalidRefreshToken"
	SESSION_REFRESH_TOKEN_EXPIRED   ForbiddenErrorCode = "SessionRefreshTokenExpired"
	INVALID_AUTH_HEADER             ForbiddenErrorCode = "InvalidAuthHeader"
	SESSION_NOT_FOUND               ForbiddenErrorCode = "SessionNotFound"
	INVALID_ARN                     ForbiddenErrorCode = "InvalidArn"
	SESSION_INVALID_AUTH_TOKEN      ForbiddenErrorCode = "SessionInvalidAuthToken"
	SESSION_AUTH_TOKEN_EXPIRED      ForbiddenErrorCode = "SessionAuthTokenExpired"
	SESSION_POSSIBLY_STOLEN_TOKEN   ForbiddenErrorCode = "SessionPossiblyStolenToken"
	MFA_DISALLOWED_IDENTITY         ForbiddenErrorCode = "MfaDisallowedIdentity"
	MFA_DISALLOWED_APPROVER         ForbiddenErrorCode = "MfaDisallowedApprover"
	MFA_TYPE_NOT_ALLOWED            ForbiddenErrorCode = "MfaTypeNotAllowed"
	MFA_NOT_APPROVED_YET            ForbiddenErrorCode = "MfaNotApprovedYet"
	MFA_CONFIRMATION_CODE_MISMATCH  ForbiddenErrorCode = "MfaConfirmationCodeMismatch"
	MFA_HTTP_REQUEST_MISMATCH       ForbiddenErrorCode = "MfaHttpRequestMismatch"
	MFA_REMOVE_BELOW_MIN            ForbiddenErrorCode = "MfaRemoveBelowMin"
	TOTP_ALREADY_CONFIGURED         ForbiddenErrorCode = "TotpAlreadyConfigured"
	TOTP_CONFIGURATION_CHANGED      ForbiddenErrorCode = "TotpConfigurationChanged"
	MFA_TOTP_BAD_CONFIGURATION      ForbiddenErrorCode = "MfaTotpBadConfiguration"
	MFA_TOTP_BAD_CODE               ForbiddenErrorCode = "MfaTotpBadCode"
	MFA_TOTP_RATE_LIMIT             ForbiddenErrorCode = "MfaTotpRateLimit"
	IMPROPER_SESSION_SCOPE          ForbiddenErrorCode = "ImproperSessionScope"
	FULL_SESSION_REQUIRED           ForbiddenErrorCode = "FullSessionRequired"
	SESSION_WITHOUT_ANY_SCOPE_UNDER ForbiddenErrorCode = "SessionWithoutAnyScopeUnder"
	USER_ROLE_UNPRIVILEGED          ForbiddenErrorCode = "UserRoleUnprivileged"
	MFA_NOT_CONFIGURED              ForbiddenErrorCode = "MfaNotConfigured"
	REMOVE_LAST_OIDC_IDENTITY       ForbiddenErrorCode = "RemoveLastOidcIdentity"
)

// All allowed values of ForbiddenErrorCode enum
var AllowedForbiddenErrorCodeEnumValues = []ForbiddenErrorCode{
	"FidoRequiredToRemoveTotp",
	"EmailOtpNotConfigured",
	"MfaChallengeExpired",
	"ChainIdNotAllowed",
	"InvalidOrg",
	"OrgIdMismatch",
	"SessionForWrongOrg",
	"SelfDelete",
	"SelfDisable",
	"UserMfaNotConfigured",
	"UserDisabled",
	"OrgDisabled",
	"OrgNotFound",
	"OrgWithoutOwner",
	"OrphanedUser",
	"OidcUserNotFound",
	"UserNotInOrg",
	"UserNotOrgOwner",
	"UserNotKeyOwner",
	"InvalidRole",
	"DisabledRole",
	"KeyDisabled",
	"RoleNotInOrg",
	"KeyNotInRole",
	"KeyNotInOrg",
	"UserExportRequestNotInOrg",
	"UserExportRequestInvalid",
	"UserNotOriginalKeyOwner",
	"UserNotInRole",
	"MustBeFullMember",
	"SessionExpired",
	"SessionChanged",
	"SessionRevoked",
	"ExpectedUserSession",
	"SessionRoleChanged",
	"ScopedNameNotFound",
	"SessionInvalidEpochToken",
	"SessionInvalidRefreshToken",
	"SessionRefreshTokenExpired",
	"InvalidAuthHeader",
	"SessionNotFound",
	"InvalidArn",
	"SessionInvalidAuthToken",
	"SessionAuthTokenExpired",
	"SessionPossiblyStolenToken",
	"MfaDisallowedIdentity",
	"MfaDisallowedApprover",
	"MfaTypeNotAllowed",
	"MfaNotApprovedYet",
	"MfaConfirmationCodeMismatch",
	"MfaHttpRequestMismatch",
	"MfaRemoveBelowMin",
	"TotpAlreadyConfigured",
	"TotpConfigurationChanged",
	"MfaTotpBadConfiguration",
	"MfaTotpBadCode",
	"MfaTotpRateLimit",
	"ImproperSessionScope",
	"FullSessionRequired",
	"SessionWithoutAnyScopeUnder",
	"UserRoleUnprivileged",
	"MfaNotConfigured",
	"RemoveLastOidcIdentity",
}

func (v *ForbiddenErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ForbiddenErrorCode(value)
	for _, existing := range AllowedForbiddenErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ForbiddenErrorCode", value)
}

// NewForbiddenErrorCodeFromValue returns a pointer to a valid ForbiddenErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewForbiddenErrorCodeFromValue(v string) (*ForbiddenErrorCode, error) {
	ev := ForbiddenErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ForbiddenErrorCode: valid values are %v", v, AllowedForbiddenErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ForbiddenErrorCode) IsValid() bool {
	for _, existing := range AllowedForbiddenErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ForbiddenErrorCode value
func (v ForbiddenErrorCode) Ptr() *ForbiddenErrorCode {
	return &v
}

type NullableForbiddenErrorCode struct {
	value *ForbiddenErrorCode
	isSet bool
}

func (v NullableForbiddenErrorCode) Get() *ForbiddenErrorCode {
	return v.value
}

func (v *NullableForbiddenErrorCode) Set(val *ForbiddenErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableForbiddenErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableForbiddenErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForbiddenErrorCode(val *ForbiddenErrorCode) *NullableForbiddenErrorCode {
	return &NullableForbiddenErrorCode{value: val, isSet: true}
}

func (v NullableForbiddenErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForbiddenErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
