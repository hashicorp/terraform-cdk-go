//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"
)

func validateTerraformSelf_GetAnyParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateTerraformSelf_GetNumberParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateTerraformSelf_GetStringParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

