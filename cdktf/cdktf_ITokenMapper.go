// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface to apply operation to tokens in a string.
//
// Interface so it can be exported via jsii.
// Experimental.
type ITokenMapper interface {
	// Replace a single token.
	// Experimental.
	MapToken(t IResolvable) interface{}
}

// The jsii proxy for ITokenMapper
type jsiiProxy_ITokenMapper struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenMapper) MapToken(t IResolvable) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"mapToken",
		[]interface{}{t},
		&returns,
	)

	return returns
}

