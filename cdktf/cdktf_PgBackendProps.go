// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type PgBackendProps struct {
	// Experimental.
	ConnStr *string `field:"required" json:"connStr" yaml:"connStr"`
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `field:"optional" json:"skipSchemaCreation" yaml:"skipSchemaCreation"`
}

