package v0

// MfaReceipt struct
type MfaReceipt struct {
	// Confirmation code the user needs to present when resuming the original request.
	Confirmation string `json:"confirmation,omitempty"`
	// The ID of the logged-in user whose action created this approval.
	FinalApprover string `json:"final_approver,omitempty"`
	// DateTime measured in seconds since unix epoch. A wrapper type for serialization
	// that encodes a `SystemTime` as a `u64` representing the number of seconds since
	// `SystemTime::UNIX_EPOCH`.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NewMfaReceipt instantiates a new MfaReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaReceipt(confirmation string, finalApprover string, timestamp *int64) *MfaReceipt {
	this := MfaReceipt{}
	this.Confirmation = confirmation
	this.FinalApprover = finalApprover
	this.Timestamp = timestamp
	return &this
}

// NewMfaReceiptWithDefaults instantiates a new MfaReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaReceiptWithDefaults() *MfaReceipt {
	this := MfaReceipt{}
	return &this
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise
func (o *MfaReceipt) GetConfirmation() string {
	if o == nil || o.Confirmation == "" {
		var ret string
		return ret
	}
	return o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaReceipt) GetConfirmationOk() (*string, bool) {
	if o == nil || o.Confirmation == "" {
		return nil, false
	}
	return &o.Confirmation, true
}

// SetConfirmation sets the Confirmation field value
func (o *MfaReceipt) SetConfirmation(v string) {
	o.Confirmation = v
}

// GetFinalApprover returns the FinalApprover field value if set, zero value otherwise
func (o *MfaReceipt) GetFinalApprover() string {
	if o == nil || o.FinalApprover == "" {
		var ret string
		return ret
	}
	return o.FinalApprover
}

// GetFinalApproverOk returns a tuple with the FinalApprover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaReceipt) GetFinalApproverOk() (*string, bool) {
	if o == nil || o.FinalApprover == "" {
		return nil, false
	}
	return &o.FinalApprover, true
}

// SetFinalApprover sets the FinalApprover field value
func (o *MfaReceipt) SetFinalApprover(v string) {
	o.FinalApprover = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise
func (o *MfaReceipt) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaReceipt) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// SetTimestamp sets the Timestamp field value
func (o *MfaReceipt) SetTimestamp(v int64) {
	o.Timestamp = &v
}
