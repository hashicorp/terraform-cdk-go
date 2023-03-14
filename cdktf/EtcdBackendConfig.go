// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Stores the state in etcd 2.x at a given path.
//
// This backend does not support state locking.
//
// Read more about this backend in the Terraform docs:
// https://developer.hashicorp.com/terraform/language/v1.2.x/settings/backends/etcd
// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
type EtcdBackendConfig struct {
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

