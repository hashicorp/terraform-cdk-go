// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdktf


// Experimental.
type AppConfig struct {
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdktf.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Default: - no additional context.
	//
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Experimental.
	HclOutput *bool `field:"optional" json:"hclOutput" yaml:"hclOutput"`
	// The directory to output Terraform resources.
	// Default: - CDKTF_OUTDIR if defined, otherwise "cdktf.out"
	//
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Whether to skip backend validation during synthesis of the app.
	// Default: - false.
	//
	// Experimental.
	SkipBackendValidation *bool `field:"optional" json:"skipBackendValidation" yaml:"skipBackendValidation"`
	// Whether to skip all validations during synthesis of the app.
	// Default: - false.
	//
	// Experimental.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// Experimental.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
}

