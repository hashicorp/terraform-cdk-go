// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"
)

func (m *jsiiProxy_Manifest) validateForStackParameters(stack TerraformStack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func validateNewManifestParameters(version *string, outdir *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if outdir == nil {
		return fmt.Errorf("parameter outdir is required, but nil was provided")
	}

	return nil
}

