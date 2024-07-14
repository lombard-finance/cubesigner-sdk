package v0

import (
	"encoding/json"
)

// TaprootSignRequest struct for TaprootSignRequest
type TaprootSignRequest struct {
	SigKind TaprootSignatureKind   `json:"sig_kind"`
	Tx      map[string]interface{} `json:"tx"`
}

// NewTaprootSignRequest instantiates a new TaprootSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaprootSignRequest(sigKind TaprootSignatureKind, tx map[string]interface{}) *TaprootSignRequest {
	this := TaprootSignRequest{}
	this.SigKind = sigKind
	this.Tx = tx
	return &this
}

// NewTaprootSignRequestWithDefaults instantiates a new TaprootSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaprootSignRequestWithDefaults() *TaprootSignRequest {
	this := TaprootSignRequest{}
	return &this
}

// GetSigKind returns the SigKind field value
func (o *TaprootSignRequest) GetSigKind() TaprootSignatureKind {
	if o == nil {
		var ret TaprootSignatureKind
		return ret
	}

	return o.SigKind
}

// GetSigKindOk returns a tuple with the SigKind field value
// and a boolean to check if the value has been set.
func (o *TaprootSignRequest) GetSigKindOk() (*TaprootSignatureKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigKind, true
}

// SetSigKind sets field value
func (o *TaprootSignRequest) SetSigKind(v TaprootSignatureKind) {
	o.SigKind = v
}

// GetTx returns the Tx field value
func (o *TaprootSignRequest) GetTx() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *TaprootSignRequest) GetTxOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tx, true
}

// SetTx sets field value
func (o *TaprootSignRequest) SetTx(v map[string]interface{}) {
	o.Tx = v
}

func (o TaprootSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sig_kind"] = o.SigKind
	}
	if true {
		toSerialize["tx"] = o.Tx
	}
	return json.Marshal(toSerialize)
}

//type NullableTaprootSignRequest struct {
//	value *TaprootSignRequest
//	isSet bool
//}
//
//func (v NullableTaprootSignRequest) Get() *TaprootSignRequest {
//	return v.value
//}
//
//func (v *NullableTaprootSignRequest) Set(val *TaprootSignRequest) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableTaprootSignRequest) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableTaprootSignRequest) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableTaprootSignRequest(val *TaprootSignRequest) *NullableTaprootSignRequest {
//	return &NullableTaprootSignRequest{value: val, isSet: true}
//}
//
//func (v NullableTaprootSignRequest) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableTaprootSignRequest) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
