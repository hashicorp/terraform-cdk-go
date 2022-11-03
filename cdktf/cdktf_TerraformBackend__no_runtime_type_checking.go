//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TerraformBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TerraformBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, fromStack *string) error {
	return nil
}

func (t *jsiiProxy_TerraformBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTerraformBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateTerraformBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTerraformBackendParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

