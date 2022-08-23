// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
//
// See {@link https://www.terraform.io/language/resources/provisioners/connection connection}
// Experimental.
type ISSHProvisionerConnection interface {
	// Set to false to disable using ssh-agent to authenticate.
	//
	// On Windows the only supported SSH authentication agent is Pageant.
	// Experimental.
	Agent() *string
	// The preferred identity from the ssh agent for authentication.
	// Experimental.
	AgentIdentity() *string
	// The contents of a signed CA Certificate.
	//
	// The certificate argument must be used in conjunction with a bastion_private_key.
	// These can be loaded from a file on disk using the the file function.
	// Experimental.
	BastionCertificate() *string
	// Setting this enables the bastion Host connection.
	//
	// The provisioner will connect to bastion_host first, and then connect from there to host.
	// Experimental.
	BastionHost() *string
	// The public key from the remote host or the signing CA, used to verify the host connection.
	// Experimental.
	BastionHostKey() *string
	// The password to use for the bastion host.
	// Experimental.
	BastionPassword() *string
	// The port to use connect to the bastion host.
	// Experimental.
	BastionPort() *float64
	// The contents of an SSH key file to use for the bastion host.
	//
	// These can be loaded from a file on disk using the file function.
	// Experimental.
	BastionPrivateKey() *string
	// The user for the connection to the bastion host.
	// Experimental.
	BastionUser() *string
	// The contents of a signed CA Certificate.
	//
	// The certificate argument must be used in conjunction with a private_key.
	// These can be loaded from a file on disk using the the file function.
	// Experimental.
	Certificate() *string
	// The address of the resource to connect to.
	// Experimental.
	Host() *string
	// The public key from the remote host or the signing CA, used to verify the connection.
	// Experimental.
	HostKey() *string
	// The password to use for the connection.
	// Experimental.
	Password() *string
	// The port to connect to.
	// Experimental.
	Port() *float64
	// The contents of an SSH key to use for the connection.
	//
	// These can be loaded from a file on disk using the file function.
	// This takes preference over password if provided.
	// Experimental.
	PrivateKey() *string
	// Setting this enables the SSH over HTTP connection.
	//
	// This host will be connected to first, and then the host or bastion_host connection will be made from there.
	// Experimental.
	ProxyHost() *string
	// The port to use connect to the proxy host.
	// Experimental.
	ProxyPort() *float64
	// The ssh connection also supports the following fields to facilitate connections by SSH over HTTP proxy.
	// Experimental.
	ProxyScheme() *string
	// The username to use connect to the private proxy host.
	//
	// This argument should be specified only if authentication is required for the HTTP Proxy server.
	// Experimental.
	ProxyUserName() *string
	// The password to use connect to the private proxy host.
	//
	// This argument should be specified only if authentication is required for the HTTP Proxy server.
	// Experimental.
	ProxyUserPassword() *string
	// The path used to copy scripts meant for remote execution.
	//
	// Refer to {@link https://www.terraform.io/language/resources/provisioners/connection#how-provisioners-execute-remote-scripts How Provisioners Execute Remote Scripts below for more details}
	// Experimental.
	ScriptPath() *string
	// The target platform to connect to.
	//
	// Valid values are "windows" and "unix".
	// If the platform is set to windows, the default script_path is c:\windows\temp\terraform_%RAND%.cmd, assuming the SSH default shell is cmd.exe.
	// If the SSH default shell is PowerShell, set script_path to "c:/windows/temp/terraform_%RAND%.ps1"
	// Experimental.
	TargetPlatform() *string
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
	// The user to use for the connection.
	// Experimental.
	User() *string
}

// The jsii proxy for ISSHProvisionerConnection
type jsiiProxy_ISSHProvisionerConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_ISSHProvisionerConnection) Agent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) AgentIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionHostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionHostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bastionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) TargetPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

