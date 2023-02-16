// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateS3Config struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Name of the S3 Bucket.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Path to the state file inside the S3 Bucket.
	//
	// When using a non-default workspace, the state path will be /workspace_key_prefix/workspace_name/key.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Optional) AWS access key.
	//
	// If configured, must also configure secret_key.
	// This can also be sourced from
	// the AWS_ACCESS_KEY_ID environment variable,
	// AWS shared credentials file (e.g. ~/.aws/credentials),
	// or AWS shared configuration file (e.g. ~/.aws/config).
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) Canned ACL to be applied to the state file.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// (Optional) IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	// Experimental.
	AssumeRolePolicy *string `field:"optional" json:"assumeRolePolicy" yaml:"assumeRolePolicy"`
	// (Optional) Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	// Experimental.
	AssumeRolePolicyArns *[]*string `field:"optional" json:"assumeRolePolicyArns" yaml:"assumeRolePolicyArns"`
	// (Optional) Map of assume role session tags.
	// Experimental.
	AssumeRoleTags *map[string]*string `field:"optional" json:"assumeRoleTags" yaml:"assumeRoleTags"`
	// (Optional) Set of assume role session tag keys to pass to any subsequent sessions.
	// Experimental.
	AssumeRoleTransitiveTagKeys *[]*string `field:"optional" json:"assumeRoleTransitiveTagKeys" yaml:"assumeRoleTransitiveTagKeys"`
	// (Optional) Custom endpoint for the AWS DynamoDB API.
	//
	// This can also be sourced from the AWS_DYNAMODB_ENDPOINT environment variable.
	// Experimental.
	DynamodbEndpoint *string `field:"optional" json:"dynamodbEndpoint" yaml:"dynamodbEndpoint"`
	// (Optional) Name of DynamoDB Table to use for state locking and consistency.
	//
	// The table must have a partition key named LockID with type of String.
	// If not configured, state locking will be disabled.
	// Experimental.
	DynamodbTable *string `field:"optional" json:"dynamodbTable" yaml:"dynamodbTable"`
	// (Optional) Enable server side encryption of the state file.
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// (Optional) Custom endpoint for the AWS S3 API.
	//
	// This can also be sourced from the AWS_S3_ENDPOINT environment variable.
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// (Optional) External identifier to use when assuming the role.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// (Optional) Enable path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>).
	// Experimental.
	ForcePathStyle *bool `field:"optional" json:"forcePathStyle" yaml:"forcePathStyle"`
	// (Optional) Custom endpoint for the AWS Identity and Access Management (IAM) API.
	//
	// This can also be sourced from the AWS_IAM_ENDPOINT environment variable.
	// Experimental.
	IamEndpoint *string `field:"optional" json:"iamEndpoint" yaml:"iamEndpoint"`
	// (Optional) Amazon Resource Name (ARN) of a Key Management Service (KMS) Key to use for encrypting the state.
	//
	// Note that if this value is specified,
	// Terraform will need kms:Encrypt, kms:Decrypt and kms:GenerateDataKey permissions on this KMS key.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// (Optional) The maximum number of times an AWS API request is retried on retryable failure.
	//
	// Defaults to 5.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// (Optional) Name of AWS profile in AWS shared credentials file (e.g. ~/.aws/credentials) or AWS shared configuration file (e.g. ~/.aws/config) to use for credentials and/or configuration. This can also be sourced from the AWS_PROFILE environment variable.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// AWS Region of the S3 Bucket and DynamoDB Table (if used).
	//
	// This can also
	// be sourced from the AWS_DEFAULT_REGION and AWS_REGION environment
	// variables.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// (Optional) Amazon Resource Name (ARN) of the IAM Role to assume.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// (Optional) AWS secret access key.
	//
	// If configured, must also configure access_key.
	// This can also be sourced from
	// the AWS_SECRET_ACCESS_KEY environment variable,
	// AWS shared credentials file (e.g. ~/.aws/credentials),
	// or AWS shared configuration file (e.g. ~/.aws/config)
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// (Optional) Session name to use when assuming the role.
	// Experimental.
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// (Optional) Path to the AWS shared credentials file.
	//
	// Defaults to ~/.aws/credentials.
	// Experimental.
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// (Optional) Skip credentials validation via the STS API.
	// Experimental.
	SkipCredentialsValidation *bool `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// (Optional) Skip usage of EC2 Metadata API.
	// Experimental.
	SkipMetadataApiCheck *bool `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// (Optional) Skip validation of provided region name.
	// Experimental.
	SkipRegionValidation *bool `field:"optional" json:"skipRegionValidation" yaml:"skipRegionValidation"`
	// (Optional) The key to use for encrypting state with Server-Side Encryption with Customer-Provided Keys (SSE-C).
	//
	// This is the base64-encoded value of the key, which must decode to 256 bits.
	// This can also be sourced from the AWS_SSE_CUSTOMER_KEY environment variable,
	// which is recommended due to the sensitivity of the value.
	// Setting it inside a terraform file will cause it to be persisted to disk in terraform.tfstate.
	// Experimental.
	SseCustomerKey *string `field:"optional" json:"sseCustomerKey" yaml:"sseCustomerKey"`
	// (Optional) Custom endpoint for the AWS Security Token Service (STS) API.
	//
	// This can also be sourced from the AWS_STS_ENDPOINT environment variable.
	// Experimental.
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
	// (Optional) Multi-Factor Authentication (MFA) token.
	//
	// This can also be sourced from the AWS_SESSION_TOKEN environment variable.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// (Optional) Prefix applied to the state path inside the bucket.
	//
	// This is only relevant when using a non-default workspace. Defaults to env:
	// Experimental.
	WorkspaceKeyPrefix *string `field:"optional" json:"workspaceKeyPrefix" yaml:"workspaceKeyPrefix"`
}

