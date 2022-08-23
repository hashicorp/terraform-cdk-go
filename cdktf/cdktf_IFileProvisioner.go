// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The file provisioner copies files or directories from the machine running Terraform to the newly created resource.
//
// The file provisioner supports both ssh and winrm type connections.
//
// See {@link https://www.terraform.io/language/resources/provisioners/file file}
// Experimental.
type IFileProvisioner interface {
	// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
	// Experimental.
	Connection() interface{}
	// The destination path to write to on the remote system.
	//
	// See Destination Paths below for more information.
	// Experimental.
	Content() *string
	// The source file or directory.
	//
	// Specify it either relative to the current working directory or as an absolute path.
	// This argument cannot be combined with content.
	// Experimental.
	Destination() *string
	// The direct content to copy on the destination.
	//
	// If destination is a file, the content will be written on that file.
	// In case of a directory, a file named tf-file-content is created inside that directory.
	// We recommend using a file as the destination when using content.
	// This argument cannot be combined with source.
	// Experimental.
	Source() *string
	// Experimental.
	Type() *string
}

// The jsii proxy for IFileProvisioner
type jsiiProxy_IFileProvisioner struct {
	_ byte // padding
}

func (j *jsiiProxy_IFileProvisioner) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

