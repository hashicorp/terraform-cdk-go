// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformResourceLifecycle struct {
	// Experimental.
	CreateBeforeDestroy *bool `field:"optional" json:"createBeforeDestroy" yaml:"createBeforeDestroy"`
	// Experimental.
	IgnoreChanges interface{} `field:"optional" json:"ignoreChanges" yaml:"ignoreChanges"`
	// Experimental.
	PreventDestroy *bool `field:"optional" json:"preventDestroy" yaml:"preventDestroy"`
}

