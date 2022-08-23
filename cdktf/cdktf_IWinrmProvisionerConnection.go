// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
//
// See {@link https://www.terraform.io/language/resources/provisioners/connection connection}
// Experimental.
type IWinrmProvisionerConnection interface {
	// The CA certificate to validate against.
	// Experimental.
	Cacert() *string
	// The address of the resource to connect to.
	// Experimental.
	Host() *string
	// Set to true to connect using HTTPS instead of HTTP.
	// Experimental.
	Https() *bool
	// Set to true to skip validating the HTTPS certificate chain.
	// Experimental.
	Insecure() *bool
	// The password to use for the connection.
	// Experimental.
	Password() *string
	// The port to connect to.
	// Experimental.
	Port() *float64
	// The path used to copy scripts meant for remote execution.
	//
	// Refer to {@link https://www.terraform.io/language/resources/provisioners/connection#how-provisioners-execute-remote-scripts How Provisioners Execute Remote Scripts below for more details}
	// Experimental.
	ScriptPath() *string
	// The timeout to wait for the connection to become available.
	//
	// Should be provided as a string (e.g., "30s" or "5m".)
	// Experimental.
	Timeout() *string
	// The connection type.
	//
	// Valid values are "ssh" and "winrm".
	// Provisioners typically assume that the remote system runs Microsoft Windows when using WinRM.
	// Behaviors based on the SSH target_platform will force Windows-specific behavior for WinRM, unless otherwise specified.
	// Experimental.
	Type() *string
	// Set to true to use NTLM authentication rather than default (basic authentication), removing the requirement for basic authentication to be enabled within the target guest.
	//
	// Refer to Authentication for Remote Connections in the Windows App Development documentation for more details.
	// Experimental.
	UseNtlm() *bool
	// The user to use for the connection.
	// Experimental.
	User() *string
}

// The jsii proxy for IWinrmProvisionerConnection
type jsiiProxy_IWinrmProvisionerConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Cacert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Https() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"https",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Insecure() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) ScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) UseNtlm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useNtlm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

