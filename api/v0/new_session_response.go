package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// NewSessionResponse Information about a new session, returned from multiple endpoints (e.g., login, refresh, etc.).
type NewSessionResponse struct {
	// Session expiration (in seconds since UNIX epoch), beyond which it cannot be refreshed.
	Expiration *int64 `json:"expiration,omitempty"`
	// Token that can be used to refresh this session.
	RefreshToken api.NullableString    `json:"refresh_token"`
	SessionInfo  api.ClientSessionInfo `json:"session_info"`
	// New token to be used for authentication. Requests to signing endpoints should include this value in the `Authorization` header.
	Token string `json:"token"`
}

// NewNewSessionResponse instantiates a new NewSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewSessionResponse(refreshToken api.NullableString, sessionInfo api.ClientSessionInfo, token string) *NewSessionResponse {
	this := NewSessionResponse{}
	this.RefreshToken = refreshToken
	this.SessionInfo = sessionInfo
	this.Token = token
	return &this
}

// NewNewSessionResponseWithDefaults instantiates a new NewSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewSessionResponseWithDefaults() *NewSessionResponse {
	this := NewSessionResponse{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *NewSessionResponse) GetExpiration() int64 {
	if o == nil || o.Expiration == nil {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSessionResponse) GetExpirationOk() (*int64, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *NewSessionResponse) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *NewSessionResponse) SetExpiration(v int64) {
	o.Expiration = &v
}

// GetRefreshToken returns the RefreshToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NewSessionResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewSessionResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// SetRefreshToken sets field value
func (o *NewSessionResponse) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}

// GetSessionInfo returns the SessionInfo field value
func (o *NewSessionResponse) GetSessionInfo() api.ClientSessionInfo {
	if o == nil {
		var ret api.ClientSessionInfo
		return ret
	}

	return o.SessionInfo
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value
// and a boolean to check if the value has been set.
func (o *NewSessionResponse) GetSessionInfoOk() (*api.ClientSessionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionInfo, true
}

// SetSessionInfo sets field value
func (o *NewSessionResponse) SetSessionInfo(v api.ClientSessionInfo) {
	o.SessionInfo = v
}

// GetToken returns the Token field value
func (o *NewSessionResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *NewSessionResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *NewSessionResponse) SetToken(v string) {
	o.Token = v
}

func (o NewSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if true {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if true {
		toSerialize["session_info"] = o.SessionInfo
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableNewSessionResponse struct {
	value *NewSessionResponse
	isSet bool
}

func (v NullableNewSessionResponse) Get() *NewSessionResponse {
	return v.value
}

func (v *NullableNewSessionResponse) Set(val *NewSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewSessionResponse(val *NewSessionResponse) *NullableNewSessionResponse {
	return &NullableNewSessionResponse{value: val, isSet: true}
}

func (v NullableNewSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
