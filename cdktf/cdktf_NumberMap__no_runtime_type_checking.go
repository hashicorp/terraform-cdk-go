//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NumberMap) validateLookupParameters(key *string) error {
	return nil
}

func (n *jsiiProxy_NumberMap) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NumberMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NumberMap) validateSetTerraformResourceParameters(val IInterpolatingParent) error {
	return nil
}

func validateNewNumberMapParameters(terraformResource IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

