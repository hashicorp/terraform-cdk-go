// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EtcdV3Backend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_EtcdV3Backend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (e *jsiiProxy_EtcdV3Backend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateEtcdV3Backend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateEtcdV3Backend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEtcdV3Backend_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewEtcdV3BackendParameters(scope constructs.Construct, props *EtcdV3BackendConfig) error {
	return nil
}

