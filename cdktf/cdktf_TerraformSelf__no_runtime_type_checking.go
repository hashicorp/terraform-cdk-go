//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func validateTerraformSelf_GetAnyParameters(key *string) error {
	return nil
}

func validateTerraformSelf_GetNumberParameters(key *string) error {
	return nil
}

func validateTerraformSelf_GetStringParameters(key *string) error {
	return nil
}

