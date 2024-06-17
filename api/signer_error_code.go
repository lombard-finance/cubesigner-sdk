package api

import (
	"encoding/json"
	"fmt"
)

// SignerErrorCode - struct for SignerErrorCode
type SignerErrorCode struct {
	AcceptedValueCode     *AcceptedValueCode
	BadGatewayErrorCode   *BadGatewayErrorCode
	BadRequestErrorCode   *BadRequestErrorCode
	ForbiddenErrorCode    *ForbiddenErrorCode
	InternalErrorCode     *InternalErrorCode
	NotFoundErrorCode     *NotFoundErrorCode
	PreconditionErrorCode *PreconditionErrorCode
	SignerErrorOwnCodes   *SignerErrorOwnCodes
	UnauthorizedErrorCode *UnauthorizedErrorCode
}

// AcceptedValueCodeAsSignerErrorCode is a convenience function that returns AcceptedValueCode wrapped in SignerErrorCode
func AcceptedValueCodeAsSignerErrorCode(v *AcceptedValueCode) SignerErrorCode {
	return SignerErrorCode{
		AcceptedValueCode: v,
	}
}

// BadGatewayErrorCodeAsSignerErrorCode is a convenience function that returns BadGatewayErrorCode wrapped in SignerErrorCode
func BadGatewayErrorCodeAsSignerErrorCode(v *BadGatewayErrorCode) SignerErrorCode {
	return SignerErrorCode{
		BadGatewayErrorCode: v,
	}
}

// BadRequestErrorCodeAsSignerErrorCode is a convenience function that returns BadRequestErrorCode wrapped in SignerErrorCode
func BadRequestErrorCodeAsSignerErrorCode(v *BadRequestErrorCode) SignerErrorCode {
	return SignerErrorCode{
		BadRequestErrorCode: v,
	}
}

// ForbiddenErrorCodeAsSignerErrorCode is a convenience function that returns ForbiddenErrorCode wrapped in SignerErrorCode
func ForbiddenErrorCodeAsSignerErrorCode(v *ForbiddenErrorCode) SignerErrorCode {
	return SignerErrorCode{
		ForbiddenErrorCode: v,
	}
}

// InternalErrorCodeAsSignerErrorCode is a convenience function that returns InternalErrorCode wrapped in SignerErrorCode
func InternalErrorCodeAsSignerErrorCode(v *InternalErrorCode) SignerErrorCode {
	return SignerErrorCode{
		InternalErrorCode: v,
	}
}

// NotFoundErrorCodeAsSignerErrorCode is a convenience function that returns NotFoundErrorCode wrapped in SignerErrorCode
func NotFoundErrorCodeAsSignerErrorCode(v *NotFoundErrorCode) SignerErrorCode {
	return SignerErrorCode{
		NotFoundErrorCode: v,
	}
}

// PreconditionErrorCodeAsSignerErrorCode is a convenience function that returns PreconditionErrorCode wrapped in SignerErrorCode
func PreconditionErrorCodeAsSignerErrorCode(v *PreconditionErrorCode) SignerErrorCode {
	return SignerErrorCode{
		PreconditionErrorCode: v,
	}
}

// SignerErrorOwnCodesAsSignerErrorCode is a convenience function that returns SignerErrorOwnCodes wrapped in SignerErrorCode
func SignerErrorOwnCodesAsSignerErrorCode(v *SignerErrorOwnCodes) SignerErrorCode {
	return SignerErrorCode{
		SignerErrorOwnCodes: v,
	}
}

// UnauthorizedErrorCodeAsSignerErrorCode is a convenience function that returns UnauthorizedErrorCode wrapped in SignerErrorCode
func UnauthorizedErrorCodeAsSignerErrorCode(v *UnauthorizedErrorCode) SignerErrorCode {
	return SignerErrorCode{
		UnauthorizedErrorCode: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SignerErrorCode) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AcceptedValueCode
	err = newStrictDecoder(data).Decode(&dst.AcceptedValueCode)
	if err == nil {
		jsonAcceptedValueCode, _ := json.Marshal(dst.AcceptedValueCode)
		if string(jsonAcceptedValueCode) == "{}" { // empty struct
			dst.AcceptedValueCode = nil
		} else {
			match++
		}
	} else {
		dst.AcceptedValueCode = nil
	}

	// try to unmarshal data into BadGatewayErrorCode
	err = newStrictDecoder(data).Decode(&dst.BadGatewayErrorCode)
	if err == nil {
		jsonBadGatewayErrorCode, _ := json.Marshal(dst.BadGatewayErrorCode)
		if string(jsonBadGatewayErrorCode) == "{}" { // empty struct
			dst.BadGatewayErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.BadGatewayErrorCode = nil
	}

	// try to unmarshal data into BadRequestErrorCode
	err = newStrictDecoder(data).Decode(&dst.BadRequestErrorCode)
	if err == nil {
		jsonBadRequestErrorCode, _ := json.Marshal(dst.BadRequestErrorCode)
		if string(jsonBadRequestErrorCode) == "{}" { // empty struct
			dst.BadRequestErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.BadRequestErrorCode = nil
	}

	// try to unmarshal data into ForbiddenErrorCode
	err = newStrictDecoder(data).Decode(&dst.ForbiddenErrorCode)
	if err == nil {
		jsonForbiddenErrorCode, _ := json.Marshal(dst.ForbiddenErrorCode)
		if string(jsonForbiddenErrorCode) == "{}" { // empty struct
			dst.ForbiddenErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.ForbiddenErrorCode = nil
	}

	// try to unmarshal data into InternalErrorCode
	err = newStrictDecoder(data).Decode(&dst.InternalErrorCode)
	if err == nil {
		jsonInternalErrorCode, _ := json.Marshal(dst.InternalErrorCode)
		if string(jsonInternalErrorCode) == "{}" { // empty struct
			dst.InternalErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.InternalErrorCode = nil
	}

	// try to unmarshal data into NotFoundErrorCode
	err = newStrictDecoder(data).Decode(&dst.NotFoundErrorCode)
	if err == nil {
		jsonNotFoundErrorCode, _ := json.Marshal(dst.NotFoundErrorCode)
		if string(jsonNotFoundErrorCode) == "{}" { // empty struct
			dst.NotFoundErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.NotFoundErrorCode = nil
	}

	// try to unmarshal data into PreconditionErrorCode
	err = newStrictDecoder(data).Decode(&dst.PreconditionErrorCode)
	if err == nil {
		jsonPreconditionErrorCode, _ := json.Marshal(dst.PreconditionErrorCode)
		if string(jsonPreconditionErrorCode) == "{}" { // empty struct
			dst.PreconditionErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.PreconditionErrorCode = nil
	}

	// try to unmarshal data into SignerErrorOwnCodes
	err = newStrictDecoder(data).Decode(&dst.SignerErrorOwnCodes)
	if err == nil {
		jsonSignerErrorOwnCodes, _ := json.Marshal(dst.SignerErrorOwnCodes)
		if string(jsonSignerErrorOwnCodes) == "{}" { // empty struct
			dst.SignerErrorOwnCodes = nil
		} else {
			match++
		}
	} else {
		dst.SignerErrorOwnCodes = nil
	}

	// try to unmarshal data into UnauthorizedErrorCode
	err = newStrictDecoder(data).Decode(&dst.UnauthorizedErrorCode)
	if err == nil {
		jsonUnauthorizedErrorCode, _ := json.Marshal(dst.UnauthorizedErrorCode)
		if string(jsonUnauthorizedErrorCode) == "{}" { // empty struct
			dst.UnauthorizedErrorCode = nil
		} else {
			match++
		}
	} else {
		dst.UnauthorizedErrorCode = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AcceptedValueCode = nil
		dst.BadGatewayErrorCode = nil
		dst.BadRequestErrorCode = nil
		dst.ForbiddenErrorCode = nil
		dst.InternalErrorCode = nil
		dst.NotFoundErrorCode = nil
		dst.PreconditionErrorCode = nil
		dst.SignerErrorOwnCodes = nil
		dst.UnauthorizedErrorCode = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SignerErrorCode)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SignerErrorCode)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SignerErrorCode) MarshalJSON() ([]byte, error) {
	if src.AcceptedValueCode != nil {
		return json.Marshal(&src.AcceptedValueCode)
	}

	if src.BadGatewayErrorCode != nil {
		return json.Marshal(&src.BadGatewayErrorCode)
	}

	if src.BadRequestErrorCode != nil {
		return json.Marshal(&src.BadRequestErrorCode)
	}

	if src.ForbiddenErrorCode != nil {
		return json.Marshal(&src.ForbiddenErrorCode)
	}

	if src.InternalErrorCode != nil {
		return json.Marshal(&src.InternalErrorCode)
	}

	if src.NotFoundErrorCode != nil {
		return json.Marshal(&src.NotFoundErrorCode)
	}

	if src.PreconditionErrorCode != nil {
		return json.Marshal(&src.PreconditionErrorCode)
	}

	if src.SignerErrorOwnCodes != nil {
		return json.Marshal(&src.SignerErrorOwnCodes)
	}

	if src.UnauthorizedErrorCode != nil {
		return json.Marshal(&src.UnauthorizedErrorCode)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SignerErrorCode) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AcceptedValueCode != nil {
		return obj.AcceptedValueCode
	}

	if obj.BadGatewayErrorCode != nil {
		return obj.BadGatewayErrorCode
	}

	if obj.BadRequestErrorCode != nil {
		return obj.BadRequestErrorCode
	}

	if obj.ForbiddenErrorCode != nil {
		return obj.ForbiddenErrorCode
	}

	if obj.InternalErrorCode != nil {
		return obj.InternalErrorCode
	}

	if obj.NotFoundErrorCode != nil {
		return obj.NotFoundErrorCode
	}

	if obj.PreconditionErrorCode != nil {
		return obj.PreconditionErrorCode
	}

	if obj.SignerErrorOwnCodes != nil {
		return obj.SignerErrorOwnCodes
	}

	if obj.UnauthorizedErrorCode != nil {
		return obj.UnauthorizedErrorCode
	}

	// all schemas are nil
	return nil
}

type NullableSignerErrorCode struct {
	value *SignerErrorCode
	isSet bool
}

func (v NullableSignerErrorCode) Get() *SignerErrorCode {
	return v.value
}

func (v *NullableSignerErrorCode) Set(val *SignerErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerErrorCode(val *SignerErrorCode) *NullableSignerErrorCode {
	return &NullableSignerErrorCode{value: val, isSet: true}
}

func (v NullableSignerErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
