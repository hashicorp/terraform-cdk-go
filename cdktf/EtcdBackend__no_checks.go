// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EtcdBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_EtcdBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (e *jsiiProxy_EtcdBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateEtcdBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateEtcdBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEtcdBackend_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewEtcdBackendParameters(scope constructs.Construct, props *EtcdBackendConfig) error {
	return nil
}

