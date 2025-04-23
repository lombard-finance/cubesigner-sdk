package v0

import (
	"encoding/json"

	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonRegistrationRequest
type BabylonRegistrationRequest struct {
	// bbn_addr, a Cosmos secp256k1 bech32 address with either 'bbn' or 'cosmos' HRP
	BbnAddress string `json:"bbn_addr"`

	// otherwise it is identical to the BabylonStakingDeposit struct
	BabylonStakingDeposit
}

// NewBabylonRegistrationRequest instantiates a new BabylonRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonRegistrationRequest(bbnAddress string, finalityProviderPk string, lockTime int32, network api.BabylonNetworkId, stakerPk string, change string, fee int64, feeType api.FeeType, psbt string, value int64) *BabylonRegistrationRequest {
	this := BabylonRegistrationRequest{}
	this.BbnAddress = bbnAddress
	this.FinalityProviderPk = finalityProviderPk
	this.LockTime = lockTime
	this.Network = network
	this.StakerPk = stakerPk
	this.Change = change
	this.Fee = fee
	this.FeeType = feeType
	this.Psbt = psbt
	this.Value = value
	return &this
}

// NewBabylonRegistrationRequestWithDefaults instantiates a new BabylonRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonRegistrationRequestWithDefaults() *BabylonRegistrationRequest {
	this := BabylonRegistrationRequest{}
	return &this
}

// GetBbnAddress returns the BbnAddress field value
func (o *BabylonRegistrationRequest) GetBbnAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BbnAddress
}

// GetBbnAddressOk returns a tuple with the BbnAddress field value
// and a boolean to check if the value has been set.
func (o *BabylonRegistrationRequest) GetBbnAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BbnAddress, true
}

// SetBbnAddress sets field value
func (o *BabylonRegistrationRequest) SetBbnAddress(v string) {
	o.BbnAddress = v
}

// Marshal data from the first non-nil pointers in the struct to JSON
// BabylonRegistrationRequest inherits MarshalJSON from BabylonStakingDeposit.
// We must override that func to get correct serialization.
func (src BabylonRegistrationRequest) MarshalJSON() ([]byte, error) {
	depJson, err := src.BabylonStakingDeposit.MarshalJSON()
	if err != nil {
		return nil, err
	}

	addrJson, err := json.Marshal(map[string]any{
		"bbn_addr": src.BbnAddress,
	})
	if err != nil {
		return nil, err
	}

	addrJson[0] = ','
	return append(depJson[:len(depJson)-1], addrJson...), nil
}
