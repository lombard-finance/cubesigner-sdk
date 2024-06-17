package api

import (
	"encoding/json"
	"fmt"
)

// InternalErrorCode the model 'InternalErrorCode'
type InternalErrorCode string

// List of InternalErrorCode
const (
	SYSTEM_TIME_ERROR                      InternalErrorCode = "SystemTimeError"
	SEND_MAIL_ERROR                        InternalErrorCode = "SendMailError"
	REQWEST_ERROR                          InternalErrorCode = "ReqwestError"
	EMAIL_CONSTRUCTION_ERROR               InternalErrorCode = "EmailConstructionError"
	DB_QUERY_ERROR                         InternalErrorCode = "DbQueryError"
	DB_GET_ERROR                           InternalErrorCode = "DbGetError"
	DB_DELETE_ERROR                        InternalErrorCode = "DbDeleteError"
	DB_PUT_ERROR                           InternalErrorCode = "DbPutError"
	DB_UPDATE_ERROR                        InternalErrorCode = "DbUpdateError"
	SERDE_ERROR                            InternalErrorCode = "SerdeError"
	TEST_AND_SET_ERROR                     InternalErrorCode = "TestAndSetError"
	CONCURRENT_LOCK_CREATION               InternalErrorCode = "ConcurrentLockCreation"
	DB_GET_ITEMS_ERROR                     InternalErrorCode = "DbGetItemsError"
	DB_WRITE_ERROR                         InternalErrorCode = "DbWriteError"
	CUBIST_SIGNER_ERROR                    InternalErrorCode = "CubistSignerError"
	CW_PUT_METRIC_DATA_ERROR               InternalErrorCode = "CwPutMetricDataError"
	KMS_GENERATE_RANDOM_ERROR              InternalErrorCode = "KmsGenerateRandomError"
	MALFORMED_TOTP_BYTES                   InternalErrorCode = "MalformedTotpBytes"
	KMS_GENERATE_RANDOM_NO_RESPONSE_ERROR  InternalErrorCode = "KmsGenerateRandomNoResponseError"
	CREATE_KEY_ERROR                       InternalErrorCode = "CreateKeyError"
	PARSE_DERIVATION_PATH_ERROR            InternalErrorCode = "ParseDerivationPathError"
	SPLIT_SIGNER_ERROR                     InternalErrorCode = "SplitSignerError"
	CREATE_IMPORT_KEY_ERROR                InternalErrorCode = "CreateImportKeyError"
	CREATE_EOTS_NONCES_ERROR               InternalErrorCode = "CreateEotsNoncesError"
	EOTS_SIGN_ERROR                        InternalErrorCode = "EotsSignError"
	COGNITO_DELETE_USER_ERROR              InternalErrorCode = "CognitoDeleteUserError"
	COGNITO_LIST_USERS_ERROR               InternalErrorCode = "CognitoListUsersError"
	COGNITO_GET_USER_ERROR                 InternalErrorCode = "CognitoGetUserError"
	MISSING_USER_EMAIL                     InternalErrorCode = "MissingUserEmail"
	COGNITO_RESEND_USER_INVITATION         InternalErrorCode = "CognitoResendUserInvitation"
	COGNITO_SET_USER_PASSWORD_ERROR        InternalErrorCode = "CognitoSetUserPasswordError"
	GENERIC_INTERNAL_ERROR                 InternalErrorCode = "GenericInternalError"
	OIDC_AUTH_WITHOUT_ORG                  InternalErrorCode = "OidcAuthWithoutOrg"
	MISSING_KEY_METADATA                   InternalErrorCode = "MissingKeyMetadata"
	KMS_KEY_WITHOUT_ID                     InternalErrorCode = "KmsKeyWithoutId"
	KMS_ENABLE_KEY_ERROR                   InternalErrorCode = "KmsEnableKeyError"
	KMS_DISABLE_KEY_ERROR                  InternalErrorCode = "KmsDisableKeyError"
	SERIALIZE_ENCRYPTED_EXPORT_KEY_ERROR   InternalErrorCode = "SerializeEncryptedExportKeyError"
	DESERIALIZE_ENCRYPTED_EXPORT_KEY_ERROR InternalErrorCode = "DeserializeEncryptedExportKeyError"
	RE_ENCRYPT_USER_EXPORT                 InternalErrorCode = "ReEncryptUserExport"
	S3_UPLOAD_ERROR                        InternalErrorCode = "S3UploadError"
	S3_DOWNLOAD_ERROR                      InternalErrorCode = "S3DownloadError"
	MANAGED_STATE_MISSING                  InternalErrorCode = "ManagedStateMissing"
	INTERNAL_HEADER_MISSING                InternalErrorCode = "InternalHeaderMissing"
	INVALID_INTERNAL_HEADER_VALUE          InternalErrorCode = "InvalidInternalHeaderValue"
	REQUEST_LOCAL_STATE_ALREADY_SET        InternalErrorCode = "RequestLocalStateAlreadySet"
	OIDC_ORG_MISMATCH                      InternalErrorCode = "OidcOrgMismatch"
	ORPHANED_ROLE_KEY_ID                   InternalErrorCode = "OrphanedRoleKeyId"
	OIDC_ISSUER_INVALID_JWK                InternalErrorCode = "OidcIssuerInvalidJwk"
	INVALID_PK_FOR_MATERIAL_ID             InternalErrorCode = "InvalidPkForMaterialId"
	UNCHECKED_ORG                          InternalErrorCode = "UncheckedOrg"
	SESSION_ORG_ID_MISSING                 InternalErrorCode = "SessionOrgIdMissing"
	AVA_SIGN_CREDS_MISSING                 InternalErrorCode = "AvaSignCredsMissing"
	AVA_SIGN_SIGNATURE_MISSING             InternalErrorCode = "AvaSignSignatureMissing"
	EXPECTED_ROLE_SESSION                  InternalErrorCode = "ExpectedRoleSession"
	INVALID_THIRD_PARTY_IDENTITY           InternalErrorCode = "InvalidThirdPartyIdentity"
	COGNITO_GET_USER                       InternalErrorCode = "CognitoGetUser"
	SNS_SUBSCRIBE_ERROR                    InternalErrorCode = "SnsSubscribeError"
	SNS_UNSUBSCRIBE_ERROR                  InternalErrorCode = "SnsUnsubscribeError"
	SNS_GET_SUBSCRIPTION_ATTRIBUTES_ERROR  InternalErrorCode = "SnsGetSubscriptionAttributesError"
	SNS_SUBSCRIPTION_ATTRIBUTES_MISSING    InternalErrorCode = "SnsSubscriptionAttributesMissing"
	SNS_SET_SUBSCRIPTION_ATTRIBUTES_ERROR  InternalErrorCode = "SnsSetSubscriptionAttributesError"
	SNS_PUBLISH_BATCH_ERROR                InternalErrorCode = "SnsPublishBatchError"
	INCONSISTENT_MULTI_VALUE_TEST_AND_SET  InternalErrorCode = "InconsistentMultiValueTestAndSet"
	MATERIAL_ID_ERROR                      InternalErrorCode = "MaterialIdError"
	INVALID_BTC_ADDRESS                    InternalErrorCode = "InvalidBtcAddress"
)

// All allowed values of InternalErrorCode enum
var AllowedInternalErrorCodeEnumValues = []InternalErrorCode{
	"SystemTimeError",
	"SendMailError",
	"ReqwestError",
	"EmailConstructionError",
	"DbQueryError",
	"DbGetError",
	"DbDeleteError",
	"DbPutError",
	"DbUpdateError",
	"SerdeError",
	"TestAndSetError",
	"ConcurrentLockCreation",
	"DbGetItemsError",
	"DbWriteError",
	"CubistSignerError",
	"CwPutMetricDataError",
	"KmsGenerateRandomError",
	"MalformedTotpBytes",
	"KmsGenerateRandomNoResponseError",
	"CreateKeyError",
	"ParseDerivationPathError",
	"SplitSignerError",
	"CreateImportKeyError",
	"CreateEotsNoncesError",
	"EotsSignError",
	"CognitoDeleteUserError",
	"CognitoListUsersError",
	"CognitoGetUserError",
	"MissingUserEmail",
	"CognitoResendUserInvitation",
	"CognitoSetUserPasswordError",
	"GenericInternalError",
	"OidcAuthWithoutOrg",
	"MissingKeyMetadata",
	"KmsKeyWithoutId",
	"KmsEnableKeyError",
	"KmsDisableKeyError",
	"SerializeEncryptedExportKeyError",
	"DeserializeEncryptedExportKeyError",
	"ReEncryptUserExport",
	"S3UploadError",
	"S3DownloadError",
	"ManagedStateMissing",
	"InternalHeaderMissing",
	"InvalidInternalHeaderValue",
	"RequestLocalStateAlreadySet",
	"OidcOrgMismatch",
	"OrphanedRoleKeyId",
	"OidcIssuerInvalidJwk",
	"InvalidPkForMaterialId",
	"UncheckedOrg",
	"SessionOrgIdMissing",
	"AvaSignCredsMissing",
	"AvaSignSignatureMissing",
	"ExpectedRoleSession",
	"InvalidThirdPartyIdentity",
	"CognitoGetUser",
	"SnsSubscribeError",
	"SnsUnsubscribeError",
	"SnsGetSubscriptionAttributesError",
	"SnsSubscriptionAttributesMissing",
	"SnsSetSubscriptionAttributesError",
	"SnsPublishBatchError",
	"InconsistentMultiValueTestAndSet",
	"MaterialIdError",
	"InvalidBtcAddress",
}

func (v *InternalErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalErrorCode(value)
	for _, existing := range AllowedInternalErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalErrorCode", value)
}

// NewInternalErrorCodeFromValue returns a pointer to a valid InternalErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalErrorCodeFromValue(v string) (*InternalErrorCode, error) {
	ev := InternalErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalErrorCode: valid values are %v", v, AllowedInternalErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalErrorCode) IsValid() bool {
	for _, existing := range AllowedInternalErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InternalErrorCode value
func (v InternalErrorCode) Ptr() *InternalErrorCode {
	return &v
}

type NullableInternalErrorCode struct {
	value *InternalErrorCode
	isSet bool
}

func (v NullableInternalErrorCode) Get() *InternalErrorCode {
	return v.value
}

func (v *NullableInternalErrorCode) Set(val *InternalErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalErrorCode(val *InternalErrorCode) *NullableInternalErrorCode {
	return &NullableInternalErrorCode{value: val, isSet: true}
}

func (v NullableInternalErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
