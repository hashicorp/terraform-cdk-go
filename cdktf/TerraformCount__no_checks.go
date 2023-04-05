// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func validateTerraformCount_IsTerraformCountParameters(x interface{}) error {
	return nil
}

func validateTerraformCount_OfParameters(count *float64) error {
	return nil
}

