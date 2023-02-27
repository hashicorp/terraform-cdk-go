// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformVariableValidationConfig struct {
	// Experimental.
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// Experimental.
	ErrorMessage *string `field:"required" json:"errorMessage" yaml:"errorMessage"`
}

