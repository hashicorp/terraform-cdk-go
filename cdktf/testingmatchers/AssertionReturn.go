// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testingmatchers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// Class representing the contents of a return by an assertion.
// Experimental.
type AssertionReturn interface {
	// - String message containing information about the result of the assertion.
	// Experimental.
	Message() *string
	// - Boolean pass denoting the success of the assertion.
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


// Create an AssertionReturn.
// Experimental.
func NewAssertionReturn(message *string, pass *bool) AssertionReturn {
	_init_.Initialize()

	if err := validateNewAssertionReturnParameters(message, pass); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssertionReturn{}

	_jsii_.Create(
		"cdktf.testingMatchers.AssertionReturn",
		[]interface{}{message, pass},
		&j,
	)

	return &j
}

// Create an AssertionReturn.
// Experimental.
func NewAssertionReturn_Override(a AssertionReturn, message *string, pass *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.testingMatchers.AssertionReturn",
		[]interface{}{message, pass},
		a,
	)
}

