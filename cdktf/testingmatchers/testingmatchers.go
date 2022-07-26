package testingmatchers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// Experimental.
type AssertionReturn interface {
	// Experimental.
	Message() *string
	// Experimental.
	Pass() *bool
}

// The jsii proxy struct for AssertionReturn
type jsiiProxy_AssertionReturn struct {
	_ byte // padding
}

func (j *jsiiProxy_AssertionReturn) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssertionReturn) Pass() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"pass",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssertionReturn(message *string, pass *bool) AssertionReturn {
	_init_.Initialize()

	j := jsiiProxy_AssertionReturn{}

	_jsii_.Create(
		"cdktf.testingMatchers.AssertionReturn",
		[]interface{}{message, pass},
		&j,
	)

	return &j
}

// Experimental.
func NewAssertionReturn_Override(a AssertionReturn, message *string, pass *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.testingMatchers.AssertionReturn",
		[]interface{}{message, pass},
		a,
	)
}

// Experimental.
type TerraformConstructor struct {
	// Experimental.
	TfResourceType *string `field:"required" json:"tfResourceType" yaml:"tfResourceType"`
}

