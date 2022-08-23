// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ITerraformAddressable interface {
	// Experimental.
	Fqn() *string
}

// The jsii proxy for ITerraformAddressable
type jsiiProxy_ITerraformAddressable struct {
	_ byte // padding
}

func (j *jsiiProxy_ITerraformAddressable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

