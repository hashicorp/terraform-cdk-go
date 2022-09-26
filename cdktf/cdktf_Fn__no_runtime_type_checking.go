//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func validateFn_AbsParameters(value *float64) error {
	return nil
}

func validateFn_AbspathParameters(value *string) error {
	return nil
}

func validateFn_AlltrueParameters(values *[]interface{}) error {
	return nil
}

func validateFn_AnytrueParameters(value *[]interface{}) error {
	return nil
}

func validateFn_Base64decodeParameters(value *string) error {
	return nil
}

func validateFn_Base64encodeParameters(value *string) error {
	return nil
}

func validateFn_Base64gzipParameters(value *string) error {
	return nil
}

func validateFn_Base64sha256Parameters(value *string) error {
	return nil
}

func validateFn_Base64sha512Parameters(value *string) error {
	return nil
}

func validateFn_BasenameParameters(value *string) error {
	return nil
}

func validateFn_BcryptParameters(value *string) error {
	return nil
}

func validateFn_CanParameters(expression interface{}) error {
	return nil
}

func validateFn_CeilParameters(value *float64) error {
	return nil
}

func validateFn_ChompParameters(value *string) error {
	return nil
}

func validateFn_ChunklistParameters(value *[]interface{}, chunkSize *float64) error {
	return nil
}

func validateFn_CidrhostParameters(prefix *string, hostnum *float64) error {
	return nil
}

func validateFn_CidrnetmaskParameters(prefix *string) error {
	return nil
}

func validateFn_CidrsubnetParameters(prefix *string, newbits *float64, netnum *float64) error {
	return nil
}

func validateFn_CidrsubnetsParameters(prefix *string, newbits *[]*float64) error {
	return nil
}

func validateFn_CoalesceParameters(value *[]interface{}) error {
	return nil
}

func validateFn_CoalescelistParameters(value *[]*[]interface{}) error {
	return nil
}

func validateFn_CompactParameters(value *[]*string) error {
	return nil
}

func validateFn_ConcatParameters(value *[]*[]interface{}) error {
	return nil
}

func validateFn_ContainsParameters(list interface{}, value interface{}) error {
	return nil
}

func validateFn_CsvdecodeParameters(value *string) error {
	return nil
}

func validateFn_DirnameParameters(value *string) error {
	return nil
}

func validateFn_DistinctParameters(list interface{}) error {
	return nil
}

func validateFn_ElementParameters(list interface{}, index *float64) error {
	return nil
}

func validateFn_FileParameters(value *string) error {
	return nil
}

func validateFn_Filebase64Parameters(value *string) error {
	return nil
}

func validateFn_Filebase64sha256Parameters(value *string) error {
	return nil
}

func validateFn_Filebase64sha512Parameters(value *string) error {
	return nil
}

func validateFn_FileexistsParameters(value *string) error {
	return nil
}

func validateFn_Filemd5Parameters(value *string) error {
	return nil
}

func validateFn_FilesetParameters(path *string, pattern *string) error {
	return nil
}

func validateFn_Filesha1Parameters(value *string) error {
	return nil
}

func validateFn_Filesha256Parameters(value *string) error {
	return nil
}

func validateFn_Filesha512Parameters(value *string) error {
	return nil
}

func validateFn_FlattenParameters(list interface{}) error {
	return nil
}

func validateFn_FloorParameters(value *float64) error {
	return nil
}

func validateFn_FormatParameters(spec *string, values *[]interface{}) error {
	return nil
}

func validateFn_FormatdateParameters(spec *string, timestamp *string) error {
	return nil
}

func validateFn_FormatlistParameters(spec *string, values *[]interface{}) error {
	return nil
}

func validateFn_IndentParameters(indentation *float64, value *string) error {
	return nil
}

func validateFn_IndexParameters(list interface{}, value interface{}) error {
	return nil
}

func validateFn_JoinParameters(separator *string, value *[]*string) error {
	return nil
}

func validateFn_JsondecodeParameters(value *string) error {
	return nil
}

func validateFn_JsonencodeParameters(value interface{}) error {
	return nil
}

func validateFn_KeysParameters(map_ interface{}) error {
	return nil
}

func validateFn_LengthOfParameters(value interface{}) error {
	return nil
}

func validateFn_LogParameters(value *float64, base *float64) error {
	return nil
}

func validateFn_LookupParameters(value interface{}, key interface{}, defaultValue interface{}) error {
	return nil
}

func validateFn_LowerParameters(value *string) error {
	return nil
}

func validateFn_MatchkeysParameters(valuesList interface{}, keysList interface{}, searchSet interface{}) error {
	return nil
}

func validateFn_MaxParameters(values *[]*float64) error {
	return nil
}

func validateFn_Md5Parameters(value *string) error {
	return nil
}

func validateFn_MergeListsParameters(values *[]interface{}) error {
	return nil
}

func validateFn_MergeMapsParameters(values *[]interface{}) error {
	return nil
}

func validateFn_MinParameters(values *[]*float64) error {
	return nil
}

func validateFn_NonsensitiveParameters(expression interface{}) error {
	return nil
}

func validateFn_OneParameters(list interface{}) error {
	return nil
}

func validateFn_ParseIntParameters(value *string, base *float64) error {
	return nil
}

func validateFn_PathexpandParameters(value *string) error {
	return nil
}

func validateFn_PowParameters(value *float64, power *float64) error {
	return nil
}

func validateFn_RangeParameters(start *float64, limit *float64) error {
	return nil
}

func validateFn_RawStringParameters(str *string) error {
	return nil
}

func validateFn_RegexParameters(pattern *string, value *string) error {
	return nil
}

func validateFn_RegexallParameters(pattern *string, value *string) error {
	return nil
}

func validateFn_ReplaceParameters(value *string, substring *string, replacement *string) error {
	return nil
}

func validateFn_ReverseParameters(values interface{}) error {
	return nil
}

func validateFn_RsadecryptParameters(ciphertext *string, privatekey *string) error {
	return nil
}

func validateFn_SensitiveParameters(expression interface{}) error {
	return nil
}

func validateFn_SetintersectionParameters(values *[]interface{}) error {
	return nil
}

func validateFn_SetproductParameters(values *[]interface{}) error {
	return nil
}

func validateFn_SetsubtractParameters(minuend interface{}, subtrahend interface{}) error {
	return nil
}

func validateFn_SetunionParameters(values *[]interface{}) error {
	return nil
}

func validateFn_Sha1Parameters(value *string) error {
	return nil
}

func validateFn_Sha256Parameters(value *string) error {
	return nil
}

func validateFn_Sha512Parameters(value *string) error {
	return nil
}

func validateFn_SignumParameters(value *float64) error {
	return nil
}

func validateFn_SliceParameters(list interface{}, startindex *float64, endindex *float64) error {
	return nil
}

func validateFn_SortParameters(list interface{}) error {
	return nil
}

func validateFn_SplitParameters(seperator *string, value *string) error {
	return nil
}

func validateFn_StrrevParameters(value *string) error {
	return nil
}

func validateFn_SubstrParameters(value *string, offset *float64, length *float64) error {
	return nil
}

func validateFn_SumParameters(list interface{}) error {
	return nil
}

func validateFn_TemplatefileParameters(path *string, vars interface{}) error {
	return nil
}

func validateFn_Textdecodebase64Parameters(value *string, encodingName *string) error {
	return nil
}

func validateFn_Textencodebase64Parameters(value *string, encodingName *string) error {
	return nil
}

func validateFn_TimeaddParameters(timestamp *string, duration *string) error {
	return nil
}

func validateFn_TitleParameters(value *string) error {
	return nil
}

func validateFn_ToboolParameters(expression interface{}) error {
	return nil
}

func validateFn_TolistParameters(expression interface{}) error {
	return nil
}

func validateFn_TomapParameters(expression interface{}) error {
	return nil
}

func validateFn_TonumberParameters(expression interface{}) error {
	return nil
}

func validateFn_TosetParameters(expression interface{}) error {
	return nil
}

func validateFn_TostringParameters(expression interface{}) error {
	return nil
}

func validateFn_TransposeParameters(value interface{}) error {
	return nil
}

func validateFn_TrimParameters(value *string, replacement *string) error {
	return nil
}

func validateFn_TrimprefixParameters(value *string, prefix *string) error {
	return nil
}

func validateFn_TrimspaceParameters(value *string) error {
	return nil
}

func validateFn_TrimsuffixParameters(value *string, suffix *string) error {
	return nil
}

func validateFn_TryParameters(expression *[]interface{}) error {
	return nil
}

func validateFn_UpperParameters(value *string) error {
	return nil
}

func validateFn_UrlencodeParameters(value *string) error {
	return nil
}

func validateFn_Uuidv5Parameters(namespace *string, name *string) error {
	return nil
}

func validateFn_ValuesParameters(value interface{}) error {
	return nil
}

func validateFn_YamldecodeParameters(value *string) error {
	return nil
}

func validateFn_YamlencodeParameters(value interface{}) error {
	return nil
}

func validateFn_ZipmapParameters(keyslist interface{}, valueslist interface{}) error {
	return nil
}

