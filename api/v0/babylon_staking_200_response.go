package v0

import (
	"encoding/json"
)

// BabylonStaking200Response struct for BabylonStaking200Response
type BabylonStaking200Response struct {
	// The transaction fee in sats
	Fee int64 `json:"fee"`
	// The PSBT in standard hex serialization, without leading \"0x\".
	Psbt string `json:"psbt"`
}

// NewBabylonStaking200Response instantiates a new BabylonStaking200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStaking200Response(fee int64, psbt string) *BabylonStaking200Response {
	this := BabylonStaking200Response{}
	this.Fee = fee
	this.Psbt = psbt
	return &this
}

// NewBabylonStaking200ResponseWithDefaults instantiates a new BabylonStaking200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStaking200ResponseWithDefaults() *BabylonStaking200Response {
	this := BabylonStaking200Response{}
	return &this
}

// GetFee returns the Fee field value
func (o *BabylonStaking200Response) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *BabylonStaking200Response) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *BabylonStaking200Response) SetFee(v int64) {
	o.Fee = v
}

// GetPsbt returns the Psbt field value
func (o *BabylonStaking200Response) GetPsbt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Psbt
}

// GetPsbtOk returns a tuple with the Psbt field value
// and a boolean to check if the value has been set.
func (o *BabylonStaking200Response) GetPsbtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Psbt, true
}

// SetPsbt sets field value
func (o *BabylonStaking200Response) SetPsbt(v string) {
	o.Psbt = v
}

func (o BabylonStaking200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["psbt"] = o.Psbt
	}
	return json.Marshal(toSerialize)
}

type NullableBabylonStaking200Response struct {
	value *BabylonStaking200Response
	isSet bool
}

func (v NullableBabylonStaking200Response) Get() *BabylonStaking200Response {
	return v.value
}

func (v *NullableBabylonStaking200Response) Set(val *BabylonStaking200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStaking200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStaking200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStaking200Response(val *BabylonStaking200Response) *NullableBabylonStaking200Response {
	return &NullableBabylonStaking200Response{value: val, isSet: true}
}

func (v NullableBabylonStaking200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStaking200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


