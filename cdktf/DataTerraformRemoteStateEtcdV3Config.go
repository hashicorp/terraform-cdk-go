// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
type DataTerraformRemoteStateEtcdV3Config struct {
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The list of 'etcd' endpoints which to connect to.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Endpoints *[]*string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Optional) The path to a PEM-encoded CA bundle with which to verify certificates of TLS-enabled etcd servers.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	CacertPath *string `field:"optional" json:"cacertPath" yaml:"cacertPath"`
	// (Optional) The path to a PEM-encoded certificate to provide to etcd for secure client identification.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	CertPath *string `field:"optional" json:"certPath" yaml:"certPath"`
	// (Optional) The path to a PEM-encoded key to provide to etcd for secure client identification.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
	// (Optional) Whether to lock state access.
	//
	// Defaults to true.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Password used to connect to the etcd cluster.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) An optional prefix to be added to keys when to storing state in etcd.
	//
	// Defaults to "".
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) Username used to connect to the etcd cluster.
	// Deprecated: CDK for Terraform no longer supports the etcdv3 backend. Terraform deprecated etcdv3 in v1.2.3 and removed it in v1.3.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

