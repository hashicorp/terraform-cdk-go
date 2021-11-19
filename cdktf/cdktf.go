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
	AnnotationMetadataEntryType_INFO AnnotationMetadataEntryType = "INFO"
	AnnotationMetadataEntryType_WARN AnnotationMetadataEntryType = "WARN"
	AnnotationMetadataEntryType_ERROR AnnotationMetadataEntryType = "ERROR"
)

// Includes API for attaching annotations such as warning messages to constructs.
// Experimental.
type Annotations interface {
	AddError(message *string)
	AddInfo(message *string)
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

// Adds an { "error": <message> } metadata entry to this construct.
//
// The toolkit will fail synthesis when errors are reported.
// Experimental.
func (a *jsiiProxy_Annotations) AddError(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addError",
		[]interface{}{message},
	)
}

// Adds an info metadata entry to this construct.
//
// The CLI will display the info message when apps are synthesized.
// Experimental.
func (a *jsiiProxy_Annotations) AddInfo(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addInfo",
		[]interface{}{message},
	)
}

// Adds a warning metadata entry to this construct.
//
// The CLI will display the warning when an app is synthesized.
// In a future release the CLI might introduce a --strict flag which
// will then fail the synthesis if it encounters a warning.
// Experimental.
func (a *jsiiProxy_Annotations) AddWarning(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addWarning",
		[]interface{}{message},
	)
}

// Represents a cdktf application.
// Experimental.
type App interface {
	constructs.Construct
	Manifest() Manifest
	Node() constructs.Node
	Outdir() *string
	SkipValidation() *bool
	TargetStackId() *string
	Synth()
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Synthesizes all resources to the output directory.
// Experimental.
func (a *jsiiProxy_App) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
// Experimental.
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
	Context *map[string]interface{} `json:"context"`
	// The directory to output Terraform resources.
	// Experimental.
	Outdir *string `json:"outdir"`
	// Whether to skip the validation during synthesis of the app.
	// Experimental.
	SkipValidation *bool `json:"skipValidation"`
	// Experimental.
	StackTraces *bool `json:"stackTraces"`
}

// Experimental.
type ArtifactoryBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type ArtifactoryBackendProps struct {
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Repo *string `json:"repo"`
	// Experimental.
	Subpath *string `json:"subpath"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	Username *string `json:"username"`
}

// Aspects can be applied to CDK tree scopes and can operate on the tree before synthesis.
// Experimental.
type Aspects interface {
	All() *[]IAspect
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

// Adds an aspect to apply this scope before synthesis.
// Experimental.
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
	AssetType_FILE AssetType = "FILE"
	AssetType_DIRECTORY AssetType = "DIRECTORY"
	AssetType_ARCHIVE AssetType = "ARCHIVE"
)

// Experimental.
type AzurermBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (a *jsiiProxy_AzurermBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AzurermBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AzurermBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type AzurermBackendProps struct {
	// Experimental.
	ContainerName *string `json:"containerName"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	StorageAccountName *string `json:"storageAccountName"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	ClientId *string `json:"clientId"`
	// Experimental.
	ClientSecret *string `json:"clientSecret"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Environment *string `json:"environment"`
	// Experimental.
	MsiEndpoint *string `json:"msiEndpoint"`
	// Experimental.
	ResourceGroupName *string `json:"resourceGroupName"`
	// Experimental.
	SasToken *string `json:"sasToken"`
	// Experimental.
	SubscriptionId *string `json:"subscriptionId"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	UseMsi *bool `json:"useMsi"`
}

// Experimental.
type BooleanMap interface {
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	Lookup(key *string) *bool
}

// The jsii proxy struct for BooleanMap
type jsiiProxy_BooleanMap struct {
	_ byte // padding
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

func (j *jsiiProxy_BooleanMap) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewBooleanMap(terraformResource ITerraformResource, terraformAttribute *string) BooleanMap {
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
func NewBooleanMap_Override(b BooleanMap, terraformResource ITerraformResource, terraformAttribute *string) {
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

func (j *jsiiProxy_BooleanMap) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BooleanMap) Lookup(key *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
type ComplexComputedList interface {
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) IResolvable
}

// The jsii proxy struct for ComplexComputedList
type jsiiProxy_ComplexComputedList struct {
	_ byte // padding
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

func (j *jsiiProxy_ComplexComputedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewComplexComputedList(terraformResource ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) ComplexComputedList {
	_init_.Initialize()

	j := jsiiProxy_ComplexComputedList{}

	_jsii_.Create(
		"cdktf.ComplexComputedList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewComplexComputedList_Override(c ComplexComputedList, terraformResource ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexComputedList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_ComplexComputedList) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ComplexComputedList) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
type ComplexObject interface {
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() IResolvable
	InterpolationForAttribute(property *string) IResolvable
}

// The jsii proxy struct for ComplexObject
type jsiiProxy_ComplexObject struct {
	_ byte // padding
}

func (j *jsiiProxy_ComplexObject) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
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

func (j *jsiiProxy_ComplexObject) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewComplexObject(terraformResource ITerraformResource, terraformAttribute *string, isSingleItem *bool) ComplexObject {
	_init_.Initialize()

	j := jsiiProxy_ComplexObject{}

	_jsii_.Create(
		"cdktf.ComplexObject",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

// Experimental.
func NewComplexObject_Override(c ComplexObject, terraformResource ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexObject",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_ComplexObject) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
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

func (j *jsiiProxy_ComplexObject) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ComplexObject) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
type ConsulBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (c *jsiiProxy_ConsulBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_ConsulBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_ConsulBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type ConsulBackendProps struct {
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	CaFile *string `json:"caFile"`
	// Experimental.
	CertFile *string `json:"certFile"`
	// Experimental.
	Datacenter *string `json:"datacenter"`
	// Experimental.
	Gzip *bool `json:"gzip"`
	// Experimental.
	HttpAuth *string `json:"httpAuth"`
	// Experimental.
	KeyFile *string `json:"keyFile"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Scheme *string `json:"scheme"`
}

// Experimental.
type CosBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (c *jsiiProxy_CosBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CosBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CosBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type CosBackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretId *string `json:"secretId"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
}

// Experimental.
type DataTerraformRemoteState interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Repo *string `json:"repo"`
	// Experimental.
	Subpath *string `json:"subpath"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateAzurerm interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	ContainerName *string `json:"containerName"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	StorageAccountName *string `json:"storageAccountName"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	ClientId *string `json:"clientId"`
	// Experimental.
	ClientSecret *string `json:"clientSecret"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Environment *string `json:"environment"`
	// Experimental.
	MsiEndpoint *string `json:"msiEndpoint"`
	// Experimental.
	ResourceGroupName *string `json:"resourceGroupName"`
	// Experimental.
	SasToken *string `json:"sasToken"`
	// Experimental.
	SubscriptionId *string `json:"subscriptionId"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	UseMsi *bool `json:"useMsi"`
}

// Experimental.
type DataTerraformRemoteStateConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
}

// Experimental.
type DataTerraformRemoteStateConsul interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	CaFile *string `json:"caFile"`
	// Experimental.
	CertFile *string `json:"certFile"`
	// Experimental.
	Datacenter *string `json:"datacenter"`
	// Experimental.
	Gzip *bool `json:"gzip"`
	// Experimental.
	HttpAuth *string `json:"httpAuth"`
	// Experimental.
	KeyFile *string `json:"keyFile"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Scheme *string `json:"scheme"`
}

// Experimental.
type DataTerraformRemoteStateCos interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretId *string `json:"secretId"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
}

// Experimental.
type DataTerraformRemoteStateEtcd interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Endpoints *string `json:"endpoints"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateEtcdV3 interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Endpoints *[]*string `json:"endpoints"`
	// Experimental.
	CacertPath *string `json:"cacertPath"`
	// Experimental.
	CertPath *string `json:"certPath"`
	// Experimental.
	KeyPath *string `json:"keyPath"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateGcs interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Credentials *string `json:"credentials"`
	// Experimental.
	EncryptionKey *string `json:"encryptionKey"`
	// Experimental.
	Prefix *string `json:"prefix"`
}

// Experimental.
type DataTerraformRemoteStateHttp interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	LockAddress *string `json:"lockAddress"`
	// Experimental.
	LockMethod *string `json:"lockMethod"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	RetryMax *float64 `json:"retryMax"`
	// Experimental.
	RetryWaitMax *float64 `json:"retryWaitMax"`
	// Experimental.
	RetryWaitMin *float64 `json:"retryWaitMin"`
	// Experimental.
	SkipCertVerification *bool `json:"skipCertVerification"`
	// Experimental.
	UnlockAddress *string `json:"unlockAddress"`
	// Experimental.
	UnlockMethod *string `json:"unlockMethod"`
	// Experimental.
	UpdateMethod *string `json:"updateMethod"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateLocal interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	WorkspaceDir *string `json:"workspaceDir"`
}

// Experimental.
type DataTerraformRemoteStateManta interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Account *string `json:"account"`
	// Experimental.
	KeyId *string `json:"keyId"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	InsecureSkipTlsVerify *bool `json:"insecureSkipTlsVerify"`
	// Experimental.
	KeyMaterial *string `json:"keyMaterial"`
	// Experimental.
	ObjectName *string `json:"objectName"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	User *string `json:"user"`
}

// Experimental.
type DataTerraformRemoteStateOss interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `json:"assumeRole"`
	// Experimental.
	EcsRoleName *string `json:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SecurityToken *string `json:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `json:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `json:"tablestoreTable"`
}

// Experimental.
type DataTerraformRemoteStatePg interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	ConnStr *string `json:"connStr"`
	// Experimental.
	SchemaName *string `json:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `json:"skipSchemaCreation"`
}

// Experimental.
type DataTerraformRemoteStateRemoteConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Organization *string `json:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `json:"workspaces"`
	// Experimental.
	Hostname *string `json:"hostname"`
	// Experimental.
	Token *string `json:"token"`
}

// Experimental.
type DataTerraformRemoteStateS3 interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRolePolicy *string `json:"assumeRolePolicy"`
	// Experimental.
	DynamodbEndpoint *string `json:"dynamodbEndpoint"`
	// Experimental.
	DynamodbTable *string `json:"dynamodbTable"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	ExternalId *string `json:"externalId"`
	// Experimental.
	ForcePathStyle *bool `json:"forcePathStyle"`
	// Experimental.
	IamEndpoint *string `json:"iamEndpoint"`
	// Experimental.
	KmsKeyId *string `json:"kmsKeyId"`
	// Experimental.
	MaxRetries *float64 `json:"maxRetries"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	RoleArn *string `json:"roleArn"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SessionName *string `json:"sessionName"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	SkipCredentialsValidation *bool `json:"skipCredentialsValidation"`
	// Experimental.
	SkipMetadataApiCheck *bool `json:"skipMetadataApiCheck"`
	// Experimental.
	SseCustomerKey *string `json:"sseCustomerKey"`
	// Experimental.
	StsEndpoint *string `json:"stsEndpoint"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	WorkspaceKeyPrefix *string `json:"workspaceKeyPrefix"`
}

// Experimental.
type DataTerraformRemoteStateSwift interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Container *string `json:"container"`
	// Experimental.
	ApplicationCredentialId *string `json:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `json:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `json:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `json:"archiveContainer"`
	// Experimental.
	AuthUrl *string `json:"authUrl"`
	// Experimental.
	CacertFile *string `json:"cacertFile"`
	// Experimental.
	Cert *string `json:"cert"`
	// Experimental.
	Cloud *string `json:"cloud"`
	// Experimental.
	DefaultDomain *string `json:"defaultDomain"`
	// Experimental.
	DomainId *string `json:"domainId"`
	// Experimental.
	DomainName *string `json:"domainName"`
	// Experimental.
	ExpireAfter *string `json:"expireAfter"`
	// Experimental.
	Insecure *bool `json:"insecure"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	ProjectDomainId *string `json:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `json:"projectDomainName"`
	// Experimental.
	RegionName *string `json:"regionName"`
	// Experimental.
	StateName *string `json:"stateName"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	TenantName *string `json:"tenantName"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	UserDomainId *string `json:"userDomainId"`
	// Experimental.
	UserDomainName *string `json:"userDomainName"`
	// Experimental.
	UserId *string `json:"userId"`
	// Experimental.
	UserName *string `json:"userName"`
}

// Default resolver implementation.
// Experimental.
type DefaultTokenResolver interface {
	ITokenResolver
	ResolveList(xs *[]*string, context IResolveContext) interface{}
	ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{}
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

// Resolve a tokenized list.
// Experimental.
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

// Resolve string fragments to Tokens.
// Experimental.
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

// Default Token resolution.
//
// Resolve the Token, recurse into whatever it returns,
// then finally post-process it.
// Experimental.
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
	DisplayHint *string `json:"displayHint"`
}

// Experimental.
type EtcdBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (e *jsiiProxy_EtcdBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EtcdBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EtcdBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type EtcdBackendProps struct {
	// Experimental.
	Endpoints *string `json:"endpoints"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type EtcdV3Backend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (e *jsiiProxy_EtcdV3Backend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type EtcdV3BackendProps struct {
	// Experimental.
	Endpoints *[]*string `json:"endpoints"`
	// Experimental.
	CacertPath *string `json:"cacertPath"`
	// Experimental.
	CertPath *string `json:"certPath"`
	// Experimental.
	KeyPath *string `json:"keyPath"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Username *string `json:"username"`
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
func Fn_Abs(value interface{}) *float64 {
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
func Fn_Abspath(value interface{}) *string {
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

// {@link https://www.terraform.io/docs/language/functions/anytrue.html anytrue} returns true if all elements in a given collection are true or "true".
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
func Fn_Base64decode(value interface{}) *string {
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
func Fn_Base64encode(value interface{}) *string {
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
func Fn_Base64gzip(value interface{}) *string {
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
func Fn_Base64sha256(value interface{}) *string {
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
func Fn_Base64sha512(value interface{}) *string {
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
func Fn_Basename(value interface{}) *string {
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
func Fn_Bcrypt(value interface{}, cost interface{}) *string {
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
func Fn_Ceil(value interface{}) *float64 {
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
func Fn_Chomp(value interface{}) *string {
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
func Fn_Chunklist(value *[]interface{}, chunkSize interface{}) *[]*string {
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
func Fn_Cidrhost(prefix interface{}, hostnum interface{}) *string {
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
func Fn_Cidrnetmask(prefix interface{}) *string {
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
func Fn_Cidrsubnet(prefix interface{}, newbits interface{}, netnum interface{}) *string {
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
func Fn_Cidrsubnets(prefix interface{}, newbits *[]interface{}) *[]*string {
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
func Fn_Contains(list *[]interface{}, value interface{}) IResolvable {
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
func Fn_Csvdecode(value interface{}) *[]*string {
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
func Fn_Dirname(value interface{}) *string {
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
func Fn_Distinct(list *[]interface{}) *[]*string {
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
func Fn_Element(list interface{}, index interface{}) interface{} {
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
func Fn_File(value interface{}) *string {
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
func Fn_Filebase64(value interface{}) *string {
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
func Fn_Filebase64sha256(value interface{}) *string {
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
func Fn_Filebase64sha512(value interface{}) *string {
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
func Fn_Fileexists(value interface{}) IResolvable {
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
func Fn_Filemd5(value interface{}) *string {
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
func Fn_Fileset(path interface{}, pattern interface{}) *[]*string {
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
func Fn_Filesha1(value interface{}) *string {
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
func Fn_Filesha256(value interface{}) *string {
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
func Fn_Filesha512(value interface{}) *string {
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
func Fn_Floor(value interface{}) *float64 {
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
func Fn_Format(spec interface{}, values *[]interface{}) *string {
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
func Fn_Formatdate(spec interface{}, timestamp interface{}) *string {
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
func Fn_Formatlist(spec interface{}, values *[]interface{}) *[]*string {
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
func Fn_Indent(indentation interface{}, value interface{}) *string {
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
func Fn_Join(separator interface{}, value *[]interface{}) *string {
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
func Fn_Jsondecode(value interface{}) interface{} {
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
func Fn_Log(value interface{}, base interface{}) *float64 {
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

// {@link https://www.terraform.io/docs/language/functions/matchkeys.html lookup} retrieves the value of a single element from a map, given its key. If the given key does not exist, the given default value is returned instead.
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
func Fn_Lower(value interface{}) *string {
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
func Fn_Max(values *[]interface{}) *float64 {
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
func Fn_Md5(value interface{}) *string {
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
func Fn_Merge(values *[]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"merge",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/min.html min} takes one or more numbers and returns the smallest number from the set.
// Experimental.
func Fn_Min(values *[]interface{}) *float64 {
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
func Fn_ParseInt(value interface{}, base interface{}) *float64 {
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
func Fn_Pathexpand(value interface{}) *string {
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
func Fn_Pow(value interface{}, power interface{}) *float64 {
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
func Fn_Range(start interface{}, limit interface{}, step *float64) *[]*string {
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

// {@link https://www.terraform.io/docs/language/functions/regexall.html regexall} applies a regular expression to a string and returns a list of all matches.
// Experimental.
func Fn_Regexall(pattern interface{}, value interface{}) *[]*string {
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
func Fn_Replace(value interface{}, substring interface{}, replacement interface{}) *string {
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
func Fn_Reverse(values *[]interface{}) *[]*string {
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
func Fn_Rsadecrypt(ciphertext interface{}, privatekey interface{}) *string {
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
func Fn_Sha1(value interface{}) *string {
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
func Fn_Sha256(value interface{}) *string {
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
func Fn_Sha512(value interface{}) *string {
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
func Fn_Signum(value interface{}) *float64 {
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
func Fn_Slice(list interface{}, startindex interface{}, endindex interface{}) *[]*string {
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
func Fn_Split(seperator interface{}, value interface{}) *[]*string {
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
func Fn_Strrev(value interface{}) *string {
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
func Fn_Substr(value interface{}, offset interface{}, length interface{}) *string {
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
func Fn_Templatefile(path interface{}, vars interface{}) *string {
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
func Fn_Textdecodebase64(value interface{}, encodingName interface{}) *string {
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
func Fn_Textencodebase64(value interface{}, encodingName interface{}) *string {
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
func Fn_Timeadd(timestamp interface{}, duration interface{}) *string {
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
func Fn_Title(value interface{}) *string {
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
func Fn_Trim(value interface{}, replacement interface{}) *string {
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
func Fn_Trimprefix(value interface{}, prefix interface{}) *string {
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
func Fn_Trimspace(value interface{}) *string {
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
func Fn_Trimsuffix(value interface{}, suffix interface{}) *string {
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
func Fn_Upper(value interface{}) *string {
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
func Fn_Urlencode(value interface{}) *string {
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
func Fn_Uuidv5(namespace interface{}, name interface{}) *string {
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
func Fn_Yamldecode(value interface{}) interface{} {
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
func Fn_Zipmap(keyslist *[]interface{}, valueslist *[]interface{}) interface{} {
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
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (g *jsiiProxy_GcsBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (g *jsiiProxy_GcsBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (g *jsiiProxy_GcsBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type GcsBackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Credentials *string `json:"credentials"`
	// Experimental.
	EncryptionKey *string `json:"encryptionKey"`
	// Experimental.
	Prefix *string `json:"prefix"`
}

// Experimental.
type HttpBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (h *jsiiProxy_HttpBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (h *jsiiProxy_HttpBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (h *jsiiProxy_HttpBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type HttpBackendProps struct {
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	LockAddress *string `json:"lockAddress"`
	// Experimental.
	LockMethod *string `json:"lockMethod"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	RetryMax *float64 `json:"retryMax"`
	// Experimental.
	RetryWaitMax *float64 `json:"retryWaitMax"`
	// Experimental.
	RetryWaitMin *float64 `json:"retryWaitMin"`
	// Experimental.
	SkipCertVerification *bool `json:"skipCertVerification"`
	// Experimental.
	UnlockAddress *string `json:"unlockAddress"`
	// Experimental.
	UnlockMethod *string `json:"unlockMethod"`
	// Experimental.
	UpdateMethod *string `json:"updateMethod"`
	// Experimental.
	Username *string `json:"username"`
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
	// True when we are still preparing, false if we're rendering the final output.
	// Experimental.
	Preparing() *bool
	// The scope from which resolution has been initiated.
	// Experimental.
	Scope() constructs.IConstruct
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
type ITerraformResource interface {
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(c interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(d *[]*string)
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

func (j *jsiiProxy_ITerraformResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetCount(val interface{}) {
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
	// Resolve a string with at least one stringified token in it.
	//
	// (May use concatenation)
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
	DisplayHint *string `json:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	// Experimental.
	OmitEmptyArray *bool `json:"omitEmptyArray"`
}

// Experimental.
type LazyBase interface {
	IResolvable
	CreationStack() *[]*string
	AddPostProcessor(postProcessor IPostProcessor)
	Resolve(context IResolveContext) interface{}
	ResolveLazy(context IResolveContext) interface{}
	ToJSON() interface{}
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

// Experimental.
func (l *jsiiProxy_LazyBase) AddPostProcessor(postProcessor IPostProcessor) {
	_jsii_.InvokeVoid(
		l,
		"addPostProcessor",
		[]interface{}{postProcessor},
	)
}

// Produce the Token's value at resolution time.
// Experimental.
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

// Experimental.
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

// Turn this Token into JSON.
//
// Called automatically when JSON.stringify() is called on a Token.
// Experimental.
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

// Return a string representation of this resolvable object.
//
// Returns a reversible string representation.
// Experimental.
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
	DisplayHint *string `json:"displayHint"`
	// If the produced list is empty, return 'undefined' instead.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty"`
}

// Options for creating a lazy string token.
// Experimental.
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `json:"displayHint"`
}

// Experimental.
type LocalBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (l *jsiiProxy_LocalBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LocalBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LocalBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type LocalBackendProps struct {
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	WorkspaceDir *string `json:"workspaceDir"`
}

// Experimental.
type Manifest interface {
	IManifest
	Outdir() *string
	Stacks() *map[string]*StackManifest
	Version() *string
	BuildManifest() IManifest
	ForStack(stack TerraformStack) *StackManifest
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

// Experimental.
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

// Experimental.
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

// Experimental.
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
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (m *jsiiProxy_MantaBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MantaBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MantaBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Account *string `json:"account"`
	// Experimental.
	KeyId *string `json:"keyId"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	InsecureSkipTlsVerify *bool `json:"insecureSkipTlsVerify"`
	// Experimental.
	KeyMaterial *string `json:"keyMaterial"`
	// Experimental.
	ObjectName *string `json:"objectName"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	User *string `json:"user"`
}

// Experimental.
type NamedRemoteWorkspace interface {
	IRemoteWorkspace
	Name() *string
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
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	Lookup(key *string) *float64
}

// The jsii proxy struct for NumberMap
type jsiiProxy_NumberMap struct {
	_ byte // padding
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

func (j *jsiiProxy_NumberMap) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewNumberMap(terraformResource ITerraformResource, terraformAttribute *string) NumberMap {
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
func NewNumberMap_Override(n NumberMap, terraformResource ITerraformResource, terraformAttribute *string) {
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

func (j *jsiiProxy_NumberMap) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
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

// Experimental.
type OssAssumeRole struct {
	// Experimental.
	RoleArn *string `json:"roleArn"`
	// Experimental.
	Policy *string `json:"policy"`
	// Experimental.
	SessionExpiration *float64 `json:"sessionExpiration"`
	// Experimental.
	SessionName *string `json:"sessionName"`
}

// Experimental.
type OssBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (o *jsiiProxy_OssBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OssBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OssBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `json:"assumeRole"`
	// Experimental.
	EcsRoleName *string `json:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SecurityToken *string `json:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `json:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `json:"tablestoreTable"`
}

// Experimental.
type PgBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (p *jsiiProxy_PgBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PgBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (p *jsiiProxy_PgBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	ConnStr *string `json:"connStr"`
	// Experimental.
	SchemaName *string `json:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `json:"skipSchemaCreation"`
}

// Experimental.
type PrefixedRemoteWorkspaces interface {
	IRemoteWorkspace
	Prefix() *string
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
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (r *jsiiProxy_RemoteBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (r *jsiiProxy_RemoteBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (r *jsiiProxy_RemoteBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Organization *string `json:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `json:"workspaces"`
	// Experimental.
	Hostname *string `json:"hostname"`
	// Experimental.
	Token *string `json:"token"`
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
	Resolver ITokenResolver `json:"resolver"`
	// The scope from which resolution is performed.
	// Experimental.
	Scope constructs.IConstruct `json:"scope"`
	// Whether the resolution is being executed during the prepare phase or not.
	// Experimental.
	Preparing *bool `json:"preparing"`
}

// A construct which represents a resource.
// Experimental.
type Resource interface {
	constructs.Construct
	IResource
	Node() constructs.Node
	Stack() TerraformStack
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Returns a string representation of this construct.
// Experimental.
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
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (s *jsiiProxy_S3Backend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_S3Backend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_S3Backend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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

// Experimental.
type S3BackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRolePolicy *string `json:"assumeRolePolicy"`
	// Experimental.
	DynamodbEndpoint *string `json:"dynamodbEndpoint"`
	// Experimental.
	DynamodbTable *string `json:"dynamodbTable"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	ExternalId *string `json:"externalId"`
	// Experimental.
	ForcePathStyle *bool `json:"forcePathStyle"`
	// Experimental.
	IamEndpoint *string `json:"iamEndpoint"`
	// Experimental.
	KmsKeyId *string `json:"kmsKeyId"`
	// Experimental.
	MaxRetries *float64 `json:"maxRetries"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	RoleArn *string `json:"roleArn"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SessionName *string `json:"sessionName"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	SkipCredentialsValidation *bool `json:"skipCredentialsValidation"`
	// Experimental.
	SkipMetadataApiCheck *bool `json:"skipMetadataApiCheck"`
	// Experimental.
	SseCustomerKey *string `json:"sseCustomerKey"`
	// Experimental.
	StsEndpoint *string `json:"stsEndpoint"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	WorkspaceKeyPrefix *string `json:"workspaceKeyPrefix"`
}

// Experimental.
type StackAnnotation struct {
	// Experimental.
	ConstructPath *string `json:"constructPath"`
	// Experimental.
	Level AnnotationMetadataEntryType `json:"level"`
	// Experimental.
	Message *string `json:"message"`
	// Experimental.
	Stacktrace *[]*string `json:"stacktrace"`
}

// Experimental.
type StackManifest struct {
	// Experimental.
	Annotations *[]*StackAnnotation `json:"annotations"`
	// Experimental.
	ConstructPath *string `json:"constructPath"`
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	SynthesizedStackPath *string `json:"synthesizedStackPath"`
	// Experimental.
	WorkingDirectory *string `json:"workingDirectory"`
}

// Converts all fragments to strings and concats those.
//
// Drops 'undefined's.
// Experimental.
type StringConcat interface {
	IFragmentConcatenator
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

// Join the fragment on the left and on the right.
// Experimental.
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
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	Lookup(key *string) *string
}

// The jsii proxy struct for StringMap
type jsiiProxy_StringMap struct {
	_ byte // padding
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

func (j *jsiiProxy_StringMap) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewStringMap(terraformResource ITerraformResource, terraformAttribute *string) StringMap {
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
func NewStringMap_Override(s StringMap, terraformResource ITerraformResource, terraformAttribute *string) {
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

func (j *jsiiProxy_StringMap) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
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

// Experimental.
type SwiftBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (s *jsiiProxy_SwiftBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_SwiftBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SwiftBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Container *string `json:"container"`
	// Experimental.
	ApplicationCredentialId *string `json:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `json:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `json:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `json:"archiveContainer"`
	// Experimental.
	AuthUrl *string `json:"authUrl"`
	// Experimental.
	CacertFile *string `json:"cacertFile"`
	// Experimental.
	Cert *string `json:"cert"`
	// Experimental.
	Cloud *string `json:"cloud"`
	// Experimental.
	DefaultDomain *string `json:"defaultDomain"`
	// Experimental.
	DomainId *string `json:"domainId"`
	// Experimental.
	DomainName *string `json:"domainName"`
	// Experimental.
	ExpireAfter *string `json:"expireAfter"`
	// Experimental.
	Insecure *bool `json:"insecure"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	ProjectDomainId *string `json:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `json:"projectDomainName"`
	// Experimental.
	RegionName *string `json:"regionName"`
	// Experimental.
	StateName *string `json:"stateName"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	TenantName *string `json:"tenantName"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	UserDomainId *string `json:"userDomainId"`
	// Experimental.
	UserDomainName *string `json:"userDomainName"`
	// Experimental.
	UserId *string `json:"userId"`
	// Experimental.
	UserName *string `json:"userName"`
}

// Experimental.
type TerraformAsset interface {
	Resource
	AssetHash() *string
	SetAssetHash(val *string)
	FileName() *string
	Node() constructs.Node
	Path() *string
	Stack() TerraformStack
	Type() AssetType
	SetType(val AssetType)
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Returns a string representation of this construct.
// Experimental.
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
	Path *string `json:"path"`
	// Experimental.
	AssetHash *string `json:"assetHash"`
	// Experimental.
	Type AssetType `json:"type"`
}

// Experimental.
type TerraformBackend interface {
	TerraformElement
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	ITerraformDependable
	ITerraformResource
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Lifecycle() *TerraformResourceLifecycle
	SetLifecycle(val *TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() TerraformProvider
	SetProvider(val TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformDataSource
type jsiiProxy_TerraformDataSource struct {
	jsiiProxy_TerraformElement
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

func (j *jsiiProxy_TerraformDataSource) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_TerraformDataSource) SetCount(val interface{}) {
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
func NewTerraformElement(scope constructs.Construct, id *string) TerraformElement {
	_init_.Initialize()

	j := jsiiProxy_TerraformElement{}

	_jsii_.Create(
		"cdktf.TerraformElement",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformElement_Override(t TerraformElement, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformElement",
		[]interface{}{scope, id},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformElement) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformElement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformElement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Path *string `json:"path"`
	// Experimental.
	StackTrace *[]*string `json:"stackTrace"`
	// Experimental.
	UniqueId *string `json:"uniqueId"`
}

// Experimental.
type TerraformHclModule interface {
	TerraformModule
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	Providers() *[]interface{}
	RawOverrides() interface{}
	Source() *string
	Variables() *map[string]interface{}
	Version() *string
	AddOverride(path *string, value interface{})
	AddProvider(provider interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	InterpolationForOutput(moduleOutput *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	Set(variable *string, value interface{})
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformHclModule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) AddProvider(provider interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addProvider",
		[]interface{}{provider},
	)
}

// Experimental.
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

// Experimental.
func (t *jsiiProxy_TerraformHclModule) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
func (t *jsiiProxy_TerraformHclModule) InterpolationForOutput(moduleOutput *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformHclModule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformHclModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) Set(variable *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"set",
		[]interface{}{variable, value},
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Source *string `json:"source"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Providers *[]interface{} `json:"providers"`
	// Experimental.
	Version *string `json:"version"`
	// Experimental.
	Variables *map[string]interface{} `json:"variables"`
}

// Experimental.
type TerraformLocal interface {
	TerraformElement
	ITerraformAddressable
	AsBoolean() IResolvable
	AsList() *[]*string
	AsNumber() *float64
	AsString() *string
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Expression() interface{}
	SetExpression(val interface{})
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformLocal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformLocal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformLocal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `json:"provider"`
}

// Experimental.
type TerraformModule interface {
	TerraformElement
	ITerraformDependable
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	Providers() *[]interface{}
	RawOverrides() interface{}
	Source() *string
	Version() *string
	AddOverride(path *string, value interface{})
	AddProvider(provider interface{})
	InterpolationForOutput(moduleOutput *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformModule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformModule) AddProvider(provider interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addProvider",
		[]interface{}{provider},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformModule) InterpolationForOutput(moduleOutput *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformModule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Source *string `json:"source"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Providers *[]interface{} `json:"providers"`
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type TerraformModuleProvider struct {
	// Experimental.
	ModuleAlias *string `json:"moduleAlias"`
	// Experimental.
	Provider TerraformProvider `json:"provider"`
}

// Experimental.
type TerraformOutput interface {
	TerraformElement
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	DependsOn() *[]ITerraformDependable
	SetDependsOn(val *[]ITerraformDependable)
	Description() *string
	SetDescription(val *string)
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	Sensitive() *bool
	SetSensitive(val *bool)
	StaticId() *bool
	SetStaticId(val *bool)
	Value() interface{}
	SetValue(val interface{})
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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
func (t *jsiiProxy_TerraformOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformOutput) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Value interface{} `json:"value"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Description *string `json:"description"`
	// Experimental.
	Sensitive *bool `json:"sensitive"`
	// If set to true the synthesized Terraform Output will be named after the `id` passed to the constructor instead of the default (TerraformOutput.friendlyUniqueId).
	// Experimental.
	StaticId *bool `json:"staticId"`
}

// Experimental.
type TerraformProvider interface {
	TerraformElement
	Alias() *string
	SetAlias(val *string)
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	MetaAttributes() *map[string]interface{}
	Node() constructs.Node
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	TerraformProviderSource() *string
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	TerraformResourceType *string `json:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `json:"terraformGeneratorMetadata"`
	// Experimental.
	TerraformProviderSource *string `json:"terraformProviderSource"`
}

// Experimental.
type TerraformProviderGeneratorMetadata struct {
	// Experimental.
	ProviderName *string `json:"providerName"`
	// Experimental.
	ProviderVersionConstraint *string `json:"providerVersionConstraint"`
}

// Experimental.
type TerraformRemoteState interface {
	TerraformElement
	ITerraformAddressable
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Node() constructs.Node
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ToMetadata() interface{}
	ToString() *string
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	ITerraformDependable
	ITerraformResource
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Lifecycle() *TerraformResourceLifecycle
	SetLifecycle(val *TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() TerraformProvider
	SetProvider(val TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformResource
type jsiiProxy_TerraformResource struct {
	jsiiProxy_TerraformElement
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

func (j *jsiiProxy_TerraformResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_TerraformResource) SetCount(val interface{}) {
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

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Adds this resource to the terraform JSON output.
// Experimental.
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
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `json:"provider"`
	// Experimental.
	TerraformResourceType *string `json:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `json:"terraformGeneratorMetadata"`
}

// Experimental.
type TerraformResourceLifecycle struct {
	// Experimental.
	CreateBeforeDestroy *bool `json:"createBeforeDestroy"`
	// Experimental.
	IgnoreChanges *[]*string `json:"ignoreChanges"`
	// Experimental.
	PreventDestroy *bool `json:"preventDestroy"`
}

// Experimental.
type TerraformStack interface {
	constructs.Construct
	Node() constructs.Node
	Synthesizer() IStackSynthesizer
	SetSynthesizer(val IStackSynthesizer)
	AddOverride(path *string, value interface{})
	AllocateLogicalId(tfElement interface{}) *string
	AllProviders() *[]TerraformProvider
	GetLogicalId(tfElement interface{}) *string
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformStack
type jsiiProxy_TerraformStack struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_TerraformStack) SetSynthesizer(val IStackSynthesizer) {
	_jsii_.Set(
		j,
		"synthesizer",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Returns the naming scheme used to allocate logical IDs.
//
// By default, uses
// the `HashedAddressingScheme` but this method can be overridden to customize
// this behavior.
// Experimental.
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

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Backend *string `json:"backend"`
	// Experimental.
	StackName *string `json:"stackName"`
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type TerraformVariable interface {
	TerraformElement
	ITerraformAddressable
	BooleanValue() *bool
	CdktfStack() TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Default() interface{}
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	ListValue() *[]*string
	Node() constructs.Node
	NumberValue() *float64
	RawOverrides() interface{}
	Sensitive() *bool
	StringValue() *string
	Type() *string
	Value() interface{}
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformVariable
type jsiiProxy_TerraformVariable struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformAddressable
}

func (j *jsiiProxy_TerraformVariable) BooleanValue() *bool {
	var returns *bool
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Experimental.
func (t *jsiiProxy_TerraformVariable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformVariable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (t *jsiiProxy_TerraformVariable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

// Experimental.
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

// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Experimental.
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
	Default interface{} `json:"default"`
	// Experimental.
	Description *string `json:"description"`
	// Experimental.
	Sensitive *bool `json:"sensitive"`
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
	Type *string `json:"type"`
}

// Testing utilities for cdktf applications.
// Experimental.
type Testing interface {
}

// The jsii proxy struct for Testing
type jsiiProxy_Testing struct {
	_ byte // padding
}

// Returns an app for testing with the following properties: - Output directory is a temp dir.
// Experimental.
func Testing_App() App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"app",
		nil, // no parameters
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
	FirstToken() IResolvable
	FirstValue() interface{}
	Intrinsic() *[]IResolvable
	Length() *float64
	Literals() *[]IResolvable
	Tokens() *[]IResolvable
	AddIntrinsic(value interface{})
	AddLiteral(lit interface{})
	AddToken(token IResolvable)
	Join(concat IFragmentConcatenator) interface{}
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

// Adds an intrinsic fragment.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) AddIntrinsic(value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addIntrinsic",
		[]interface{}{value},
	)
}

// Adds a literal fragment.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) AddLiteral(lit interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addLiteral",
		[]interface{}{lit},
	)
}

// Adds a token fragment.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) AddToken(token IResolvable) {
	_jsii_.InvokeVoid(
		t,
		"addToken",
		[]interface{}{token},
	)
}

// Combine the string fragments using the given joiner.
//
// If there are any
// Experimental.
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

// Apply a transformation function to all tokens in the string.
// Experimental.
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

