// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateArtifactoryConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) - The password.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// (Required) - The repository name.
	// Experimental.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// (Required) - Path within the repository.
	// Experimental.
	Subpath *string `field:"required" json:"subpath" yaml:"subpath"`
	// (Required) - The URL.
	//
	// Note that this is the base url to artifactory not the full repo and subpath.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// (Required) - The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

