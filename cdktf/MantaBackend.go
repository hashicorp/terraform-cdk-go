// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
type MantaBackend interface {
	TerraformBackend
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	CdktfStack() TerraformStack
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	ConstructNodeMetadata() *map[string]interface{}
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	Fqn() *string
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	FriendlyUniqueId() *string
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	Name() *string
	// The tree node.
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	Node() constructs.Node
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	RawOverrides() interface{}
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	AddOverride(path *string, value interface{})
	// Creates a TerraformRemoteState resource that accesses this backend.
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState
	// Overrides the auto-generated logical ID with a specific ID.
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	ResetOverrideLogicalId()
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	SynthesizeAttributes() *map[string]interface{}
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
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


// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
func NewMantaBackend(scope constructs.Construct, props *MantaBackendConfig) MantaBackend {
	_init_.Initialize()

	if err := validateNewMantaBackendParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MantaBackend{}

	_jsii_.Create(
		"cdktf.MantaBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
func NewMantaBackend_Override(m MantaBackend, scope constructs.Construct, props *MantaBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.MantaBackend",
		[]interface{}{scope, props},
		m,
	)
}

// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
func MantaBackend_IsBackend(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMantaBackend_IsBackendParameters(x); err != nil {
		panic(err)
	}
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
// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
func MantaBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMantaBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.MantaBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Deprecated: CDK for Terraform no longer supports the manta backend. Terraform deprecated manta in v1.2.3 and removed it in v1.3.
func MantaBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMantaBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.MantaBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MantaBackend) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MantaBackend) GetRemoteStateDataSource(scope constructs.Construct, name *string, _fromStack *string) TerraformRemoteState {
	if err := m.validateGetRemoteStateDataSourceParameters(scope, name, _fromStack); err != nil {
		panic(err)
	}
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
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
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

