// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformOutputConfig struct {
	// Experimental.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	Sensitive *bool `field:"optional" json:"sensitive" yaml:"sensitive"`
	// If set to true the synthesized Terraform Output will be named after the `id` passed to the constructor instead of the default (TerraformOutput.friendlyUniqueId).
	// Experimental.
	StaticId *bool `field:"optional" json:"staticId" yaml:"staticId"`
}

