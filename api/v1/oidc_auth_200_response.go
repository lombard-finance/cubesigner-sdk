package v1

import (
	"encoding/json"
)

// OidcAuth200Response Information about a new session, returned from multiple endpoints (e.g., login, refresh, etc.).
type OidcAuth200Response struct {
	// Session expiration (in seconds since UNIX epoch), beyond which it cannot be refreshed.
	Expiration  *int64            `json:"expiration,omitempty"`
	SessionInfo ClientSessionInfo `json:"session_info"`
	// New token to be used for authentication. Requests to signing endpoints should include this value in the `Authorization` header
	Token string `json:"token"`
}

// NewOidcAuth200Response instantiates a new OidcAuth200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcAuth200Response(sessionInfo ClientSessionInfo, token string) *OidcAuth200Response {
	this := OidcAuth200Response{}
	this.SessionInfo = sessionInfo
	this.Token = token
	return &this
}

// NewOidcAuth200ResponseWithDefaults instantiates a new OidcAuth200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcAuth200ResponseWithDefaults() *OidcAuth200Response {
	this := OidcAuth200Response{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *OidcAuth200Response) GetExpiration() int64 {
	if o == nil || o.Expiration == nil {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcAuth200Response) GetExpirationOk() (*int64, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *OidcAuth200Response) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *OidcAuth200Response) SetExpiration(v int64) {
	o.Expiration = &v
}

// GetSessionInfo returns the SessionInfo field value
func (o *OidcAuth200Response) GetSessionInfo() ClientSessionInfo {
	if o == nil {
		var ret ClientSessionInfo
		return ret
	}

	return o.SessionInfo
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value
// and a boolean to check if the value has been set.
func (o *OidcAuth200Response) GetSessionInfoOk() (*ClientSessionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionInfo, true
}

// SetSessionInfo sets field value
func (o *OidcAuth200Response) SetSessionInfo(v ClientSessionInfo) {
	o.SessionInfo = v
}

// GetToken returns the Token field value
func (o *OidcAuth200Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OidcAuth200Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *OidcAuth200Response) SetToken(v string) {
	o.Token = v
}

func (o OidcAuth200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if true {
		toSerialize["session_info"] = o.SessionInfo
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

//
//type NullableOidcAuth200Response struct {
//	value *OidcAuth200Response
//	isSet bool
//}
//
//func (v NullableOidcAuth200Response) Get() *OidcAuth200Response {
//	return v.value
//}
//
//func (v *NullableOidcAuth200Response) Set(val *OidcAuth200Response) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableOidcAuth200Response) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableOidcAuth200Response) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableOidcAuth200Response(val *OidcAuth200Response) *NullableOidcAuth200Response {
//	return &NullableOidcAuth200Response{value: val, isSet: true}
//}
//
//func (v NullableOidcAuth200Response) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableOidcAuth200Response) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
