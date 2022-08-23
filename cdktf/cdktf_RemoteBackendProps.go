// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type RemoteBackendProps struct {
	// Experimental.
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `field:"required" json:"workspaces" yaml:"workspaces"`
	// Experimental.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

