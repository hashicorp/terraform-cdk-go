// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformStackMetadata struct {
	// Experimental.
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Experimental.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Experimental.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
}

