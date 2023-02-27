// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
type MantaBackendConfig struct {
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	InsecureSkipTlsVerify *bool `field:"optional" json:"insecureSkipTlsVerify" yaml:"insecureSkipTlsVerify"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	KeyMaterial *string `field:"optional" json:"keyMaterial" yaml:"keyMaterial"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	User *string `field:"optional" json:"user" yaml:"user"`
}

