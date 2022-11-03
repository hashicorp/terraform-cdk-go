//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PgBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PgBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (p *jsiiProxy_PgBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validatePgBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validatePgBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPgBackendParameters(scope constructs.Construct, props *PgBackendProps) error {
	return nil
}

