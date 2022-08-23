// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The remote-exec provisioner invokes a script on a remote resource after it is created.
//
// This can be used to run a configuration management tool, bootstrap into a cluster, etc
// The remote-exec provisioner requires a connection and supports both ssh and winrm.
//
// See {@link https://www.terraform.io/language/resources/provisioners/remote-exec remote-exec}
// Experimental.
type IRemoteExecProvisioner interface {
	// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
	//
	// A connection must be provided here or in the parent resource.
	// Experimental.
	Connection() interface{}
	// This is a list of command strings.
	//
	// They are executed in the order they are provided.
	// This cannot be provided with script or scripts.
	// Experimental.
	Inline() *[]*string
	// This is a path (relative or absolute) to a local script that will be copied to the remote resource and then executed.
	//
	// This cannot be provided with inline or scripts.
	// Experimental.
	Script() *string
	// This is a list of paths (relative or absolute) to local scripts that will be copied to the remote resource and then executed.
	//
	// They are executed in the order they are provided.
	// This cannot be provided with inline or script.
	// Experimental.
	Scripts() *[]*string
	// Experimental.
	Type() *string
}

// The jsii proxy for IRemoteExecProvisioner
type jsiiProxy_IRemoteExecProvisioner struct {
	_ byte // padding
}

func (j *jsiiProxy_IRemoteExecProvisioner) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Inline() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Scripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

