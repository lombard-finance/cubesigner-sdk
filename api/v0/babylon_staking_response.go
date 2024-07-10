package v0

import (
	"encoding/json"
)

// BabylonStakingResponse struct for BabylonStakingResponse
type BabylonStakingResponse struct {
	// The transaction fee in sats
	Fee int64 `json:"fee"`
	// The PSBT in standard hex serialization, without leading \"0x\".
	Psbt string `json:"psbt"`
}

// NewBabylonStakingResponse instantiates a new BabylonStakingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingResponse(fee int64, psbt string) *BabylonStakingResponse {
	this := BabylonStakingResponse{}
	this.Fee = fee
	this.Psbt = psbt
	return &this
}

// NewBabylonStakingResponseWithDefaults instantiates a new BabylonStakingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingResponseWithDefaults() *BabylonStakingResponse {
	this := BabylonStakingResponse{}
	return &this
}

// GetFee returns the Fee field value
func (o *BabylonStakingResponse) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingResponse) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *BabylonStakingResponse) SetFee(v int64) {
	o.Fee = v
}

// GetPsbt returns the Psbt field value
func (o *BabylonStakingResponse) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingResponse) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *BabylonStakingResponse) SetPsbt(v string) {
	o.Psbt = v
}

func (o BabylonStakingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["psbt"] = o.Psbt
	}
	return json.Marshal(toSerialize)
}

type NullableBabylonStakingResponse struct {
	value *BabylonStakingResponse
	isSet bool
}

func (v NullableBabylonStakingResponse) Get() *BabylonStakingResponse {
	return v.value
}

func (v *NullableBabylonStakingResponse) Set(val *BabylonStakingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingResponse(val *BabylonStakingResponse) *NullableBabylonStakingResponse {
	return &NullableBabylonStakingResponse{value: val, isSet: true}
}

func (v NullableBabylonStakingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


