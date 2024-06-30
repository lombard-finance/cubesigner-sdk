package v0

// BabylonStakingResponse struct for BabylonStakingResponse
type BabylonStakingResponse struct {
	// The transaction fee in sats.
	Fee int64 `json:"fee"`
	// The PSBT in standard hex serialization, without leading "0x".
	PSBT string `json:"psbt"`
}

func NewBabylonStakingResponse(fee int64, psbt string) *BabylonStakingResponse {
	this := BabylonStakingResponse{}
	this.Fee = fee
	this.PSBT = psbt
	return &this
}

// NewBabylonStakingResponseWithDefaults instantiates a new BabylonStakingResponse object with defaults
func NewBabylonStakingResponseWithDefaults() *BabylonStakingResponse {
	this := BabylonStakingResponse{}
	return &this
}

func (o *BabylonStakingResponse) GetFee() int64 {
	if o == nil {
		return 0
	}
	return o.Fee
}

func (o *BabylonStakingResponse) GetFeeOk() (int64, bool) {
	if o == nil {
		return 0, false
	}
	return o.Fee, true
}

func (o *BabylonStakingResponse) SetFee(v int64) {
	o.Fee = v
}

func (o *BabylonStakingResponse) GetPSBT() string {
	if o == nil {
		return ""
	}
	return o.PSBT
}

func (o *BabylonStakingResponse) GetPSBTOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.PSBT, true
}

func (o *BabylonStakingResponse) SetPSBT(v string) {
	o.PSBT = v
}
