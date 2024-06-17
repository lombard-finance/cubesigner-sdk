package api

import (
	"encoding/json"
	"fmt"
)

// NotFoundErrorCode the model 'NotFoundErrorCode'
type NotFoundErrorCode string

// List of NotFoundErrorCode
const (
	URI_SEGMENT_MISSING              NotFoundErrorCode = "UriSegmentMissing"
	URI_SEGMENT_INVALID              NotFoundErrorCode = "UriSegmentInvalid"
	TOTP_NOT_CONFIGURED              NotFoundErrorCode = "TotpNotConfigured"
	FIDO_KEY_NOT_FOUND               NotFoundErrorCode = "FidoKeyNotFound"
	FIDO_CHALLENGE_NOT_FOUND         NotFoundErrorCode = "FidoChallengeNotFound"
	TOTP_CHALLENGE_NOT_FOUND         NotFoundErrorCode = "TotpChallengeNotFound"
	USER_EXPORT_REQUEST_NOT_FOUND    NotFoundErrorCode = "UserExportRequestNotFound"
	USER_EXPORT_CIPHERTEXT_NOT_FOUND NotFoundErrorCode = "UserExportCiphertextNotFound"
)

// All allowed values of NotFoundErrorCode enum
var AllowedNotFoundErrorCodeEnumValues = []NotFoundErrorCode{
	"UriSegmentMissing",
	"UriSegmentInvalid",
	"TotpNotConfigured",
	"FidoKeyNotFound",
	"FidoChallengeNotFound",
	"TotpChallengeNotFound",
	"UserExportRequestNotFound",
	"UserExportCiphertextNotFound",
}

func (v *NotFoundErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotFoundErrorCode(value)
	for _, existing := range AllowedNotFoundErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotFoundErrorCode", value)
}

// NewNotFoundErrorCodeFromValue returns a pointer to a valid NotFoundErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotFoundErrorCodeFromValue(v string) (*NotFoundErrorCode, error) {
	ev := NotFoundErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotFoundErrorCode: valid values are %v", v, AllowedNotFoundErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotFoundErrorCode) IsValid() bool {
	for _, existing := range AllowedNotFoundErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotFoundErrorCode value
func (v NotFoundErrorCode) Ptr() *NotFoundErrorCode {
	return &v
}

type NullableNotFoundErrorCode struct {
	value *NotFoundErrorCode
	isSet bool
}

func (v NullableNotFoundErrorCode) Get() *NotFoundErrorCode {
	return v.value
}

func (v *NullableNotFoundErrorCode) Set(val *NotFoundErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundErrorCode(val *NotFoundErrorCode) *NullableNotFoundErrorCode {
	return &NullableNotFoundErrorCode{value: val, isSet: true}
}

func (v NullableNotFoundErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
