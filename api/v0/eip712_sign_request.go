package v0

import (
	"encoding/json"
)

// Eip712SignRequest struct for Eip712SignRequest
type Eip712SignRequest struct {
	// The chain-id to which this typed data will be sent
	ChainId   int64     `json:"chain_id"`
	TypedData TypedData `json:"typed_data"`
}

// NewEip712SignRequest instantiates a new Eip712SignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEip712SignRequest(chainId int64, typedData TypedData) *Eip712SignRequest {
	this := Eip712SignRequest{}
	this.ChainId = chainId
	this.TypedData = typedData
	return &this
}

// NewEip712SignRequestWithDefaults instantiates a new Eip712SignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEip712SignRequestWithDefaults() *Eip712SignRequest {
	this := Eip712SignRequest{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *Eip712SignRequest) GetChainId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *Eip712SignRequest) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *Eip712SignRequest) SetChainId(v int64) {
	o.ChainId = v
}

// GetTypedData returns the TypedData field value
func (o *Eip712SignRequest) GetTypedData() TypedData {
	if o == nil {
		var ret TypedData
		return ret
	}

	return o.TypedData
}

// GetTypedDataOk returns a tuple with the TypedData field value
// and a boolean to check if the value has been set.
func (o *Eip712SignRequest) GetTypedDataOk() (*TypedData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypedData, true
}

// SetTypedData sets field value
func (o *Eip712SignRequest) SetTypedData(v TypedData) {
	o.TypedData = v
}

func (o Eip712SignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain_id"] = o.ChainId
	}
	if true {
		toSerialize["typed_data"] = o.TypedData
	}
	return json.Marshal(toSerialize)
}

type NullableEip712SignRequest struct {
	value *Eip712SignRequest
	isSet bool
}

func (v NullableEip712SignRequest) Get() *Eip712SignRequest {
	return v.value
}

func (v *NullableEip712SignRequest) Set(val *Eip712SignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEip712SignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEip712SignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEip712SignRequest(val *Eip712SignRequest) *NullableEip712SignRequest {
	return &NullableEip712SignRequest{value: val, isSet: true}
}

func (v NullableEip712SignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEip712SignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
