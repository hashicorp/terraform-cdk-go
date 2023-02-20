// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
type DataTerraformRemoteStateEtcdConfig struct {
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) A space-separated list of the etcd endpoints.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Endpoints *string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Required) The path where to store the state.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) The password.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) The username.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

