//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzurermBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AzurermBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (a *jsiiProxy_AzurermBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAzurermBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateAzurermBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAzurermBackendParameters(scope constructs.Construct, props *AzurermBackendProps) error {
	return nil
}

