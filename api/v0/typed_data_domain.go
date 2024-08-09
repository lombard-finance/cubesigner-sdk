package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// TypedDataDomain Represents the [EIP-712](https://eips.ethereum.org/EIPS/eip-712) EIP712Domain object.  EIP712Domain is a JSON object with one or more of the below fields. Protocol designers only need to include the fields that make sense for their signing domain.
type TypedDataDomain struct {
	// The EIP-155 chain id. The user-agent should refuse signing if it does not match the currently active chain.
	ChainId api.NullableString `json:"chainId,omitempty"`
	// The user readable name of signing domain, i.e., the name of the DApp or the protocol.
	Name api.NullableString   `json:"name,omitempty"`
	Salt *TypedDataDomainSalt `json:"salt,omitempty"`
	// The address of the contract that will verify the signature.
	VerifyingContract api.NullableString `json:"verifyingContract,omitempty"`
	// The current major version of the signing domain. Signatures from different versions are not compatible.
	Version api.NullableString `json:"version,omitempty"`
}

// NewTypedDataDomain instantiates a new TypedDataDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypedDataDomain() *TypedDataDomain {
	this := TypedDataDomain{}
	return &this
}

// NewTypedDataDomainWithDefaults instantiates a new TypedDataDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypedDataDomainWithDefaults() *TypedDataDomain {
	this := TypedDataDomain{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypedDataDomain) GetChainId() string {
	if o == nil || o.ChainId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypedDataDomain) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *TypedDataDomain) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableString and assigns it to the ChainId field.
func (o *TypedDataDomain) SetChainId(v string) {
	o.ChainId.Set(&v)
}

// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *TypedDataDomain) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *TypedDataDomain) UnsetChainId() {
	o.ChainId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypedDataDomain) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypedDataDomain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TypedDataDomain) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TypedDataDomain) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TypedDataDomain) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TypedDataDomain) UnsetName() {
	o.Name.Unset()
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *TypedDataDomain) GetSalt() TypedDataDomainSalt {
	if o == nil || o.Salt == nil {
		var ret TypedDataDomainSalt
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypedDataDomain) GetSaltOk() (*TypedDataDomainSalt, bool) {
	if o == nil || o.Salt == nil {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *TypedDataDomain) HasSalt() bool {
	if o != nil && o.Salt != nil {
		return true
	}

	return false
}

// SetSalt gets a reference to the given TypedDataDomainSalt and assigns it to the Salt field.
func (o *TypedDataDomain) SetSalt(v TypedDataDomainSalt) {
	o.Salt = &v
}

// GetVerifyingContract returns the VerifyingContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypedDataDomain) GetVerifyingContract() string {
	if o == nil || o.VerifyingContract.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerifyingContract.Get()
}

// GetVerifyingContractOk returns a tuple with the VerifyingContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypedDataDomain) GetVerifyingContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerifyingContract.Get(), o.VerifyingContract.IsSet()
}

// HasVerifyingContract returns a boolean if a field has been set.
func (o *TypedDataDomain) HasVerifyingContract() bool {
	if o != nil && o.VerifyingContract.IsSet() {
		return true
	}

	return false
}

// SetVerifyingContract gets a reference to the given NullableString and assigns it to the VerifyingContract field.
func (o *TypedDataDomain) SetVerifyingContract(v string) {
	o.VerifyingContract.Set(&v)
}

// SetVerifyingContractNil sets the value for VerifyingContract to be an explicit nil
func (o *TypedDataDomain) SetVerifyingContractNil() {
	o.VerifyingContract.Set(nil)
}

// UnsetVerifyingContract ensures that no value is present for VerifyingContract, not even an explicit nil
func (o *TypedDataDomain) UnsetVerifyingContract() {
	o.VerifyingContract.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypedDataDomain) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypedDataDomain) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *TypedDataDomain) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *TypedDataDomain) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *TypedDataDomain) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *TypedDataDomain) UnsetVersion() {
	o.Version.Unset()
}

func (o TypedDataDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chainId"] = o.ChainId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Salt != nil {
		toSerialize["salt"] = o.Salt
	}
	if o.VerifyingContract.IsSet() {
		toSerialize["verifyingContract"] = o.VerifyingContract.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTypedDataDomain struct {
	value *TypedDataDomain
	isSet bool
}

func (v NullableTypedDataDomain) Get() *TypedDataDomain {
	return v.value
}

func (v *NullableTypedDataDomain) Set(val *TypedDataDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedDataDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedDataDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedDataDomain(val *TypedDataDomain) *NullableTypedDataDomain {
	return &NullableTypedDataDomain{value: val, isSet: true}
}

func (v NullableTypedDataDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedDataDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
