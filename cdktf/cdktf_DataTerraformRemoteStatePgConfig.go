// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStatePgConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Experimental.
	ConnStr *string `field:"required" json:"connStr" yaml:"connStr"`
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `field:"optional" json:"skipSchemaCreation" yaml:"skipSchemaCreation"`
}

