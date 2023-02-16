// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

