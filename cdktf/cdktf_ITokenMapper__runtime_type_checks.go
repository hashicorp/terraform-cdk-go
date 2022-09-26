//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"
)

func (i *jsiiProxy_ITokenMapper) validateMapTokenParameters(t IResolvable) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

