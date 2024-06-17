package api

import (
	"encoding/json"
	"fmt"
)

// BadGatewayErrorCode the model 'BadGatewayErrorCode'
type BadGatewayErrorCode string

// List of BadGatewayErrorCode
const (
	O_AUTH_PROVIDER_ERROR                BadGatewayErrorCode = "OAuthProviderError"
	OIDC_DISOVERY_FAILED                 BadGatewayErrorCode = "OidcDisoveryFailed"
	OIDC_ISSUER_JWK_ENDPOINT_UNAVAILABLE BadGatewayErrorCode = "OidcIssuerJwkEndpointUnavailable"
	SMTP_SERVER_UNAVAILABLE              BadGatewayErrorCode = "SmtpServerUnavailable"
)

// All allowed values of BadGatewayErrorCode enum
var AllowedBadGatewayErrorCodeEnumValues = []BadGatewayErrorCode{
	"OAuthProviderError",
	"OidcDisoveryFailed",
	"OidcIssuerJwkEndpointUnavailable",
	"SmtpServerUnavailable",
}

func (v *BadGatewayErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BadGatewayErrorCode(value)
	for _, existing := range AllowedBadGatewayErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BadGatewayErrorCode", value)
}

// NewBadGatewayErrorCodeFromValue returns a pointer to a valid BadGatewayErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBadGatewayErrorCodeFromValue(v string) (*BadGatewayErrorCode, error) {
	ev := BadGatewayErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BadGatewayErrorCode: valid values are %v", v, AllowedBadGatewayErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BadGatewayErrorCode) IsValid() bool {
	for _, existing := range AllowedBadGatewayErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BadGatewayErrorCode value
func (v BadGatewayErrorCode) Ptr() *BadGatewayErrorCode {
	return &v
}

type NullableBadGatewayErrorCode struct {
	value *BadGatewayErrorCode
	isSet bool
}

func (v NullableBadGatewayErrorCode) Get() *BadGatewayErrorCode {
	return v.value
}

func (v *NullableBadGatewayErrorCode) Set(val *BadGatewayErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableBadGatewayErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableBadGatewayErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadGatewayErrorCode(val *BadGatewayErrorCode) *NullableBadGatewayErrorCode {
	return &NullableBadGatewayErrorCode{value: val, isSet: true}
}

func (v NullableBadGatewayErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadGatewayErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
