package v1

import (
	"encoding/json"
)

// AuthData Data required for both `authenticate` and `refresh`.
type AuthData struct {
	EpochNum int32 `json:"epoch_num"`
	// Wrapper around a zeroizing 32-byte fixed-size array
	EpochToken string `json:"epoch_token"`
	OtherToken string `json:"other_token"`
}

// NewAuthData instantiates a new AuthData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthData(epochNum int32, epochToken string, otherToken string) *AuthData {
	this := AuthData{}
	this.EpochNum = epochNum
	this.EpochToken = epochToken
	this.OtherToken = otherToken
	return &this
}

// NewAuthDataWithDefaults instantiates a new AuthData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDataWithDefaults() *AuthData {
	this := AuthData{}
	return &this
}

// GetEpochNum returns the EpochNum field value
func (o *AuthData) GetEpochNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EpochNum
}

// GetEpochNumOk returns a tuple with the EpochNum field value
// and a boolean to check if the value has been set.
func (o *AuthData) GetEpochNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochNum, true
}

// SetEpochNum sets field value
func (o *AuthData) SetEpochNum(v int32) {
	o.EpochNum = v
}

// GetEpochToken returns the EpochToken field value
func (o *AuthData) GetEpochToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EpochToken
}

// GetEpochTokenOk returns a tuple with the EpochToken field value
// and a boolean to check if the value has been set.
func (o *AuthData) GetEpochTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpochToken, true
}

// SetEpochToken sets field value
func (o *AuthData) SetEpochToken(v string) {
	o.EpochToken = v
}

// GetOtherToken returns the OtherToken field value
func (o *AuthData) GetOtherToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OtherToken
}

// GetOtherTokenOk returns a tuple with the OtherToken field value
// and a boolean to check if the value has been set.
func (o *AuthData) GetOtherTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherToken, true
}

// SetOtherToken sets field value
func (o *AuthData) SetOtherToken(v string) {
	o.OtherToken = v
}

func (o AuthData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["epoch_num"] = o.EpochNum
	}
	if true {
		toSerialize["epoch_token"] = o.EpochToken
	}
	if true {
		toSerialize["other_token"] = o.OtherToken
	}
	return json.Marshal(toSerialize)
}

//
//type NullableAuthData struct {
//	value *AuthData
//	isSet bool
//}
//
//func (v NullableAuthData) Get() *AuthData {
//	return v.value
//}
//
//func (v *NullableAuthData) Set(val *AuthData) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableAuthData) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableAuthData) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableAuthData(val *AuthData) *NullableAuthData {
//	return &NullableAuthData{value: val, isSet: true}
//}
//
//func (v NullableAuthData) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableAuthData) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
