// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The local-exec provisioner invokes a local executable after a resource is created.
//
// This invokes a process on the machine running Terraform, not on the resource.
//
// See {@link https://www.terraform.io/language/resources/provisioners/local-exec local-exec}
// Experimental.
type ILocalExecProvisioner interface {
	// This is the command to execute.
	//
	// It can be provided as a relative path to the current working directory or as an absolute path.
	// It is evaluated in a shell, and can use environment variables or Terraform variables.
	// Experimental.
	Command() *string
	// A record of key value pairs representing the environment of the executed command.
	//
	// It inherits the current process environment.
	// Experimental.
	Environment() *map[string]*string
	// If provided, this is a list of interpreter arguments used to execute the command.
	//
	// The first argument is the interpreter itself.
	// It can be provided as a relative path to the current working directory or as an absolute path
	// The remaining arguments are appended prior to the command.
	// This allows building command lines of the form "/bin/bash", "-c", "echo foo".
	// If interpreter is unspecified, sensible defaults will be chosen based on the system OS.
	// Experimental.
	Interpreter() *[]*string
	// Experimental.
	Type() *string
	// If provided, specifies when Terraform will execute the command.
	//
	// For example, when = destroy specifies that the provisioner will run when the associated resource is destroyed.
	// Experimental.
	When() *string
	// If provided, specifies the working directory where command will be executed.
	//
	// It can be provided as a relative path to the current working directory or as an absolute path.
	// The directory must exist.
	// Experimental.
	WorkingDir() *string
}

// The jsii proxy for ILocalExecProvisioner
type jsiiProxy_ILocalExecProvisioner struct {
	_ byte // padding
}

func (j *jsiiProxy_ILocalExecProvisioner) Command() *string {
	var returns *string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) Interpreter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"interpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) When() *string {
	var returns *string
	_jsii_.Get(
		j,
		"when",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

