//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GcsBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GcsBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (g *jsiiProxy_GcsBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGcsBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateGcsBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGcsBackendParameters(scope constructs.Construct, props *GcsBackendProps) error {
	return nil
}

