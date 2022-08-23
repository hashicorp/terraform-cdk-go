// Cloud Development Kit for Terraform
package cdktf


// Stores the state as an artifact in a given repository in Artifactory.
//
// Generic HTTP repositories are supported, and state from different configurations
// may be kept at different subpaths within the repository.
//
// Note: The URL must include the path to the Artifactory installation.
// It will likely end in /artifactory.
//
// This backend does not support state locking.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/artifactory
// Experimental.
type ArtifactoryBackendProps struct {
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

