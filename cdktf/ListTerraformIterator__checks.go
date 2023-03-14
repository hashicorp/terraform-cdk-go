// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (l *jsiiProxy_ListTerraformIterator) validateDynamicParameters(attributes *map[string]interface{}) error {
	if attributes == nil {
		return fmt.Errorf("parameter attributes is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetAnyParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetAnyMapParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetBooleanParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetBooleanMapParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetListParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetMapParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetNumberParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetNumberListParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetNumberMapParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetStringParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListTerraformIterator) validateGetStringMapParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func validateListTerraformIterator_FromListParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]*string:
		// ok
	case []*string:
		// ok
	case IResolvable:
		// ok
	case *[]*float64:
		// ok
	case []*float64:
		// ok
	case ComplexList:
		// ok
	case StringMapList:
		// ok
	case NumberMapList:
		// ok
	case BooleanMapList:
		// ok
	case AnyMapList:
		// ok
	case *[]interface{}:
		list := list.(*[]interface{})
		for idx_a33039, v := range *list {
			switch v.(type) {
			case *bool:
				// ok
			case bool:
				// ok
			case IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter list[%#v] must be one of the allowed types: *bool, IResolvable; received %#v (a %T)", idx_a33039, v, v)
				}
			}
		}
	case []interface{}:
		list_ := list.([]interface{})
		list := &list_
		for idx_a33039, v := range *list {
			switch v.(type) {
			case *bool:
				// ok
			case bool:
				// ok
			case IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter list[%#v] must be one of the allowed types: *bool, IResolvable; received %#v (a %T)", idx_a33039, v, v)
				}
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]*string, IResolvable, *[]*float64, ComplexList, StringMapList, NumberMapList, BooleanMapList, AnyMapList, *[]interface{}; received %#v (a %T)", list, list)
		}
	}

	return nil
}

func validateListTerraformIterator_FromMapParameters(map_ interface{}) error {
	if map_ == nil {
		return fmt.Errorf("parameter map_ is required, but nil was provided")
	}
	switch map_.(type) {
	case ComplexMap:
		// ok
	case *map[string]interface{}:
		// ok
	case map[string]interface{}:
		// ok
	case *map[string]*string:
		// ok
	case map[string]*string:
		// ok
	case *map[string]*float64:
		// ok
	case map[string]*float64:
		// ok
	case *map[string]*bool:
		// ok
	case map[string]*bool:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(map_) {
			return fmt.Errorf("parameter map_ must be one of the allowed types: ComplexMap, *map[string]interface{}, *map[string]*string, *map[string]*float64, *map[string]*bool; received %#v (a %T)", map_, map_)
		}
	}

	return nil
}

func validateNewListTerraformIteratorParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]*string:
		// ok
	case []*string:
		// ok
	case IResolvable:
		// ok
	case *[]*float64:
		// ok
	case []*float64:
		// ok
	case ComplexList:
		// ok
	case StringMapList:
		// ok
	case NumberMapList:
		// ok
	case BooleanMapList:
		// ok
	case AnyMapList:
		// ok
	case *[]interface{}:
		list := list.(*[]interface{})
		for idx_a33039, v := range *list {
			switch v.(type) {
			case *bool:
				// ok
			case bool:
				// ok
			case IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter list[%#v] must be one of the allowed types: *bool, IResolvable; received %#v (a %T)", idx_a33039, v, v)
				}
			}
		}
	case []interface{}:
		list_ := list.([]interface{})
		list := &list_
		for idx_a33039, v := range *list {
			switch v.(type) {
			case *bool:
				// ok
			case bool:
				// ok
			case IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter list[%#v] must be one of the allowed types: *bool, IResolvable; received %#v (a %T)", idx_a33039, v, v)
				}
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]*string, IResolvable, *[]*float64, ComplexList, StringMapList, NumberMapList, BooleanMapList, AnyMapList, *[]interface{}; received %#v (a %T)", list, list)
		}
	}

	return nil
}

