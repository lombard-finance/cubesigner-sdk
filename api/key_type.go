package api

import (
	"encoding/json"
	"fmt"
)

// KeyType the model 'KeyType'
type KeyType string

// List of KeyType
const (
	SECP_ETH_ADDR           KeyType = "SecpEthAddr"
	SECP_BTC                KeyType = "SecpBtc"
	SECP_BTC_TEST           KeyType = "SecpBtcTest"
	SECP_AVA_ADDR           KeyType = "SecpAvaAddr"
	SECP_AVA_TEST_ADDR      KeyType = "SecpAvaTestAddr"
	BLS_PUB                 KeyType = "BlsPub"
	BLS_INACTIVE            KeyType = "BlsInactive"
	ED25519_SOLANA_ADDR     KeyType = "Ed25519SolanaAddr"
	ED25519_SUI_ADDR        KeyType = "Ed25519SuiAddr"
	ED25519_APTOS_ADDR      KeyType = "Ed25519AptosAddr"
	ED25519_CARDANO_ADDR_VK KeyType = "Ed25519CardanoAddrVk"
	ED25519_STELLAR_ADDR    KeyType = "Ed25519StellarAddr"
	MNEMONIC                KeyType = "Mnemonic"
	STARK                   KeyType = "Stark"
	BABYLON_EOTS            KeyType = "BabylonEots"
	TAPROOT_BTC             KeyType = "TaprootBtc"
	TAPROOT_BTC_TEST        KeyType = "TaprootBtcTest"
)

// All allowed values of KeyType enum
var AllowedKeyTypeEnumValues = []KeyType{
	"SecpEthAddr",
	"SecpBtc",
	"SecpBtcTest",
	"SecpAvaAddr",
	"SecpAvaTestAddr",
	"BlsPub",
	"BlsInactive",
	"Ed25519SolanaAddr",
	"Ed25519SuiAddr",
	"Ed25519AptosAddr",
	"Ed25519CardanoAddrVk",
	"Ed25519StellarAddr",
	"Mnemonic",
	"Stark",
	"BabylonEots",
	"TaprootBtc",
	"TaprootBtcTest",
}

func (v *KeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyType(value)
	for _, existing := range AllowedKeyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyType", value)
}

// NewKeyTypeFromValue returns a pointer to a valid KeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyTypeFromValue(v string) (*KeyType, error) {
	ev := KeyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyType: valid values are %v", v, AllowedKeyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyType) IsValid() bool {
	for _, existing := range AllowedKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyType value
func (v KeyType) Ptr() *KeyType {
	return &v
}

type NullableKeyType struct {
	value *KeyType
	isSet bool
}

func (v NullableKeyType) Get() *KeyType {
	return v.value
}

func (v *NullableKeyType) Set(val *KeyType) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyType) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyType(val *KeyType) *NullableKeyType {
	return &NullableKeyType{value: val, isSet: true}
}

func (v NullableKeyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
