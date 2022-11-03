//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CosBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CosBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (c *jsiiProxy_CosBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCosBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateCosBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCosBackendParameters(scope constructs.Construct, props *CosBackendProps) error {
	return nil
}

