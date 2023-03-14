// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// The name of a single Terraform Cloud workspace.
//
// You will only be able to use the workspace specified in the configuration with this working directory, and cannot manage workspaces from the CLI (e.g. terraform workspace select or terraform workspace new).
// Experimental.
type NamedCloudWorkspace interface {
	CloudWorkspace
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NamedCloudWorkspace
type jsiiProxy_NamedCloudWorkspace struct {
	jsiiProxy_CloudWorkspace
}

func (j *jsiiProxy_NamedCloudWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewNamedCloudWorkspace(name *string) NamedCloudWorkspace {
	_init_.Initialize()

	if err := validateNewNamedCloudWorkspaceParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_NamedCloudWorkspace{}

	_jsii_.Create(
		"cdktf.NamedCloudWorkspace",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Experimental.
func NewNamedCloudWorkspace_Override(n NamedCloudWorkspace, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.NamedCloudWorkspace",
		[]interface{}{name},
		n,
	)
}

func (j *jsiiProxy_NamedCloudWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (n *jsiiProxy_NamedCloudWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

