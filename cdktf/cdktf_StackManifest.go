// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type StackManifest struct {
	// Experimental.
	Annotations *[]*StackAnnotation `field:"required" json:"annotations" yaml:"annotations"`
	// Experimental.
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// Experimental.
	Dependencies *[]*string `field:"required" json:"dependencies" yaml:"dependencies"`
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	SynthesizedStackPath *string `field:"required" json:"synthesizedStackPath" yaml:"synthesizedStackPath"`
	// Experimental.
	WorkingDirectory *string `field:"required" json:"workingDirectory" yaml:"workingDirectory"`
}

