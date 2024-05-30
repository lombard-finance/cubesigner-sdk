package v0

import (
	"encoding/json"
)

// AddKeysToRole200Rsponse struct for AddKeysToRole200Rsponse
type AddKeysToRole200Rsponse struct {
	// All keys in a role
	KeysIDs []string `json:"key_ids,omitempty"`
	// Policies that are checked before key in role is used
	Policy []map[string]interface{} `json:"policy,omitempty"`
	Status string                   `json:"status,omitempty"`
}

func (o *AddKeysToRole200Rsponse) GetKeyIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KeysIDs
}

func (o *AddKeysToRole200Rsponse) GetKeyIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeysIDs, true
}

func (o *AddKeysToRole200Rsponse) SetKeyIDs(v []string) {
	o.KeysIDs = v
}

// GetPolicy returns the Policy field value
func (o *AddKeysToRole200Rsponse) GetPolicy() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *AddKeysToRole200Rsponse) GetPolicyOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy, true
}

// SetPolicy sets field value
func (o *AddKeysToRole200Rsponse) SetPolicy(v []map[string]interface{}) {
	o.Policy = v
}

func (o AddKeysToRole200Rsponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key_ids"] = o.KeysIDs
	toSerialize["policy"] = o.Policy
	return json.Marshal(toSerialize)
}
