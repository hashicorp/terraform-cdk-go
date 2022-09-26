//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Aspects) validateAddParameters(aspect IAspect) error {
	return nil
}

func validateAspects_OfParameters(scope constructs.IConstruct) error {
	return nil
}

