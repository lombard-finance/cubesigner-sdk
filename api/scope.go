package api

import (
	"encoding/json"
	"fmt"
)

// Scope the model 'Scope'
type Scope string

// List of Scope
const (
	SIGN                              Scope = "sign:*"
	SIGNAVA                           Scope = "sign:ava"
	SIGNBLOB                          Scope = "sign:blob"
	SIGNBTC                           Scope = "sign:btc:*"
	SIGNBTCSEGWIT                     Scope = "sign:btc:segwit"
	SIGNBTCTAPROOT                    Scope = "sign:btc:taproot"
	SIGNBABYLON                       Scope = "sign:babylon:*"
	SIGNBABYLONEOTS                   Scope = "sign:babylon:eots:*"
	SIGNBABYLONEOTSNONCES             Scope = "sign:babylon:eots:nonces"
	SIGNBABYLONEOTSSIGN               Scope = "sign:babylon:eots:sign"
	SIGNBABYLONSTAKING                Scope = "sign:babylon:staking:*"
	SIGNBABYLONSTAKINGDEPOSIT         Scope = "sign:babylon:staking:deposit"
	SIGNBABYLONSTAKINGUNBOND          Scope = "sign:babylon:staking:unbond"
	SIGNBABYLONSTAKINGWITHDRAW        Scope = "sign:babylon:staking:withdraw"
	SIGNEVM                           Scope = "sign:evm:*"
	SIGNEVMTX                         Scope = "sign:evm:tx"
	SIGNEVMEIP191                     Scope = "sign:evm:eip191"
	SIGNEVMEIP712                     Scope = "sign:evm:eip712"
	SIGNETH2                          Scope = "sign:eth2:*"
	SIGNETH2VALIDATE                  Scope = "sign:eth2:validate"
	SIGNETH2STAKE                     Scope = "sign:eth2:stake"
	SIGNETH2UNSTAKE                   Scope = "sign:eth2:unstake"
	SIGNSOLANA                        Scope = "sign:solana"
	MANAGE                            Scope = "manage:*"
	MANAGEEMAIL                       Scope = "manage:email"
	MANAGEMFA                         Scope = "manage:mfa:*"
	MANAGEMFALIST                     Scope = "manage:mfa:list"
	MANAGEMFAVOTE                     Scope = "manage:mfa:vote:*"
	MANAGEMFAVOTECS                   Scope = "manage:mfa:vote:cs"
	MANAGEMFAVOTEFIDO                 Scope = "manage:mfa:vote:fido"
	MANAGEMFAVOTETOTP                 Scope = "manage:mfa:vote:totp"
	MANAGEMFAREGISTER                 Scope = "manage:mfa:register:*"
	MANAGEMFAREGISTERFIDO             Scope = "manage:mfa:register:fido"
	MANAGEMFAREGISTERTOTP             Scope = "manage:mfa:register:totp"
	MANAGEMFAUNREGISTER               Scope = "manage:mfa:unregister:*"
	MANAGEMFAUNREGISTERFIDO           Scope = "manage:mfa:unregister:fido"
	MANAGEMFAUNREGISTERTOTP           Scope = "manage:mfa:unregister:totp"
	MANAGEMFAVERIFY                   Scope = "manage:mfa:verify:*"
	MANAGEMFAVERIFYTOTP               Scope = "manage:mfa:verify:totp"
	MANAGEKEY                         Scope = "manage:key:*"
	MANAGEKEYGET                      Scope = "manage:key:get"
	MANAGEKEYLIST_ROLES               Scope = "manage:key:listRoles"
	MANAGEKEYLIST                     Scope = "manage:key:list"
	MANAGEKEYCREATE                   Scope = "manage:key:create"
	MANAGEKEYIMPORT                   Scope = "manage:key:import"
	MANAGEKEYUPDATE                   Scope = "manage:key:update:*"
	MANAGEKEYUPDATEOWNER              Scope = "manage:key:update:owner"
	MANAGEKEYUPDATEPOLICY             Scope = "manage:key:update:policy"
	MANAGEKEYUPDATEENABLED            Scope = "manage:key:update:enabled"
	MANAGEKEYUPDATEMETADATA           Scope = "manage:key:update:metadata"
	MANAGEKEYUPDATEPOLICY_ON_UPDATES  Scope = "manage:key:update:policyOnUpdates"
	MANAGEKEYDELETE                   Scope = "manage:key:delete"
	MANAGEROLE                        Scope = "manage:role:*"
	MANAGEROLECREATE                  Scope = "manage:role:create"
	MANAGEROLEDELETE                  Scope = "manage:role:delete"
	MANAGEROLEGET                     Scope = "manage:role:get:*"
	MANAGEROLEGETKEYS                 Scope = "manage:role:get:keys"
	MANAGEROLEGETUSERS                Scope = "manage:role:get:users"
	MANAGEROLELIST                    Scope = "manage:role:list"
	MANAGEROLEUPDATE                  Scope = "manage:role:update:*"
	MANAGEROLEUPDATEENABLED           Scope = "manage:role:update:enabled"
	MANAGEROLEUPDATEPOLICY            Scope = "manage:role:update:policy"
	MANAGEROLEUPDATEPOLICY_ON_UPDATES Scope = "manage:role:update:policyOnUpdates"
	MANAGEROLEUPDATEKEYADD            Scope = "manage:role:update:key:add"
	MANAGEROLEUPDATEKEYREMOVE         Scope = "manage:role:update:key:remove"
	MANAGEROLEUPDATEUSERADD           Scope = "manage:role:update:user:add"
	MANAGEROLEUPDATEUSERREMOVE        Scope = "manage:role:update:user:remove"
	MANAGEIDENTITY                    Scope = "manage:identity:*"
	MANAGEIDENTITYVERIFY              Scope = "manage:identity:verify"
	MANAGEIDENTITYADD                 Scope = "manage:identity:add"
	MANAGEIDENTITYREMOVE              Scope = "manage:identity:remove"
	MANAGEIDENTITYLIST                Scope = "manage:identity:list"
	MANAGEORG                         Scope = "manage:org:*"
	MANAGEORGADD_USER                 Scope = "manage:org:addUser"
	MANAGEORGINVITE_USER              Scope = "manage:org:inviteUser"
	MANAGEORGUPDATE_MEMBERSHIP        Scope = "manage:org:updateMembership"
	MANAGEORGLIST_USERS               Scope = "manage:org:listUsers"
	MANAGEORGDELETE_USER              Scope = "manage:org:deleteUser"
	MANAGEORGGET                      Scope = "manage:org:get"
	MANAGESESSION                     Scope = "manage:session:*"
	MANAGESESSIONCREATE               Scope = "manage:session:create"
	MANAGEEXPORT                      Scope = "manage:export:*"
	MANAGEEXPORTUSER                  Scope = "manage:export:user:*"
	MANAGEEXPORTUSERDELETE            Scope = "manage:export:user:delete"
	MANAGEEXPORTUSERLIST              Scope = "manage:export:user:list"
	EXPORT                            Scope = "export:*"
	EXPORTUSER                        Scope = "export:user:*"
	EXPORTUSERINIT                    Scope = "export:user:init"
	EXPORTUSERCOMPLETE                Scope = "export:user:complete"
	MMI                               Scope = "mmi:*"
)

// All allowed values of Scope enum
var AllowedScopeEnumValues = []Scope{
	"sign:*",
	"sign:ava",
	"sign:blob",
	"sign:btc:*",
	"sign:btc:segwit",
	"sign:btc:taproot",
	"sign:babylon:*",
	"sign:babylon:eots:*",
	"sign:babylon:eots:nonces",
	"sign:babylon:eots:sign",
	"sign:evm:*",
	"sign:evm:tx",
	"sign:evm:eip191",
	"sign:evm:eip712",
	"sign:eth2:*",
	"sign:eth2:validate",
	"sign:eth2:stake",
	"sign:eth2:unstake",
	"sign:solana",
	"manage:*",
	"manage:email",
	"manage:mfa:*",
	"manage:mfa:list",
	"manage:mfa:vote:*",
	"manage:mfa:vote:cs",
	"manage:mfa:vote:fido",
	"manage:mfa:vote:totp",
	"manage:mfa:register:*",
	"manage:mfa:register:fido",
	"manage:mfa:register:totp",
	"manage:mfa:unregister:*",
	"manage:mfa:unregister:fido",
	"manage:mfa:unregister:totp",
	"manage:mfa:verify:*",
	"manage:mfa:verify:totp",
	"manage:key:*",
	"manage:key:get",
	"manage:key:listRoles",
	"manage:key:list",
	"manage:key:create",
	"manage:key:import",
	"manage:key:update:*",
	"manage:key:update:owner",
	"manage:key:update:policy",
	"manage:key:update:enabled",
	"manage:key:update:metadata",
	"manage:key:update:policyOnUpdates",
	"manage:key:delete",
	"manage:role:*",
	"manage:role:create",
	"manage:role:delete",
	"manage:role:get:*",
	"manage:role:get:keys",
	"manage:role:get:users",
	"manage:role:list",
	"manage:role:update:*",
	"manage:role:update:enabled",
	"manage:role:update:policy",
	"manage:role:update:policyOnUpdates",
	"manage:role:update:key:add",
	"manage:role:update:key:remove",
	"manage:role:update:user:add",
	"manage:role:update:user:remove",
	"manage:identity:*",
	"manage:identity:verify",
	"manage:identity:add",
	"manage:identity:remove",
	"manage:identity:list",
	"manage:org:*",
	"manage:org:addUser",
	"manage:org:inviteUser",
	"manage:org:updateMembership",
	"manage:org:listUsers",
	"manage:org:deleteUser",
	"manage:org:get",
	"manage:session:*",
	"manage:session:create",
	"manage:export:*",
	"manage:export:user:*",
	"manage:export:user:delete",
	"manage:export:user:list",
	"export:*",
	"export:user:*",
	"export:user:init",
	"export:user:complete",
	"mmi:*",
}

func (v *Scope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Scope(value)
	for _, existing := range AllowedScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Scope", value)
}

// NewScopeFromValue returns a pointer to a valid Scope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopeFromValue(v string) (*Scope, error) {
	ev := Scope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Scope: valid values are %v", v, AllowedScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Scope) IsValid() bool {
	for _, existing := range AllowedScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Scope value
func (v Scope) Ptr() *Scope {
	return &v
}

type NullableScope struct {
	value *Scope
	isSet bool
}

func (v NullableScope) Get() *Scope {
	return v.value
}

func (v *NullableScope) Set(val *Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope(val *Scope) *NullableScope {
	return &NullableScope{value: val, isSet: true}
}

func (v NullableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
