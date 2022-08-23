// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateEtcdV3Config struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The list of 'etcd' endpoints which to connect to.
	// Experimental.
	Endpoints *[]*string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Optional) The path to a PEM-encoded CA bundle with which to verify certificates of TLS-enabled etcd servers.
	// Experimental.
	CacertPath *string `field:"optional" json:"cacertPath" yaml:"cacertPath"`
	// (Optional) The path to a PEM-encoded certificate to provide to etcd for secure client identification.
	// Experimental.
	CertPath *string `field:"optional" json:"certPath" yaml:"certPath"`
	// (Optional) The path to a PEM-encoded key to provide to etcd for secure client identification.
	// Experimental.
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
	// (Optional) Whether to lock state access.
	//
	// Defaults to true.
	// Experimental.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Password used to connect to the etcd cluster.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) An optional prefix to be added to keys when to storing state in etcd.
	//
	// Defaults to "".
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) Username used to connect to the etcd cluster.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

