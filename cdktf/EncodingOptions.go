// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Properties to string encodings.
// Experimental.
type EncodingOptions struct {
	// A hint for the Token's purpose when stringifying it.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

