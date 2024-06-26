package v0

// MfaReceiptHeader struct
type MfaReceiptHeader struct {
	// Approval request ID.
	Id string `json:"id"`
	// Confirmation code the user needs to present when resuming the original request.
	Confirmation string `json:"confirmation"`
	// Organization ID.
	OrganizationId string `json:"organization"`
}

// NewMfaReceiptHeader instantiates a new MfaReceiptHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaReceiptHeader(id string, confirmation string, organization string) *MfaReceiptHeader {
	this := MfaReceiptHeader{}
	this.Id = id
	this.Confirmation = confirmation
	this.OrganizationId = organization
	return &this
}

// NewMfaReceiptHeaderWithDefaults instantiates a new MfaReceiptHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaReceiptHeaderWithDefaults() *MfaReceiptHeader {
	this := MfaReceiptHeader{}
	return &this
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise
func (o *MfaReceiptHeader) GetConfirmation() string {
	if o == nil || o.Confirmation == "" {
		var ret string
		return ret
	}
	return o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaReceiptHeader) GetConfirmationOk() (*string, bool) {
	if o == nil || o.Confirmation == "" {
		return nil, false
	}
	return &o.Confirmation, true
}

// SetConfirmation sets the Confirmation field value
func (o *MfaReceiptHeader) SetConfirmation(v string) {
	o.Confirmation = v
}

// GetId returns the Id field value
func (o *MfaReceiptHeader) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MfaReceiptHeader) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets the Id field value
func (o *MfaReceiptHeader) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *MfaReceiptHeader) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *MfaReceiptHeader) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetId sets the OrganizationId field value
func (o *MfaReceiptHeader) SetOrganizationId(v string) {
	o.OrganizationId = v
}
