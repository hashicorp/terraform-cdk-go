// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

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

