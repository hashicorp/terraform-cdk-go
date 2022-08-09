// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf/internal"
)

// Experimental.
type AnnotationMetadataEntryType string

const (
	// Experimental.
	AnnotationMetadataEntryType_INFO AnnotationMetadataEntryType = "INFO"
	// Experimental.
	AnnotationMetadataEntryType_WARN AnnotationMetadataEntryType = "WARN"
	// Experimental.
	AnnotationMetadataEntryType_ERROR AnnotationMetadataEntryType = "ERROR"
)

// Includes API for attaching annotations such as warning messages to constructs.
// Experimental.
type Annotations interface {
	// Adds an { "error": <message> } metadata entry to this construct.
	//
	// The toolkit will fail synthesis when errors are reported.
	// Experimental.
	AddError(message *string)
	// Adds an info metadata entry to this construct.
	//
	// The CLI will display the info message when apps are synthesized.
	// Experimental.
	AddInfo(message *string)
	// Adds a warning metadata entry to this construct.
	//
	// The CLI will display the warning when an app is synthesized.
	// In a future release the CLI might introduce a --strict flag which
	// will then fail the synthesis if it encounters a warning.
	// Experimental.
	AddWarning(message *string)
}

// The jsii proxy struct for Annotations
type jsiiProxy_Annotations struct {
	_ byte // padding
}

// Returns the annotations API for a construct scope.
// Experimental.
func Annotations_Of(scope constructs.IConstruct) Annotations {
	_init_.Initialize()

	var returns Annotations

	_jsii_.StaticInvoke(
		"cdktf.Annotations",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) AddError(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addError",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddInfo(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addInfo",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddWarning(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addWarning",
		[]interface{}{message},
	)
}

// Experimental.
type AnyMap interface {
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Lookup(key *string) interface{}
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AnyMap
type jsiiProxy_AnyMap struct {
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_AnyMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMap) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAnyMap(terraformResource IInterpolatingParent, terraformAttribute *string) AnyMap {
	_init_.Initialize()

	j := jsiiProxy_AnyMap{}

	_jsii_.Create(
		"cdktf.AnyMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAnyMap_Override(a AnyMap, terraformResource IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.AnyMap",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AnyMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AnyMap) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AnyMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMap) Lookup(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMap) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type AnyMapList interface {
	IInterpolatingParent
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) AnyMap
	// Experimental.
	InterpolationForAttribute(property *string) IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AnyMapList
type jsiiProxy_AnyMapList struct {
	jsiiProxy_IInterpolatingParent
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_AnyMapList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMapList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMapList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMapList) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnyMapList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAnyMapList(terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AnyMapList {
	_init_.Initialize()

	j := jsiiProxy_AnyMapList{}

	_jsii_.Create(
		"cdktf.AnyMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAnyMapList_Override(a AnyMapList, terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.AnyMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AnyMapList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AnyMapList) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AnyMapList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AnyMapList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMapList) Get(index *float64) AnyMap {
	var returns AnyMap

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMapList) InterpolationForAttribute(property *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMapList) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnyMapList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a cdktf application.
// Experimental.
type App interface {
	constructs.Construct
	// Experimental.
	Manifest() Manifest
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The output directory into which resources will be synthesized.
	// Experimental.
	Outdir() *string
	// Whether to skip the validation during synthesis of the app.
	// Experimental.
	SkipValidation() *bool
	// The stack which will be synthesized.
	//
	// If not set, all stacks will be synthesized.
	// Experimental.
	TargetStackId() *string
	// Creates a reference from one stack to another, invoked on prepareStack since it creates extra resources.
	// Experimental.
	CrossStackReference(fromStack TerraformStack, toStack TerraformStack, identifier *string) *string
	// Synthesizes all resources to the output directory.
	// Experimental.
	Synth()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_App) Manifest() Manifest {
	var returns Manifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) SkipValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TargetStackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStackId",
		&returns,
	)
	return returns
}


// Defines an app.
// Experimental.
func NewApp(options *AppOptions) App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"cdktf.App",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Defines an app.
// Experimental.
func NewApp_Override(a App, options *AppOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.App",
		[]interface{}{options},
		a,
	)
}

// Experimental.
func App_IsApp(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.App",
		"isApp",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func App_Of(construct constructs.IConstruct) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.App",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) CrossStackReference(fromStack TerraformStack, toStack TerraformStack, identifier *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"crossStackReference",
		[]interface{}{fromStack, toStack, identifier},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type AppOptions struct {
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdktf.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// The directory to output Terraform resources.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Whether to skip the validation during synthesis of the app.
	// Experimental.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// Experimental.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
}

// Experimental.
type ArtifactoryBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ArtifactoryBackend
type jsiiProxy_ArtifactoryBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_ArtifactoryBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewArtifactoryBackend(scope constructs.Construct, props *ArtifactoryBackendProps) ArtifactoryBackend {
	_init_.Initialize()

	j := jsiiProxy_ArtifactoryBackend{}

	_jsii_.Create(
		"cdktf.ArtifactoryBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewArtifactoryBackend_Override(a ArtifactoryBackend, scope constructs.Construct, props *ArtifactoryBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ArtifactoryBackend",
		[]interface{}{scope, props},
		a,
	)
}

// Experimental.
func ArtifactoryBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.ArtifactoryBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func ArtifactoryBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.ArtifactoryBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactoryBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArtifactoryBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		a,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactoryBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArtifactoryBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactoryBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactoryBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactoryBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactoryBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state as an artifact in a given repository in Artifactory.
//
// Generic HTTP repositories are supported, and state from different configurations
// may be kept at different subpaths within the repository.
//
// Note: The URL must include the path to the Artifactory installation.
// It will likely end in /artifactory.
//
// This backend does not support state locking.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/artifactory
// Experimental.
type ArtifactoryBackendProps struct {
	// (Required) - The password.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// (Required) - The repository name.
	// Experimental.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// (Required) - Path within the repository.
	// Experimental.
	Subpath *string `field:"required" json:"subpath" yaml:"subpath"`
	// (Required) - The URL.
	//
	// Note that this is the base url to artifactory not the full repo and subpath.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// (Required) - The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

// Aspects can be applied to CDK tree scopes and can operate on the tree before synthesis.
// Experimental.
type Aspects interface {
	// The list of aspects which were directly applied on this scope.
	// Experimental.
	All() *[]IAspect
	// Adds an aspect to apply this scope before synthesis.
	// Experimental.
	Add(aspect IAspect)
}

// The jsii proxy struct for Aspects
type jsiiProxy_Aspects struct {
	_ byte // padding
}

func (j *jsiiProxy_Aspects) All() *[]IAspect {
	var returns *[]IAspect
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}


// Returns the `Aspects` object associated with a construct scope.
// Experimental.
func Aspects_Of(scope constructs.IConstruct) Aspects {
	_init_.Initialize()

	var returns Aspects

	_jsii_.StaticInvoke(
		"cdktf.Aspects",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Aspects) Add(aspect IAspect) {
	_jsii_.InvokeVoid(
		a,
		"add",
		[]interface{}{aspect},
	)
}

// Experimental.
type AssetType string

const (
	// Experimental.
	AssetType_FILE AssetType = "FILE"
	// Experimental.
	AssetType_DIRECTORY AssetType = "DIRECTORY"
	// Experimental.
	AssetType_ARCHIVE AssetType = "ARCHIVE"
)

// Experimental.
type AzurermBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AzurermBackend
type jsiiProxy_AzurermBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_AzurermBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewAzurermBackend(scope constructs.Construct, props *AzurermBackendProps) AzurermBackend {
	_init_.Initialize()

	j := jsiiProxy_AzurermBackend{}

	_jsii_.Create(
		"cdktf.AzurermBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAzurermBackend_Override(a AzurermBackend, scope constructs.Construct, props *AzurermBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.AzurermBackend",
		[]interface{}{scope, props},
		a,
	)
}

// Experimental.
func AzurermBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.AzurermBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AzurermBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.AzurermBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzurermBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		a,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzurermBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state as a Blob with the given Key within the Blob Container within the Blob Storage Account.
//
// This backend supports state locking and consistency checking
// with Azure Blob Storage native capabilities.
//
// Note: By default the Azure Backend uses ADAL for authentication which is deprecated
// in favour of MSAL - MSAL can be used by setting use_microsoft_graph to true.
// The default for this will change in Terraform 1.2,
// so that MSAL authentication is used by default.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/azurerm
// Experimental.
type AzurermBackendProps struct {
	// (Required) The Name of the Storage Container within the Storage Account.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// (Required) The name of the Blob used to retrieve/store Terraform's State file inside the Storage Container.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Required) The Name of the Storage Account.
	// Experimental.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// access_key - (Optional) The Access Key used to access the Blob Storage Account.
	//
	// This can also be sourced from the ARM_ACCESS_KEY environment variable.
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) The Client ID of the Service Principal.
	//
	// This can also be sourced from the ARM_CLIENT_ID environment variable.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// (Optional) The Client Secret of the Service Principal.
	//
	// This can also be sourced from the ARM_CLIENT_SECRET environment variable.
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// (Optional) The Custom Endpoint for Azure Resource Manager. This can also be sourced from the ARM_ENDPOINT environment variable.
	//
	// NOTE: An endpoint should only be configured when using Azure Stack.
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// (Optional) The Azure Environment which should be used.
	//
	// This can also be sourced from the ARM_ENVIRONMENT environment variable.
	//   Possible values are public, china, german, stack and usgovernment. Defaults to public.
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// (Optional) The path to a custom Managed Service Identity endpoint which is automatically determined if not specified.
	//
	// This can also be sourced from the ARM_MSI_ENDPOINT environment variable.
	// Experimental.
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// (Required) The Name of the Resource Group in which the Storage Account exists.
	// Experimental.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// (Optional) The SAS Token used to access the Blob Storage Account.
	//
	// This can also be sourced from the ARM_SAS_TOKEN environment variable.
	// Experimental.
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// (Optional) The Subscription ID in which the Storage Account exists.
	//
	// This can also be sourced from the ARM_SUBSCRIPTION_ID environment variable.
	// Experimental.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// (Optional) The Tenant ID in which the Subscription exists.
	//
	// This can also be sourced from the ARM_TENANT_ID environment variable.
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// (Optional) Should Managed Service Identity authentication be used?
	//
	// This can also be sourced from the ARM_USE_MSI environment variable.
	// Experimental.
	UseMsi *bool `field:"optional" json:"useMsi" yaml:"useMsi"`
}

// Experimental.
type BooleanMap interface {
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Lookup(key *string) IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BooleanMap
type jsiiProxy_BooleanMap struct {
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_BooleanMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMap) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewBooleanMap(terraformResource IInterpolatingParent, terraformAttribute *string) BooleanMap {
	_init_.Initialize()

	j := jsiiProxy_BooleanMap{}

	_jsii_.Create(
		"cdktf.BooleanMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewBooleanMap_Override(b BooleanMap, terraformResource IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.BooleanMap",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BooleanMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BooleanMap) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BooleanMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMap) Lookup(key *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		b,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMap) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type BooleanMapList interface {
	IInterpolatingParent
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) BooleanMap
	// Experimental.
	InterpolationForAttribute(property *string) IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BooleanMapList
type jsiiProxy_BooleanMapList struct {
	jsiiProxy_IInterpolatingParent
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_BooleanMapList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMapList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMapList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMapList) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMapList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewBooleanMapList(terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BooleanMapList {
	_init_.Initialize()

	j := jsiiProxy_BooleanMapList{}

	_jsii_.Create(
		"cdktf.BooleanMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewBooleanMapList_Override(b BooleanMapList, terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.BooleanMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BooleanMapList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BooleanMapList) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BooleanMapList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BooleanMapList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMapList) Get(index *float64) BooleanMap {
	var returns BooleanMap

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMapList) InterpolationForAttribute(property *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMapList) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BooleanMapList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Deprecated: Going to be replaced by Array of ComplexListItem
// and will be removed in the future.
type ComplexComputedList interface {
	IInterpolatingParent
	IResolvable
	ITerraformAddressable
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	ComplexComputedListIndex() *string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	SetComplexComputedListIndex(val *string)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	CreationStack() *[]*string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	Fqn() *string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	TerraformAttribute() *string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	SetTerraformAttribute(val *string)
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	TerraformResource() IInterpolatingParent
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	SetTerraformResource(val IInterpolatingParent)
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	WrapsSet() *bool
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	SetWrapsSet(val *bool)
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	ComputeFqn() *string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetBooleanAttribute(terraformAttribute *string) IResolvable
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetStringAttribute(terraformAttribute *string) *string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	InterpolationForAttribute(property *string) IResolvable
	// Produce the Token's value at resolution time.
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Deprecated: Going to be replaced by Array of ComplexListItem
	// and will be removed in the future.
	ToString() *string
}

// The jsii proxy struct for ComplexComputedList
type jsiiProxy_ComplexComputedList struct {
	jsiiProxy_IInterpolatingParent
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_ComplexComputedList) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Deprecated: Going to be replaced by Array of ComplexListItem
// and will be removed in the future.
func NewComplexComputedList(terraformResource IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) ComplexComputedList {
	_init_.Initialize()

	j := jsiiProxy_ComplexComputedList{}

	_jsii_.Create(
		"cdktf.ComplexComputedList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Deprecated: Going to be replaced by Array of ComplexListItem
// and will be removed in the future.
func NewComplexComputedList_Override(c ComplexComputedList, terraformResource IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexComputedList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComplexComputedList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetBooleanAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) InterpolationForAttribute(property *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexComputedList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ComplexList interface {
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComplexList
type jsiiProxy_ComplexList struct {
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_ComplexList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexList) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewComplexList_Override(c ComplexList, terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComplexList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComplexList) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComplexList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComplexList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexList) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ComplexMap interface {
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComplexMap
type jsiiProxy_ComplexMap struct {
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_ComplexMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexMap) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewComplexMap_Override(c ComplexMap, terraformResource IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexMap",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComplexMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComplexMap) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComplexMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexMap) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ComplexObject interface {
	IInterpolatingParent
	IResolvable
	ITerraformAddressable
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComplexObject
type jsiiProxy_ComplexObject struct {
	jsiiProxy_IInterpolatingParent
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_ComplexObject) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexObject) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexObject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexObject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexObject) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexObject) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewComplexObject(terraformResource IInterpolatingParent, terraformAttribute *string, complexObjectIsFromSet *bool, complexObjectIndex interface{}) ComplexObject {
	_init_.Initialize()

	j := jsiiProxy_ComplexObject{}

	_jsii_.Create(
		"cdktf.ComplexObject",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIsFromSet, complexObjectIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewComplexObject_Override(c ComplexObject, terraformResource IInterpolatingParent, terraformAttribute *string, complexObjectIsFromSet *bool, complexObjectIndex interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexObject",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIsFromSet, complexObjectIndex},
		c,
	)
}

func (j *jsiiProxy_ComplexObject) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComplexObject) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComplexObject) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComplexObject) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComplexObject) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetBooleanAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) InterpolationAsList() IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) InterpolationForAttribute(property *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComplexObject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ConsulBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ConsulBackend
type jsiiProxy_ConsulBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_ConsulBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewConsulBackend(scope constructs.Construct, props *ConsulBackendProps) ConsulBackend {
	_init_.Initialize()

	j := jsiiProxy_ConsulBackend{}

	_jsii_.Create(
		"cdktf.ConsulBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewConsulBackend_Override(c ConsulBackend, scope constructs.Construct, props *ConsulBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ConsulBackend",
		[]interface{}{scope, props},
		c,
	)
}

// Experimental.
func ConsulBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.ConsulBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func ConsulBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.ConsulBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConsulBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		c,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConsulBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state in the Consul KV store at a given path. This backend supports state locking.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/consul
// Experimental.
type ConsulBackendProps struct {
	// (Required) Access token.
	// Experimental.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// (Required) Path in the Consul KV store.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) DNS name and port of your Consul endpoint specified in the format dnsname:port.
	//
	// Defaults to the local agent HTTP listener.
	// Experimental.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// (Optional) A path to a PEM-encoded certificate authority used to verify the remote agent's certificate.
	// Experimental.
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// (Optional) A path to a PEM-encoded certificate provided to the remote agent;
	//
	// requires use of key_file.
	// Experimental.
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// (Optional) The datacenter to use.
	//
	// Defaults to that of the agent.
	// Experimental.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// (Optional) true to compress the state data using gzip, or false (the default) to leave it uncompressed.
	// Experimental.
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// (Optional) HTTP Basic Authentication credentials to be used when communicating with Consul, in the format of either user or user:pass.
	// Experimental.
	HttpAuth *string `field:"optional" json:"httpAuth" yaml:"httpAuth"`
	// (Optional) A path to a PEM-encoded private key, required if cert_file is specified.
	// Experimental.
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// (Optional) false to disable locking.
	//
	// This defaults to true, but will require session permissions with Consul and
	// at least kv write permissions on $path/.lock to perform locking.
	// Experimental.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Specifies what protocol to use when talking to the given address,either http or https.
	//
	// SSL support can also be triggered by setting then environment variable CONSUL_HTTP_SSL to true.
	// Experimental.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

// Experimental.
type CosBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CosBackend
type jsiiProxy_CosBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_CosBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewCosBackend(scope constructs.Construct, props *CosBackendProps) CosBackend {
	_init_.Initialize()

	j := jsiiProxy_CosBackend{}

	_jsii_.Create(
		"cdktf.CosBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCosBackend_Override(c CosBackend, scope constructs.Construct, props *CosBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.CosBackend",
		[]interface{}{scope, props},
		c,
	)
}

// Experimental.
func CosBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.CosBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func CosBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.CosBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		c,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state as an object in a configurable prefix in a given bucket on Tencent Cloud Object Storage (COS).
//
// This backend supports state locking.
//
// Warning! It is highly recommended that you enable Object Versioning on the COS bucket to allow for state recovery in the case of accidental deletions and human error.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/cos
// Experimental.
type CosBackendProps struct {
	// (Required) The name of the COS bucket.
	//
	// You shall manually create it first.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) Object ACL to be applied to the state file, allows private and public-read.
	//
	// Defaults to private.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// (Optional) Whether to enable server side encryption of the state file.
	//
	// If it is true, COS will use 'AES256' encryption algorithm to encrypt state file.
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// (Optional) The path for saving the state file in bucket.
	//
	// Defaults to terraform.tfstate.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// (Optional) The directory for saving the state file in bucket.
	//
	// Default to "env:".
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) The region of the COS bucket.
	//
	// It supports environment variables TENCENTCLOUD_REGION.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// (Optional) Secret id of Tencent Cloud.
	//
	// It supports environment variables TENCENTCLOUD_SECRET_ID.
	// Experimental.
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// (Optional) Secret key of Tencent Cloud.
	//
	// It supports environment variables TENCENTCLOUD_SECRET_KEY.
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
}

// Experimental.
type DataTerraformRemoteState interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteState
type jsiiProxy_DataTerraformRemoteState struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteState) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteState(scope constructs.Construct, id *string, config *DataTerraformRemoteStateRemoteConfig) DataTerraformRemoteState {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteState{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteState",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteState_Override(d DataTerraformRemoteState, scope constructs.Construct, id *string, config *DataTerraformRemoteStateRemoteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteState",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteState_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteState",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteState_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteState",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteState) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteState) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteState) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteState) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateArtifactory interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateArtifactory
type jsiiProxy_DataTerraformRemoteStateArtifactory struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateArtifactory(scope constructs.Construct, id *string, config *DataTerraformRemoteStateArtifactoryConfig) DataTerraformRemoteStateArtifactory {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateArtifactory{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateArtifactory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateArtifactory_Override(d DataTerraformRemoteStateArtifactory, scope constructs.Construct, id *string, config *DataTerraformRemoteStateArtifactoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateArtifactory",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateArtifactory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateArtifactory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateArtifactory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateArtifactory",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateArtifactoryConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) - The password.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// (Required) - The repository name.
	// Experimental.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// (Required) - Path within the repository.
	// Experimental.
	Subpath *string `field:"required" json:"subpath" yaml:"subpath"`
	// (Required) - The URL.
	//
	// Note that this is the base url to artifactory not the full repo and subpath.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// (Required) - The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

// Experimental.
type DataTerraformRemoteStateAzurerm interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateAzurerm
type jsiiProxy_DataTerraformRemoteStateAzurerm struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateAzurerm(scope constructs.Construct, id *string, config *DataTerraformRemoteStateAzurermConfig) DataTerraformRemoteStateAzurerm {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateAzurerm{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateAzurerm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateAzurerm_Override(d DataTerraformRemoteStateAzurerm, scope constructs.Construct, id *string, config *DataTerraformRemoteStateAzurermConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateAzurerm",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateAzurerm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateAzurerm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateAzurerm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateAzurerm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateAzurermConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The Name of the Storage Container within the Storage Account.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// (Required) The name of the Blob used to retrieve/store Terraform's State file inside the Storage Container.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Required) The Name of the Storage Account.
	// Experimental.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// access_key - (Optional) The Access Key used to access the Blob Storage Account.
	//
	// This can also be sourced from the ARM_ACCESS_KEY environment variable.
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) The Client ID of the Service Principal.
	//
	// This can also be sourced from the ARM_CLIENT_ID environment variable.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// (Optional) The Client Secret of the Service Principal.
	//
	// This can also be sourced from the ARM_CLIENT_SECRET environment variable.
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// (Optional) The Custom Endpoint for Azure Resource Manager. This can also be sourced from the ARM_ENDPOINT environment variable.
	//
	// NOTE: An endpoint should only be configured when using Azure Stack.
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// (Optional) The Azure Environment which should be used.
	//
	// This can also be sourced from the ARM_ENVIRONMENT environment variable.
	//   Possible values are public, china, german, stack and usgovernment. Defaults to public.
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// (Optional) The path to a custom Managed Service Identity endpoint which is automatically determined if not specified.
	//
	// This can also be sourced from the ARM_MSI_ENDPOINT environment variable.
	// Experimental.
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// (Required) The Name of the Resource Group in which the Storage Account exists.
	// Experimental.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// (Optional) The SAS Token used to access the Blob Storage Account.
	//
	// This can also be sourced from the ARM_SAS_TOKEN environment variable.
	// Experimental.
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// (Optional) The Subscription ID in which the Storage Account exists.
	//
	// This can also be sourced from the ARM_SUBSCRIPTION_ID environment variable.
	// Experimental.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// (Optional) The Tenant ID in which the Subscription exists.
	//
	// This can also be sourced from the ARM_TENANT_ID environment variable.
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// (Optional) Should Managed Service Identity authentication be used?
	//
	// This can also be sourced from the ARM_USE_MSI environment variable.
	// Experimental.
	UseMsi *bool `field:"optional" json:"useMsi" yaml:"useMsi"`
}

// Experimental.
type DataTerraformRemoteStateConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

// Experimental.
type DataTerraformRemoteStateConsul interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateConsul
type jsiiProxy_DataTerraformRemoteStateConsul struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateConsul(scope constructs.Construct, id *string, config *DataTerraformRemoteStateConsulConfig) DataTerraformRemoteStateConsul {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateConsul{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateConsul",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateConsul_Override(d DataTerraformRemoteStateConsul, scope constructs.Construct, id *string, config *DataTerraformRemoteStateConsulConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateConsul",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateConsul_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateConsul",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateConsul_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateConsul",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateConsul) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateConsulConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) Access token.
	// Experimental.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// (Required) Path in the Consul KV store.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) DNS name and port of your Consul endpoint specified in the format dnsname:port.
	//
	// Defaults to the local agent HTTP listener.
	// Experimental.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// (Optional) A path to a PEM-encoded certificate authority used to verify the remote agent's certificate.
	// Experimental.
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// (Optional) A path to a PEM-encoded certificate provided to the remote agent;
	//
	// requires use of key_file.
	// Experimental.
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// (Optional) The datacenter to use.
	//
	// Defaults to that of the agent.
	// Experimental.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// (Optional) true to compress the state data using gzip, or false (the default) to leave it uncompressed.
	// Experimental.
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// (Optional) HTTP Basic Authentication credentials to be used when communicating with Consul, in the format of either user or user:pass.
	// Experimental.
	HttpAuth *string `field:"optional" json:"httpAuth" yaml:"httpAuth"`
	// (Optional) A path to a PEM-encoded private key, required if cert_file is specified.
	// Experimental.
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// (Optional) false to disable locking.
	//
	// This defaults to true, but will require session permissions with Consul and
	// at least kv write permissions on $path/.lock to perform locking.
	// Experimental.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Specifies what protocol to use when talking to the given address,either http or https.
	//
	// SSL support can also be triggered by setting then environment variable CONSUL_HTTP_SSL to true.
	// Experimental.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

// Experimental.
type DataTerraformRemoteStateCos interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateCos
type jsiiProxy_DataTerraformRemoteStateCos struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateCos(scope constructs.Construct, id *string, config *DataTerraformRemoteStateCosConfig) DataTerraformRemoteStateCos {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateCos{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateCos",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateCos_Override(d DataTerraformRemoteStateCos, scope constructs.Construct, id *string, config *DataTerraformRemoteStateCosConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateCos",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateCos_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateCos",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateCos_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateCos",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateCos) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateCosConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The name of the COS bucket.
	//
	// You shall manually create it first.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) Object ACL to be applied to the state file, allows private and public-read.
	//
	// Defaults to private.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// (Optional) Whether to enable server side encryption of the state file.
	//
	// If it is true, COS will use 'AES256' encryption algorithm to encrypt state file.
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// (Optional) The path for saving the state file in bucket.
	//
	// Defaults to terraform.tfstate.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// (Optional) The directory for saving the state file in bucket.
	//
	// Default to "env:".
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) The region of the COS bucket.
	//
	// It supports environment variables TENCENTCLOUD_REGION.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// (Optional) Secret id of Tencent Cloud.
	//
	// It supports environment variables TENCENTCLOUD_SECRET_ID.
	// Experimental.
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// (Optional) Secret key of Tencent Cloud.
	//
	// It supports environment variables TENCENTCLOUD_SECRET_KEY.
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
}

// Experimental.
type DataTerraformRemoteStateEtcd interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateEtcd
type jsiiProxy_DataTerraformRemoteStateEtcd struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateEtcd(scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdConfig) DataTerraformRemoteStateEtcd {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateEtcd{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcd",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateEtcd_Override(d DataTerraformRemoteStateEtcd, scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcd",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateEtcd_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateEtcd",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateEtcd_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateEtcd",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateEtcdConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) A space-separated list of the etcd endpoints.
	// Experimental.
	Endpoints *string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Required) The path where to store the state.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) The password.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) The username.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Experimental.
type DataTerraformRemoteStateEtcdV3 interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateEtcdV3
type jsiiProxy_DataTerraformRemoteStateEtcdV3 struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateEtcdV3(scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdV3Config) DataTerraformRemoteStateEtcdV3 {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateEtcdV3{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateEtcdV3_Override(d DataTerraformRemoteStateEtcdV3, scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateEtcdV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateEtcdV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateEtcdV3Config struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The list of 'etcd' endpoints which to connect to.
	// Experimental.
	Endpoints *[]*string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Optional) The path to a PEM-encoded CA bundle with which to verify certificates of TLS-enabled etcd servers.
	// Experimental.
	CacertPath *string `field:"optional" json:"cacertPath" yaml:"cacertPath"`
	// (Optional) The path to a PEM-encoded certificate to provide to etcd for secure client identification.
	// Experimental.
	CertPath *string `field:"optional" json:"certPath" yaml:"certPath"`
	// (Optional) The path to a PEM-encoded key to provide to etcd for secure client identification.
	// Experimental.
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
	// (Optional) Whether to lock state access.
	//
	// Defaults to true.
	// Experimental.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Password used to connect to the etcd cluster.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) An optional prefix to be added to keys when to storing state in etcd.
	//
	// Defaults to "".
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) Username used to connect to the etcd cluster.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Experimental.
type DataTerraformRemoteStateGcs interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateGcs
type jsiiProxy_DataTerraformRemoteStateGcs struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateGcs(scope constructs.Construct, id *string, config *DataTerraformRemoteStateGcsConfig) DataTerraformRemoteStateGcs {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateGcs{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateGcs",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateGcs_Override(d DataTerraformRemoteStateGcs, scope constructs.Construct, id *string, config *DataTerraformRemoteStateGcsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateGcs",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateGcs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateGcs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateGcs_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateGcs",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateGcs) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateGcsConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The name of the GCS bucket.
	//
	// This name must be globally unique.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) A temporary [OAuth 2.0 access token] obtained from the Google Authorization server, i.e. the Authorization: Bearer token used to authenticate HTTP requests to GCP APIs. This is an alternative to credentials. If both are specified, access_token will be used over the credentials field.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// (Optional) Local path to Google Cloud Platform account credentials in JSON format.
	//
	// If unset, Google Application Default Credentials are used.
	// The provided credentials must have Storage Object Admin role on the bucket.
	//
	// Warning: if using the Google Cloud Platform provider as well,
	// it will also pick up the GOOGLE_CREDENTIALS environment variable.
	// Experimental.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// (Optional) A 32 byte base64 encoded 'customer supplied encryption key' used to encrypt all state.
	// Experimental.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// (Optional) The service account to impersonate for accessing the State Bucket.
	//
	// You must have roles/iam.serviceAccountTokenCreator role on that account for the impersonation to succeed.
	// If you are using a delegation chain, you can specify that using the impersonate_service_account_delegates field.
	// Alternatively, this can be specified using the GOOGLE_IMPERSONATE_SERVICE_ACCOUNT environment variable.
	// Experimental.
	ImpersonateServiceAccount *string `field:"optional" json:"impersonateServiceAccount" yaml:"impersonateServiceAccount"`
	// (Optional) The delegation chain for an impersonating a service account.
	// Experimental.
	ImpersonateServiceAccountDelegates *[]*string `field:"optional" json:"impersonateServiceAccountDelegates" yaml:"impersonateServiceAccountDelegates"`
	// (Optional) GCS prefix inside the bucket.
	//
	// Named states for workspaces are stored in an object called <prefix>/<name>.tfstate.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Experimental.
type DataTerraformRemoteStateHttp interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateHttp
type jsiiProxy_DataTerraformRemoteStateHttp struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateHttp(scope constructs.Construct, id *string, config *DataTerraformRemoteStateHttpConfig) DataTerraformRemoteStateHttp {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateHttp{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateHttp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateHttp_Override(d DataTerraformRemoteStateHttp, scope constructs.Construct, id *string, config *DataTerraformRemoteStateHttpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateHttp",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateHttp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateHttp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateHttp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateHttp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateHttp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateHttpConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The address of the REST endpoint.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// (Optional) The address of the lock REST endpoint.
	//
	// Defaults to disabled.
	// Experimental.
	LockAddress *string `field:"optional" json:"lockAddress" yaml:"lockAddress"`
	// (Optional) The HTTP method to use when locking.
	//
	// Defaults to LOCK.
	// Experimental.
	LockMethod *string `field:"optional" json:"lockMethod" yaml:"lockMethod"`
	// (Optional) The password for HTTP basic authentication.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) The number of HTTP request retries.
	//
	// Defaults to 2.
	// Experimental.
	RetryMax *float64 `field:"optional" json:"retryMax" yaml:"retryMax"`
	// (Optional) The maximum time in seconds to wait between HTTP request attempts.
	//
	// Defaults to 30.
	// Experimental.
	RetryWaitMax *float64 `field:"optional" json:"retryWaitMax" yaml:"retryWaitMax"`
	// (Optional) The minimum time in seconds to wait between HTTP request attempts.
	//
	// Defaults to 1.
	// Experimental.
	RetryWaitMin *float64 `field:"optional" json:"retryWaitMin" yaml:"retryWaitMin"`
	// (Optional) Whether to skip TLS verification.
	//
	// Defaults to false.
	// Experimental.
	SkipCertVerification *bool `field:"optional" json:"skipCertVerification" yaml:"skipCertVerification"`
	// (Optional) The address of the unlock REST endpoint.
	//
	// Defaults to disabled.
	// Experimental.
	UnlockAddress *string `field:"optional" json:"unlockAddress" yaml:"unlockAddress"`
	// (Optional) The HTTP method to use when unlocking.
	//
	// Defaults to UNLOCK.
	// Experimental.
	UnlockMethod *string `field:"optional" json:"unlockMethod" yaml:"unlockMethod"`
	// (Optional) HTTP method to use when updating state.
	//
	// Defaults to POST.
	// Experimental.
	UpdateMethod *string `field:"optional" json:"updateMethod" yaml:"updateMethod"`
	// (Optional) The username for HTTP basic authentication.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Experimental.
type DataTerraformRemoteStateLocal interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateLocal
type jsiiProxy_DataTerraformRemoteStateLocal struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateLocal(scope constructs.Construct, id *string, config *DataTerraformRemoteStateLocalConfig) DataTerraformRemoteStateLocal {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateLocal{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateLocal",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateLocal_Override(d DataTerraformRemoteStateLocal, scope constructs.Construct, id *string, config *DataTerraformRemoteStateLocalConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateLocal",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateLocal_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateLocal",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateLocal_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateLocal",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateLocal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateLocalConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Path where the state file is stored.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// (Optional) The path to non-default workspaces.
	// Experimental.
	WorkspaceDir *string `field:"optional" json:"workspaceDir" yaml:"workspaceDir"`
}

// Experimental.
type DataTerraformRemoteStateManta interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateManta
type jsiiProxy_DataTerraformRemoteStateManta struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateManta(scope constructs.Construct, id *string, config *DataTerraformRemoteStateMantaConfig) DataTerraformRemoteStateManta {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateManta{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateManta",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateManta_Override(d DataTerraformRemoteStateManta, scope constructs.Construct, id *string, config *DataTerraformRemoteStateMantaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateManta",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateManta_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateManta",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateManta_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateManta",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateManta) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateMantaConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
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

// Experimental.
type DataTerraformRemoteStateOss interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateOss
type jsiiProxy_DataTerraformRemoteStateOss struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateOss(scope constructs.Construct, id *string, config *DataTerraformRemoteStateOssConfig) DataTerraformRemoteStateOss {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateOss{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateOss",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateOss_Override(d DataTerraformRemoteStateOss, scope constructs.Construct, id *string, config *DataTerraformRemoteStateOssConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateOss",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateOss_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateOss",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateOss_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateOss",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateOss) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateOssConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Experimental.
	EcsRoleName *string `field:"optional" json:"ecsRoleName" yaml:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Experimental.
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `field:"optional" json:"tablestoreEndpoint" yaml:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `field:"optional" json:"tablestoreTable" yaml:"tablestoreTable"`
}

// Experimental.
type DataTerraformRemoteStatePg interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStatePg
type jsiiProxy_DataTerraformRemoteStatePg struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStatePg(scope constructs.Construct, id *string, config *DataTerraformRemoteStatePgConfig) DataTerraformRemoteStatePg {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStatePg{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStatePg",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStatePg_Override(d DataTerraformRemoteStatePg, scope constructs.Construct, id *string, config *DataTerraformRemoteStatePgConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStatePg",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStatePg_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStatePg",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStatePg_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStatePg",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStatePg) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

// Experimental.
type DataTerraformRemoteStateRemoteConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Experimental.
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `field:"required" json:"workspaces" yaml:"workspaces"`
	// Experimental.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

// Experimental.
type DataTerraformRemoteStateS3 interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateS3
type jsiiProxy_DataTerraformRemoteStateS3 struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateS3(scope constructs.Construct, id *string, config *DataTerraformRemoteStateS3Config) DataTerraformRemoteStateS3 {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateS3{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateS3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateS3_Override(d DataTerraformRemoteStateS3, scope constructs.Construct, id *string, config *DataTerraformRemoteStateS3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateS3",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateS3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateS3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateS3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateS3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateS3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateS3Config struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Name of the S3 Bucket.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Path to the state file inside the S3 Bucket.
	//
	// When using a non-default workspace, the state path will be /workspace_key_prefix/workspace_name/key.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Optional) AWS access key.
	//
	// If configured, must also configure secret_key.
	// This can also be sourced from
	// the AWS_ACCESS_KEY_ID environment variable,
	// AWS shared credentials file (e.g. ~/.aws/credentials),
	// or AWS shared configuration file (e.g. ~/.aws/config).
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) Canned ACL to be applied to the state file.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// (Optional) IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	// Experimental.
	AssumeRolePolicy *string `field:"optional" json:"assumeRolePolicy" yaml:"assumeRolePolicy"`
	// (Optional) Custom endpoint for the AWS DynamoDB API.
	//
	// This can also be sourced from the AWS_DYNAMODB_ENDPOINT environment variable.
	// Experimental.
	DynamodbEndpoint *string `field:"optional" json:"dynamodbEndpoint" yaml:"dynamodbEndpoint"`
	// (Optional) Name of DynamoDB Table to use for state locking and consistency.
	//
	// The table must have a partition key named LockID with type of String.
	// If not configured, state locking will be disabled.
	// Experimental.
	DynamodbTable *string `field:"optional" json:"dynamodbTable" yaml:"dynamodbTable"`
	// (Optional) Enable server side encryption of the state file.
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// (Optional) Custom endpoint for the AWS S3 API.
	//
	// This can also be sourced from the AWS_S3_ENDPOINT environment variable.
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// (Optional) External identifier to use when assuming the role.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// (Optional) Enable path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>).
	// Experimental.
	ForcePathStyle *bool `field:"optional" json:"forcePathStyle" yaml:"forcePathStyle"`
	// (Optional) Custom endpoint for the AWS Identity and Access Management (IAM) API.
	//
	// This can also be sourced from the AWS_IAM_ENDPOINT environment variable.
	// Experimental.
	IamEndpoint *string `field:"optional" json:"iamEndpoint" yaml:"iamEndpoint"`
	// (Optional) Amazon Resource Name (ARN) of a Key Management Service (KMS) Key to use for encrypting the state.
	//
	// Note that if this value is specified,
	// Terraform will need kms:Encrypt, kms:Decrypt and kms:GenerateDataKey permissions on this KMS key.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// (Optional) The maximum number of times an AWS API request is retried on retryable failure.
	//
	// Defaults to 5.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// (Optional) Name of AWS profile in AWS shared credentials file (e.g. ~/.aws/credentials) or AWS shared configuration file (e.g. ~/.aws/config) to use for credentials and/or configuration. This can also be sourced from the AWS_PROFILE environment variable.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// AWS Region of the S3 Bucket and DynamoDB Table (if used).
	//
	// This can also
	// be sourced from the AWS_DEFAULT_REGION and AWS_REGION environment
	// variables.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// (Optional) Amazon Resource Name (ARN) of the IAM Role to assume.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// (Optional) AWS secret access key.
	//
	// If configured, must also configure access_key.
	// This can also be sourced from
	// the AWS_SECRET_ACCESS_KEY environment variable,
	// AWS shared credentials file (e.g. ~/.aws/credentials),
	// or AWS shared configuration file (e.g. ~/.aws/config)
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// (Optional) Session name to use when assuming the role.
	// Experimental.
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// (Optional) Path to the AWS shared credentials file.
	//
	// Defaults to ~/.aws/credentials.
	// Experimental.
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// (Optional) Skip credentials validation via the STS API.
	// Experimental.
	SkipCredentialsValidation *bool `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// (Optional) Skip usage of EC2 Metadata API.
	// Experimental.
	SkipMetadataApiCheck *bool `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// (Optional) The key to use for encrypting state with Server-Side Encryption with Customer-Provided Keys (SSE-C).
	//
	// This is the base64-encoded value of the key, which must decode to 256 bits.
	// This can also be sourced from the AWS_SSE_CUSTOMER_KEY environment variable,
	// which is recommended due to the sensitivity of the value.
	// Setting it inside a terraform file will cause it to be persisted to disk in terraform.tfstate.
	// Experimental.
	SseCustomerKey *string `field:"optional" json:"sseCustomerKey" yaml:"sseCustomerKey"`
	// (Optional) Custom endpoint for the AWS Security Token Service (STS) API.
	//
	// This can also be sourced from the AWS_STS_ENDPOINT environment variable.
	// Experimental.
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
	// (Optional) Multi-Factor Authentication (MFA) token.
	//
	// This can also be sourced from the AWS_SESSION_TOKEN environment variable.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// (Optional) Prefix applied to the state path inside the bucket.
	//
	// This is only relevant when using a non-default workspace. Defaults to env:
	// Experimental.
	WorkspaceKeyPrefix *string `field:"optional" json:"workspaceKeyPrefix" yaml:"workspaceKeyPrefix"`
}

// Experimental.
type DataTerraformRemoteStateSwift interface {
	TerraformRemoteState
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateSwift
type jsiiProxy_DataTerraformRemoteStateSwift struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateSwift(scope constructs.Construct, id *string, config *DataTerraformRemoteStateSwiftConfig) DataTerraformRemoteStateSwift {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateSwift{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateSwift",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateSwift_Override(d DataTerraformRemoteStateSwift, scope constructs.Construct, id *string, config *DataTerraformRemoteStateSwiftConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateSwift",
		[]interface{}{scope, id, config},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataTerraformRemoteStateSwift_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateSwift",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTerraformRemoteStateSwift_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.DataTerraformRemoteStateSwift",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTerraformRemoteStateSwift) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateSwiftConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// Experimental.
	Container *string `field:"required" json:"container" yaml:"container"`
	// Experimental.
	ApplicationCredentialId *string `field:"optional" json:"applicationCredentialId" yaml:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `field:"optional" json:"applicationCredentialName" yaml:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `field:"optional" json:"applicationCredentialSecret" yaml:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `field:"optional" json:"archiveContainer" yaml:"archiveContainer"`
	// Experimental.
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// Experimental.
	CacertFile *string `field:"optional" json:"cacertFile" yaml:"cacertFile"`
	// Experimental.
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// Experimental.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// Experimental.
	DefaultDomain *string `field:"optional" json:"defaultDomain" yaml:"defaultDomain"`
	// Experimental.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Experimental.
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
	// Experimental.
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	ProjectDomainId *string `field:"optional" json:"projectDomainId" yaml:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `field:"optional" json:"projectDomainName" yaml:"projectDomainName"`
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Experimental.
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Experimental.
	TenantName *string `field:"optional" json:"tenantName" yaml:"tenantName"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Experimental.
	UserDomainId *string `field:"optional" json:"userDomainId" yaml:"userDomainId"`
	// Experimental.
	UserDomainName *string `field:"optional" json:"userDomainName" yaml:"userDomainName"`
	// Experimental.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Experimental.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

// Default resolver implementation.
// Experimental.
type DefaultTokenResolver interface {
	ITokenResolver
	// Resolve a tokenized list.
	// Experimental.
	ResolveList(xs *[]*string, context IResolveContext) interface{}
	// Resolve a tokenized map.
	// Experimental.
	ResolveMap(xs *map[string]interface{}, context IResolveContext) interface{}
	// Resolve a tokenized number list.
	// Experimental.
	ResolveNumberList(xs *[]*float64, context IResolveContext) interface{}
	// Resolve string fragments to Tokens.
	// Experimental.
	ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{}
	// Default Token resolution.
	//
	// Resolve the Token, recurse into whatever it returns,
	// then finally post-process it.
	// Experimental.
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy struct for DefaultTokenResolver
type jsiiProxy_DefaultTokenResolver struct {
	jsiiProxy_ITokenResolver
}

// Experimental.
func NewDefaultTokenResolver(concat IFragmentConcatenator) DefaultTokenResolver {
	_init_.Initialize()

	j := jsiiProxy_DefaultTokenResolver{}

	_jsii_.Create(
		"cdktf.DefaultTokenResolver",
		[]interface{}{concat},
		&j,
	)

	return &j
}

// Experimental.
func NewDefaultTokenResolver_Override(d DefaultTokenResolver, concat IFragmentConcatenator) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DefaultTokenResolver",
		[]interface{}{concat},
		d,
	)
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveList(xs *[]*string, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveList",
		[]interface{}{xs, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveMap(xs *map[string]interface{}, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveMap",
		[]interface{}{xs, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveNumberList(xs *[]*float64, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveNumberList",
		[]interface{}{xs, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveString",
		[]interface{}{fragments, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

// Properties to string encodings.
// Experimental.
type EncodingOptions struct {
	// A hint for the Token's purpose when stringifying it.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

// Experimental.
type EtcdBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EtcdBackend
type jsiiProxy_EtcdBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_EtcdBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewEtcdBackend(scope constructs.Construct, props *EtcdBackendProps) EtcdBackend {
	_init_.Initialize()

	j := jsiiProxy_EtcdBackend{}

	_jsii_.Create(
		"cdktf.EtcdBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEtcdBackend_Override(e EtcdBackend, scope constructs.Construct, props *EtcdBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.EtcdBackend",
		[]interface{}{scope, props},
		e,
	)
}

// Experimental.
func EtcdBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.EtcdBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func EtcdBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.EtcdBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EtcdBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		e,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EtcdBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EtcdBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state in etcd 2.x at a given path.
//
// This backend does not support state locking.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/etcd
// Experimental.
type EtcdBackendProps struct {
	// (Required) A space-separated list of the etcd endpoints.
	// Experimental.
	Endpoints *string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Required) The path where to store the state.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) The password.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) The username.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Experimental.
type EtcdV3Backend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EtcdV3Backend
type jsiiProxy_EtcdV3Backend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_EtcdV3Backend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewEtcdV3Backend(scope constructs.Construct, props *EtcdV3BackendProps) EtcdV3Backend {
	_init_.Initialize()

	j := jsiiProxy_EtcdV3Backend{}

	_jsii_.Create(
		"cdktf.EtcdV3Backend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEtcdV3Backend_Override(e EtcdV3Backend, scope constructs.Construct, props *EtcdV3BackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.EtcdV3Backend",
		[]interface{}{scope, props},
		e,
	)
}

// Experimental.
func EtcdV3Backend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.EtcdV3Backend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func EtcdV3Backend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.EtcdV3Backend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdV3Backend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EtcdV3Backend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		e,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdV3Backend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EtcdV3Backend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EtcdV3Backend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdV3Backend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdV3Backend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdV3Backend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state in the etcd KV store with a given prefix.
//
// This backend supports state locking.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/etcdv3
// Experimental.
type EtcdV3BackendProps struct {
	// (Required) The list of 'etcd' endpoints which to connect to.
	// Experimental.
	Endpoints *[]*string `field:"required" json:"endpoints" yaml:"endpoints"`
	// (Optional) The path to a PEM-encoded CA bundle with which to verify certificates of TLS-enabled etcd servers.
	// Experimental.
	CacertPath *string `field:"optional" json:"cacertPath" yaml:"cacertPath"`
	// (Optional) The path to a PEM-encoded certificate to provide to etcd for secure client identification.
	// Experimental.
	CertPath *string `field:"optional" json:"certPath" yaml:"certPath"`
	// (Optional) The path to a PEM-encoded key to provide to etcd for secure client identification.
	// Experimental.
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
	// (Optional) Whether to lock state access.
	//
	// Defaults to true.
	// Experimental.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Password used to connect to the etcd cluster.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) An optional prefix to be added to keys when to storing state in etcd.
	//
	// Defaults to "".
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// (Optional) Username used to connect to the etcd cluster.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Experimental.
type Fn interface {
}

// The jsii proxy struct for Fn
type jsiiProxy_Fn struct {
	_ byte // padding
}

// Experimental.
func NewFn() Fn {
	_init_.Initialize()

	j := jsiiProxy_Fn{}

	_jsii_.Create(
		"cdktf.Fn",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFn_Override(f Fn) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Fn",
		nil, // no parameters
		f,
	)
}

// {@link https://www.terraform.io/docs/language/functions/abs.html abs} returns the absolute value of the given number.
// Experimental.
func Fn_Abs(value *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"abs",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/abspath.html abspath} takes a string containing a filesystem path and converts it to an absolute path.
// Experimental.
func Fn_Abspath(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"abspath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/alltrue.html alltrue} returns true if all elements in a given collection are true or "true".
// Experimental.
func Fn_Alltrue(values *[]interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"alltrue",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/anytrue.html anytrue} returns true if any element in a given collection is true or "true".
// Experimental.
func Fn_Anytrue(value *[]interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"anytrue",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64decode.html base64decode} takes a string containing a Base64 character sequence and returns the original string.
// Experimental.
func Fn_Base64decode(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64decode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64encode.html base64encode} takes a string containing a Base64 character sequence and returns the original string.
// Experimental.
func Fn_Base64encode(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64encode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64gzip.html base64gzip} compresses a string with gzip and then encodes the result in Base64 encoding.
// Experimental.
func Fn_Base64gzip(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64gzip",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64sha256.html base64sha256} computes the SHA256 hash of a given string and encodes it with Base64.
// Experimental.
func Fn_Base64sha256(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64sha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64sha512.html base64sha512} computes the SHA512 hash of a given string and encodes it with Base64.
// Experimental.
func Fn_Base64sha512(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64sha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/basename.html basename} takes a string containing a filesystem path and removes all except the last portion from it.
// Experimental.
func Fn_Basename(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"basename",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/bcrypt.html bcrypt} computes a hash of the given string using the Blowfish cipher, returning a string in the Modular Crypt Format usually expected in the shadow password file on many Unix systems.
// Experimental.
func Fn_Bcrypt(value *string, cost *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"bcrypt",
		[]interface{}{value, cost},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/can.html can} evaluates the given expression and returns a boolean value indicating whether the expression produced a result without any errors.
// Experimental.
func Fn_Can(expression interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"can",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/ceil.html ceil} returns the closest whole number that is greater than or equal to the given value, which may be a fraction.
// Experimental.
func Fn_Ceil(value *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"ceil",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/chomp.html chomp} removes newline characters at the end of a string.
// Experimental.
func Fn_Chomp(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"chomp",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/chunklist.html chunklist} splits a single list into fixed-size chunks, returning a list of lists.
// Experimental.
func Fn_Chunklist(value *[]interface{}, chunkSize *float64) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"chunklist",
		[]interface{}{value, chunkSize},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrhost.html cidrhost} calculates a full host IP address for a given host number within a given IP network address prefix.
// Experimental.
func Fn_Cidrhost(prefix *string, hostnum *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrhost",
		[]interface{}{prefix, hostnum},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrnetmask.html cidrnetmask} converts an IPv4 address prefix given in CIDR notation into a subnet mask address.
// Experimental.
func Fn_Cidrnetmask(prefix *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrnetmask",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrsubnet.html cidrsubnet} calculates a subnet address within given IP network address prefix.
// Experimental.
func Fn_Cidrsubnet(prefix *string, newbits *float64, netnum *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrsubnet",
		[]interface{}{prefix, newbits, netnum},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrsubnets.html cidrsubnets} calculates a sequence of consecutive IP address ranges within a particular CIDR prefix.
// Experimental.
func Fn_Cidrsubnets(prefix *string, newbits *[]*float64) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrsubnets",
		[]interface{}{prefix, newbits},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/coalesce.html coalesce} takes any number of arguments and returns the first one that isn't null or an empty string.
// Experimental.
func Fn_Coalesce(value *[]interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"coalesce",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/coalescelist.html coalescelist} takes any number of list arguments and returns the first one that isn't empty.
// Experimental.
func Fn_Coalescelist(value *[]*[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"coalescelist",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/compact.html compact} takes a list of strings and returns a new list with any empty string elements removed.
// Experimental.
func Fn_Compact(value *[]*string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"compact",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/concat.html concat} takes two or more lists and combines them into a single list.
// Experimental.
func Fn_Concat(value *[]*[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"concat",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/contains.html contains} determines whether a given list or set contains a given single value as one of its elements.
// Experimental.
func Fn_Contains(list interface{}, value interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"contains",
		[]interface{}{list, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/csvdecode.html csvdecode} decodes a string containing CSV-formatted data and produces a list of maps representing that data.
// Experimental.
func Fn_Csvdecode(value *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"csvdecode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/dirname.html dirname} takes a string containing a filesystem path and removes the last portion from it.
// Experimental.
func Fn_Dirname(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"dirname",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/distinct.html distinct} takes a list and returns a new list with any duplicate elements removed.
// Experimental.
func Fn_Distinct(list interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"distinct",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/element.html element} retrieves a single element from a list.
// Experimental.
func Fn_Element(list interface{}, index *float64) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"element",
		[]interface{}{list, index},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/file.html file} takes a string containing a filesystem path and removes all except the last portion from it.
// Experimental.
func Fn_File(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"file",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filebase64.html filebase64} reads the contents of a file at the given path and returns them as a base64-encoded string.
// Experimental.
func Fn_Filebase64(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filebase64",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filebase64sha256.html filebase64sha256} is a variant of base64sha256 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filebase64sha256(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filebase64sha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filebase64sha512.html filebase64sha512} is a variant of base64sha512 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filebase64sha512(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filebase64sha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/fileexists.html fileexists} determines whether a file exists at a given path.
// Experimental.
func Fn_Fileexists(value *string) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"fileexists",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filemd5.html filemd5} is a variant of md5 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filemd5(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filemd5",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/fileset.html fileset} enumerates a set of regular file names given a path and pattern.
// Experimental.
func Fn_Fileset(path *string, pattern *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"fileset",
		[]interface{}{path, pattern},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filesha1.html filesha1} is a variant of sha1 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filesha1(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filesha1",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filesha256.html filesha256} is a variant of sha256 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filesha256(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filesha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filesha512.html filesha512} is a variant of sha512 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filesha512(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filesha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/flatten.html flatten} takes a list and replaces any elements that are lists with a flattened sequence of the list contents.
// Experimental.
func Fn_Flatten(list interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"flatten",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/floor.html floor} returns the closest whole number that is less than or equal to the given value, which may be a fraction.
// Experimental.
func Fn_Floor(value *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"floor",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/format.html format} produces a string by formatting a number of other values according to a specification string.
// Experimental.
func Fn_Format(spec *string, values *[]interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"format",
		[]interface{}{spec, values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/formatdate.html formatdate} converts a timestamp into a different time format.
// Experimental.
func Fn_Formatdate(spec *string, timestamp *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"formatdate",
		[]interface{}{spec, timestamp},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/formatlist.html formatlist} produces a list of strings by formatting a number of other values according to a specification string.
// Experimental.
func Fn_Formatlist(spec *string, values *[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"formatlist",
		[]interface{}{spec, values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/indent.html indent} adds a given number of spaces to the beginnings of all but the first line in a given multi-line string.
// Experimental.
func Fn_Indent(indentation *float64, value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"indent",
		[]interface{}{indentation, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/index.html index} finds the element index for a given value in a list.
// Experimental.
func Fn_Index(list interface{}, value interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"index",
		[]interface{}{list, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/join.html join} produces a string by concatenating together all elements of a given list of strings with the given delimiter.
// Experimental.
func Fn_Join(separator *string, value *[]*string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"join",
		[]interface{}{separator, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/jsondecode.html jsondecode} interprets a given string as JSON, returning a representation of the result of decoding that string.
// Experimental.
func Fn_Jsondecode(value *string) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"jsondecode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/jsonencode.html jsonencode} encodes a given value to a string using JSON syntax.
// Experimental.
func Fn_Jsonencode(value interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"jsonencode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/keys.html keys} takes a map and returns a list containing the keys from that map.
// Experimental.
func Fn_Keys(map_ interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"keys",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/length.html length} determines the length of a given list, map, or string.
// Experimental.
func Fn_LengthOf(value interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"lengthOf",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/log.html log} returns the logarithm of a given number in a given base.
// Experimental.
func Fn_Log(value *float64, base *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"log",
		[]interface{}{value, base},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/lookup.html lookup} retrieves the value of a single element from a map, given its key. If the given key does not exist, the given default value is returned instead.
// Experimental.
func Fn_Lookup(value interface{}, key interface{}, defaultValue interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"lookup",
		[]interface{}{value, key, defaultValue},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/lower.html lower} converts all cased letters in the given string to lowercase.
// Experimental.
func Fn_Lower(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"lower",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/matchkeys.html matchkeys} constructs a new list by taking a subset of elements from one list whose indexes match the corresponding indexes of values in another list.
// Experimental.
func Fn_Matchkeys(valuesList interface{}, keysList interface{}, searchSet interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"matchkeys",
		[]interface{}{valuesList, keysList, searchSet},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/max.html max} takes one or more numbers and returns the greatest number from the set.
// Experimental.
func Fn_Max(values *[]*float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"max",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/md5.html md5} computes the MD5 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Md5(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"md5",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/merge.html merge} takes an arbitrary number of maps or objects, and returns a single map or object that contains a merged set of elements from all arguments.
// Experimental.
func Fn_MergeLists(values *[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"mergeLists",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/merge.html merge} takes an arbitrary number of maps or objects, and returns a single map or object that contains a merged set of elements from all arguments.
// Experimental.
func Fn_MergeMaps(values *[]interface{}) *map[string]*string {
	_init_.Initialize()

	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"mergeMaps",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/min.html min} takes one or more numbers and returns the smallest number from the set.
// Experimental.
func Fn_Min(values *[]*float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"min",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/nonsensitive.html nonsensitive} takes a sensitive value and returns a copy of that value with the sensitive marking removed, thereby exposing the sensitive value.
// Experimental.
func Fn_Nonsensitive(expression interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"nonsensitive",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/one.html one} takes a list, set, or tuple value with either zero or one elements.
// Experimental.
func Fn_One(list interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"one",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/parseint.html parseInt} parses the given string as a representation of an integer in the specified base and returns the resulting number. The base must be between 2 and 62 inclusive.
// Experimental.
func Fn_ParseInt(value *string, base *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"parseInt",
		[]interface{}{value, base},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/pathexpand.html pathexpand} takes a string containing a filesystem path and removes the last portion from it.
// Experimental.
func Fn_Pathexpand(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"pathexpand",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/pow.html pow} calculates an exponent, by raising its first argument to the power of the second argument.
// Experimental.
func Fn_Pow(value *float64, power *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"pow",
		[]interface{}{value, power},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/range.html range} generates a list of numbers using a start value, a limit value, and a step value.
// Experimental.
func Fn_Range(start *float64, limit *float64, step *float64) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"range",
		[]interface{}{start, limit, step},
		&returns,
	)

	return returns
}

// Use this function to wrap a string and escape it properly for the use in Terraform This is only needed in certain scenarios (e.g., if you have unescaped double quotes in the string).
// Experimental.
func Fn_RawString(str *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"rawString",
		[]interface{}{str},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/regex.html regex} applies a regular expression to a string and returns the matching substrings in pattern.
// Experimental.
func Fn_Regex(pattern *string, value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"regex",
		[]interface{}{pattern, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/regexall.html regexall} applies a regular expression to a string and returns a list of all matches.
// Experimental.
func Fn_Regexall(pattern *string, value *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"regexall",
		[]interface{}{pattern, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/replace.html replace} searches a given string for another given substring, and replaces each occurrence with a given replacement string.
// Experimental.
func Fn_Replace(value *string, substring *string, replacement *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"replace",
		[]interface{}{value, substring, replacement},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/reverse.html reverse} takes a sequence and produces a new sequence of the same length with all of the same elements as the given sequence but in reverse order.
// Experimental.
func Fn_Reverse(values interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"reverse",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/rsadecrypt.html rsadecrypt} decrypts an RSA-encrypted ciphertext, returning the corresponding cleartext.
// Experimental.
func Fn_Rsadecrypt(ciphertext *string, privatekey *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"rsadecrypt",
		[]interface{}{ciphertext, privatekey},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sensitive.html sensitive} takes any value and returns a copy of it marked so that Terraform will treat it as sensitive, with the same meaning and behavior as for sensitive input variables.
// Experimental.
func Fn_Sensitive(expression interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sensitive",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/setintersection.html setintersection} function takes multiple sets and produces a single set containing only the elements that all of the given sets have in common.
// Experimental.
func Fn_Setintersection(values *[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setintersection",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/setproduct.html setproduct} function finds all of the possible combinations of elements from all of the given sets by computing the Cartesian product.
// Experimental.
func Fn_Setproduct(values *[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setproduct",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/slice.html setsubtract} function returns a new set containing the elements from the first set that are not present in the second set.
// Experimental.
func Fn_Setsubtract(minuend interface{}, subtrahend interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setsubtract",
		[]interface{}{minuend, subtrahend},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/setunion.html setunion} function takes multiple sets and produces a single set containing the elements from all of the given sets.
// Experimental.
func Fn_Setunion(values *[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setunion",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sha1.html sha1} computes the SHA1 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Sha1(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sha1",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sha256.html sha256} computes the SHA256 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Sha256(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sha512.html sha512} computes the SHA512 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Sha512(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/signum.html signum} determines the sign of a number, returning a number between -1 and 1 to represent the sign.
// Experimental.
func Fn_Signum(value *float64) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"signum",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/slice.html slice} extracts some consecutive elements from within a list.
// Experimental.
func Fn_Slice(list interface{}, startindex *float64, endindex *float64) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"slice",
		[]interface{}{list, startindex, endindex},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sort.html sort} takes a list of strings and returns a new list with those strings sorted lexicographically.
// Experimental.
func Fn_Sort(list interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sort",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/split.html split} produces a list by dividing a given string at all occurrences of a given separator.
// Experimental.
func Fn_Split(seperator *string, value *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"split",
		[]interface{}{seperator, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/strrev.html strrev} reverses the characters in a string.
// Experimental.
func Fn_Strrev(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"strrev",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/substr.html substr} extracts a substring from a given string by offset and length.
// Experimental.
func Fn_Substr(value *string, offset *float64, length *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"substr",
		[]interface{}{value, offset, length},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sum.html sum} takes a list or set of numbers and returns the sum of those numbers.
// Experimental.
func Fn_Sum(list interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sum",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/templatefile.html templatefile} reads the file at the given path and renders its content as a template using a supplied set of template variables.
// Experimental.
func Fn_Templatefile(path *string, vars interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"templatefile",
		[]interface{}{path, vars},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/textdecodebase64.html textdecodebase64} function decodes a string that was previously Base64-encoded, and then interprets the result as characters in a specified character encoding.
// Experimental.
func Fn_Textdecodebase64(value *string, encodingName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"textdecodebase64",
		[]interface{}{value, encodingName},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/textencodebase64.html textencodebase64}  encodes the unicode characters in a given string using a specified character encoding, returning the result base64 encoded because Terraform language strings are always sequences of unicode characters.
// Experimental.
func Fn_Textencodebase64(value *string, encodingName *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"textencodebase64",
		[]interface{}{value, encodingName},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/timeadd.html timeadd} adds a duration to a timestamp, returning a new timestamp.
// Experimental.
func Fn_Timeadd(timestamp *string, duration *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"timeadd",
		[]interface{}{timestamp, duration},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/timestamp.html timestamp} returns a UTC timestamp string in RFC 3339 format.
// Experimental.
func Fn_Timestamp() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"timestamp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/title.html title} converts the first letter of each word in the given string to uppercase.
// Experimental.
func Fn_Title(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"title",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tobool.html tobool} converts its argument to a boolean value.
// Experimental.
func Fn_Tobool(expression interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tobool",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tolist.html tolist} converts its argument to a list value.
// Experimental.
func Fn_Tolist(expression interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tolist",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tomap.html tomap} converts its argument to a map value.
// Experimental.
func Fn_Tomap(expression interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tomap",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tonumber.html tonumber} converts its argument to a number value.
// Experimental.
func Fn_Tonumber(expression interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tonumber",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/toset.html toset} converts its argument to a set value.
// Experimental.
func Fn_Toset(expression interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"toset",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tostring.html tostring} converts its argument to a string value.
// Experimental.
func Fn_Tostring(expression interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tostring",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/transpose.html transpose} takes a map of lists of strings and swaps the keys and values to produce a new map of lists of strings.
// Experimental.
func Fn_Transpose(value interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"transpose",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trim.html trim} removes the specified characters from the start and end of the given string.
// Experimental.
func Fn_Trim(value *string, replacement *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trim",
		[]interface{}{value, replacement},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trimprefix.html trimprefix} removes the specified prefix from the start of the given string.
// Experimental.
func Fn_Trimprefix(value *string, prefix *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trimprefix",
		[]interface{}{value, prefix},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trimspace.html trimspace} removes any space characters from the start and end of the given string.
// Experimental.
func Fn_Trimspace(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trimspace",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trimsuffix.html trimsuffix} removes the specified suffix from the end of the given string.
// Experimental.
func Fn_Trimsuffix(value *string, suffix *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trimsuffix",
		[]interface{}{value, suffix},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/try.html try} evaluates all of its argument expressions in turn and returns the result of the first one that does not produce any errors.
// Experimental.
func Fn_Try(expression *[]interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"try",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/upper.html upper} converts all cased letters in the given string to uppercase.
// Experimental.
func Fn_Upper(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"upper",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/urlencode.html urlencode} applies URL encoding to a given string.
// Experimental.
func Fn_Urlencode(value *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"urlencode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/uuid.html uuid} generates a unique identifier string.
// Experimental.
func Fn_Uuid() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"uuid",
		nil, // no parameters
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/uuidv5.html uuidv5} generates a unique identifier string.
// Experimental.
func Fn_Uuidv5(namespace *string, name *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"uuidv5",
		[]interface{}{namespace, name},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/values.html values} takes a map and returns a list containing the values of the elements in that map.
// Experimental.
func Fn_Values(value interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"values",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/yamldecode.html yamldecode} parses a string as a subset of YAML, and produces a representation of its value.
// Experimental.
func Fn_Yamldecode(value *string) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"yamldecode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/yamlencode.html yamlencode} encodes a given value to a string using JSON syntax.
// Experimental.
func Fn_Yamlencode(value interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"yamlencode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/zipmap.html zipmap} constructs a map from a list of keys and a corresponding list of values.
// Experimental.
func Fn_Zipmap(keyslist interface{}, valueslist interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"zipmap",
		[]interface{}{keyslist, valueslist},
		&returns,
	)

	return returns
}

// Experimental.
type GcsBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GcsBackend
type jsiiProxy_GcsBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_GcsBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewGcsBackend(scope constructs.Construct, props *GcsBackendProps) GcsBackend {
	_init_.Initialize()

	j := jsiiProxy_GcsBackend{}

	_jsii_.Create(
		"cdktf.GcsBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGcsBackend_Override(g GcsBackend, scope constructs.Construct, props *GcsBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.GcsBackend",
		[]interface{}{scope, props},
		g,
	)
}

// Experimental.
func GcsBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.GcsBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func GcsBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.GcsBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcsBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GcsBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		g,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcsBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GcsBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcsBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcsBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcsBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcsBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state as an object in a configurable prefix in a pre-existing bucket on Google Cloud Storage (GCS).
//
// The bucket must exist prior to configuring the backend.
//
// This backend supports state locking.
//
// Warning! It is highly recommended that you enable Object Versioning on the GCS bucket
// to allow for state recovery in the case of accidental deletions and human error.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/gcs
// Experimental.
type GcsBackendProps struct {
	// (Required) The name of the GCS bucket.
	//
	// This name must be globally unique.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// (Optional) A temporary [OAuth 2.0 access token] obtained from the Google Authorization server, i.e. the Authorization: Bearer token used to authenticate HTTP requests to GCP APIs. This is an alternative to credentials. If both are specified, access_token will be used over the credentials field.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// (Optional) Local path to Google Cloud Platform account credentials in JSON format.
	//
	// If unset, Google Application Default Credentials are used.
	// The provided credentials must have Storage Object Admin role on the bucket.
	//
	// Warning: if using the Google Cloud Platform provider as well,
	// it will also pick up the GOOGLE_CREDENTIALS environment variable.
	// Experimental.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// (Optional) A 32 byte base64 encoded 'customer supplied encryption key' used to encrypt all state.
	// Experimental.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// (Optional) The service account to impersonate for accessing the State Bucket.
	//
	// You must have roles/iam.serviceAccountTokenCreator role on that account for the impersonation to succeed.
	// If you are using a delegation chain, you can specify that using the impersonate_service_account_delegates field.
	// Alternatively, this can be specified using the GOOGLE_IMPERSONATE_SERVICE_ACCOUNT environment variable.
	// Experimental.
	ImpersonateServiceAccount *string `field:"optional" json:"impersonateServiceAccount" yaml:"impersonateServiceAccount"`
	// (Optional) The delegation chain for an impersonating a service account.
	// Experimental.
	ImpersonateServiceAccountDelegates *[]*string `field:"optional" json:"impersonateServiceAccountDelegates" yaml:"impersonateServiceAccountDelegates"`
	// (Optional) GCS prefix inside the bucket.
	//
	// Named states for workspaces are stored in an object called <prefix>/<name>.tfstate.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Experimental.
type HttpBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HttpBackend
type jsiiProxy_HttpBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_HttpBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpBackend(scope constructs.Construct, props *HttpBackendProps) HttpBackend {
	_init_.Initialize()

	j := jsiiProxy_HttpBackend{}

	_jsii_.Create(
		"cdktf.HttpBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpBackend_Override(h HttpBackend, scope constructs.Construct, props *HttpBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.HttpBackend",
		[]interface{}{scope, props},
		h,
	)
}

// Experimental.
func HttpBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.HttpBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func HttpBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.HttpBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HttpBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		h,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HttpBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HttpBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state using a simple REST client.
//
// State will be fetched via GET, updated via POST, and purged with DELETE.
// The method used for updating is configurable.
//
// This backend optionally supports state locking.
// When locking support is enabled it will use LOCK and UNLOCK requests providing the lock info in the body.
// The endpoint should return a 423: Locked or 409: Conflict with the holding lock info when
// it's already taken, 200: OK for success. Any other status will be considered an error.
// The ID of the holding lock info will be added as a query parameter to state updates requests.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/http
// Experimental.
type HttpBackendProps struct {
	// (Required) The address of the REST endpoint.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// (Optional) The address of the lock REST endpoint.
	//
	// Defaults to disabled.
	// Experimental.
	LockAddress *string `field:"optional" json:"lockAddress" yaml:"lockAddress"`
	// (Optional) The HTTP method to use when locking.
	//
	// Defaults to LOCK.
	// Experimental.
	LockMethod *string `field:"optional" json:"lockMethod" yaml:"lockMethod"`
	// (Optional) The password for HTTP basic authentication.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// (Optional) The number of HTTP request retries.
	//
	// Defaults to 2.
	// Experimental.
	RetryMax *float64 `field:"optional" json:"retryMax" yaml:"retryMax"`
	// (Optional) The maximum time in seconds to wait between HTTP request attempts.
	//
	// Defaults to 30.
	// Experimental.
	RetryWaitMax *float64 `field:"optional" json:"retryWaitMax" yaml:"retryWaitMax"`
	// (Optional) The minimum time in seconds to wait between HTTP request attempts.
	//
	// Defaults to 1.
	// Experimental.
	RetryWaitMin *float64 `field:"optional" json:"retryWaitMin" yaml:"retryWaitMin"`
	// (Optional) Whether to skip TLS verification.
	//
	// Defaults to false.
	// Experimental.
	SkipCertVerification *bool `field:"optional" json:"skipCertVerification" yaml:"skipCertVerification"`
	// (Optional) The address of the unlock REST endpoint.
	//
	// Defaults to disabled.
	// Experimental.
	UnlockAddress *string `field:"optional" json:"unlockAddress" yaml:"unlockAddress"`
	// (Optional) The HTTP method to use when unlocking.
	//
	// Defaults to UNLOCK.
	// Experimental.
	UnlockMethod *string `field:"optional" json:"unlockMethod" yaml:"unlockMethod"`
	// (Optional) HTTP method to use when updating state.
	//
	// Defaults to POST.
	// Experimental.
	UpdateMethod *string `field:"optional" json:"updateMethod" yaml:"updateMethod"`
	// (Optional) The username for HTTP basic authentication.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Interface for lazy untyped value producers.
// Experimental.
type IAnyProducer interface {
	// Produce the value.
	// Experimental.
	Produce(context IResolveContext) interface{}
}

// The jsii proxy for IAnyProducer
type jsiiProxy_IAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IAnyProducer) Produce(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Represents an Aspect.
// Experimental.
type IAspect interface {
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy for IAspect
type jsiiProxy_IAspect struct {
	_ byte // padding
}

func (i *jsiiProxy_IAspect) Visit(node constructs.IConstruct) {
	_jsii_.InvokeVoid(
		i,
		"visit",
		[]interface{}{node},
	)
}

// The file provisioner copies files or directories from the machine running Terraform to the newly created resource.
//
// The file provisioner supports both ssh and winrm type connections.
//
// See {@link https://www.terraform.io/language/resources/provisioners/file file}
// Experimental.
type IFileProvisioner interface {
	// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
	// Experimental.
	Connection() interface{}
	// The destination path to write to on the remote system.
	//
	// See Destination Paths below for more information.
	// Experimental.
	Content() *string
	// The source file or directory.
	//
	// Specify it either relative to the current working directory or as an absolute path.
	// This argument cannot be combined with content.
	// Experimental.
	Destination() *string
	// The direct content to copy on the destination.
	//
	// If destination is a file, the content will be written on that file.
	// In case of a directory, a file named tf-file-content is created inside that directory.
	// We recommend using a file as the destination when using content.
	// This argument cannot be combined with source.
	// Experimental.
	Source() *string
	// Experimental.
	Type() *string
}

// The jsii proxy for IFileProvisioner
type jsiiProxy_IFileProvisioner struct {
	_ byte // padding
}

func (j *jsiiProxy_IFileProvisioner) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileProvisioner) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Function used to concatenate symbols in the target document language.
//
// Interface so it could potentially be exposed over jsii.
// Experimental.
type IFragmentConcatenator interface {
	// Join the fragment on the left and on the right.
	// Experimental.
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy for IFragmentConcatenator
type jsiiProxy_IFragmentConcatenator struct {
	_ byte // padding
}

func (i *jsiiProxy_IFragmentConcatenator) Join(left interface{}, right interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Experimental.
type IInterpolatingParent interface {
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
}

// The jsii proxy for IInterpolatingParent
type jsiiProxy_IInterpolatingParent struct {
	_ byte // padding
}

func (i *jsiiProxy_IInterpolatingParent) InterpolationForAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Interface for lazy list producers.
// Experimental.
type IListProducer interface {
	// Produce the list value.
	// Experimental.
	Produce(context IResolveContext) *[]*string
}

// The jsii proxy for IListProducer
type jsiiProxy_IListProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IListProducer) Produce(context IResolveContext) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// The local-exec provisioner invokes a local executable after a resource is created.
//
// This invokes a process on the machine running Terraform, not on the resource.
//
// See {@link https://www.terraform.io/language/resources/provisioners/local-exec local-exec}
// Experimental.
type ILocalExecProvisioner interface {
	// This is the command to execute.
	//
	// It can be provided as a relative path to the current working directory or as an absolute path.
	// It is evaluated in a shell, and can use environment variables or Terraform variables.
	// Experimental.
	Command() *string
	// A record of key value pairs representing the environment of the executed command.
	//
	// It inherits the current process environment.
	// Experimental.
	Environment() *map[string]*string
	// If provided, this is a list of interpreter arguments used to execute the command.
	//
	// The first argument is the interpreter itself.
	// It can be provided as a relative path to the current working directory or as an absolute path
	// The remaining arguments are appended prior to the command.
	// This allows building command lines of the form "/bin/bash", "-c", "echo foo".
	// If interpreter is unspecified, sensible defaults will be chosen based on the system OS.
	// Experimental.
	Interpreter() *[]*string
	// Experimental.
	Type() *string
	// If provided, specifies when Terraform will execute the command.
	//
	// For example, when = destroy specifies that the provisioner will run when the associated resource is destroyed.
	// Experimental.
	When() *string
	// If provided, specifies the working directory where command will be executed.
	//
	// It can be provided as a relative path to the current working directory or as an absolute path.
	// The directory must exist.
	// Experimental.
	WorkingDir() *string
}

// The jsii proxy for ILocalExecProvisioner
type jsiiProxy_ILocalExecProvisioner struct {
	_ byte // padding
}

func (j *jsiiProxy_ILocalExecProvisioner) Command() *string {
	var returns *string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) Interpreter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"interpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) When() *string {
	var returns *string
	_jsii_.Get(
		j,
		"when",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalExecProvisioner) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

// Experimental.
type IManifest interface {
	// Experimental.
	Stacks() *map[string]*StackManifest
	// Experimental.
	Version() *string
}

// The jsii proxy for IManifest
type jsiiProxy_IManifest struct {
	_ byte // padding
}

func (j *jsiiProxy_IManifest) Stacks() *map[string]*StackManifest {
	var returns *map[string]*StackManifest
	_jsii_.Get(
		j,
		"stacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManifest) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Interface for lazy number producers.
// Experimental.
type INumberProducer interface {
	// Produce the number value.
	// Experimental.
	Produce(context IResolveContext) *float64
}

// The jsii proxy for INumberProducer
type jsiiProxy_INumberProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_INumberProducer) Produce(context IResolveContext) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// A Token that can post-process the complete resolved value, after resolve() has recursed over it.
// Experimental.
type IPostProcessor interface {
	// Process the completely resolved value, after full recursion/resolution has happened.
	// Experimental.
	PostProcess(input interface{}, context IResolveContext) interface{}
}

// The jsii proxy for IPostProcessor
type jsiiProxy_IPostProcessor struct {
	_ byte // padding
}

func (i *jsiiProxy_IPostProcessor) PostProcess(input interface{}, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"postProcess",
		[]interface{}{input, context},
		&returns,
	)

	return returns
}

// The remote-exec provisioner invokes a script on a remote resource after it is created.
//
// This can be used to run a configuration management tool, bootstrap into a cluster, etc
// The remote-exec provisioner requires a connection and supports both ssh and winrm.
//
// See {@link https://www.terraform.io/language/resources/provisioners/remote-exec remote-exec}
// Experimental.
type IRemoteExecProvisioner interface {
	// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
	//
	// A connection must be provided here or in the parent resource.
	// Experimental.
	Connection() interface{}
	// This is a list of command strings.
	//
	// They are executed in the order they are provided.
	// This cannot be provided with script or scripts.
	// Experimental.
	Inline() *[]*string
	// This is a path (relative or absolute) to a local script that will be copied to the remote resource and then executed.
	//
	// This cannot be provided with inline or scripts.
	// Experimental.
	Script() *string
	// This is a list of paths (relative or absolute) to local scripts that will be copied to the remote resource and then executed.
	//
	// They are executed in the order they are provided.
	// This cannot be provided with inline or script.
	// Experimental.
	Scripts() *[]*string
	// Experimental.
	Type() *string
}

// The jsii proxy for IRemoteExecProvisioner
type jsiiProxy_IRemoteExecProvisioner struct {
	_ byte // padding
}

func (j *jsiiProxy_IRemoteExecProvisioner) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Inline() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Scripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemoteExecProvisioner) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Experimental.
type IRemoteWorkspace interface {
}

// The jsii proxy for IRemoteWorkspace
type jsiiProxy_IRemoteWorkspace struct {
	_ byte // padding
}

// Interface for values that can be resolvable later.
//
// Tokens are special objects that participate in synthesis.
// Experimental.
type IResolvable interface {
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
}

// The jsii proxy for IResolvable
type jsiiProxy_IResolvable struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolvable) Resolve(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResolvable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolvable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

// Current resolution context for tokens.
// Experimental.
type IResolveContext interface {
	// Use this postprocessor after the entire token structure has been resolved.
	// Experimental.
	RegisterPostProcessor(postProcessor IPostProcessor)
	// Resolve an inner object.
	// Experimental.
	Resolve(x interface{}) interface{}
	// TerraformIterators can be passed for block attributes and normal list attributes both require different handling when the iterable variable is accessed e.g. a dynamic block needs each.key while a for expression just needs key.
	// Experimental.
	IteratorContext() *string
	// Experimental.
	SetIteratorContext(i *string)
	// True when we are still preparing, false if we're rendering the final output.
	// Experimental.
	Preparing() *bool
	// The scope from which resolution has been initiated.
	// Experimental.
	Scope() constructs.IConstruct
	// True when ${} should be ommitted (because already inside them), false otherwise.
	// Experimental.
	SuppressBraces() *bool
	// Experimental.
	SetSuppressBraces(s *bool)
}

// The jsii proxy for IResolveContext
type jsiiProxy_IResolveContext struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolveContext) RegisterPostProcessor(postProcessor IPostProcessor) {
	_jsii_.InvokeVoid(
		i,
		"registerPostProcessor",
		[]interface{}{postProcessor},
	)
}

func (i *jsiiProxy_IResolveContext) Resolve(x interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolveContext) IteratorContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iteratorContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) SetIteratorContext(val *string) {
	_jsii_.Set(
		j,
		"iteratorContext",
		val,
	)
}

func (j *jsiiProxy_IResolveContext) Preparing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preparing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Scope() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) SuppressBraces() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"suppressBraces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) SetSuppressBraces(val *bool) {
	_jsii_.Set(
		j,
		"suppressBraces",
		val,
	)
}

// Experimental.
type IResource interface {
	constructs.IConstruct
	// The stack in which this resource is defined.
	// Experimental.
	Stack() TerraformStack
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResource) Stack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Experimental.
type IResourceConstructor interface {
}

// The jsii proxy for IResourceConstructor
type jsiiProxy_IResourceConstructor struct {
	_ byte // padding
}

// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
//
// See {@link https://www.terraform.io/language/resources/provisioners/connection connection}
// Experimental.
type ISSHProvisionerConnection interface {
	// Set to false to disable using ssh-agent to authenticate.
	//
	// On Windows the only supported SSH authentication agent is Pageant.
	// Experimental.
	Agent() *string
	// The preferred identity from the ssh agent for authentication.
	// Experimental.
	AgentIdentity() *string
	// The contents of a signed CA Certificate.
	//
	// The certificate argument must be used in conjunction with a bastion_private_key.
	// These can be loaded from a file on disk using the the file function.
	// Experimental.
	BastionCertificate() *string
	// Setting this enables the bastion Host connection.
	//
	// The provisioner will connect to bastion_host first, and then connect from there to host.
	// Experimental.
	BastionHost() *string
	// The public key from the remote host or the signing CA, used to verify the host connection.
	// Experimental.
	BastionHostKey() *string
	// The password to use for the bastion host.
	// Experimental.
	BastionPassword() *string
	// The port to use connect to the bastion host.
	// Experimental.
	BastionPort() *float64
	// The contents of an SSH key file to use for the bastion host.
	//
	// These can be loaded from a file on disk using the file function.
	// Experimental.
	BastionPrivateKey() *string
	// The user for the connection to the bastion host.
	// Experimental.
	BastionUser() *string
	// The contents of a signed CA Certificate.
	//
	// The certificate argument must be used in conjunction with a private_key.
	// These can be loaded from a file on disk using the the file function.
	// Experimental.
	Certificate() *string
	// The address of the resource to connect to.
	// Experimental.
	Host() *string
	// The public key from the remote host or the signing CA, used to verify the connection.
	// Experimental.
	HostKey() *string
	// The password to use for the connection.
	// Experimental.
	Password() *string
	// The port to connect to.
	// Experimental.
	Port() *float64
	// The contents of an SSH key to use for the connection.
	//
	// These can be loaded from a file on disk using the file function.
	// This takes preference over password if provided.
	// Experimental.
	PrivateKey() *string
	// Setting this enables the SSH over HTTP connection.
	//
	// This host will be connected to first, and then the host or bastion_host connection will be made from there.
	// Experimental.
	ProxyHost() *string
	// The port to use connect to the proxy host.
	// Experimental.
	ProxyPort() *float64
	// The ssh connection also supports the following fields to facilitate connections by SSH over HTTP proxy.
	// Experimental.
	ProxyScheme() *string
	// The username to use connect to the private proxy host.
	//
	// This argument should be specified only if authentication is required for the HTTP Proxy server.
	// Experimental.
	ProxyUserName() *string
	// The password to use connect to the private proxy host.
	//
	// This argument should be specified only if authentication is required for the HTTP Proxy server.
	// Experimental.
	ProxyUserPassword() *string
	// The path used to copy scripts meant for remote execution.
	//
	// Refer to {@link https://www.terraform.io/language/resources/provisioners/connection#how-provisioners-execute-remote-scripts How Provisioners Execute Remote Scripts below for more details}
	// Experimental.
	ScriptPath() *string
	// The target platform to connect to.
	//
	// Valid values are "windows" and "unix".
	// If the platform is set to windows, the default script_path is c:\windows\temp\terraform_%RAND%.cmd, assuming the SSH default shell is cmd.exe.
	// If the SSH default shell is PowerShell, set script_path to "c:/windows/temp/terraform_%RAND%.ps1"
	// Experimental.
	TargetPlatform() *string
	// The timeout to wait for the connection to become available.
	//
	// Should be provided as a string (e.g., "30s" or "5m".)
	// Experimental.
	Timeout() *string
	// The connection type.
	//
	// Valid values are "ssh" and "winrm".
	// Provisioners typically assume that the remote system runs Microsoft Windows when using WinRM.
	// Behaviors based on the SSH target_platform will force Windows-specific behavior for WinRM, unless otherwise specified.
	// Experimental.
	Type() *string
	// The user to use for the connection.
	// Experimental.
	User() *string
}

// The jsii proxy for ISSHProvisionerConnection
type jsiiProxy_ISSHProvisionerConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_ISSHProvisionerConnection) Agent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) AgentIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionHostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionHostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bastionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) BastionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ProxyUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) ScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) TargetPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISSHProvisionerConnection) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

// Experimental.
type IScopeCallback interface {
}

// The jsii proxy for IScopeCallback
type jsiiProxy_IScopeCallback struct {
	_ byte // padding
}

// Encodes information how a certain Stack should be deployed inspired by AWS CDK v2 implementation (synth functionality was removed in constructs v10).
// Experimental.
type IStackSynthesizer interface {
	// Synthesize the associated stack to the session.
	// Experimental.
	Synthesize(session ISynthesisSession)
}

// The jsii proxy for IStackSynthesizer
type jsiiProxy_IStackSynthesizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

// Interface for lazy string producers.
// Experimental.
type IStringProducer interface {
	// Produce the string value.
	// Experimental.
	Produce(context IResolveContext) *string
}

// The jsii proxy for IStringProducer
type jsiiProxy_IStringProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStringProducer) Produce(context IResolveContext) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Represents a single session of synthesis.
//
// Passed into `TerraformStack.onSynthesize()` methods.
// originally from aws/constructs lib v3.3.126 (synth functionality was removed in constructs v10)
// Experimental.
type ISynthesisSession interface {
	// Experimental.
	Manifest() Manifest
	// The output directory for this synthesis session.
	// Experimental.
	Outdir() *string
	// Experimental.
	SkipValidation() *bool
}

// The jsii proxy for ISynthesisSession
type jsiiProxy_ISynthesisSession struct {
	_ byte // padding
}

func (j *jsiiProxy_ISynthesisSession) Manifest() Manifest {
	var returns Manifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) SkipValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}

// Experimental.
type ITerraformAddressable interface {
	// Experimental.
	Fqn() *string
}

// The jsii proxy for ITerraformAddressable
type jsiiProxy_ITerraformAddressable struct {
	_ byte // padding
}

func (j *jsiiProxy_ITerraformAddressable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

// Experimental.
type ITerraformDependable interface {
	ITerraformAddressable
}

// The jsii proxy for ITerraformDependable
type jsiiProxy_ITerraformDependable struct {
	jsiiProxy_ITerraformAddressable
}

// Experimental.
type ITerraformIterator interface {
}

// The jsii proxy for ITerraformIterator
type jsiiProxy_ITerraformIterator struct {
	_ byte // padding
}

// Experimental.
type ITerraformResource interface {
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(c *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(d *[]*string)
	// Experimental.
	ForEach() ITerraformIterator
	// Experimental.
	SetForEach(f ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(l *TerraformResourceLifecycle)
	// Experimental.
	Provider() TerraformProvider
	// Experimental.
	SetProvider(p TerraformProvider)
	// Experimental.
	TerraformResourceType() *string
}

// The jsii proxy for ITerraformResource
type jsiiProxy_ITerraformResource struct {
	_ byte // padding
}

func (i *jsiiProxy_ITerraformResource) InterpolationForAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITerraformResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) ForEach() ITerraformIterator {
	var returns ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetForEach(val ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) Lifecycle() *TerraformResourceLifecycle {
	var returns *TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetLifecycle(val *TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Interface to apply operation to tokens in a string.
//
// Interface so it can be exported via jsii.
// Experimental.
type ITokenMapper interface {
	// Replace a single token.
	// Experimental.
	MapToken(t IResolvable) interface{}
}

// The jsii proxy for ITokenMapper
type jsiiProxy_ITokenMapper struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenMapper) MapToken(t IResolvable) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"mapToken",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// How to resolve tokens.
// Experimental.
type ITokenResolver interface {
	// Resolve a tokenized list.
	// Experimental.
	ResolveList(l *[]*string, context IResolveContext) interface{}
	// Resolve a tokenized map.
	// Experimental.
	ResolveMap(m *map[string]interface{}, context IResolveContext) interface{}
	// Resolve a tokenized number list.
	// Experimental.
	ResolveNumberList(l *[]*float64, context IResolveContext) interface{}
	// Resolve a string with at least one stringified token in it.
	//
	// (May use concatenation).
	// Experimental.
	ResolveString(s TokenizedStringFragments, context IResolveContext) interface{}
	// Resolve a single token.
	// Experimental.
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy for ITokenResolver
type jsiiProxy_ITokenResolver struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenResolver) ResolveList(l *[]*string, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveList",
		[]interface{}{l, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveMap(m *map[string]interface{}, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveMap",
		[]interface{}{m, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveNumberList(l *[]*float64, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveNumberList",
		[]interface{}{l, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveString(s TokenizedStringFragments, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveString",
		[]interface{}{s, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

// Most provisioners require access to the remote resource via SSH or WinRM and expect a nested connection block with details about how to connect.
//
// See {@link https://www.terraform.io/language/resources/provisioners/connection connection}
// Experimental.
type IWinrmProvisionerConnection interface {
	// The CA certificate to validate against.
	// Experimental.
	Cacert() *string
	// The address of the resource to connect to.
	// Experimental.
	Host() *string
	// Set to true to connect using HTTPS instead of HTTP.
	// Experimental.
	Https() *bool
	// Set to true to skip validating the HTTPS certificate chain.
	// Experimental.
	Insecure() *bool
	// The password to use for the connection.
	// Experimental.
	Password() *string
	// The port to connect to.
	// Experimental.
	Port() *float64
	// The path used to copy scripts meant for remote execution.
	//
	// Refer to {@link https://www.terraform.io/language/resources/provisioners/connection#how-provisioners-execute-remote-scripts How Provisioners Execute Remote Scripts below for more details}
	// Experimental.
	ScriptPath() *string
	// The timeout to wait for the connection to become available.
	//
	// Should be provided as a string (e.g., "30s" or "5m".)
	// Experimental.
	Timeout() *string
	// The connection type.
	//
	// Valid values are "ssh" and "winrm".
	// Provisioners typically assume that the remote system runs Microsoft Windows when using WinRM.
	// Behaviors based on the SSH target_platform will force Windows-specific behavior for WinRM, unless otherwise specified.
	// Experimental.
	Type() *string
	// Set to true to use NTLM authentication rather than default (basic authentication), removing the requirement for basic authentication to be enabled within the target guest.
	//
	// Refer to Authentication for Remote Connections in the Windows App Development documentation for more details.
	// Experimental.
	UseNtlm() *bool
	// The user to use for the connection.
	// Experimental.
	User() *string
}

// The jsii proxy for IWinrmProvisionerConnection
type jsiiProxy_IWinrmProvisionerConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Cacert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Https() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"https",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Insecure() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) ScriptPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) UseNtlm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useNtlm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWinrmProvisionerConnection) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

// Lazily produce a value.
//
// Can be used to return a string, list or numeric value whose actual value
// will only be calculated later, during synthesis.
// Experimental.
type Lazy interface {
}

// The jsii proxy struct for Lazy
type jsiiProxy_Lazy struct {
	_ byte // padding
}

// Experimental.
func NewLazy() Lazy {
	_init_.Initialize()

	j := jsiiProxy_Lazy{}

	_jsii_.Create(
		"cdktf.Lazy",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewLazy_Override(l Lazy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Lazy",
		nil, // no parameters
		l,
	)
}

// Produces a lazy token from an untyped value.
// Experimental.
func Lazy_AnyValue(producer IAnyProducer, options *LazyAnyValueOptions) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"anyValue",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Returns a list-ified token for a lazy value.
// Experimental.
func Lazy_ListValue(producer IListProducer, options *LazyListValueOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"listValue",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Returns a numberified token for a lazy value.
// Experimental.
func Lazy_NumberValue(producer INumberProducer) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"numberValue",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

// Returns a stringified token for a lazy value.
// Experimental.
func Lazy_StringValue(producer IStringProducer, options *LazyStringValueOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"stringValue",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Options for creating lazy untyped tokens.
// Experimental.
type LazyAnyValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	// Experimental.
	OmitEmptyArray *bool `field:"optional" json:"omitEmptyArray" yaml:"omitEmptyArray"`
}

// Experimental.
type LazyBase interface {
	IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	AddPostProcessor(postProcessor IPostProcessor)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context IResolveContext) interface{}
	// Experimental.
	ResolveLazy(context IResolveContext) interface{}
	// Turn this Token into JSON.
	//
	// Called automatically when JSON.stringify() is called on a Token.
	// Experimental.
	ToJSON() interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LazyBase
type jsiiProxy_LazyBase struct {
	jsiiProxy_IResolvable
}

func (j *jsiiProxy_LazyBase) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


// Experimental.
func NewLazyBase_Override(l LazyBase) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.LazyBase",
		nil, // no parameters
		l,
	)
}

func (l *jsiiProxy_LazyBase) AddPostProcessor(postProcessor IPostProcessor) {
	_jsii_.InvokeVoid(
		l,
		"addPostProcessor",
		[]interface{}{postProcessor},
	)
}

func (l *jsiiProxy_LazyBase) Resolve(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LazyBase) ResolveLazy(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolveLazy",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LazyBase) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LazyBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for creating a lazy list token.
// Experimental.
type LazyListValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced list is empty, return 'undefined' instead.
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
}

// Options for creating a lazy string token.
// Experimental.
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

// Experimental.
type ListTerraformIterator interface {
	TerraformIterator
	// Returns the currenty entry in the list or set that is being iterated over.
	//
	// For lists this is the same as `iterator.value`. If you need the index,
	// use count using the escape hatch:
	// https://www.terraform.io/cdktf/concepts/resources#escape-hatch
	// Experimental.
	Key() interface{}
	// Returns the value of the current item iterated over.
	// Experimental.
	Value() interface{}
	// Experimental.
	Dynamic(attributes *map[string]interface{}) IResolvable
	// Returns: the given attribute of the current item iterated over as any.
	// Experimental.
	GetAny(attribute *string) IResolvable
	// Returns: the given attribute of the current item iterated over as a map of any.
	// Experimental.
	GetAnyMap(attribute *string) *map[string]interface{}
	// Returns: the given attribute of the current item iterated over as a boolean.
	// Experimental.
	GetBoolean(attribute *string) IResolvable
	// Returns: the given attribute of the current item iterated over as a map of booleans.
	// Experimental.
	GetBooleanMap(attribute *string) *map[string]*bool
	// Returns: the given attribute of the current item iterated over as a (string) list.
	// Experimental.
	GetList(attribute *string) *[]*string
	// Returns: the given attribute of the current item iterated over as a map.
	// Experimental.
	GetMap(attribute *string) *map[string]interface{}
	// Returns: the given attribute of the current item iterated over as a number.
	// Experimental.
	GetNumber(attribute *string) *float64
	// Returns: the given attribute of the current item iterated over as a number list.
	// Experimental.
	GetNumberList(attribute *string) *[]*float64
	// Returns: the given attribute of the current item iterated over as a map of numbers.
	// Experimental.
	GetNumberMap(attribute *string) *map[string]*float64
	// Returns: the given attribute of the current item iterated over as a string.
	// Experimental.
	GetString(attribute *string) *string
	// Returns: the given attribute of the current item iterated over as a map of strings.
	// Experimental.
	GetStringMap(attribute *string) *map[string]*string
}

// The jsii proxy struct for ListTerraformIterator
type jsiiProxy_ListTerraformIterator struct {
	jsiiProxy_TerraformIterator
}

func (j *jsiiProxy_ListTerraformIterator) Key() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListTerraformIterator) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewListTerraformIterator(list interface{}) ListTerraformIterator {
	_init_.Initialize()

	j := jsiiProxy_ListTerraformIterator{}

	_jsii_.Create(
		"cdktf.ListTerraformIterator",
		[]interface{}{list},
		&j,
	)

	return &j
}

// Experimental.
func NewListTerraformIterator_Override(l ListTerraformIterator, list interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ListTerraformIterator",
		[]interface{}{list},
		l,
	)
}

// Creates a new iterator from a list.
// Experimental.
func ListTerraformIterator_FromList(list interface{}) ListTerraformIterator {
	_init_.Initialize()

	var returns ListTerraformIterator

	_jsii_.StaticInvoke(
		"cdktf.ListTerraformIterator",
		"fromList",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// Creates a new iterator from a map.
// Experimental.
func ListTerraformIterator_FromMap(map_ interface{}) MapTerraformIterator {
	_init_.Initialize()

	var returns MapTerraformIterator

	_jsii_.StaticInvoke(
		"cdktf.ListTerraformIterator",
		"fromMap",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) Dynamic(attributes *map[string]interface{}) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		l,
		"dynamic",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetAny(attribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		l,
		"getAny",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetAnyMap(attribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetBoolean(attribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		l,
		"getBoolean",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetBooleanMap(attribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetList(attribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getList",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetMap(attribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetNumber(attribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumber",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetNumberList(attribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberList",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetNumberMap(attribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetString(attribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getString",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListTerraformIterator) GetStringMap(attribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

// Experimental.
type LocalBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LocalBackend
type jsiiProxy_LocalBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_LocalBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewLocalBackend(scope constructs.Construct, props *LocalBackendProps) LocalBackend {
	_init_.Initialize()

	j := jsiiProxy_LocalBackend{}

	_jsii_.Create(
		"cdktf.LocalBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLocalBackend_Override(l LocalBackend, scope constructs.Construct, props *LocalBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.LocalBackend",
		[]interface{}{scope, props},
		l,
	)
}

// Experimental.
func LocalBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.LocalBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func LocalBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.LocalBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocalBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LocalBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		l,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, fromStack},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocalBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LocalBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocalBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocalBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocalBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocalBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The local backend stores state on the local filesystem, locks that state using system APIs, and performs operations locally.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/local
// Experimental.
type LocalBackendProps struct {
	// Path where the state file is stored.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// (Optional) The path to non-default workspaces.
	// Experimental.
	WorkspaceDir *string `field:"optional" json:"workspaceDir" yaml:"workspaceDir"`
}

// Experimental.
type Manifest interface {
	IManifest
	// Experimental.
	Outdir() *string
	// Experimental.
	Stacks() *map[string]*StackManifest
	// Experimental.
	Version() *string
	// Experimental.
	BuildManifest() IManifest
	// Experimental.
	ForStack(stack TerraformStack) *StackManifest
	// Experimental.
	WriteToFile()
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	jsiiProxy_IManifest
}

func (j *jsiiProxy_Manifest) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Manifest) Stacks() *map[string]*StackManifest {
	var returns *map[string]*StackManifest
	_jsii_.Get(
		j,
		"stacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Manifest) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewManifest(version *string, outdir *string) Manifest {
	_init_.Initialize()

	j := jsiiProxy_Manifest{}

	_jsii_.Create(
		"cdktf.Manifest",
		[]interface{}{version, outdir},
		&j,
	)

	return &j
}

// Experimental.
func NewManifest_Override(m Manifest, version *string, outdir *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Manifest",
		[]interface{}{version, outdir},
		m,
	)
}

func Manifest_FileName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Manifest",
		"fileName",
		&returns,
	)
	return returns
}

func Manifest_StackFileName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Manifest",
		"stackFileName",
		&returns,
	)
	return returns
}

func Manifest_StacksFolder() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Manifest",
		"stacksFolder",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Manifest) BuildManifest() IManifest {
	var returns IManifest

	_jsii_.Invoke(
		m,
		"buildManifest",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Manifest) ForStack(stack TerraformStack) *StackManifest {
	var returns *StackManifest

	_jsii_.Invoke(
		m,
		"forStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Manifest) WriteToFile() {
	_jsii_.InvokeVoid(
		m,
		"writeToFile",
		nil, // no parameters
	)
}

// Experimental.
type MantaBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MantaBackend
type jsiiProxy_MantaBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_MantaBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewMantaBackend(scope constructs.Construct, props *MantaBackendProps) MantaBackend {
	_init_.Initialize()

	j := jsiiProxy_MantaBackend{}

	_jsii_.Create(
		"cdktf.MantaBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMantaBackend_Override(m MantaBackend, scope constructs.Construct, props *MantaBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.MantaBackend",
		[]interface{}{scope, props},
		m,
	)
}

// Experimental.
func MantaBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.MantaBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func MantaBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.MantaBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MantaBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MantaBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		m,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MantaBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MantaBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MantaBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MantaBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MantaBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MantaBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

// Experimental.
type MapTerraformIterator interface {
	TerraformIterator
	// Returns the key of the current entry in the map that is being iterated over.
	// Experimental.
	Key() *string
	// Returns the value of the current item iterated over.
	// Experimental.
	Value() interface{}
	// Experimental.
	Dynamic(attributes *map[string]interface{}) IResolvable
	// Returns: the given attribute of the current item iterated over as any.
	// Experimental.
	GetAny(attribute *string) IResolvable
	// Returns: the given attribute of the current item iterated over as a map of any.
	// Experimental.
	GetAnyMap(attribute *string) *map[string]interface{}
	// Returns: the given attribute of the current item iterated over as a boolean.
	// Experimental.
	GetBoolean(attribute *string) IResolvable
	// Returns: the given attribute of the current item iterated over as a map of booleans.
	// Experimental.
	GetBooleanMap(attribute *string) *map[string]*bool
	// Returns: the given attribute of the current item iterated over as a (string) list.
	// Experimental.
	GetList(attribute *string) *[]*string
	// Returns: the given attribute of the current item iterated over as a map.
	// Experimental.
	GetMap(attribute *string) *map[string]interface{}
	// Returns: the given attribute of the current item iterated over as a number.
	// Experimental.
	GetNumber(attribute *string) *float64
	// Returns: the given attribute of the current item iterated over as a number list.
	// Experimental.
	GetNumberList(attribute *string) *[]*float64
	// Returns: the given attribute of the current item iterated over as a map of numbers.
	// Experimental.
	GetNumberMap(attribute *string) *map[string]*float64
	// Returns: the given attribute of the current item iterated over as a string.
	// Experimental.
	GetString(attribute *string) *string
	// Returns: the given attribute of the current item iterated over as a map of strings.
	// Experimental.
	GetStringMap(attribute *string) *map[string]*string
}

// The jsii proxy struct for MapTerraformIterator
type jsiiProxy_MapTerraformIterator struct {
	jsiiProxy_TerraformIterator
}

func (j *jsiiProxy_MapTerraformIterator) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapTerraformIterator) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewMapTerraformIterator(map_ interface{}) MapTerraformIterator {
	_init_.Initialize()

	j := jsiiProxy_MapTerraformIterator{}

	_jsii_.Create(
		"cdktf.MapTerraformIterator",
		[]interface{}{map_},
		&j,
	)

	return &j
}

// Experimental.
func NewMapTerraformIterator_Override(m MapTerraformIterator, map_ interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.MapTerraformIterator",
		[]interface{}{map_},
		m,
	)
}

// Creates a new iterator from a list.
// Experimental.
func MapTerraformIterator_FromList(list interface{}) ListTerraformIterator {
	_init_.Initialize()

	var returns ListTerraformIterator

	_jsii_.StaticInvoke(
		"cdktf.MapTerraformIterator",
		"fromList",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// Creates a new iterator from a map.
// Experimental.
func MapTerraformIterator_FromMap(map_ interface{}) MapTerraformIterator {
	_init_.Initialize()

	var returns MapTerraformIterator

	_jsii_.StaticInvoke(
		"cdktf.MapTerraformIterator",
		"fromMap",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) Dynamic(attributes *map[string]interface{}) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		m,
		"dynamic",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetAny(attribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		m,
		"getAny",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetAnyMap(attribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetBoolean(attribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		m,
		"getBoolean",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetBooleanMap(attribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetList(attribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getList",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetMap(attribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetNumber(attribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumber",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetNumberList(attribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberList",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetNumberMap(attribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetString(attribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getString",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapTerraformIterator) GetStringMap(attribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

// Experimental.
type NamedRemoteWorkspace interface {
	IRemoteWorkspace
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
}

// The jsii proxy struct for NamedRemoteWorkspace
type jsiiProxy_NamedRemoteWorkspace struct {
	jsiiProxy_IRemoteWorkspace
}

func (j *jsiiProxy_NamedRemoteWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewNamedRemoteWorkspace(name *string) NamedRemoteWorkspace {
	_init_.Initialize()

	j := jsiiProxy_NamedRemoteWorkspace{}

	_jsii_.Create(
		"cdktf.NamedRemoteWorkspace",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Experimental.
func NewNamedRemoteWorkspace_Override(n NamedRemoteWorkspace, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.NamedRemoteWorkspace",
		[]interface{}{name},
		n,
	)
}

func (j *jsiiProxy_NamedRemoteWorkspace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Experimental.
type NumberMap interface {
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Lookup(key *string) *float64
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NumberMap
type jsiiProxy_NumberMap struct {
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_NumberMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMap) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewNumberMap(terraformResource IInterpolatingParent, terraformAttribute *string) NumberMap {
	_init_.Initialize()

	j := jsiiProxy_NumberMap{}

	_jsii_.Create(
		"cdktf.NumberMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewNumberMap_Override(n NumberMap, terraformResource IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.NumberMap",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NumberMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NumberMap) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NumberMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMap) Lookup(key *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMap) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type NumberMapList interface {
	IInterpolatingParent
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) NumberMap
	// Experimental.
	InterpolationForAttribute(property *string) IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NumberMapList
type jsiiProxy_NumberMapList struct {
	jsiiProxy_IInterpolatingParent
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_NumberMapList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMapList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMapList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMapList) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMapList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewNumberMapList(terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NumberMapList {
	_init_.Initialize()

	j := jsiiProxy_NumberMapList{}

	_jsii_.Create(
		"cdktf.NumberMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewNumberMapList_Override(n NumberMapList, terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.NumberMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NumberMapList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NumberMapList) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NumberMapList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NumberMapList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMapList) Get(index *float64) NumberMap {
	var returns NumberMap

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMapList) InterpolationForAttribute(property *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMapList) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NumberMapList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type OssAssumeRole struct {
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Experimental.
	SessionExpiration *float64 `field:"optional" json:"sessionExpiration" yaml:"sessionExpiration"`
	// Experimental.
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
}

// Experimental.
type OssBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OssBackend
type jsiiProxy_OssBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_OssBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewOssBackend(scope constructs.Construct, props *OssBackendProps) OssBackend {
	_init_.Initialize()

	j := jsiiProxy_OssBackend{}

	_jsii_.Create(
		"cdktf.OssBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOssBackend_Override(o OssBackend, scope constructs.Construct, props *OssBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.OssBackend",
		[]interface{}{scope, props},
		o,
	)
}

// Experimental.
func OssBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.OssBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func OssBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.OssBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OssBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OssBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		o,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OssBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OssBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OssBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OssBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OssBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OssBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type OssBackendProps struct {
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Experimental.
	EcsRoleName *string `field:"optional" json:"ecsRoleName" yaml:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Experimental.
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `field:"optional" json:"tablestoreEndpoint" yaml:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `field:"optional" json:"tablestoreTable" yaml:"tablestoreTable"`
}

// Experimental.
type PgBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PgBackend
type jsiiProxy_PgBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_PgBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewPgBackend(scope constructs.Construct, props *PgBackendProps) PgBackend {
	_init_.Initialize()

	j := jsiiProxy_PgBackend{}

	_jsii_.Create(
		"cdktf.PgBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPgBackend_Override(p PgBackend, scope constructs.Construct, props *PgBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.PgBackend",
		[]interface{}{scope, props},
		p,
	)
}

// Experimental.
func PgBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.PgBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func PgBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.PgBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PgBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		p,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PgBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PgBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PgBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type PgBackendProps struct {
	// Experimental.
	ConnStr *string `field:"required" json:"connStr" yaml:"connStr"`
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `field:"optional" json:"skipSchemaCreation" yaml:"skipSchemaCreation"`
}

// Experimental.
type PrefixedRemoteWorkspaces interface {
	IRemoteWorkspace
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
}

// The jsii proxy struct for PrefixedRemoteWorkspaces
type jsiiProxy_PrefixedRemoteWorkspaces struct {
	jsiiProxy_IRemoteWorkspace
}

func (j *jsiiProxy_PrefixedRemoteWorkspaces) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrefixedRemoteWorkspaces(prefix *string) PrefixedRemoteWorkspaces {
	_init_.Initialize()

	j := jsiiProxy_PrefixedRemoteWorkspaces{}

	_jsii_.Create(
		"cdktf.PrefixedRemoteWorkspaces",
		[]interface{}{prefix},
		&j,
	)

	return &j
}

// Experimental.
func NewPrefixedRemoteWorkspaces_Override(p PrefixedRemoteWorkspaces, prefix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.PrefixedRemoteWorkspaces",
		[]interface{}{prefix},
		p,
	)
}

func (j *jsiiProxy_PrefixedRemoteWorkspaces) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

// Experimental.
type RemoteBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RemoteBackend
type jsiiProxy_RemoteBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_RemoteBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewRemoteBackend(scope constructs.Construct, props *RemoteBackendProps) RemoteBackend {
	_init_.Initialize()

	j := jsiiProxy_RemoteBackend{}

	_jsii_.Create(
		"cdktf.RemoteBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRemoteBackend_Override(r RemoteBackend, scope constructs.Construct, props *RemoteBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.RemoteBackend",
		[]interface{}{scope, props},
		r,
	)
}

// Experimental.
func RemoteBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.RemoteBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func RemoteBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.RemoteBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RemoteBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RemoteBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		r,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RemoteBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RemoteBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RemoteBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RemoteBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RemoteBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RemoteBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type RemoteBackendProps struct {
	// Experimental.
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `field:"required" json:"workspaces" yaml:"workspaces"`
	// Experimental.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

// Options to the resolve() operation.
//
// NOT the same as the ResolveContext; ResolveContext is exposed to Token
// implementors and resolution hooks, whereas this struct is just to bundle
// a number of things that would otherwise be arguments to resolve() in a
// readable way.
// Experimental.
type ResolveOptions struct {
	// The resolver to apply to any resolvable tokens found.
	// Experimental.
	Resolver ITokenResolver `field:"required" json:"resolver" yaml:"resolver"`
	// The scope from which resolution is performed.
	// Experimental.
	Scope constructs.IConstruct `field:"required" json:"scope" yaml:"scope"`
	// Whether the resolution is being executed during the prepare phase or not.
	// Experimental.
	Preparing *bool `field:"optional" json:"preparing" yaml:"preparing"`
}

// A construct which represents a resource.
// Experimental.
type Resource interface {
	constructs.Construct
	IResource
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The stack in which this resource is defined.
	// Experimental.
	Stack() TerraformStack
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	internal.Type__constructsConstruct
	jsiiProxy_IResource
}

func (j *jsiiProxy_Resource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resource) Stack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewResource_Override(r Resource, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Resource",
		[]interface{}{scope, id},
		r,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func Resource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Resource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type S3Backend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for S3Backend
type jsiiProxy_S3Backend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_S3Backend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3Backend(scope constructs.Construct, props *S3BackendProps) S3Backend {
	_init_.Initialize()

	j := jsiiProxy_S3Backend{}

	_jsii_.Create(
		"cdktf.S3Backend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Backend_Override(s S3Backend, scope constructs.Construct, props *S3BackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.S3Backend",
		[]interface{}{scope, props},
		s,
	)
}

// Experimental.
func S3Backend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.S3Backend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func S3Backend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.S3Backend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Backend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_S3Backend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		s,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Backend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_S3Backend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Backend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Backend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Backend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Backend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Stores the state as a given key in a given bucket on Amazon S3.
//
// This backend
// also supports state locking and consistency checking via Dynamo DB, which
// can be enabled by setting the dynamodb_table field to an existing DynamoDB
// table name. A single DynamoDB table can be used to lock multiple remote
// state files. Terraform generates key names that include the values of the
// bucket and key variables.
//
// Warning! It is highly recommended that you enable Bucket Versioning on the
// S3 bucket to allow for state recovery in the case of accidental deletions
// and human error.
//
// Read more about this backend in the Terraform docs:
// https://www.terraform.io/language/settings/backends/s3
// Experimental.
type S3BackendProps struct {
	// Name of the S3 Bucket.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Path to the state file inside the S3 Bucket.
	//
	// When using a non-default workspace, the state path will be /workspace_key_prefix/workspace_name/key.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Optional) AWS access key.
	//
	// If configured, must also configure secret_key.
	// This can also be sourced from
	// the AWS_ACCESS_KEY_ID environment variable,
	// AWS shared credentials file (e.g. ~/.aws/credentials),
	// or AWS shared configuration file (e.g. ~/.aws/config).
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) Canned ACL to be applied to the state file.
	// Experimental.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// (Optional) IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	// Experimental.
	AssumeRolePolicy *string `field:"optional" json:"assumeRolePolicy" yaml:"assumeRolePolicy"`
	// (Optional) Custom endpoint for the AWS DynamoDB API.
	//
	// This can also be sourced from the AWS_DYNAMODB_ENDPOINT environment variable.
	// Experimental.
	DynamodbEndpoint *string `field:"optional" json:"dynamodbEndpoint" yaml:"dynamodbEndpoint"`
	// (Optional) Name of DynamoDB Table to use for state locking and consistency.
	//
	// The table must have a partition key named LockID with type of String.
	// If not configured, state locking will be disabled.
	// Experimental.
	DynamodbTable *string `field:"optional" json:"dynamodbTable" yaml:"dynamodbTable"`
	// (Optional) Enable server side encryption of the state file.
	// Experimental.
	Encrypt *bool `field:"optional" json:"encrypt" yaml:"encrypt"`
	// (Optional) Custom endpoint for the AWS S3 API.
	//
	// This can also be sourced from the AWS_S3_ENDPOINT environment variable.
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// (Optional) External identifier to use when assuming the role.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// (Optional) Enable path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>).
	// Experimental.
	ForcePathStyle *bool `field:"optional" json:"forcePathStyle" yaml:"forcePathStyle"`
	// (Optional) Custom endpoint for the AWS Identity and Access Management (IAM) API.
	//
	// This can also be sourced from the AWS_IAM_ENDPOINT environment variable.
	// Experimental.
	IamEndpoint *string `field:"optional" json:"iamEndpoint" yaml:"iamEndpoint"`
	// (Optional) Amazon Resource Name (ARN) of a Key Management Service (KMS) Key to use for encrypting the state.
	//
	// Note that if this value is specified,
	// Terraform will need kms:Encrypt, kms:Decrypt and kms:GenerateDataKey permissions on this KMS key.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// (Optional) The maximum number of times an AWS API request is retried on retryable failure.
	//
	// Defaults to 5.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// (Optional) Name of AWS profile in AWS shared credentials file (e.g. ~/.aws/credentials) or AWS shared configuration file (e.g. ~/.aws/config) to use for credentials and/or configuration. This can also be sourced from the AWS_PROFILE environment variable.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// AWS Region of the S3 Bucket and DynamoDB Table (if used).
	//
	// This can also
	// be sourced from the AWS_DEFAULT_REGION and AWS_REGION environment
	// variables.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// (Optional) Amazon Resource Name (ARN) of the IAM Role to assume.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// (Optional) AWS secret access key.
	//
	// If configured, must also configure access_key.
	// This can also be sourced from
	// the AWS_SECRET_ACCESS_KEY environment variable,
	// AWS shared credentials file (e.g. ~/.aws/credentials),
	// or AWS shared configuration file (e.g. ~/.aws/config)
	// Experimental.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// (Optional) Session name to use when assuming the role.
	// Experimental.
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// (Optional) Path to the AWS shared credentials file.
	//
	// Defaults to ~/.aws/credentials.
	// Experimental.
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// (Optional) Skip credentials validation via the STS API.
	// Experimental.
	SkipCredentialsValidation *bool `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// (Optional) Skip usage of EC2 Metadata API.
	// Experimental.
	SkipMetadataApiCheck *bool `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// (Optional) The key to use for encrypting state with Server-Side Encryption with Customer-Provided Keys (SSE-C).
	//
	// This is the base64-encoded value of the key, which must decode to 256 bits.
	// This can also be sourced from the AWS_SSE_CUSTOMER_KEY environment variable,
	// which is recommended due to the sensitivity of the value.
	// Setting it inside a terraform file will cause it to be persisted to disk in terraform.tfstate.
	// Experimental.
	SseCustomerKey *string `field:"optional" json:"sseCustomerKey" yaml:"sseCustomerKey"`
	// (Optional) Custom endpoint for the AWS Security Token Service (STS) API.
	//
	// This can also be sourced from the AWS_STS_ENDPOINT environment variable.
	// Experimental.
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
	// (Optional) Multi-Factor Authentication (MFA) token.
	//
	// This can also be sourced from the AWS_SESSION_TOKEN environment variable.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// (Optional) Prefix applied to the state path inside the bucket.
	//
	// This is only relevant when using a non-default workspace. Defaults to env:
	// Experimental.
	WorkspaceKeyPrefix *string `field:"optional" json:"workspaceKeyPrefix" yaml:"workspaceKeyPrefix"`
}

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

// Converts all fragments to strings and concats those.
//
// Drops 'undefined's.
// Experimental.
type StringConcat interface {
	IFragmentConcatenator
	// Join the fragment on the left and on the right.
	// Experimental.
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy struct for StringConcat
type jsiiProxy_StringConcat struct {
	jsiiProxy_IFragmentConcatenator
}

// Experimental.
func NewStringConcat() StringConcat {
	_init_.Initialize()

	j := jsiiProxy_StringConcat{}

	_jsii_.Create(
		"cdktf.StringConcat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStringConcat_Override(s StringConcat) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.StringConcat",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StringConcat) Join(left interface{}, right interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Experimental.
type StringMap interface {
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Lookup(key *string) *string
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StringMap
type jsiiProxy_StringMap struct {
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_StringMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMap) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewStringMap(terraformResource IInterpolatingParent, terraformAttribute *string) StringMap {
	_init_.Initialize()

	j := jsiiProxy_StringMap{}

	_jsii_.Create(
		"cdktf.StringMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewStringMap_Override(s StringMap, terraformResource IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.StringMap",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StringMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StringMap) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StringMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMap) Lookup(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMap) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type StringMapList interface {
	IInterpolatingParent
	IResolvable
	ITerraformAddressable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() IInterpolatingParent
	// Experimental.
	SetTerraformResource(val IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) StringMap
	// Experimental.
	InterpolationForAttribute(property *string) IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StringMapList
type jsiiProxy_StringMapList struct {
	jsiiProxy_IInterpolatingParent
	jsiiProxy_IResolvable
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_StringMapList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMapList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMapList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMapList) TerraformResource() IInterpolatingParent {
	var returns IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMapList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewStringMapList(terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StringMapList {
	_init_.Initialize()

	j := jsiiProxy_StringMapList{}

	_jsii_.Create(
		"cdktf.StringMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewStringMapList_Override(s StringMapList, terraformResource IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.StringMapList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StringMapList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StringMapList) SetTerraformResource(val IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StringMapList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StringMapList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMapList) Get(index *float64) StringMap {
	var returns StringMap

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMapList) InterpolationForAttribute(property *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMapList) Resolve(_context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringMapList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type SwiftBackend interface {
	TerraformBackend
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SwiftBackend
type jsiiProxy_SwiftBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_SwiftBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewSwiftBackend(scope constructs.Construct, props *SwiftBackendProps) SwiftBackend {
	_init_.Initialize()

	j := jsiiProxy_SwiftBackend{}

	_jsii_.Create(
		"cdktf.SwiftBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSwiftBackend_Override(s SwiftBackend, scope constructs.Construct, props *SwiftBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.SwiftBackend",
		[]interface{}{scope, props},
		s,
	)
}

// Experimental.
func SwiftBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.SwiftBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func SwiftBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.SwiftBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SwiftBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SwiftBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		s,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, _fromStack},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SwiftBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SwiftBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SwiftBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SwiftBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SwiftBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SwiftBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type SwiftBackendProps struct {
	// Experimental.
	Container *string `field:"required" json:"container" yaml:"container"`
	// Experimental.
	ApplicationCredentialId *string `field:"optional" json:"applicationCredentialId" yaml:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `field:"optional" json:"applicationCredentialName" yaml:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `field:"optional" json:"applicationCredentialSecret" yaml:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `field:"optional" json:"archiveContainer" yaml:"archiveContainer"`
	// Experimental.
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// Experimental.
	CacertFile *string `field:"optional" json:"cacertFile" yaml:"cacertFile"`
	// Experimental.
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// Experimental.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// Experimental.
	DefaultDomain *string `field:"optional" json:"defaultDomain" yaml:"defaultDomain"`
	// Experimental.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Experimental.
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
	// Experimental.
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	ProjectDomainId *string `field:"optional" json:"projectDomainId" yaml:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `field:"optional" json:"projectDomainName" yaml:"projectDomainName"`
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Experimental.
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Experimental.
	TenantName *string `field:"optional" json:"tenantName" yaml:"tenantName"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Experimental.
	UserDomainId *string `field:"optional" json:"userDomainId" yaml:"userDomainId"`
	// Experimental.
	UserDomainName *string `field:"optional" json:"userDomainName" yaml:"userDomainName"`
	// Experimental.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Experimental.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

// Experimental.
type TerraformAsset interface {
	Resource
	// Experimental.
	AssetHash() *string
	// Experimental.
	SetAssetHash(val *string)
	// Name of the asset.
	// Experimental.
	FileName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The path relative to the root of the terraform directory in posix format Use this property to reference the asset.
	// Experimental.
	Path() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() TerraformStack
	// Experimental.
	Type() AssetType
	// Experimental.
	SetType(val AssetType)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TerraformAsset
type jsiiProxy_TerraformAsset struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_TerraformAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Stack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Type() AssetType {
	var returns AssetType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// A Terraform Asset takes a file or directory outside of the CDK for Terraform context and moves it into it.
//
// Assets copy referenced files into the stacks context for further usage in other resources.
// Experimental.
func NewTerraformAsset(scope constructs.Construct, id *string, config *TerraformAssetConfig) TerraformAsset {
	_init_.Initialize()

	j := jsiiProxy_TerraformAsset{}

	_jsii_.Create(
		"cdktf.TerraformAsset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// A Terraform Asset takes a file or directory outside of the CDK for Terraform context and moves it into it.
//
// Assets copy referenced files into the stacks context for further usage in other resources.
// Experimental.
func NewTerraformAsset_Override(t TerraformAsset, scope constructs.Construct, id *string, config *TerraformAssetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformAsset",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformAsset) SetAssetHash(val *string) {
	_jsii_.Set(
		j,
		"assetHash",
		val,
	)
}

func (j *jsiiProxy_TerraformAsset) SetType(val AssetType) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformAssetConfig struct {
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Experimental.
	Type AssetType `field:"optional" json:"type" yaml:"type"`
}

// Experimental.
type TerraformBackend interface {
	TerraformElement
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Experimental.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformBackend
type jsiiProxy_TerraformBackend struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformBackend_Override(t TerraformBackend, scope constructs.Construct, id *string, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformBackend",
		[]interface{}{scope, id, name},
		t,
	)
}

// Experimental.
func TerraformBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformBackend",
		"isBackend",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, fromStack *string) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		t,
		"getRemoteStateDataSource",
		[]interface{}{scope, name, fromStack},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformDataSource interface {
	TerraformElement
	IInterpolatingParent
	ITerraformDependable
	ITerraformResource
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() ITerraformIterator
	// Experimental.
	SetForEach(val ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Provider() TerraformProvider
	// Experimental.
	SetProvider(val TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformDataSource
type jsiiProxy_TerraformDataSource struct {
	jsiiProxy_TerraformElement
	jsiiProxy_IInterpolatingParent
	jsiiProxy_ITerraformDependable
	jsiiProxy_ITerraformResource
}

func (j *jsiiProxy_TerraformDataSource) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) ForEach() ITerraformIterator {
	var returns ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Lifecycle() *TerraformResourceLifecycle {
	var returns *TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata {
	var returns *TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformDataSource(scope constructs.Construct, id *string, config *TerraformResourceConfig) TerraformDataSource {
	_init_.Initialize()

	j := jsiiProxy_TerraformDataSource{}

	_jsii_.Create(
		"cdktf.TerraformDataSource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformDataSource_Override(t TerraformDataSource, scope constructs.Construct, id *string, config *TerraformResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformDataSource",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetForEach(val ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetLifecycle(val *TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformDataSource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetBooleanAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) InterpolationForAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformDataSource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformDataSource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformDataSource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformElement interface {
	constructs.Construct
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformElement
type jsiiProxy_TerraformElement struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TerraformElement) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformElement(scope constructs.Construct, id *string, elementType *string) TerraformElement {
	_init_.Initialize()

	j := jsiiProxy_TerraformElement{}

	_jsii_.Create(
		"cdktf.TerraformElement",
		[]interface{}{scope, id, elementType},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformElement_Override(t TerraformElement, scope constructs.Construct, id *string, elementType *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformElement",
		[]interface{}{scope, id, elementType},
		t,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformElement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformElement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformElement) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformElement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformElement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformElement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformElement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformElement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformElementMetadata struct {
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Experimental.
	StackTrace *[]*string `field:"required" json:"stackTrace" yaml:"stackTrace"`
	// Experimental.
	UniqueId *string `field:"required" json:"uniqueId" yaml:"uniqueId"`
}

// Experimental.
type TerraformHclModule interface {
	TerraformModule
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() ITerraformIterator
	// Experimental.
	SetForEach(val ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	// Experimental.
	Variables() *map[string]interface{}
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	Get(output *string) interface{}
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	Set(variable *string, value interface{})
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformHclModule
type jsiiProxy_TerraformHclModule struct {
	jsiiProxy_TerraformModule
}

func (j *jsiiProxy_TerraformHclModule) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) ForEach() ITerraformIterator {
	var returns ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Variables() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformHclModule(scope constructs.Construct, id *string, options *TerraformHclModuleOptions) TerraformHclModule {
	_init_.Initialize()

	j := jsiiProxy_TerraformHclModule{}

	_jsii_.Create(
		"cdktf.TerraformHclModule",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformHclModule_Override(t TerraformHclModule, scope constructs.Construct, id *string, options *TerraformHclModuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformHclModule",
		[]interface{}{scope, id, options},
		t,
	)
}

func (j *jsiiProxy_TerraformHclModule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformHclModule) SetForEach(val ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformHclModule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformHclModule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformHclModule) AddProvider(provider interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addProvider",
		[]interface{}{provider},
	)
}

func (t *jsiiProxy_TerraformHclModule) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) InterpolationForOutput(moduleOutput *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformHclModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformHclModule) Set(variable *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"set",
		[]interface{}{variable, value},
	)
}

func (t *jsiiProxy_TerraformHclModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformHclModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformHclModuleOptions struct {
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Experimental.
	Variables *map[string]interface{} `field:"optional" json:"variables" yaml:"variables"`
}

// Experimental.
type TerraformIterator interface {
	ITerraformIterator
	// Experimental.
	Dynamic(attributes *map[string]interface{}) IResolvable
	// Returns: the given attribute of the current item iterated over as any.
	// Experimental.
	GetAny(attribute *string) IResolvable
	// Returns: the given attribute of the current item iterated over as a map of any.
	// Experimental.
	GetAnyMap(attribute *string) *map[string]interface{}
	// Returns: the given attribute of the current item iterated over as a boolean.
	// Experimental.
	GetBoolean(attribute *string) IResolvable
	// Returns: the given attribute of the current item iterated over as a map of booleans.
	// Experimental.
	GetBooleanMap(attribute *string) *map[string]*bool
	// Returns: the given attribute of the current item iterated over as a (string) list.
	// Experimental.
	GetList(attribute *string) *[]*string
	// Returns: the given attribute of the current item iterated over as a map.
	// Experimental.
	GetMap(attribute *string) *map[string]interface{}
	// Returns: the given attribute of the current item iterated over as a number.
	// Experimental.
	GetNumber(attribute *string) *float64
	// Returns: the given attribute of the current item iterated over as a number list.
	// Experimental.
	GetNumberList(attribute *string) *[]*float64
	// Returns: the given attribute of the current item iterated over as a map of numbers.
	// Experimental.
	GetNumberMap(attribute *string) *map[string]*float64
	// Returns: the given attribute of the current item iterated over as a string.
	// Experimental.
	GetString(attribute *string) *string
	// Returns: the given attribute of the current item iterated over as a map of strings.
	// Experimental.
	GetStringMap(attribute *string) *map[string]*string
}

// The jsii proxy struct for TerraformIterator
type jsiiProxy_TerraformIterator struct {
	jsiiProxy_ITerraformIterator
}

// Experimental.
func NewTerraformIterator_Override(t TerraformIterator) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformIterator",
		nil, // no parameters
		t,
	)
}

// Creates a new iterator from a list.
// Experimental.
func TerraformIterator_FromList(list interface{}) ListTerraformIterator {
	_init_.Initialize()

	var returns ListTerraformIterator

	_jsii_.StaticInvoke(
		"cdktf.TerraformIterator",
		"fromList",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// Creates a new iterator from a map.
// Experimental.
func TerraformIterator_FromMap(map_ interface{}) MapTerraformIterator {
	_init_.Initialize()

	var returns MapTerraformIterator

	_jsii_.StaticInvoke(
		"cdktf.TerraformIterator",
		"fromMap",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) Dynamic(attributes *map[string]interface{}) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"dynamic",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetAny(attribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"getAny",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetAnyMap(attribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetBoolean(attribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetBooleanMap(attribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetList(attribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getList",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetMap(attribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetNumber(attribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumber",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetNumberList(attribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberList",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetNumberMap(attribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetString(attribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getString",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformIterator) GetStringMap(attribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

// Experimental.
type TerraformLocal interface {
	TerraformElement
	ITerraformAddressable
	// Experimental.
	AsBoolean() IResolvable
	// Experimental.
	AsList() *[]*string
	// Experimental.
	AsNumber() *float64
	// Experimental.
	AsString() *string
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Expression() interface{}
	// Experimental.
	SetExpression(val interface{})
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformLocal
type jsiiProxy_TerraformLocal struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_TerraformLocal) AsBoolean() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"asBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) AsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) AsNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) AsString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) Expression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformLocal(scope constructs.Construct, id *string, expression interface{}) TerraformLocal {
	_init_.Initialize()

	j := jsiiProxy_TerraformLocal{}

	_jsii_.Create(
		"cdktf.TerraformLocal",
		[]interface{}{scope, id, expression},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformLocal_Override(t TerraformLocal, scope constructs.Construct, id *string, expression interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformLocal",
		[]interface{}{scope, id, expression},
		t,
	)
}

func (j *jsiiProxy_TerraformLocal) SetExpression(val interface{}) {
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformLocal_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformLocal",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformLocal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformLocal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformLocal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformLocal) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformLocal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformLocal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformMetaArguments struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
}

// Experimental.
type TerraformModule interface {
	TerraformElement
	ITerraformDependable
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() ITerraformIterator
	// Experimental.
	SetForEach(val ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformModule
type jsiiProxy_TerraformModule struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformDependable
}

func (j *jsiiProxy_TerraformModule) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) ForEach() ITerraformIterator {
	var returns ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformModule_Override(t TerraformModule, scope constructs.Construct, id *string, options *TerraformModuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformModule",
		[]interface{}{scope, id, options},
		t,
	)
}

func (j *jsiiProxy_TerraformModule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformModule) SetForEach(val ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformModule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformModule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformModule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformModule) AddProvider(provider interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addProvider",
		[]interface{}{provider},
	)
}

func (t *jsiiProxy_TerraformModule) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformModule) InterpolationForOutput(moduleOutput *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformModule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformModule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformModuleOptions struct {
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Experimental.
type TerraformModuleProvider struct {
	// Experimental.
	ModuleAlias *string `field:"required" json:"moduleAlias" yaml:"moduleAlias"`
	// Experimental.
	Provider TerraformProvider `field:"required" json:"provider" yaml:"provider"`
}

// Experimental.
type TerraformModuleUserOptions struct {
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
}

// Experimental.
type TerraformOutput interface {
	TerraformElement
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]ITerraformDependable
	// Experimental.
	SetDependsOn(val *[]ITerraformDependable)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Sensitive() *bool
	// Experimental.
	SetSensitive(val *bool)
	// Experimental.
	StaticId() *bool
	// Experimental.
	SetStaticId(val *bool)
	// Experimental.
	Value() interface{}
	// Experimental.
	SetValue(val interface{})
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformOutput
type jsiiProxy_TerraformOutput struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformOutput) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) DependsOn() *[]ITerraformDependable {
	var returns *[]ITerraformDependable
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Sensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) StaticId() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"staticId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformOutput(scope constructs.Construct, id *string, config *TerraformOutputConfig) TerraformOutput {
	_init_.Initialize()

	j := jsiiProxy_TerraformOutput{}

	_jsii_.Create(
		"cdktf.TerraformOutput",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformOutput_Override(t TerraformOutput, scope constructs.Construct, id *string, config *TerraformOutputConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformOutput",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformOutput) SetDependsOn(val *[]ITerraformDependable) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetSensitive(val *bool) {
	_jsii_.Set(
		j,
		"sensitive",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetStaticId(val *bool) {
	_jsii_.Set(
		j,
		"staticId",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetValue(val interface{}) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TerraformOutput_IsTerrafromOutput(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformOutput",
		"isTerrafromOutput",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformOutput) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformOutput) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformOutput) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformOutput) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

// Experimental.
type TerraformProvider interface {
	TerraformElement
	// Experimental.
	Alias() *string
	// Experimental.
	SetAlias(val *string)
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformProvider
type jsiiProxy_TerraformProvider struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata {
	var returns *TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformProvider_Override(t TerraformProvider, scope constructs.Construct, id *string, config *TerraformProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformProvider",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformProviderConfig struct {
	// Experimental.
	TerraformResourceType *string `field:"required" json:"terraformResourceType" yaml:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `field:"optional" json:"terraformGeneratorMetadata" yaml:"terraformGeneratorMetadata"`
	// Experimental.
	TerraformProviderSource *string `field:"optional" json:"terraformProviderSource" yaml:"terraformProviderSource"`
}

// Experimental.
type TerraformProviderGeneratorMetadata struct {
	// Experimental.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Experimental.
	ProviderVersion *string `field:"optional" json:"providerVersion" yaml:"providerVersion"`
	// Experimental.
	ProviderVersionConstraint *string `field:"optional" json:"providerVersionConstraint" yaml:"providerVersionConstraint"`
}

// Experimental.
type TerraformRemoteState interface {
	TerraformElement
	ITerraformAddressable
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	Get(output *string) IResolvable
	// Experimental.
	GetBoolean(output *string) IResolvable
	// Experimental.
	GetList(output *string) *[]*string
	// Experimental.
	GetNumber(output *string) *float64
	// Experimental.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformRemoteState
type jsiiProxy_TerraformRemoteState struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_TerraformRemoteState) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformRemoteState_Override(t TerraformRemoteState, scope constructs.Construct, id *string, backend *string, config *DataTerraformRemoteStateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformRemoteState",
		[]interface{}{scope, id, backend, config},
		t,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformRemoteState_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformRemoteState",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TerraformRemoteState_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.TerraformRemoteState",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TerraformRemoteState) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformRemoteState) Get(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) GetBoolean(output *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformRemoteState) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformRemoteState) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformRemoteState) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformResource interface {
	TerraformElement
	IInterpolatingParent
	ITerraformDependable
	ITerraformResource
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() ITerraformIterator
	// Experimental.
	SetForEach(val ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Provider() TerraformProvider
	// Experimental.
	SetProvider(val TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformResource
type jsiiProxy_TerraformResource struct {
	jsiiProxy_TerraformElement
	jsiiProxy_IInterpolatingParent
	jsiiProxy_ITerraformDependable
	jsiiProxy_ITerraformResource
}

func (j *jsiiProxy_TerraformResource) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) ForEach() ITerraformIterator {
	var returns ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Lifecycle() *TerraformResourceLifecycle {
	var returns *TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata {
	var returns *TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformResource(scope constructs.Construct, id *string, config *TerraformResourceConfig) TerraformResource {
	_init_.Initialize()

	j := jsiiProxy_TerraformResource{}

	_jsii_.Create(
		"cdktf.TerraformResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformResource_Override(t TerraformResource, scope constructs.Construct, id *string, config *TerraformResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformResource",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformResource) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetForEach(val ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetLifecycle(val *TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetBooleanAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) InterpolationForAttribute(terraformAttribute *string) IResolvable {
	var returns IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformResourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Experimental.
	TerraformResourceType *string `field:"required" json:"terraformResourceType" yaml:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `field:"optional" json:"terraformGeneratorMetadata" yaml:"terraformGeneratorMetadata"`
}

// Experimental.
type TerraformResourceLifecycle struct {
	// Experimental.
	CreateBeforeDestroy *bool `field:"optional" json:"createBeforeDestroy" yaml:"createBeforeDestroy"`
	// Experimental.
	IgnoreChanges interface{} `field:"optional" json:"ignoreChanges" yaml:"ignoreChanges"`
	// Experimental.
	PreventDestroy *bool `field:"optional" json:"preventDestroy" yaml:"preventDestroy"`
}

// Expressions in connection blocks cannot refer to their parent resource by name.
//
// References create dependencies, and referring to a resource by name within its own block would create a dependency cycle.
// Instead, expressions can use the self object, which represents the connection's parent resource and has all of that resource's attributes.
// For example, use self.public_ip to reference an aws_instance's public_ip attribute.
// Experimental.
type TerraformSelf interface {
}

// The jsii proxy struct for TerraformSelf
type jsiiProxy_TerraformSelf struct {
	_ byte // padding
}

// Experimental.
func NewTerraformSelf() TerraformSelf {
	_init_.Initialize()

	j := jsiiProxy_TerraformSelf{}

	_jsii_.Create(
		"cdktf.TerraformSelf",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformSelf_Override(t TerraformSelf) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformSelf",
		nil, // no parameters
		t,
	)
}

// Only usable within a connection block to reference the connections parent resource.
//
// Access a property on the resource like this: `getAny("hostPort")`.
// Experimental.
func TerraformSelf_GetAny(key *string) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.TerraformSelf",
		"getAny",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Only usable within a connection block to reference the connections parent resource.
//
// Access a property on the resource like this: `getNumber("hostPort")`.
// Experimental.
func TerraformSelf_GetNumber(key *string) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.TerraformSelf",
		"getNumber",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Only usable within a connection block to reference the connections parent resource.
//
// Access a property on the resource like this: `getString("publicIp")`.
// Experimental.
func TerraformSelf_GetString(key *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.TerraformSelf",
		"getString",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
type TerraformStack interface {
	constructs.Construct
	// Experimental.
	Dependencies() *[]TerraformStack
	// Experimental.
	SetDependencies(val *[]TerraformStack)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Synthesizer() IStackSynthesizer
	// Experimental.
	SetSynthesizer(val IStackSynthesizer)
	// Experimental.
	AddDependency(dependency TerraformStack)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Returns the naming scheme used to allocate logical IDs.
	//
	// By default, uses
	// the `HashedAddressingScheme` but this method can be overridden to customize
	// this behavior.
	// Experimental.
	AllocateLogicalId(tfElement interface{}) *string
	// Experimental.
	AllProviders() *[]TerraformProvider
	// Experimental.
	DependsOn(stack TerraformStack) *bool
	// Experimental.
	EnsureBackendExists() TerraformBackend
	// Experimental.
	GetLogicalId(tfElement interface{}) *string
	// Experimental.
	PrepareStack()
	// Experimental.
	RegisterIncomingCrossStackReference(fromStack TerraformStack) TerraformRemoteState
	// Experimental.
	RegisterOutgoingCrossStackReference(identifier *string) TerraformOutput
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformStack
type jsiiProxy_TerraformStack struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TerraformStack) Dependencies() *[]TerraformStack {
	var returns *[]TerraformStack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformStack) Synthesizer() IStackSynthesizer {
	var returns IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformStack(scope constructs.Construct, id *string) TerraformStack {
	_init_.Initialize()

	j := jsiiProxy_TerraformStack{}

	_jsii_.Create(
		"cdktf.TerraformStack",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformStack_Override(t TerraformStack, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformStack",
		[]interface{}{scope, id},
		t,
	)
}

func (j *jsiiProxy_TerraformStack) SetDependencies(val *[]TerraformStack) {
	_jsii_.Set(
		j,
		"dependencies",
		val,
	)
}

func (j *jsiiProxy_TerraformStack) SetSynthesizer(val IStackSynthesizer) {
	_jsii_.Set(
		j,
		"synthesizer",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TerraformStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TerraformStack_Of(construct constructs.IConstruct) TerraformStack {
	_init_.Initialize()

	var returns TerraformStack

	_jsii_.StaticInvoke(
		"cdktf.TerraformStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) AddDependency(dependency TerraformStack) {
	_jsii_.InvokeVoid(
		t,
		"addDependency",
		[]interface{}{dependency},
	)
}

func (t *jsiiProxy_TerraformStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformStack) AllocateLogicalId(tfElement interface{}) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"allocateLogicalId",
		[]interface{}{tfElement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) AllProviders() *[]TerraformProvider {
	var returns *[]TerraformProvider

	_jsii_.Invoke(
		t,
		"allProviders",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) DependsOn(stack TerraformStack) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"dependsOn",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) EnsureBackendExists() TerraformBackend {
	var returns TerraformBackend

	_jsii_.Invoke(
		t,
		"ensureBackendExists",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) GetLogicalId(tfElement interface{}) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getLogicalId",
		[]interface{}{tfElement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) PrepareStack() {
	_jsii_.InvokeVoid(
		t,
		"prepareStack",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformStack) RegisterIncomingCrossStackReference(fromStack TerraformStack) TerraformRemoteState {
	var returns TerraformRemoteState

	_jsii_.Invoke(
		t,
		"registerIncomingCrossStackReference",
		[]interface{}{fromStack},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) RegisterOutgoingCrossStackReference(identifier *string) TerraformOutput {
	var returns TerraformOutput

	_jsii_.Invoke(
		t,
		"registerOutgoingCrossStackReference",
		[]interface{}{identifier},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformStackMetadata struct {
	// Experimental.
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Experimental.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
}

// Experimental.
type TerraformVariable interface {
	TerraformElement
	ITerraformAddressable
	// Experimental.
	BooleanValue() IResolvable
	// Experimental.
	CdktfStack() TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Default() interface{}
	// Experimental.
	Description() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	ListValue() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Nullable() *bool
	// Experimental.
	NumberValue() *float64
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Sensitive() *bool
	// Experimental.
	StringValue() *string
	// Experimental.
	Type() *string
	// Experimental.
	Validation() *[]*TerraformVariableValidationConfig
	// Experimental.
	Value() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddValidation(validation *TerraformVariableValidationConfig)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformVariable
type jsiiProxy_TerraformVariable struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_TerraformVariable) BooleanValue() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"booleanValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) ListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Nullable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Sensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Validation() *[]*TerraformVariableValidationConfig {
	var returns *[]*TerraformVariableValidationConfig
	_jsii_.Get(
		j,
		"validation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformVariable(scope constructs.Construct, id *string, config *TerraformVariableConfig) TerraformVariable {
	_init_.Initialize()

	j := jsiiProxy_TerraformVariable{}

	_jsii_.Create(
		"cdktf.TerraformVariable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformVariable_Override(t TerraformVariable, scope constructs.Construct, id *string, config *TerraformVariableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformVariable",
		[]interface{}{scope, id, config},
		t,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TerraformVariable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformVariable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformVariable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TerraformVariable) AddValidation(validation *TerraformVariableValidationConfig) {
	_jsii_.InvokeVoid(
		t,
		"addValidation",
		[]interface{}{validation},
	)
}

func (t *jsiiProxy_TerraformVariable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TerraformVariable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TerraformVariable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformVariable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformVariable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TerraformVariable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformVariableConfig struct {
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	Nullable *bool `field:"optional" json:"nullable" yaml:"nullable"`
	// Experimental.
	Sensitive *bool `field:"optional" json:"sensitive" yaml:"sensitive"`
	// The type argument in a variable block allows you to restrict the type of value that will be accepted as the value for a variable.
	//
	// If no type constraint is set then a value of any type is accepted.
	//
	// While type constraints are optional, we recommend specifying them; they serve as easy reminders for users of the module, and allow Terraform to return a helpful error message if the wrong type is used.
	//
	// Type constraints are created from a mixture of type keywords and type constructors. The supported type keywords are:
	//
	// - string
	// - number
	// - bool
	//
	// The type constructors allow you to specify complex types such as collections:
	//
	// - list(\<TYPE\>)
	// - set(\<TYPE\>)
	// - map(\<TYPE\>)
	// - object({\<ATTR NAME\> = \<TYPE\>, ... })
	// - tuple([\<TYPE\>, ...])
	//
	// The keyword any may be used to indicate that any type is acceptable. For more information on the meaning and behavior of these different types, as well as detailed information about automatic conversion of complex types, see {@link https://www.terraform.io/docs/configuration/types.html|Type Constraints}.
	//
	// If both the type and default arguments are specified, the given default value must be convertible to the specified type.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specify arbitrary custom validation rules for a particular variable using a validation block nested within the corresponding variable block.
	// Experimental.
	Validation *[]*TerraformVariableValidationConfig `field:"optional" json:"validation" yaml:"validation"`
}

// Experimental.
type TerraformVariableValidationConfig struct {
	// Experimental.
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// Experimental.
	ErrorMessage *string `field:"required" json:"errorMessage" yaml:"errorMessage"`
}

// Testing utilities for cdktf applications.
// Experimental.
type Testing interface {
}

// The jsii proxy struct for Testing
type jsiiProxy_Testing struct {
	_ byte // padding
}

// Experimental.
func NewTesting() Testing {
	_init_.Initialize()

	j := jsiiProxy_Testing{}

	_jsii_.Create(
		"cdktf.Testing",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTesting_Override(t Testing) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Testing",
		nil, // no parameters
		t,
	)
}

// Returns an app for testing with the following properties: - Output directory is a temp dir.
// Experimental.
func Testing_App(options *TestingAppOptions) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"app",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_EnableFutureFlags(app App) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"enableFutureFlags",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_FakeCdktfJsonPath(app App) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"fakeCdktfJsonPath",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_FullSynth(stack TerraformStack) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"fullSynth",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_RenderConstructTree(construct constructs.IConstruct) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"renderConstructTree",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_SetupJest() {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"cdktf.Testing",
		"setupJest",
		nil, // no parameters
	)
}

// Experimental.
func Testing_StubVersion(app App) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"stubVersion",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// Returns the Terraform synthesized JSON.
// Experimental.
func Testing_Synth(stack TerraformStack) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"synth",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_SynthScope(fn IScopeCallback) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"synthScope",
		[]interface{}{fn},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_ToBeValidTerraform(received *string) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"toBeValidTerraform",
		[]interface{}{received},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_ToHaveDataSource(received *string, resourceType *string) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"toHaveDataSource",
		[]interface{}{received, resourceType},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_ToHaveDataSourceWithProperties(received *string, resourceType *string, properties *map[string]interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"toHaveDataSourceWithProperties",
		[]interface{}{received, resourceType, properties},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_ToHaveResource(received *string, resourceType *string) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"toHaveResource",
		[]interface{}{received, resourceType},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_ToHaveResourceWithProperties(received *string, resourceType *string, properties *map[string]interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"toHaveResourceWithProperties",
		[]interface{}{received, resourceType, properties},
		&returns,
	)

	return returns
}

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

// Represents a special or lazily-evaluated value.
//
// Can be used to delay evaluation of a certain value in case, for example,
// that it requires some context or late-bound data. Can also be used to
// mark values that need special processing at document rendering time.
//
// Tokens can be embedded into strings while retaining their original
// semantics.
// Experimental.
type Token interface {
}

// The jsii proxy struct for Token
type jsiiProxy_Token struct {
	_ byte // padding
}

// Experimental.
func NewToken() Token {
	_init_.Initialize()

	j := jsiiProxy_Token{}

	_jsii_.Create(
		"cdktf.Token",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewToken_Override(t Token) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Token",
		nil, // no parameters
		t,
	)
}

// Return a resolvable representation of the given value.
// Experimental.
func Token_AsAny(value interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asAny",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible map representation of this token.
// Experimental.
func Token_AsAnyMap(value interface{}, options *EncodingOptions) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asAnyMap",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible map representation of this token.
// Experimental.
func Token_AsBooleanMap(value interface{}, options *EncodingOptions) *map[string]*bool {
	_init_.Initialize()

	var returns *map[string]*bool

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asBooleanMap",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible list representation of this token.
// Experimental.
func Token_AsList(value interface{}, options *EncodingOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asList",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible map representation of this token.
// Experimental.
func Token_AsMap(value interface{}, mapValue interface{}, options *EncodingOptions) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asMap",
		[]interface{}{value, mapValue, options},
		&returns,
	)

	return returns
}

// Return a reversible number representation of this token.
// Experimental.
func Token_AsNumber(value interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible list representation of this token.
// Experimental.
func Token_AsNumberList(value interface{}) *[]*float64 {
	_init_.Initialize()

	var returns *[]*float64

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asNumberList",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible map representation of this token.
// Experimental.
func Token_AsNumberMap(value interface{}, options *EncodingOptions) *map[string]*float64 {
	_init_.Initialize()

	var returns *map[string]*float64

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asNumberMap",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible string representation of this token.
//
// If the Token is initialized with a literal, the stringified value of the
// literal is returned. Otherwise, a special quoted string representation
// of the Token is returned that can be embedded into other strings.
//
// Strings with quoted Tokens in them can be restored back into
// complex values with the Tokens restored by calling `resolve()`
// on the string.
// Experimental.
func Token_AsString(value interface{}, options *EncodingOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asString",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible map representation of this token.
// Experimental.
func Token_AsStringMap(value interface{}, options *EncodingOptions) *map[string]*string {
	_init_.Initialize()

	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asStringMap",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Returns true if obj represents an unresolved value.
//
// One of these must be true:
//
// - `obj` is an IResolvable
// - `obj` is a string containing at least one encoded `IResolvable`
// - `obj` is either an encoded number or list
//
// This does NOT recurse into lists or objects to see if they
// containing resolvables.
// Experimental.
func Token_IsUnresolved(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"isUnresolved",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func Token_ANY_MAP_TOKEN_VALUE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Token",
		"ANY_MAP_TOKEN_VALUE",
		&returns,
	)
	return returns
}

func Token_NUMBER_MAP_TOKEN_VALUE() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"cdktf.Token",
		"NUMBER_MAP_TOKEN_VALUE",
		&returns,
	)
	return returns
}

func Token_STRING_MAP_TOKEN_VALUE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Token",
		"STRING_MAP_TOKEN_VALUE",
		&returns,
	)
	return returns
}

// Less oft-needed functions to manipulate Tokens.
// Experimental.
type Tokenization interface {
}

// The jsii proxy struct for Tokenization
type jsiiProxy_Tokenization struct {
	_ byte // padding
}

// Experimental.
func NewTokenization() Tokenization {
	_init_.Initialize()

	j := jsiiProxy_Tokenization{}

	_jsii_.Create(
		"cdktf.Tokenization",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTokenization_Override(t Tokenization) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Tokenization",
		nil, // no parameters
		t,
	)
}

// Return whether the given object is an IResolvable object.
//
// This is different from Token.isUnresolved() which will also check for
// encoded Tokens, whereas this method will only do a type check on the given
// object.
// Experimental.
func Tokenization_IsResolvable(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"isResolvable",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Resolves an object by evaluating all tokens and removing any undefined or empty objects or arrays.
//
// Values can only be primitives, arrays or tokens. Other objects (i.e. with methods) will be rejected.
// Experimental.
func Tokenization_Resolve(obj interface{}, options *ResolveOptions) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"resolve",
		[]interface{}{obj, options},
		&returns,
	)

	return returns
}

// Reverse any value into Resolvables, if possible.
// Experimental.
func Tokenization_Reverse(x interface{}) *[]IResolvable {
	_init_.Initialize()

	var returns *[]IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverse",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a list.
// Experimental.
func Tokenization_ReverseList(l *[]*string) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseList",
		[]interface{}{l},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a map.
// Experimental.
func Tokenization_ReverseMap(m *map[string]interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseMap",
		[]interface{}{m},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a number.
// Experimental.
func Tokenization_ReverseNumber(n *float64) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseNumber",
		[]interface{}{n},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a list.
// Experimental.
func Tokenization_ReverseNumberList(l *[]*float64) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseNumberList",
		[]interface{}{l},
		&returns,
	)

	return returns
}

// Un-encode a string potentially containing encoded tokens.
// Experimental.
func Tokenization_ReverseString(s *string) TokenizedStringFragments {
	_init_.Initialize()

	var returns TokenizedStringFragments

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Stringify a number directly or lazily if it's a Token.
//
// If it is an object (i.e., { Ref: 'SomeLogicalId' }), return it as-is.
// Experimental.
func Tokenization_StringifyNumber(x *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"stringifyNumber",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Fragments of a concatenated string containing stringified Tokens.
// Experimental.
type TokenizedStringFragments interface {
	// Returns the first token.
	// Experimental.
	FirstToken() IResolvable
	// Returns the first value.
	// Experimental.
	FirstValue() interface{}
	// Return all intrinsic fragments from this string.
	// Experimental.
	Intrinsic() *[]IResolvable
	// Returns the number of fragments.
	// Experimental.
	Length() *float64
	// Return all literals from this string.
	// Experimental.
	Literals() *[]IResolvable
	// Return all Tokens from this string.
	// Experimental.
	Tokens() *[]IResolvable
	// Adds an intrinsic fragment.
	// Experimental.
	AddIntrinsic(value interface{})
	// Adds a literal fragment.
	// Experimental.
	AddLiteral(lit interface{})
	// Adds a token fragment.
	// Experimental.
	AddToken(token IResolvable)
	// Combine the string fragments using the given joiner.
	//
	// If there are any.
	// Experimental.
	Join(concat IFragmentConcatenator) interface{}
	// Apply a transformation function to all tokens in the string.
	// Experimental.
	MapTokens(mapper ITokenMapper) TokenizedStringFragments
}

// The jsii proxy struct for TokenizedStringFragments
type jsiiProxy_TokenizedStringFragments struct {
	_ byte // padding
}

func (j *jsiiProxy_TokenizedStringFragments) FirstToken() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"firstToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) FirstValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Intrinsic() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"intrinsic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Literals() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"literals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Tokens() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


// Experimental.
func NewTokenizedStringFragments() TokenizedStringFragments {
	_init_.Initialize()

	j := jsiiProxy_TokenizedStringFragments{}

	_jsii_.Create(
		"cdktf.TokenizedStringFragments",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTokenizedStringFragments_Override(t TokenizedStringFragments) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TokenizedStringFragments",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddIntrinsic(value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addIntrinsic",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddLiteral(lit interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addLiteral",
		[]interface{}{lit},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddToken(token IResolvable) {
	_jsii_.InvokeVoid(
		t,
		"addToken",
		[]interface{}{token},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) Join(concat IFragmentConcatenator) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"join",
		[]interface{}{concat},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenizedStringFragments) MapTokens(mapper ITokenMapper) TokenizedStringFragments {
	var returns TokenizedStringFragments

	_jsii_.Invoke(
		t,
		"mapTokens",
		[]interface{}{mapper},
		&returns,
	)

	return returns
}

// Experimental.
type VariableType interface {
}

// The jsii proxy struct for VariableType
type jsiiProxy_VariableType struct {
	_ byte // padding
}

// Experimental.
func NewVariableType_Override(v VariableType) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.VariableType",
		nil, // no parameters
		v,
	)
}

// Experimental.
func VariableType_List(type_ *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"list",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Map(type_ *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"map",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Object(attributes *map[string]*string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"object",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Set(type_ *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"set",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Tuple(elements ...*string) *string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range elements {
		args = append(args, a)
	}

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"tuple",
		args,
		&returns,
	)

	return returns
}

func VariableType_ANY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"ANY",
		&returns,
	)
	return returns
}

func VariableType_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"BOOL",
		&returns,
	)
	return returns
}

func VariableType_LIST() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST",
		&returns,
	)
	return returns
}

func VariableType_LIST_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST_BOOL",
		&returns,
	)
	return returns
}

func VariableType_LIST_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST_NUMBER",
		&returns,
	)
	return returns
}

func VariableType_LIST_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST_STRING",
		&returns,
	)
	return returns
}

func VariableType_MAP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP",
		&returns,
	)
	return returns
}

func VariableType_MAP_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP_BOOL",
		&returns,
	)
	return returns
}

func VariableType_MAP_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP_NUMBER",
		&returns,
	)
	return returns
}

func VariableType_MAP_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP_STRING",
		&returns,
	)
	return returns
}

func VariableType_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"NUMBER",
		&returns,
	)
	return returns
}

func VariableType_SET() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET",
		&returns,
	)
	return returns
}

func VariableType_SET_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET_BOOL",
		&returns,
	)
	return returns
}

func VariableType_SET_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET_NUMBER",
		&returns,
	)
	return returns
}

func VariableType_SET_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET_STRING",
		&returns,
	)
	return returns
}

func VariableType_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"STRING",
		&returns,
	)
	return returns
}

