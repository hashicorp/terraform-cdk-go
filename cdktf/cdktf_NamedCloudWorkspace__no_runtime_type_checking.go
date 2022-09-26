//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_NamedCloudWorkspace) validateSetNameParameters(val *string) error {
	return nil
}

func validateNewNamedCloudWorkspaceParameters(name *string) error {
	return nil
}

