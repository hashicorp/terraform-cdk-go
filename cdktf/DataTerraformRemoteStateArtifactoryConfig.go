// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
type DataTerraformRemoteStateArtifactoryConfig struct {
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) - The password.
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Password *string `field:"required" json:"password" yaml:"password"`
	// (Required) - The repository name.
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// (Required) - Path within the repository.
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Subpath *string `field:"required" json:"subpath" yaml:"subpath"`
	// (Required) - The URL.
	//
	// Note that this is the base url to artifactory not the full repo and subpath.
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Url *string `field:"required" json:"url" yaml:"url"`
	// (Required) - The username.
	// Deprecated: CDK for Terraform no longer supports the artifactory backend. Terraform deprecated artifactory in v1.2.3 and removed it in v1.3.
	Username *string `field:"required" json:"username" yaml:"username"`
}

