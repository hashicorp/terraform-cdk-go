// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformModuleProvider struct {
	// Experimental.
	ModuleAlias *string `field:"required" json:"moduleAlias" yaml:"moduleAlias"`
	// Experimental.
	Provider TerraformProvider `field:"required" json:"provider" yaml:"provider"`
}

