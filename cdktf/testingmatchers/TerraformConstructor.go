// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testingmatchers


// Experimental.
type TerraformConstructor struct {
	// Experimental.
	TfResourceType *string `field:"required" json:"tfResourceType" yaml:"tfResourceType"`
}

