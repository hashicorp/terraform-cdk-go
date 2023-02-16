// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateRemoteConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Experimental.
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `field:"required" json:"workspaces" yaml:"workspaces"`
	// Experimental.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

