package v0

import (
	"encoding/json"
)

// TypedData Represents the [EIP-712](https://eips.ethereum.org/EIPS/eip-712) typed data object.  Typed data is a JSON object containing type information, domain separator parameters and the message object.
type TypedData struct {
	Domain TypedDataDomain `json:"domain"`
	// The message to be signed.
	Message map[string]interface{} `json:"message"`
	// The type of the message.
	PrimaryType string `json:"primaryType"`
	// The custom types used by this message.
	Types map[string][]TypedDataTypesValueInner `json:"types"`
}

// NewTypedData instantiates a new TypedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypedData(domain TypedDataDomain, message map[string]interface{}, primaryType string, types map[string][]TypedDataTypesValueInner) *TypedData {
	this := TypedData{}
	this.Domain = domain
	this.Message = message
	this.PrimaryType = primaryType
	this.Types = types
	return &this
}

// NewTypedDataWithDefaults instantiates a new TypedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypedDataWithDefaults() *TypedData {
	this := TypedData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *TypedData) GetDomain() TypedDataDomain {
	if o == nil {
		var ret TypedDataDomain
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *TypedData) GetDomainOk() (*TypedDataDomain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *TypedData) SetDomain(v TypedDataDomain) {
	o.Domain = v
}

// GetMessage returns the Message field value
func (o *TypedData) GetMessage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TypedData) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *TypedData) SetMessage(v map[string]interface{}) {
	o.Message = v
}

// GetPrimaryType returns the PrimaryType field value
func (o *TypedData) GetPrimaryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryType
}

// GetPrimaryTypeOk returns a tuple with the PrimaryType field value
// and a boolean to check if the value has been set.
func (o *TypedData) GetPrimaryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryType, true
}

// SetPrimaryType sets field value
func (o *TypedData) SetPrimaryType(v string) {
	o.PrimaryType = v
}

// GetTypes returns the Types field value
func (o *TypedData) GetTypes() map[string][]TypedDataTypesValueInner {
	if o == nil {
		var ret map[string][]TypedDataTypesValueInner
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *TypedData) GetTypesOk() (*map[string][]TypedDataTypesValueInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Types, true
}

// SetTypes sets field value
func (o *TypedData) SetTypes(v map[string][]TypedDataTypesValueInner) {
	o.Types = v
}

func (o TypedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["primaryType"] = o.PrimaryType
	}
	if true {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableTypedData struct {
	value *TypedData
	isSet bool
}

func (v NullableTypedData) Get() *TypedData {
	return v.value
}

func (v *NullableTypedData) Set(val *TypedData) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedData) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedData(val *TypedData) *NullableTypedData {
	return &NullableTypedData{value: val, isSet: true}
}

func (v NullableTypedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
