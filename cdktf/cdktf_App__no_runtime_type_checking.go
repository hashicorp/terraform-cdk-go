//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_App) validateCrossStackReferenceParameters(fromStack TerraformStack, toStack TerraformStack, identifier *string) error {
	return nil
}

func validateApp_IsAppParameters(x interface{}) error {
	return nil
}

func validateApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApp_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAppParameters(options *AppOptions) error {
	return nil
}

