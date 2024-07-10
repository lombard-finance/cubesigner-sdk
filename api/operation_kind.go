package api

import (
	"encoding/json"
	"fmt"
)

// OperationKind All different kinds of sensitive operations
type OperationKind string

// List of OperationKind
const (
	OperationKindAvaSign        OperationKind = "AvaSign"
	OperationKindAvaChainTxSign OperationKind = "AvaChainTxSign"
	OperationKindBabylonStaking OperationKind = "BabylonStaking"
	OperationKindBlobSign       OperationKind = "BlobSign"
	OperationKindBtcSign        OperationKind = "BtcSign"
	OperationKindTaprootSign    OperationKind = "TaprootSign"
	OperationKindEip191Sign     OperationKind = "Eip191Sign"
	OperationKindEip712Sign     OperationKind = "Eip712Sign"
	OperationKindEotsNonces     OperationKind = "EotsNonces"
	OperationKindEotsSign       OperationKind = "EotsSign"
	OperationKindEth1Sign       OperationKind = "Eth1Sign"
	OperationKindEth2Sign       OperationKind = "Eth2Sign"
	OperationKindEth2Stake      OperationKind = "Eth2Stake"
	OperationKindEth2Unstake    OperationKind = "Eth2Unstake"
	OperationKindSolanaSign     OperationKind = "SolanaSign"
	OperationKindRoleUpdate     OperationKind = "RoleUpdate"
)

// All allowed values of OperationKind enum
var AllowedOperationKindEnumValues = []OperationKind{
	"AvaSign",
	"AvaChainTxSign",
	"BabylonStaking",
	"BlobSign",
	"BtcSign",
	"TaprootSign",
	"Eip191Sign",
	"Eip712Sign",
	"EotsNonces",
	"EotsSign",
	"Eth1Sign",
	"Eth2Sign",
	"Eth2Stake",
	"Eth2Unstake",
	"SolanaSign",
	"RoleUpdate",
}

func (v *OperationKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationKind(value)
	for _, existing := range AllowedOperationKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationKind", value)
}

// NewOperationKindFromValue returns a pointer to a valid OperationKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationKindFromValue(v string) (*OperationKind, error) {
	ev := OperationKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationKind: valid values are %v", v, AllowedOperationKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationKind) IsValid() bool {
	for _, existing := range AllowedOperationKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationKind value
func (v OperationKind) Ptr() *OperationKind {
	return &v
}

type NullableOperationKind struct {
	value *OperationKind
	isSet bool
}

func (v NullableOperationKind) Get() *OperationKind {
	return v.value
}

func (v *NullableOperationKind) Set(val *OperationKind) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationKind) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationKind(val *OperationKind) *NullableOperationKind {
	return &NullableOperationKind{value: val, isSet: true}
}

func (v NullableOperationKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
