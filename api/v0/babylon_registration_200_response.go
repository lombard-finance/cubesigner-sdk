package v0

import (
	"encoding/json"
)

// BabylonRegistration200Response
type BabylonRegistration200Response struct {
	// The deposit PSBT. Its encoding will be either hex or base64, matching the
	// encoding of the PSBT in the request (which can be either hex or base64)
	Deposit string `json:"deposit"`
	// The fee computed for the deposit in sats
	DepositFee uint64 `json:"deposit_fee"`
	// The unsigned unbonding transaction in Bitcoin consensus hex encoding
	UnbondTx string `json:"unbond"`
	// The slash-deposit transaction in Bitcoin consensus hex encoding
	SlashDeposit string `json:"slash_deposit"`
	// The slash-unbond transaction in Bitcoin consensus hex encoding
	SlashUnbond string `json:"slash_unbond"`
	// The slash-deposit Taproot signature in hex-encoded BIP340 serialization
	SlashDepositSig string `json:"slash_deposit_sig"`
	// The slash-unbond Taproot signature in hex-encoded BIP340 serialization
	SlashUnbondSig string `json:"slash_unbond_sig"`
	// The proof of possession, a BIP322 signature on the staker's Babylon
	// address signed by the staker Taproot key. This is a Bitcoin witness
	// stack in Bitcoin consensus hex encoding
	ProofOfPossession string `json:"pop"`
	// The Babylon address that will receive the staking rewards. This is the
	// same as the `bbn_addr` value specified in the request, but it always
	// has a `bbn` human-readable part.
	BbnAddress string `json:"bbn_addr"`
}

// NewBabylonRegistrationResponse instantiates a new BabylonRegistrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonRegistrationResponse(deposit string, depositFee uint64, unbondTx string, slashDeposit string, slashUnbond string, slashDepositSig string, slashUnbondSig string, proofOfPossession string, bbnAddress string) *BabylonRegistration200Response {
	this := BabylonRegistration200Response{}
	this.Deposit = deposit
	this.DepositFee = depositFee
	this.UnbondTx = unbondTx
	this.SlashDeposit = slashDeposit
	this.SlashUnbond = slashUnbond
	this.SlashDepositSig = slashDepositSig
	this.SlashUnbondSig = slashUnbondSig
	this.ProofOfPossession = proofOfPossession
	this.BbnAddress = bbnAddress
	return &this
}

// NewBabylonRegistrationResponseWithDefaults instantiates a new BabylonRegistrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonRegistrationResponseWithDefaults() *BabylonRegistration200Response {
	this := BabylonRegistration200Response{}
	return &this
}

// GetDeposit returns the Deposit field value
func (o *BabylonRegistration200Response) GetDeposit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetDepositOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deposit, true
}

// SetDeposit sets field value
func (o *BabylonRegistration200Response) SetDeposit(v string) {
	o.Deposit = v
}

// GetDepositFee returns the DepositFee field value
func (o *BabylonRegistration200Response) GetDepositFee() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.DepositFee
}

// GetDepositFeeOk returns a tuple with the DepositFee field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetDepositFeeOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositFee, true
}

// SetDepositFee sets field value
func (o *BabylonRegistration200Response) SetDepositFee(v uint64) {
	o.DepositFee = v
}

// GetUnbondTx returns the UnbondTx field value
func (o *BabylonRegistration200Response) GetUnbondTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnbondTx
}

// GetUnbondTxOk returns a tuple with the UnbondTx field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetUnbondTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbondTx, true
}

// SetUnbondTx sets field value
func (o *BabylonRegistration200Response) SetUnbondTx(v string) {
	o.UnbondTx = v
}

// GetSlashDeposit returns the SlashDeposit field value
func (o *BabylonRegistration200Response) GetSlashDeposit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlashDeposit
}

// GetSlashDepositOk returns a tuple with the SlashDeposit field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetSlashDepositOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashDeposit, true
}

// SetSlashDeposit sets field value
func (o *BabylonRegistration200Response) SetSlashDeposit(v string) {
	o.SlashDeposit = v
}

// GetSlashUnbond returns the SlashUnbond field value
func (o *BabylonRegistration200Response) GetSlashUnbond() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlashUnbond
}

// GetSlashUnbondOk returns a tuple with the SlashUnbond field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetSlashUnbondOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashUnbond, true
}

// SetSlashUnbond sets field value
func (o *BabylonRegistration200Response) SetSlashUnbond(v string) {
	o.SlashUnbond = v
}

// GetSlashDepositSig returns the SlashDepositSig field value
func (o *BabylonRegistration200Response) GetSlashDepositSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlashDepositSig
}

// GetSlashDepositSigOk returns a tuple with the SlashDepositSig field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetSlashDepositSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashDepositSig, true
}

// SetSlashDepositSig sets field value
func (o *BabylonRegistration200Response) SetSlashDepositSig(v string) {
	o.SlashDepositSig = v
}

// GetSlashUnbondSig returns the SlashUnbondSig field value
func (o *BabylonRegistration200Response) GetSlashUnbondSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlashUnbondSig
}

// GetSlashUnbondSigOk returns a tuple with the SlashUnbondSig field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetSlashUnbondSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlashUnbondSig, true
}

// SetSlashUnbondSig sets field value
func (o *BabylonRegistration200Response) SetSlashUnbondSig(v string) {
	o.SlashUnbondSig = v
}

// GetProofOfPossession returns the ProofOfPossession field value
func (o *BabylonRegistration200Response) GetProofOfPossession() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProofOfPossession
}

// GetProofOfPossessionOk returns a tuple with the ProofOfPossession field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetProofOfPossessionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProofOfPossession, true
}

// SetProofOfPossession sets field value
func (o *BabylonRegistration200Response) SetProofOfPossession(v string) {
	o.ProofOfPossession = v
}

// GetBbnAddress returns the BbnAddress field value
func (o *BabylonRegistration200Response) GetBbnAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BbnAddress
}

// GetBbnAddressOk returns a tuple with the BbnAddress field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistration200Response) GetBbnAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BbnAddress, true
}

// SetBbnAddress sets field value
func (o *BabylonRegistration200Response) SetBbnAddress(v string) {
	o.BbnAddress = v
}

// MarshalJSON serializes the struct using spec logic
func (o BabylonRegistration200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deposit"] = o.Deposit
	toSerialize["deposit_fee"] = o.DepositFee
	toSerialize["unbond"] = o.UnbondTx
	toSerialize["slash_deposit"] = o.SlashDeposit
	toSerialize["slash_unbond"] = o.SlashUnbond
	toSerialize["slash_deposit_sig"] = o.SlashDepositSig
	toSerialize["slash_unbond_sig"] = o.SlashUnbondSig
	toSerialize["pop"] = o.ProofOfPossession
	toSerialize["bbn_addr"] = o.BbnAddress
	return json.Marshal(toSerialize)
}
