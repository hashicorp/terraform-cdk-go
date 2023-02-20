// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type PgBackendConfig struct {
	// Experimental.
	ConnStr *string `field:"required" json:"connStr" yaml:"connStr"`
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `field:"optional" json:"skipSchemaCreation" yaml:"skipSchemaCreation"`
}

