// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// Experimental.
type ListTerraformIterator interface {
	TerraformIterator
	// Returns the currenty entry in the list or set that is being iterated over.
	//
	// For lists this is the same as `iterator.value`. If you need the index,
	// use count using the escape hatch:
	// https://developer.hashicorp.com/terraform/cdktf/concepts/resources#escape-hatch
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

	if err := validateNewListTerraformIteratorParameters(list); err != nil {
		panic(err)
	}
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

	if err := validateListTerraformIterator_FromListParameters(list); err != nil {
		panic(err)
	}
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

	if err := validateListTerraformIterator_FromMapParameters(map_); err != nil {
		panic(err)
	}
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
	if err := l.validateDynamicParameters(attributes); err != nil {
		panic(err)
	}
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
	if err := l.validateGetAnyParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetAnyMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetBooleanParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetBooleanMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetListParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetNumberParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetNumberListParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetNumberMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetStringParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := l.validateGetStringMapParameters(attribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

