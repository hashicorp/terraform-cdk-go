// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IInterpolatingParent interface {
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
}

// The jsii proxy for IInterpolatingParent
type jsiiProxy_IInterpolatingParent struct {
	_ byte // padding
}

func (i *jsiiProxy_IInterpolatingParent) InterpolationForAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

