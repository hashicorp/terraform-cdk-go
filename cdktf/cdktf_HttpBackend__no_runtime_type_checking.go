//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HttpBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (h *jsiiProxy_HttpBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHttpBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateHttpBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHttpBackendParameters(scope constructs.Construct, props *HttpBackendProps) error {
	return nil
}

