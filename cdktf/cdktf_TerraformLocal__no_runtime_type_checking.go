//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TerraformLocal) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TerraformLocal) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTerraformLocal_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TerraformLocal) validateSetExpressionParameters(val interface{}) error {
	return nil
}

func validateNewTerraformLocalParameters(scope constructs.Construct, id *string, expression interface{}) error {
	return nil
}

