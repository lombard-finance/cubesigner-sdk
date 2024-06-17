package api

import (
	"encoding/json"
	"fmt"
)

// SignerErrorOwnCodes the model 'SignerErrorOwnCodes'
type SignerErrorOwnCodes string

// List of SignerErrorOwnCodes
const (
	JRPC_ERROR        SignerErrorOwnCodes = "JrpcError"
	UNHANDLED_ERROR   SignerErrorOwnCodes = "UnhandledError"
	PROXY_START_ERROR SignerErrorOwnCodes = "ProxyStartError"
	ENCLAVE_ERROR     SignerErrorOwnCodes = "EnclaveError"
)

// All allowed values of SignerErrorOwnCodes enum
var AllowedSignerErrorOwnCodesEnumValues = []SignerErrorOwnCodes{
	"JrpcError",
	"UnhandledError",
	"ProxyStartError",
	"EnclaveError",
}

func (v *SignerErrorOwnCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignerErrorOwnCodes(value)
	for _, existing := range AllowedSignerErrorOwnCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignerErrorOwnCodes", value)
}

// NewSignerErrorOwnCodesFromValue returns a pointer to a valid SignerErrorOwnCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignerErrorOwnCodesFromValue(v string) (*SignerErrorOwnCodes, error) {
	ev := SignerErrorOwnCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignerErrorOwnCodes: valid values are %v", v, AllowedSignerErrorOwnCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignerErrorOwnCodes) IsValid() bool {
	for _, existing := range AllowedSignerErrorOwnCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignerErrorOwnCodes value
func (v SignerErrorOwnCodes) Ptr() *SignerErrorOwnCodes {
	return &v
}

type NullableSignerErrorOwnCodes struct {
	value *SignerErrorOwnCodes
	isSet bool
}

func (v NullableSignerErrorOwnCodes) Get() *SignerErrorOwnCodes {
	return v.value
}

func (v *NullableSignerErrorOwnCodes) Set(val *SignerErrorOwnCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerErrorOwnCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerErrorOwnCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerErrorOwnCodes(val *SignerErrorOwnCodes) *NullableSignerErrorOwnCodes {
	return &NullableSignerErrorOwnCodes{value: val, isSet: true}
}

func (v NullableSignerErrorOwnCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerErrorOwnCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
