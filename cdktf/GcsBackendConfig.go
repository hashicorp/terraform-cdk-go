// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Stores the state as an object in a configurable prefix in a pre-existing bucket on Google Cloud Storage (GCS).
//
// The bucket must exist prior to configuring the backend.
//
// This backend supports state locking.
//
// Warning! It is highly recommended that you enable Object Versioning on the GCS bucket
// to allow for state recovery in the case of accidental deletions and human error.
//
// Read more about this backend in the Terraform docs:
// https://developer.hashicorp.com/terraform/language/settings/backends/gcs
// Experimental.
type GcsBackendConfig struct {
	// (Required) The name of the GCS bucket.
	//
	// This name must be globally unique.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) A temporary [OAuth 2.0 access token] obtained from the Google Authorization server, i.e. the Authorization: Bearer token used to authenticate HTTP requests to GCP APIs. This is an alternative to credentials. If both are specified, access_token will be used over the credentials field.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// (Optional) Local path to Google Cloud Platform account credentials in JSON format.
	//
	// If unset, Google Application Default Credentials are used.
	// The provided credentials must have Storage Object Admin role on the bucket.
	//
	// Warning: if using the Google Cloud Platform provider as well,
	// it will also pick up the GOOGLE_CREDENTIALS environment variable.
	// Experimental.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// (Optional) A 32 byte base64 encoded 'customer supplied encryption key' used to encrypt all state.
	// Experimental.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// (Optional) The service account to impersonate for accessing the State Bucket.
	//
	// You must have roles/iam.serviceAccountTokenCreator role on that account for the impersonation to succeed.
	// If you are using a delegation chain, you can specify that using the impersonate_service_account_delegates field.
	// Alternatively, this can be specified using the GOOGLE_IMPERSONATE_SERVICE_ACCOUNT environment variable.
	// Experimental.
	ImpersonateServiceAccount *string `field:"optional" json:"impersonateServiceAccount" yaml:"impersonateServiceAccount"`
	// (Optional) The delegation chain for an impersonating a service account.
	// Experimental.
	ImpersonateServiceAccountDelegates *[]*string `field:"optional" json:"impersonateServiceAccountDelegates" yaml:"impersonateServiceAccountDelegates"`
	// (Optional) GCS prefix inside the bucket.
	//
	// Named states for workspaces are stored in an object called <prefix>/<name>.tfstate.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

