// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ArtifactoryBackend) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ArtifactoryBackend) validateGetRemoteStateDataSourceParameters(scope constructs.Construct, name *string, _fromStack *string) error {
	return nil
}

func (a *jsiiProxy_ArtifactoryBackend) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateArtifactoryBackend_IsBackendParameters(x interface{}) error {
	return nil
}

func validateArtifactoryBackend_IsConstructParameters(x interface{}) error {
	return nil
}

func validateArtifactoryBackend_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewArtifactoryBackendParameters(scope constructs.Construct, props *ArtifactoryBackendConfig) error {
	return nil
}

