// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for lazy untyped value producers.
// Experimental.
type IAnyProducer interface {
	// Produce the value.
	// Experimental.
	Produce(context IResolveContext) interface{}
}

// The jsii proxy for IAnyProducer
type jsiiProxy_IAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IAnyProducer) Produce(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

