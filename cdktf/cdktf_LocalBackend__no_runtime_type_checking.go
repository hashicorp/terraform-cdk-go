//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocalBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LocalBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, fromStack *string) error {
	return nil
}

func (l *jsiiProxy_LocalBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateLocalBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateLocalBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLocalBackendParameters(scope constructs.Construct, props *LocalBackendProps) error {
	return nil
}

