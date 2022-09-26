//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_ITerraformResource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ITerraformResource) validateSetLifecycleParameters(val *TerraformResourceLifecycle) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

