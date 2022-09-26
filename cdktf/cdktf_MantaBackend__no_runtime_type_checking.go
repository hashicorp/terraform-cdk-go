//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MantaBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_MantaBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (m *jsiiProxy_MantaBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateMantaBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateMantaBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMantaBackendParameters(scope constructs.Construct, props *MantaBackendProps) error {
	return nil
}

