// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

