// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OssBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_OssBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (o *jsiiProxy_OssBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateOssBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateOssBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOssBackend_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewOssBackendParameters(scope constructs.Construct, props *OssBackendConfig) error {
	return nil
}

