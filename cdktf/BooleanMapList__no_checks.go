// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BooleanMapList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BooleanMapList) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (b *jsiiProxy_BooleanMapList) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BooleanMapList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BooleanMapList) validateSetTerraformResourceParameters(val IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BooleanMapList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBooleanMapListParameters(terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

