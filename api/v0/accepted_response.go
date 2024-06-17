package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// AcceptedResponse struct for AcceptedResponse
type AcceptedResponse struct {
	Accepted  NullableErrorResponseAccepted `json:"accepted,omitempty"`
	ErrorCode api.SignerErrorCode           `json:"error_code"`
	// Error message
	Message string `json:"message"`
	// Optional request identifier
	RequestId *string `json:"request_id,omitempty"`
}

// NewAcceptedResponse instantiates a new AcceptedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptedResponse(errorCode api.SignerErrorCode, message string) *AcceptedResponse {
	this := AcceptedResponse{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewAcceptedResponseWithDefaults instantiates a new AcceptedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedResponseWithDefaults() *AcceptedResponse {
	this := AcceptedResponse{}
	return &this
}

// GetAccepted returns the Accepted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceptedResponse) GetAccepted() ErrorResponseAccepted {
	if o == nil || o.Accepted.Get() == nil {
		var ret ErrorResponseAccepted
		return ret
	}
	return *o.Accepted.Get()
}

// GetAcceptedOk returns a tuple with the Accepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceptedResponse) GetAcceptedOk() (*ErrorResponseAccepted, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accepted.Get(), o.Accepted.IsSet()
}

// HasAccepted returns a boolean if a field has been set.
func (o *AcceptedResponse) HasAccepted() bool {
	if o != nil && o.Accepted.IsSet() {
		return true
	}

	return false
}

// SetAccepted gets a reference to the given NullableErrorResponseAccepted and assigns it to the Accepted field.
func (o *AcceptedResponse) SetAccepted(v ErrorResponseAccepted) {
	o.Accepted.Set(&v)
}

// SetAcceptedNil sets the value for Accepted to be an explicit nil
func (o *AcceptedResponse) SetAcceptedNil() {
	o.Accepted.Set(nil)
}

// UnsetAccepted ensures that no value is present for Accepted, not even an explicit nil
func (o *AcceptedResponse) UnsetAccepted() {
	o.Accepted.Unset()
}

// GetErrorCode returns the ErrorCode field value
func (o *AcceptedResponse) GetErrorCode() api.SignerErrorCode {
	if o == nil {
		var ret api.SignerErrorCode
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *AcceptedResponse) GetErrorCodeOk() (*api.SignerErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *AcceptedResponse) SetErrorCode(v api.SignerErrorCode) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *AcceptedResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AcceptedResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AcceptedResponse) SetMessage(v string) {
	o.Message = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AcceptedResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AcceptedResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AcceptedResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o AcceptedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accepted.IsSet() {
		toSerialize["accepted"] = o.Accepted.Get()
	}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptedResponse struct {
	value *AcceptedResponse
	isSet bool
}

func (v NullableAcceptedResponse) Get() *AcceptedResponse {
	return v.value
}

func (v *NullableAcceptedResponse) Set(val *AcceptedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedResponse(val *AcceptedResponse) *NullableAcceptedResponse {
	return &NullableAcceptedResponse{value: val, isSet: true}
}

func (v NullableAcceptedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
