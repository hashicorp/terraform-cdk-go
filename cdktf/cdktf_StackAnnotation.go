// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type StackAnnotation struct {
	// Experimental.
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// Experimental.
	Level AnnotationMetadataEntryType `field:"required" json:"level" yaml:"level"`
	// Experimental.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Experimental.
	Stacktrace *[]*string `field:"optional" json:"stacktrace" yaml:"stacktrace"`
}

