package api

import (
	"encoding/json"
)

// ClientSessionInfo Session information sent to the client. This struct works in tandem with its server-side counterpart [`SessionData`].
type ClientSessionInfo struct {
	// Token to use for authorization.
	AuthToken string `json:"auth_token"`
	// DateTime measured in seconds since unix epoch. A wrapper type for serialization that encodes a [`SystemTime`] as a [`u64`] representing the number of seconds since [`SystemTime::UNIX_EPOCH`].
	AuthTokenExp int64 `json:"auth_token_exp"`
	// Epoch at which the token was last refreshed
	Epoch int32 `json:"epoch"`
	// Wrapper around a zeroizing 32-byte fixed-size array
	EpochToken string `json:"epoch_token"`
	// Token to use for refreshing the `(auth, refresh)` token pair
	RefreshToken string `json:"refresh_token"`
	// DateTime measured in seconds since unix epoch. A wrapper type for serialization that encodes a [`SystemTime`] as a [`u64`] representing the number of seconds since [`SystemTime::UNIX_EPOCH`].
	RefreshTokenExp int64 `json:"refresh_token_exp"`
	// Session ID
	SessionId string `json:"session_id"`
}

// NewClientSessionInfo instantiates a new ClientSessionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSessionInfo(authToken string, authTokenExp int64, epoch int32, epochToken string, refreshToken string, refreshTokenExp int64, sessionId string) *ClientSessionInfo {
	this := ClientSessionInfo{}
	this.AuthToken = authToken
	this.AuthTokenExp = authTokenExp
	this.Epoch = epoch
	this.EpochToken = epochToken
	this.RefreshToken = refreshToken
	this.RefreshTokenExp = refreshTokenExp
	this.SessionId = sessionId
	return &this
}

// NewClientSessionInfoWithDefaults instantiates a new ClientSessionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSessionInfoWithDefaults() *ClientSessionInfo {
	this := ClientSessionInfo{}
	return &this
}

// GetAuthToken returns the AuthToken field value
func (o *ClientSessionInfo) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *ClientSessionInfo) SetAuthToken(v string) {
	o.AuthToken = v
}

// GetAuthTokenExp returns the AuthTokenExp field value
func (o *ClientSessionInfo) GetAuthTokenExp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AuthTokenExp
}

// GetAuthTokenExpOk returns a tuple with the AuthTokenExp field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetAuthTokenExpOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthTokenExp, true
}

// SetAuthTokenExp sets field value
func (o *ClientSessionInfo) SetAuthTokenExp(v int64) {
	o.AuthTokenExp = v
}

// GetEpoch returns the Epoch field value
func (o *ClientSessionInfo) GetEpoch() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetEpochOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *ClientSessionInfo) SetEpoch(v int32) {
	o.Epoch = v
}

// GetEpochToken returns the EpochToken field value
func (o *ClientSessionInfo) GetEpochToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EpochToken
}

// GetEpochTokenOk returns a tuple with the EpochToken field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetEpochTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochToken, true
}

// SetEpochToken sets field value
func (o *ClientSessionInfo) SetEpochToken(v string) {
	o.EpochToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *ClientSessionInfo) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *ClientSessionInfo) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetRefreshTokenExp returns the RefreshTokenExp field value
func (o *ClientSessionInfo) GetRefreshTokenExp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefreshTokenExp
}

// GetRefreshTokenExpOk returns a tuple with the RefreshTokenExp field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetRefreshTokenExpOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshTokenExp, true
}

// SetRefreshTokenExp sets field value
func (o *ClientSessionInfo) SetRefreshTokenExp(v int64) {
	o.RefreshTokenExp = v
}

// GetSessionId returns the SessionId field value
func (o *ClientSessionInfo) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ClientSessionInfo) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ClientSessionInfo) SetSessionId(v string) {
	o.SessionId = v
}

func (o ClientSessionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth_token"] = o.AuthToken
	}
	if true {
		toSerialize["auth_token_exp"] = o.AuthTokenExp
	}
	if true {
		toSerialize["epoch"] = o.Epoch
	}
	if true {
		toSerialize["epoch_token"] = o.EpochToken
	}
	if true {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if true {
		toSerialize["refresh_token_exp"] = o.RefreshTokenExp
	}
	if true {
		toSerialize["session_id"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableClientSessionInfo struct {
	value *ClientSessionInfo
	isSet bool
}

func (v NullableClientSessionInfo) Get() *ClientSessionInfo {
	return v.value
}

func (v *NullableClientSessionInfo) Set(val *ClientSessionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSessionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSessionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSessionInfo(val *ClientSessionInfo) *NullableClientSessionInfo {
	return &NullableClientSessionInfo{value: val, isSet: true}
}

func (v NullableClientSessionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSessionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
