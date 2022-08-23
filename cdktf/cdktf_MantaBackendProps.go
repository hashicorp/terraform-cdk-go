// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type MantaBackendProps struct {
	// Experimental.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Experimental.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Experimental.
	InsecureSkipTlsVerify *bool `field:"optional" json:"insecureSkipTlsVerify" yaml:"insecureSkipTlsVerify"`
	// Experimental.
	KeyMaterial *string `field:"optional" json:"keyMaterial" yaml:"keyMaterial"`
	// Experimental.
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
}

