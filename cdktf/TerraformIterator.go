// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

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

	if err := validateTerraformIterator_FromListParameters(list); err != nil {
		panic(err)
	}
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

	if err := validateTerraformIterator_FromMapParameters(map_); err != nil {
		panic(err)
	}
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
	if err := t.validateDynamicParameters(attributes); err != nil {
		panic(err)
	}
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
	if err := t.validateGetAnyParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetAnyMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetBooleanParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetBooleanMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetListParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetNumberParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetNumberListParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetNumberMapParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetStringParameters(attribute); err != nil {
		panic(err)
	}
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
	if err := t.validateGetStringMapParameters(attribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMap",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

