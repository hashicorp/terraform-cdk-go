// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TerraformAssetConfig struct {
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Experimental.
	Type AssetType `field:"optional" json:"type" yaml:"type"`
}

