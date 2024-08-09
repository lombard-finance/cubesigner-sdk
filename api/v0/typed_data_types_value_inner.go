package v0

import (
	"encoding/json"
)

// TypedDataTypesValueInner Represents the name and type pair
type TypedDataTypesValueInner struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewTypedDataTypesValueInner instantiates a new TypedDataTypesValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypedDataTypesValueInner(name string, type_ string) *TypedDataTypesValueInner {
	this := TypedDataTypesValueInner{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewTypedDataTypesValueInnerWithDefaults instantiates a new TypedDataTypesValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypedDataTypesValueInnerWithDefaults() *TypedDataTypesValueInner {
	this := TypedDataTypesValueInner{}
	return &this
}

// GetName returns the Name field value
func (o *TypedDataTypesValueInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TypedDataTypesValueInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TypedDataTypesValueInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TypedDataTypesValueInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TypedDataTypesValueInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TypedDataTypesValueInner) SetType(v string) {
	o.Type = v
}

func (o TypedDataTypesValueInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTypedDataTypesValueInner struct {
	value *TypedDataTypesValueInner
	isSet bool
}

func (v NullableTypedDataTypesValueInner) Get() *TypedDataTypesValueInner {
	return v.value
}

func (v *NullableTypedDataTypesValueInner) Set(val *TypedDataTypesValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedDataTypesValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedDataTypesValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedDataTypesValueInner(val *TypedDataTypesValueInner) *NullableTypedDataTypesValueInner {
	return &NullableTypedDataTypesValueInner{value: val, isSet: true}
}

func (v NullableTypedDataTypesValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedDataTypesValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
