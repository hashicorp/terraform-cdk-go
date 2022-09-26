//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TerraformHclModule) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateGetParameters(output *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateGetBooleanParameters(output *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateGetListParameters(output *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateGetNumberParameters(output *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateGetStringParameters(output *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (t *jsiiProxy_TerraformHclModule) validateSetParameters(variable *string, value interface{}) error {
	return nil
}

func validateTerraformHclModule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTerraformHclModuleParameters(scope constructs.Construct, id *string, options *TerraformHclModuleOptions) error {
	return nil
}

