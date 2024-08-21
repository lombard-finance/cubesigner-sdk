package v0

import (
	"encoding/json"
	"fmt"

	"github.com/lombard-finance/cubesigner-sdk/api"
)

// ConfiguredMfa - struct for ConfiguredMfa
type ConfiguredMfa struct {
	ConfiguredMfaOneOf  *ConfiguredMfaOneOf
	ConfiguredMfaOneOf1 *ConfiguredMfaOneOf1
}

// ConfiguredMfaOneOfAsConfiguredMfa is a convenience function that returns ConfiguredMfaOneOf wrapped in ConfiguredMfa
func ConfiguredMfaOneOfAsConfiguredMfa(v *ConfiguredMfaOneOf) ConfiguredMfa {
	return ConfiguredMfa{
		ConfiguredMfaOneOf: v,
	}
}

// ConfiguredMfaOneOf1AsConfiguredMfa is a convenience function that returns ConfiguredMfaOneOf1 wrapped in ConfiguredMfa
func ConfiguredMfaOneOf1AsConfiguredMfa(v *ConfiguredMfaOneOf1) ConfiguredMfa {
	return ConfiguredMfa{
		ConfiguredMfaOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConfiguredMfa) UnmarshalJSON(data []byte) error {
	var err error
	match := 0

	// Define a temporary struct to check the value of the "type" field
	var temp map[string]interface{}
	if err = json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Check the value of the "type" field
	if mfaType, ok := temp["type"].(string); ok {
		switch mfaType {
		case "totp", "":
			// If type is "totp" or empty, it should match ConfiguredMfaOneOf
			err = api.NewStrictDecoder(data).Decode(&dst.ConfiguredMfaOneOf)
			if err == nil {
				match++
			} else {
				dst.ConfiguredMfaOneOf = nil
			}
		case "fido":
			// If type is "fido", it should match ConfiguredMfaOneOf1
			err = api.NewStrictDecoder(data).Decode(&dst.ConfiguredMfaOneOf1)
			if err == nil {
				match++
			} else {
				dst.ConfiguredMfaOneOf1 = nil
			}
		default:
			// Handle unknown types, if necessary
			return fmt.Errorf("unknown type: %s", mfaType)
		}
	} else {
		// If "type" is missing, treat it as an empty string
		err = api.NewStrictDecoder(data).Decode(&dst.ConfiguredMfaOneOf)
		if err == nil {
			match++
		} else {
			dst.ConfiguredMfaOneOf = nil
		}
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ConfiguredMfaOneOf = nil
		dst.ConfiguredMfaOneOf1 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ConfiguredMfa)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ConfiguredMfa)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConfiguredMfa) MarshalJSON() ([]byte, error) {
	if src.ConfiguredMfaOneOf != nil {
		return json.Marshal(&src.ConfiguredMfaOneOf)
	}

	if src.ConfiguredMfaOneOf1 != nil {
		return json.Marshal(&src.ConfiguredMfaOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConfiguredMfa) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ConfiguredMfaOneOf != nil {
		return obj.ConfiguredMfaOneOf
	}

	if obj.ConfiguredMfaOneOf1 != nil {
		return obj.ConfiguredMfaOneOf1
	}

	// all schemas are nil
	return nil
}

//
//type NullableConfiguredMfa struct {
//	value *ConfiguredMfa
//	isSet bool
//}
//
//func (v NullableConfiguredMfa) Get() *ConfiguredMfa {
//	return v.value
//}
//
//func (v *NullableConfiguredMfa) Set(val *ConfiguredMfa) {
//	v.value = val
//	v.isSet = true
//}
//
//func (v NullableConfiguredMfa) IsSet() bool {
//	return v.isSet
//}
//
//func (v *NullableConfiguredMfa) Unset() {
//	v.value = nil
//	v.isSet = false
//}
//
//func NewNullableConfiguredMfa(val *ConfiguredMfa) *NullableConfiguredMfa {
//	return &NullableConfiguredMfa{value: val, isSet: true}
//}
//
//func (v NullableConfiguredMfa) MarshalJSON() ([]byte, error) {
//	return json.Marshal(v.value)
//}
//
//func (v *NullableConfiguredMfa) UnmarshalJSON(src []byte) error {
//	v.isSet = true
//	return json.Unmarshal(src, &v.value)
//}
