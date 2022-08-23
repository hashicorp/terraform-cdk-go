// Cloud Development Kit for Terraform
package cdktf


// Stores the state as an object in a configurable prefix in a given bucket on Tencent Cloud Object Storage (COS).
//
// This backend supports state locking.
//
// Warning! It is highly recommended that you enable Object Versioning on the COS bucket to allow for state recovery in the case of accidental deletions and human error.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/cos
// Experimental.
type CosBackendProps struct {
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

