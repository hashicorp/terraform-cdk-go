// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

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

	if err := validateNewMapTerraformIteratorParameters(map_); err != nil {
		panic(err)
	}
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

	if err := validateMapTerraformIterator_FromListParameters(list); err != nil {
		panic(err)
	}
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

	if err := validateMapTerraformIterator_FromMapParameters(map_); err != nil {
		panic(err)
	}
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
	if err := m.validateDynamicParameters(attributes); err != nil {
		panic(err)
	}
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
	if err := m.validateGetAnyParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetAnyMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetBooleanParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetBooleanMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetListParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetNumberParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetNumberListParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetNumberMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetStringParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := m.validateGetStringMapParameters(attribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

