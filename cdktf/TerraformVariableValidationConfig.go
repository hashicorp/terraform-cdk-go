// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformVariableValidationConfig struct {
	// Experimental.
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// Experimental.
	ErrorMessage *string `field:"required" json:"errorMessage" yaml:"errorMessage"`
}

