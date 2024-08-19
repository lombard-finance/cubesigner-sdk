package v0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMfaConfig(t *testing.T) {
	tests := []struct {
		name         string
		jsonData     string
		expectEmpty  bool
		expectOneOf  bool
		expectOneOf1 bool
	}{
		{
			name: "Empty Mfa Array",
			jsonData: `{
				"email": "infra@appliedstudios.com",
				"mfa": [],
				"name": "Applied",
				"org_ids": [
					"Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c"
				],
				"orgs": [
					{
					"membership": "Member",
					"org_id": "Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c",
					"status": "enabled"
					}
				],
				"user_id": "User#f319b4e8-2853-4e21-ac12-4984d13cf51f"
				}`,
			expectEmpty: true,
		},
		{
			name: "Mfa with ConfiguredMfaOneOf",
			jsonData: `{
				"email": "infra@appliedstudios.com",
				"mfa": [
					{
						"type": "fido",
						"name": "Default Key",
					}
				],
				"name": "Applied",
				"org_ids": [
					"Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c"
				],
				"orgs": [
					{
					"membership": "Member",
					"org_id": "Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c",
					"status": "enabled"
					}
				],
				"user_id": "User#f319b4e8-2853-4e21-ac12-4984d13cf51f"
				}`,
			expectOneOf: true,
		},
		{
			name: "Mfa with ConfiguredMfaOneOf1",
			jsonData: `{
				"email": "infra@appliedstudios.com",
				"mfa": [
					{
					"type": "totp"
					}
				],
				"name": "Applied",
				"org_ids": [
					"Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c"
				],
				"orgs": [
					{
					"membership": "Member",
					"org_id": "Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c",
					"status": "enabled"
					}
				],
				"user_id": "User#f319b4e8-2853-4e21-ac12-4984d13cf51f"
				}`,
			expectOneOf1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result AboutMeLegacy200Response
			err := json.Unmarshal([]byte(tt.jsonData), &result)
			assert.NoError(t, err, "Unmarshal should not return an error")
		})
	}
}
