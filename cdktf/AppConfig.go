// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type AppConfig struct {
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdktf.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// The directory to output Terraform resources.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Whether to skip the validation during synthesis of the app.
	// Experimental.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// Experimental.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
}

