// Cloud Development Kit for Terraform
package cdktf


// Stores the state in etcd 2.x at a given path.
//
// This backend does not support state locking.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/etcd
// Experimental.
type EtcdBackendProps struct {
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

