// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
type DataTerraformRemoteStateEtcd interface {
	TerraformRemoteState
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	CdktfStack() TerraformStack
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ConstructNodeMetadata() *map[string]interface{}
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Fqn() *string
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	FriendlyUniqueId() *string
	// The tree node.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Node() constructs.Node
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	RawOverrides() interface{}
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	AddOverride(path *string, value interface{})
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Get(output *string) IResolvable
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	GetBoolean(output *string) IResolvable
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	GetList(output *string) *[]*string
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	GetNumber(output *string) *float64
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	GetString(output *string) *string
	// Overrides the auto-generated logical ID with a specific ID.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ResetOverrideLogicalId()
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
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


// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func NewDataTerraformRemoteStateEtcd(scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdConfig) DataTerraformRemoteStateEtcd {
	_init_.Initialize()

	if err := validateNewDataTerraformRemoteStateEtcdParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTerraformRemoteStateEtcd{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcd",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
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
// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func DataTerraformRemoteStateEtcd_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTerraformRemoteStateEtcd_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateEtcd",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func DataTerraformRemoteStateEtcd_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTerraformRemoteStateEtcd_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.DataTerraformRemoteStateEtcd",
		"isTerraformElement",
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
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTerraformRemoteStateEtcd) Get(output *string) IResolvable {
	if err := d.validateGetParameters(output); err != nil {
		panic(err)
	}
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
	if err := d.validateGetBooleanParameters(output); err != nil {
		panic(err)
	}
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
	if err := d.validateGetListParameters(output); err != nil {
		panic(err)
	}
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
	if err := d.validateGetNumberParameters(output); err != nil {
		panic(err)
	}
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
	if err := d.validateGetStringParameters(output); err != nil {
		panic(err)
	}
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
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
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

