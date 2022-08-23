// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateEtcdConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) A space-separated list of the etcd endpoints.
	// Experimental.
	Endpoints *string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Required) The path where to store the state.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) The password.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) The username.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

