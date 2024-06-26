package v0

// MfaHttpRequest struct
type MfaHttpRequest struct {
	// 	Information about the request. Captures all the relevant info (including the
	// request body) about requests that require MFA. We use this to verify that when
	// a request is resumed (after obtaining necessary MFA approvals) it is exactly
	// the same as it originally was.
	Body map[string]any `json:"body,omitempty"`
	// HTTP method of the request
	Method string `json:"method"`
	// HTTP path of the request
	Path string `json:"path"`
}

// NewMfaHttpRequest instantiates a new MfaHttpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaHttpRequest(body map[string]any, method string, path string) *MfaHttpRequest {
	this := MfaHttpRequest{}
	this.Body = body
	this.Method = method
	this.Path = path
	return &this
}

// NewMfaHttpRequestWithDefaults instantiates a new MfaHttpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaHttpRequestWithDefaults() *MfaHttpRequest {
	this := MfaHttpRequest{}
	return &this
}

// GetBody returns the Body field value
func (o *MfaHttpRequest) GetBody() map[string]any {
	if o == nil {
		var ret map[string]any
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *MfaHttpRequest) GetBodyOk() (*map[string]any, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets the Body field value
func (o *MfaHttpRequest) SetBody(v map[string]any) {
	o.Body = v
}

// GetMethod returns the Method field value
func (o *MfaHttpRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *MfaHttpRequest) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets the Method field value
func (o *MfaHttpRequest) SetMethod(v string) {
	o.Method = v
}

// GetPath returns the Path field value
func (o *MfaHttpRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *MfaHttpRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets the Path field value
func (o *MfaHttpRequest) SetPath(v string) {
	o.Path = v
}
