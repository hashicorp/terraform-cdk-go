// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package testingmatchers

// Building without runtime type checking enabled, so all the below just return nil

func validateNewAssertionReturnParameters(message *string, pass *bool) error {
	return nil
}

