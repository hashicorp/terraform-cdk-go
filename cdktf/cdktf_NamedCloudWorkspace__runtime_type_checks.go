//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"
)

func (j *jsiiProxy_NamedCloudWorkspace) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNamedCloudWorkspaceParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

