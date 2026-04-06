// Copyright 2024, Pulumi Corporation.  All rights reserved.
package examples

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

// akeylessIntegrationOpts returns Config/Secrets so the bridged provider authenticates during
// ProgramTest. Relying only on process env is brittle (IDE test runners often omit shell exports).
func akeylessIntegrationOpts() integration.ProgramTestOptions {
	accessID := os.Getenv("AKEYLESS_ACCESS_ID")
	accessKey := os.Getenv("AKEYLESS_ACCESS_KEY")
	if accessID == "" || accessKey == "" {
		return integration.ProgramTestOptions{}
	}

	logins, err := json.Marshal([]map[string]string{
		{"accessId": accessID, "accessKey": accessKey},
	})
	if err != nil {
		return integration.ProgramTestOptions{}
	}

	out := integration.ProgramTestOptions{
		Secrets: map[string]string{
			"akeyless:apiKeyLogins": string(logins),
		},
		// Also forward explicitly; provider subprocess should inherit os.Environ, but this
		// duplicates credentials at the end of the env slice (last wins if duplicated).
		Env: []string{
			"AKEYLESS_ACCESS_ID=" + accessID,
			"AKEYLESS_ACCESS_KEY=" + accessKey,
		},
	}
	if gw := os.Getenv("AKEYLESS_GATEWAY"); gw != "" {
		out.Config = map[string]string{
			"akeyless:apiGatewayAddress": gw,
		}
		out.Env = append(out.Env, "AKEYLESS_GATEWAY="+gw)
	}
	// So Pulumi language hosts see the same path as manual runs (IDE tests often omit env).
	if p := strings.TrimSpace(os.Getenv("AKEYLESS_SMOKE_PARENT_PATH")); p != "" {
		out.Env = append(out.Env, "AKEYLESS_SMOKE_PARENT_PATH="+p)
	}
	return out
}
