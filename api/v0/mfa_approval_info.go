package v0

// MfaApprovalInfo struct
type MfaApprovalInfo struct {
	// DateTime measured in seconds since unix epoch. A wrapper type for serialization
	// that encodes a `SystemTime` as a `u64` representing the number of seconds since
	// `SystemTime::UNIX_EPOCH`.
	Timestamp int64 `json:"timestamp"`
}

// NewMfaApprovalInfo instantiates a new MfaApprovalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaApprovalInfo(timestamp int64) *MfaApprovalInfo {
	this := MfaApprovalInfo{}
	this.Timestamp = timestamp

	return &this
}

// NewMfaApprovalInfoWithDefaults instantiates a new MfaApprovalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaApprovalInfoWithDefaults() *MfaApprovalInfo {
	this := MfaApprovalInfo{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *MfaApprovalInfo) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MfaApprovalInfo) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets the Timestamp field value
func (o *MfaApprovalInfo) SetTimestamp(v int64) {
	o.Timestamp = v
}
