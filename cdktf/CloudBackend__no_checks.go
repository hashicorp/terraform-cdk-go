// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CloudBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (c *jsiiProxy_CloudBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCloudBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateCloudBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudBackend_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewCloudBackendParameters(scope constructs.Construct, props *CloudBackendConfig) error {
	return nil
}

