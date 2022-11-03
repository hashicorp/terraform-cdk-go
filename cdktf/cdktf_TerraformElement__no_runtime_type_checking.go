//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TerraformElement) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TerraformElement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTerraformElement_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTerraformElementParameters(scope constructs.Construct, id *string) error {
	return nil
}

