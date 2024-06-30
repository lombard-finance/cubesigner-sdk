package v0

// BabylonScriptData struct for BabylonScriptData
type BabylonScriptData struct {
	// Optional explicit parameters.
	ExplicitParams *BabylonStakingParams `json:"explicit_params,omitempty"`
	// The Schnorr public key (i.e., 32-byte X-coordinate) of the finality provider to which the stake is delegated.
	FinalityProviderPK string `json:"finality_provider_pk"`
	// The lock time used for the withdrawal output in the staking deposit transaction.
	LockTime int32 `json:"lock_time"`
	// The network ID for Babylon.
	Network BabylonNetworkId `json:"network"`
	// The Schnorr public key (i.e., 32-byte X-coordinate) of the staker. This is the key that signs the slashing, withdrawal, and unbonding scripts.
	StakerPK string `json:"staker_pk"`
	// The parameter version to use. If `None`, uses the latest version.
	Version *int32 `json:"version,omitempty"`
}

func NewBabylonScriptData(finalityProviderPK string, lockTime int32, network BabylonNetworkId, stakerPK string) *BabylonScriptData {
	this := BabylonScriptData{}
	this.FinalityProviderPK = finalityProviderPK
	this.LockTime = lockTime
	this.Network = network
	this.StakerPK = stakerPK
	return &this
}

func NewBabylonScriptDataWithDefaults() *BabylonScriptData {
	this := BabylonScriptData{}
	return &this
}

func (o *BabylonScriptData) GetExplicitParams() *BabylonStakingParams {
	if o == nil {
		return nil
	}
	return o.ExplicitParams
}

func (o *BabylonScriptData) GetExplicitParamsOk() (*BabylonStakingParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExplicitParams, true
}

func (o *BabylonScriptData) SetExplicitParams(v *BabylonStakingParams) {
	o.ExplicitParams = v
}

func (o *BabylonScriptData) GetFinalityProviderPK() string {
	if o == nil {
		return ""
	}
	return o.FinalityProviderPK
}

func (o *BabylonScriptData) GetFinalityProviderPKOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.FinalityProviderPK, true
}

func (o *BabylonScriptData) SetFinalityProviderPK(v string) {
	o.FinalityProviderPK = v
}

func (o *BabylonScriptData) GetLockTime() int32 {
	if o == nil {
		return 0
	}
	return o.LockTime
}

func (o *BabylonScriptData) GetLockTimeOk() (int32, bool) {
	if o == nil {
		return 0, false
	}
	return o.LockTime, true
}

func (o *BabylonScriptData) SetLockTime(v int32) {
	o.LockTime = v
}

func (o *BabylonScriptData) GetNetwork() BabylonNetworkId {
	if o == nil {
		return ""
	}
	return o.Network
}

func (o *BabylonScriptData) GetNetworkOk() (BabylonNetworkId, bool) {
	if o == nil {
		return "", false
	}
	return o.Network, true
}

func (o *BabylonScriptData) SetNetwork(v BabylonNetworkId) {
	o.Network = v
}

func (o *BabylonScriptData) GetStakerPK() string {
	if o == nil {
		return ""
	}
	return o.StakerPK
}

func (o *BabylonScriptData) GetStakerPKOk() (string, bool) {
	if o == nil {
		return "", false
	}
	return o.StakerPK, true
}

func (o *BabylonScriptData) SetStakerPK(v string) {
	o.StakerPK = v
}

func (o *BabylonScriptData) GetVersion() *int32 {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *BabylonScriptData) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version, true
}

func (o *BabylonScriptData) SetVersion(v *int32) {
	o.Version = v
}
