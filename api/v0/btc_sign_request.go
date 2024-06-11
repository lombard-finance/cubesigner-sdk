package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BtcSignRequest struct for BtcSignRequest
type BtcSignRequest struct {
	SigKind api.BtcSignatureKind   `json:"sig_kind"`
	Tx      map[string]interface{} `json:"tx"`
}

// NewBtcSignRequest instantiates a new BtcSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtcSignRequest(sigKind api.BtcSignatureKind, tx map[string]interface{}) *BtcSignRequest {
	this := BtcSignRequest{}
	this.SigKind = sigKind
	this.Tx = tx
	return &this
}

// NewBtcSignRequestWithDefaults instantiates a new BtcSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtcSignRequestWithDefaults() *BtcSignRequest {
	this := BtcSignRequest{}
	return &this
}

// GetSigKind returns the SigKind field value
func (o *BtcSignRequest) GetSigKind() api.BtcSignatureKind {
	if o == nil {
		var ret api.BtcSignatureKind
		return ret
	}

	return o.SigKind
}

// GetSigKindOk returns a tuple with the SigKind field value
// and a boolean to check if the value has been set.
func (o *BtcSignRequest) GetSigKindOk() (*api.BtcSignatureKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigKind, true
}

// SetSigKind sets field value
func (o *BtcSignRequest) SetSigKind(v api.BtcSignatureKind) {
	o.SigKind = v
}

// GetTx returns the Tx field value
func (o *BtcSignRequest) GetTx() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *BtcSignRequest) GetTxOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tx, true
}

// SetTx sets field value
func (o *BtcSignRequest) SetTx(v map[string]interface{}) {
	o.Tx = v
}

func (o BtcSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sig_kind"] = o.SigKind
	}
	if true {
		toSerialize["tx"] = o.Tx
	}
	return json.Marshal(toSerialize)
}
