// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateCosConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The name of the COS bucket.
	//
	// You shall manually create it first.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) Object ACL to be applied to the state file, allows private and public-read.
	//
	// Defaults to private.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// (Optional) Whether to enable server side encryption of the state file.
	//
	// If it is true, COS will use 'AES256' encryption algorithm to encrypt state file.
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// (Optional) The path for saving the state file in bucket.
	//
	// Defaults to terraform.tfstate.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// (Optional) The directory for saving the state file in bucket.
	//
	// Default to "env:".
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) The region of the COS bucket.
	//
	// It supports environment variables TENCENTCLOUD_REGION.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// (Optional) Secret id of Tencent Cloud.
	//
	// It supports environment variables TENCENTCLOUD_SECRET_ID.
	// Experimental.
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// (Optional) Secret key of Tencent Cloud.
	//
	// It supports environment variables TENCENTCLOUD_SECRET_KEY.
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
}

