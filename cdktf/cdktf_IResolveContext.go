// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Current resolution context for tokens.
// Experimental.
type IResolveContext interface {
	// Use this postprocessor after the entire token structure has been resolved.
	// Experimental.
	RegisterPostProcessor(postProcessor IPostProcessor)
	// Resolve an inner object.
	// Experimental.
	Resolve(x interface{}) interface{}
	// TerraformIterators can be passed for block attributes and normal list attributes both require different handling when the iterable variable is accessed e.g. a dynamic block needs each.key while a for expression just needs key.
	// Experimental.
	IteratorContext() *string
	// Experimental.
	SetIteratorContext(i *string)
	// True when we are still preparing, false if we're rendering the final output.
	// Experimental.
	Preparing() *bool
	// The scope from which resolution has been initiated.
	// Experimental.
	Scope() constructs.IConstruct
	// True when ${} should be ommitted (because already inside them), false otherwise.
	// Experimental.
	SuppressBraces() *bool
	// Experimental.
	SetSuppressBraces(s *bool)
}

// The jsii proxy for IResolveContext
type jsiiProxy_IResolveContext struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolveContext) RegisterPostProcessor(postProcessor IPostProcessor) {
	if err := i.validateRegisterPostProcessorParameters(postProcessor); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerPostProcessor",
		[]interface{}{postProcessor},
	)
}

func (i *jsiiProxy_IResolveContext) Resolve(x interface{}) interface{} {
	if err := i.validateResolveParameters(x); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolveContext) IteratorContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iteratorContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext)SetIteratorContext(val *string) {
	_jsii_.Set(
		j,
		"iteratorContext",
		val,
	)
}

func (j *jsiiProxy_IResolveContext) Preparing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preparing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Scope() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) SuppressBraces() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suppressBraces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext)SetSuppressBraces(val *bool) {
	_jsii_.Set(
		j,
		"suppressBraces",
		val,
	)
}

