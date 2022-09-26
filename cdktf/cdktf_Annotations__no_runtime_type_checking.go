//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Annotations) validateAddErrorParameters(message *string) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateAddInfoParameters(message *string) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateAddWarningParameters(message *string) error {
	return nil
}

func validateAnnotations_OfParameters(scope constructs.IConstruct) error {
	return nil
}

