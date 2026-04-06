// Copyright 2024, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestBasicTs(t *testing.T) {
	if os.Getenv("AKEYLESS_ACCESS_ID") == "" || os.Getenv("AKEYLESS_ACCESS_KEY") == "" {
		t.Skip("set AKEYLESS_ACCESS_ID and AKEYLESS_ACCESS_KEY to run this integration test")
	}

	opts := getJSBaseOptions(t).
		With(akeylessIntegrationOpts()).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "basic-ts"),
		})

	integration.ProgramTest(t, &opts)
}
