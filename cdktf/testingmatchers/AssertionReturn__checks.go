// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package testingmatchers

import (
	"fmt"
)

func validateNewAssertionReturnParameters(message *string, pass *bool) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	if pass == nil {
		return fmt.Errorf("parameter pass is required, but nil was provided")
	}

	return nil
}

