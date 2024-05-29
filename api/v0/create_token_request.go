package v0

import (
	"encoding/json"
)

// CreateTokenRequest struct for CreateTokenRequest
type CreateTokenRequest struct {
	AuthLifetime    *int64 `json:"auth_lifetime,omitempty"`
	GraceLifetime   *int64 `json:"grace_lifetime,omitempty"`
	RefreshLifetime *int64 `json:"refresh_lifetime,omitempty"`
	SessionLifetime *int64 `json:"session_lifetime,omitempty"`
	// A human readable description of the purpose of the key
	Purpose string `json:"purpose"`
	// Controls what capabilities this session will have. By default, it has all signing capabilities, i.e., just the 'sign:*' scope.
	Scopes []string `json:"scopes,omitempty"`
}

// NewCreateTokenRequest instantiates a new CreateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenRequest(purpose string) *CreateTokenRequest {
	this := CreateTokenRequest{}
	this.Purpose = purpose
	return &this
}

// NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenRequestWithDefaults() *CreateTokenRequest {
	this := CreateTokenRequest{}
	return &this
}

// GetAuthLifetime returns the AuthLifetime field value if set, zero value otherwise.
func (o *CreateTokenRequest) GetAuthLifetime() int64 {
	if o == nil || o.AuthLifetime == nil {
		var ret int64
		return ret
	}
	return *o.AuthLifetime
}

// GetAuthLifetimeOk returns a tuple with the AuthLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetAuthLifetimeOk() (*int64, bool) {
	if o == nil || o.AuthLifetime == nil {
		return nil, false
	}
	return o.AuthLifetime, true
}

// HasAuthLifetime returns a boolean if a field has been set.
func (o *CreateTokenRequest) HasAuthLifetime() bool {
	if o != nil && o.AuthLifetime != nil {
		return true
	}

	return false
}

// SetAuthLifetime gets a reference to the given int64 and assigns it to the AuthLifetime field.
func (o *CreateTokenRequest) SetAuthLifetime(v int64) {
	o.AuthLifetime = &v
}

// GetGraceLifetime returns the GraceLifetime field value if set, zero value otherwise.
func (o *CreateTokenRequest) GetGraceLifetime() int64 {
	if o == nil || o.GraceLifetime == nil {
		var ret int64
		return ret
	}
	return *o.GraceLifetime
}

// GetGraceLifetimeOk returns a tuple with the GraceLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetGraceLifetimeOk() (*int64, bool) {
	if o == nil || o.GraceLifetime == nil {
		return nil, false
	}
	return o.GraceLifetime, true
}

// HasGraceLifetime returns a boolean if a field has been set.
func (o *CreateTokenRequest) HasGraceLifetime() bool {
	if o != nil && o.GraceLifetime != nil {
		return true
	}

	return false
}

// SetGraceLifetime gets a reference to the given int64 and assigns it to the GraceLifetime field.
func (o *CreateTokenRequest) SetGraceLifetime(v int64) {
	o.GraceLifetime = &v
}

// GetRefreshLifetime returns the RefreshLifetime field value if set, zero value otherwise.
func (o *CreateTokenRequest) GetRefreshLifetime() int64 {
	if o == nil || o.RefreshLifetime == nil {
		var ret int64
		return ret
	}
	return *o.RefreshLifetime
}

// GetRefreshLifetimeOk returns a tuple with the RefreshLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetRefreshLifetimeOk() (*int64, bool) {
	if o == nil || o.RefreshLifetime == nil {
		return nil, false
	}
	return o.RefreshLifetime, true
}

// HasRefreshLifetime returns a boolean if a field has been set.
func (o *CreateTokenRequest) HasRefreshLifetime() bool {
	if o != nil && o.RefreshLifetime != nil {
		return true
	}

	return false
}

// SetRefreshLifetime gets a reference to the given int64 and assigns it to the RefreshLifetime field.
func (o *CreateTokenRequest) SetRefreshLifetime(v int64) {
	o.RefreshLifetime = &v
}

// GetSessionLifetime returns the SessionLifetime field value if set, zero value otherwise.
func (o *CreateTokenRequest) GetSessionLifetime() int64 {
	if o == nil || o.SessionLifetime == nil {
		var ret int64
		return ret
	}
	return *o.SessionLifetime
}

// GetSessionLifetimeOk returns a tuple with the SessionLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetSessionLifetimeOk() (*int64, bool) {
	if o == nil || o.SessionLifetime == nil {
		return nil, false
	}
	return o.SessionLifetime, true
}

// HasSessionLifetime returns a boolean if a field has been set.
func (o *CreateTokenRequest) HasSessionLifetime() bool {
	if o != nil && o.SessionLifetime != nil {
		return true
	}

	return false
}

// SetSessionLifetime gets a reference to the given int64 and assigns it to the SessionLifetime field.
func (o *CreateTokenRequest) SetSessionLifetime(v int64) {
	o.SessionLifetime = &v
}

// GetPurpose returns the Purpose field value
func (o *CreateTokenRequest) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *CreateTokenRequest) SetPurpose(v string) {
	o.Purpose = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *CreateTokenRequest) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *CreateTokenRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o CreateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthLifetime != nil {
		toSerialize["auth_lifetime"] = o.AuthLifetime
	}
	if o.GraceLifetime != nil {
		toSerialize["grace_lifetime"] = o.GraceLifetime
	}
	if o.RefreshLifetime != nil {
		toSerialize["refresh_lifetime"] = o.RefreshLifetime
	}
	if o.SessionLifetime != nil {
		toSerialize["session_lifetime"] = o.SessionLifetime
	}
	if true {
		toSerialize["purpose"] = o.Purpose
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return json.Marshal(toSerialize)
}
