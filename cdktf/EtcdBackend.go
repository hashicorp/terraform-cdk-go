// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
type EtcdBackend interface {
	TerraformBackend
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	CdktfStack() TerraformStack
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ConstructNodeMetadata() *map[string]interface{}
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Fqn() *string
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	FriendlyUniqueId() *string
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Name() *string
	// The tree node.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	Node() constructs.Node
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	RawOverrides() interface{}
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ResetOverrideLogicalId()
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	SynthesizeAttributes() *map[string]interface{}
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
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


// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func NewEtcdBackend(scope constructs.Construct, props *EtcdBackendConfig) EtcdBackend {
	_init_.Initialize()

	if err := validateNewEtcdBackendParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EtcdBackend{}

	_jsii_.Create(
		"cdktf.EtcdBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func NewEtcdBackend_Override(e EtcdBackend, scope constructs.Construct, props *EtcdBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.EtcdBackend",
		[]interface{}{scope, props},
		e,
	)
}

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func EtcdBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEtcdBackend_IsBackendParameters(x); err != nil {
		panic(err)
	}
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
// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func EtcdBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEtcdBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.EtcdBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Deprecated: CDK for Terraform no longer supports the etcd backend. Terraform deprecated etcd in v1.2.3 and removed it in v1.3.
func EtcdBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEtcdBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.EtcdBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EtcdBackend) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EtcdBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	if err := e.validateGetRemoteStateDataSourceParameters(scope, name, _fromStack); err != nil {
		panic(err)
	}
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
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
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

