// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_PrefixedRemoteWorkspaces) validateSetPrefixParameters(val *string) error {
	return nil
}

func validateNewPrefixedRemoteWorkspacesParameters(prefix *string) error {
	return nil
}

