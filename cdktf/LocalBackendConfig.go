// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// The local backend stores state on the local filesystem, locks that state using system APIs, and performs operations locally.
//
// Read more about this backend in the Terraform docs:
// https://developer.hashicorp.com/terraform/language/settings/backends/local
// Experimental.
type LocalBackendConfig struct {
	// Path where the state file is stored.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// (Optional) The path to non-default workspaces.
	// Experimental.
	WorkspaceDir *string `field:"optional" json:"workspaceDir" yaml:"workspaceDir"`
}

