// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateOssConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Experimental.
	EcsRoleName *string `field:"optional" json:"ecsRoleName" yaml:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Experimental.
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `field:"optional" json:"tablestoreEndpoint" yaml:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `field:"optional" json:"tablestoreTable" yaml:"tablestoreTable"`
}

