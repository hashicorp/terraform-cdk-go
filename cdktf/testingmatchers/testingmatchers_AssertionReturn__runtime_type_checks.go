//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

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

