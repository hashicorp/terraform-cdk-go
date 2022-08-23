// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type TestingAppOptions struct {
	// Experimental.
	EnableFutureFlags *bool `field:"optional" json:"enableFutureFlags" yaml:"enableFutureFlags"`
	// Experimental.
	FakeCdktfJsonPath *bool `field:"optional" json:"fakeCdktfJsonPath" yaml:"fakeCdktfJsonPath"`
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Experimental.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
	// Experimental.
	StubVersion *bool `field:"optional" json:"stubVersion" yaml:"stubVersion"`
}

