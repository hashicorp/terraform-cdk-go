//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RemoteBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_RemoteBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (r *jsiiProxy_RemoteBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateRemoteBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateRemoteBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRemoteBackendParameters(scope constructs.Construct, props *RemoteBackendProps) error {
	return nil
}

