//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateFn_AbsParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_AbspathParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_AlltrueParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_AnytrueParameters(value *[]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Base64decodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Base64encodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Base64gzipParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Base64sha256Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Base64sha512Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_BasenameParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_BcryptParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_CanParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_CeilParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ChompParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ChunklistParameters(value *[]interface{}, chunkSize *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if chunkSize == nil {
		return fmt.Errorf("parameter chunkSize is required, but nil was provided")
	}

	return nil
}

func validateFn_CidrhostParameters(prefix *string, hostnum *float64) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	if hostnum == nil {
		return fmt.Errorf("parameter hostnum is required, but nil was provided")
	}

	return nil
}

func validateFn_CidrnetmaskParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func validateFn_CidrsubnetParameters(prefix *string, newbits *float64, netnum *float64) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	if newbits == nil {
		return fmt.Errorf("parameter newbits is required, but nil was provided")
	}

	if netnum == nil {
		return fmt.Errorf("parameter netnum is required, but nil was provided")
	}

	return nil
}

func validateFn_CidrsubnetsParameters(prefix *string, newbits *[]*float64) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	if newbits == nil {
		return fmt.Errorf("parameter newbits is required, but nil was provided")
	}

	return nil
}

func validateFn_CoalesceParameters(value *[]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_CoalescelistParameters(value *[]*[]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_CompactParameters(value *[]*string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ConcatParameters(value *[]*[]interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ContainsParameters(list interface{}, value interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_CsvdecodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_DirnameParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_DistinctParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	return nil
}

func validateFn_ElementParameters(list interface{}, index *float64) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func validateFn_FileParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Filebase64Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Filebase64sha256Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Filebase64sha512Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_FileexistsParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Filemd5Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_FilesetParameters(path *string, pattern *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateFn_Filesha1Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Filesha256Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Filesha512Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_FlattenParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	return nil
}

func validateFn_FloorParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_FormatParameters(spec *string, values *[]interface{}) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_FormatdateParameters(spec *string, timestamp *string) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	if timestamp == nil {
		return fmt.Errorf("parameter timestamp is required, but nil was provided")
	}

	return nil
}

func validateFn_FormatlistParameters(spec *string, values *[]interface{}) error {
	if spec == nil {
		return fmt.Errorf("parameter spec is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_IndentParameters(indentation *float64, value *string) error {
	if indentation == nil {
		return fmt.Errorf("parameter indentation is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_IndexParameters(list interface{}, value interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_JoinParameters(separator *string, value *[]*string) error {
	if separator == nil {
		return fmt.Errorf("parameter separator is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_JsondecodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_JsonencodeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_KeysParameters(map_ interface{}) error {
	if map_ == nil {
		return fmt.Errorf("parameter map_ is required, but nil was provided")
	}
	switch map_.(type) {
	case IResolvable:
		// ok
	case *map[string]interface{}:
		// ok
	case map[string]interface{}:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(map_) {
			return fmt.Errorf("parameter map_ must be one of the allowed types: IResolvable, *map[string]interface{}; received %#v (a %T)", map_, map_)
		}
	}

	return nil
}

func validateFn_LengthOfParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_LogParameters(value *float64, base *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if base == nil {
		return fmt.Errorf("parameter base is required, but nil was provided")
	}

	return nil
}

func validateFn_LookupParameters(value interface{}, key interface{}, defaultValue interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if defaultValue == nil {
		return fmt.Errorf("parameter defaultValue is required, but nil was provided")
	}

	return nil
}

func validateFn_LowerParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_MatchkeysParameters(valuesList interface{}, keysList interface{}, searchSet interface{}) error {
	if valuesList == nil {
		return fmt.Errorf("parameter valuesList is required, but nil was provided")
	}
	switch valuesList.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(valuesList) {
			return fmt.Errorf("parameter valuesList must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", valuesList, valuesList)
		}
	}

	if keysList == nil {
		return fmt.Errorf("parameter keysList is required, but nil was provided")
	}
	switch keysList.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(keysList) {
			return fmt.Errorf("parameter keysList must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", keysList, keysList)
		}
	}

	if searchSet == nil {
		return fmt.Errorf("parameter searchSet is required, but nil was provided")
	}
	switch searchSet.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(searchSet) {
			return fmt.Errorf("parameter searchSet must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", searchSet, searchSet)
		}
	}

	return nil
}

func validateFn_MaxParameters(values *[]*float64) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_Md5Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_MergeListsParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_MergeMapsParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_MinParameters(values *[]*float64) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_NonsensitiveParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_OneParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	return nil
}

func validateFn_ParseIntParameters(value *string, base *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if base == nil {
		return fmt.Errorf("parameter base is required, but nil was provided")
	}

	return nil
}

func validateFn_PathexpandParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_PowParameters(value *float64, power *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if power == nil {
		return fmt.Errorf("parameter power is required, but nil was provided")
	}

	return nil
}

func validateFn_RangeParameters(start *float64, limit *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if limit == nil {
		return fmt.Errorf("parameter limit is required, but nil was provided")
	}

	return nil
}

func validateFn_RawStringParameters(str *string) error {
	if str == nil {
		return fmt.Errorf("parameter str is required, but nil was provided")
	}

	return nil
}

func validateFn_RegexParameters(pattern *string, value *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_RegexallParameters(pattern *string, value *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ReplaceParameters(value *string, substring *string, replacement *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if substring == nil {
		return fmt.Errorf("parameter substring is required, but nil was provided")
	}

	if replacement == nil {
		return fmt.Errorf("parameter replacement is required, but nil was provided")
	}

	return nil
}

func validateFn_ReverseParameters(values interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}
	switch values.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(values) {
			return fmt.Errorf("parameter values must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", values, values)
		}
	}

	return nil
}

func validateFn_RsadecryptParameters(ciphertext *string, privatekey *string) error {
	if ciphertext == nil {
		return fmt.Errorf("parameter ciphertext is required, but nil was provided")
	}

	if privatekey == nil {
		return fmt.Errorf("parameter privatekey is required, but nil was provided")
	}

	return nil
}

func validateFn_SensitiveParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_SetintersectionParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_SetproductParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_SetsubtractParameters(minuend interface{}, subtrahend interface{}) error {
	if minuend == nil {
		return fmt.Errorf("parameter minuend is required, but nil was provided")
	}
	switch minuend.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(minuend) {
			return fmt.Errorf("parameter minuend must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", minuend, minuend)
		}
	}

	if subtrahend == nil {
		return fmt.Errorf("parameter subtrahend is required, but nil was provided")
	}
	switch subtrahend.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(subtrahend) {
			return fmt.Errorf("parameter subtrahend must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", subtrahend, subtrahend)
		}
	}

	return nil
}

func validateFn_SetunionParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateFn_Sha1Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Sha256Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Sha512Parameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_SignumParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_SliceParameters(list interface{}, startindex *float64, endindex *float64) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	if startindex == nil {
		return fmt.Errorf("parameter startindex is required, but nil was provided")
	}

	if endindex == nil {
		return fmt.Errorf("parameter endindex is required, but nil was provided")
	}

	return nil
}

func validateFn_SortParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	return nil
}

func validateFn_SplitParameters(seperator *string, value *string) error {
	if seperator == nil {
		return fmt.Errorf("parameter seperator is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_StrrevParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_SubstrParameters(value *string, offset *float64, length *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if offset == nil {
		return fmt.Errorf("parameter offset is required, but nil was provided")
	}

	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

func validateFn_SumParameters(list interface{}) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}
	switch list.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(list) {
			return fmt.Errorf("parameter list must be one of the allowed types: *string, *[]interface{}, IResolvable; received %#v (a %T)", list, list)
		}
	}

	return nil
}

func validateFn_TemplatefileParameters(path *string, vars interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if vars == nil {
		return fmt.Errorf("parameter vars is required, but nil was provided")
	}

	return nil
}

func validateFn_Textdecodebase64Parameters(value *string, encodingName *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if encodingName == nil {
		return fmt.Errorf("parameter encodingName is required, but nil was provided")
	}

	return nil
}

func validateFn_Textencodebase64Parameters(value *string, encodingName *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if encodingName == nil {
		return fmt.Errorf("parameter encodingName is required, but nil was provided")
	}

	return nil
}

func validateFn_TimeaddParameters(timestamp *string, duration *string) error {
	if timestamp == nil {
		return fmt.Errorf("parameter timestamp is required, but nil was provided")
	}

	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func validateFn_TitleParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ToboolParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_TolistParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_TomapParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_TonumberParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_TosetParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_TostringParameters(expression interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_TransposeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_TrimParameters(value *string, replacement *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if replacement == nil {
		return fmt.Errorf("parameter replacement is required, but nil was provided")
	}

	return nil
}

func validateFn_TrimprefixParameters(value *string, prefix *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func validateFn_TrimspaceParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_TrimsuffixParameters(value *string, suffix *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if suffix == nil {
		return fmt.Errorf("parameter suffix is required, but nil was provided")
	}

	return nil
}

func validateFn_TryParameters(expression *[]interface{}) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateFn_UpperParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_UrlencodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_Uuidv5Parameters(namespace *string, name *string) error {
	if namespace == nil {
		return fmt.Errorf("parameter namespace is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateFn_ValuesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_YamldecodeParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_YamlencodeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFn_ZipmapParameters(keyslist interface{}, valueslist interface{}) error {
	if keyslist == nil {
		return fmt.Errorf("parameter keyslist is required, but nil was provided")
	}
	switch keyslist.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(keyslist) {
			return fmt.Errorf("parameter keyslist must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", keyslist, keyslist)
		}
	}

	if valueslist == nil {
		return fmt.Errorf("parameter valueslist is required, but nil was provided")
	}
	switch valueslist.(type) {
	case *[]interface{}:
		// ok
	case []interface{}:
		// ok
	case IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(valueslist) {
			return fmt.Errorf("parameter valueslist must be one of the allowed types: *[]interface{}, IResolvable; received %#v (a %T)", valueslist, valueslist)
		}
	}

	return nil
}

