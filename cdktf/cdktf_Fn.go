// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// Experimental.
type Fn interface {
}

// The jsii proxy struct for Fn
type jsiiProxy_Fn struct {
	_ byte // padding
}

// Experimental.
func NewFn() Fn {
	_init_.Initialize()

	j := jsiiProxy_Fn{}

	_jsii_.Create(
		"cdktf.Fn",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFn_Override(f Fn) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Fn",
		nil, // no parameters
		f,
	)
}

// {@link https://www.terraform.io/docs/language/functions/abs.html abs} returns the absolute value of the given number.
// Experimental.
func Fn_Abs(value *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_AbsParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"abs",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/abspath.html abspath} takes a string containing a filesystem path and converts it to an absolute path.
// Experimental.
func Fn_Abspath(value *string) *string {
	_init_.Initialize()

	if err := validateFn_AbspathParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"abspath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/alltrue.html alltrue} returns true if all elements in a given collection are true or "true".
// Experimental.
func Fn_Alltrue(values *[]interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_AlltrueParameters(values); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"alltrue",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/anytrue.html anytrue} returns true if any element in a given collection is true or "true".
// Experimental.
func Fn_Anytrue(value *[]interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_AnytrueParameters(value); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"anytrue",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64decode.html base64decode} takes a string containing a Base64 character sequence and returns the original string.
// Experimental.
func Fn_Base64decode(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Base64decodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64decode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64encode.html base64encode} takes a string containing a Base64 character sequence and returns the original string.
// Experimental.
func Fn_Base64encode(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Base64encodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64encode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64gzip.html base64gzip} compresses a string with gzip and then encodes the result in Base64 encoding.
// Experimental.
func Fn_Base64gzip(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Base64gzipParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64gzip",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64sha256.html base64sha256} computes the SHA256 hash of a given string and encodes it with Base64.
// Experimental.
func Fn_Base64sha256(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Base64sha256Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64sha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/base64sha512.html base64sha512} computes the SHA512 hash of a given string and encodes it with Base64.
// Experimental.
func Fn_Base64sha512(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Base64sha512Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"base64sha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/basename.html basename} takes a string containing a filesystem path and removes all except the last portion from it.
// Experimental.
func Fn_Basename(value *string) *string {
	_init_.Initialize()

	if err := validateFn_BasenameParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"basename",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/bcrypt.html bcrypt} computes a hash of the given string using the Blowfish cipher, returning a string in the Modular Crypt Format usually expected in the shadow password file on many Unix systems.
// Experimental.
func Fn_Bcrypt(value *string, cost *float64) *string {
	_init_.Initialize()

	if err := validateFn_BcryptParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"bcrypt",
		[]interface{}{value, cost},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/can.html can} evaluates the given expression and returns a boolean value indicating whether the expression produced a result without any errors.
// Experimental.
func Fn_Can(expression interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_CanParameters(expression); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"can",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/ceil.html ceil} returns the closest whole number that is greater than or equal to the given value, which may be a fraction.
// Experimental.
func Fn_Ceil(value *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_CeilParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"ceil",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/chomp.html chomp} removes newline characters at the end of a string.
// Experimental.
func Fn_Chomp(value *string) *string {
	_init_.Initialize()

	if err := validateFn_ChompParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"chomp",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/chunklist.html chunklist} splits a single list into fixed-size chunks, returning a list of lists.
// Experimental.
func Fn_Chunklist(value *[]interface{}, chunkSize *float64) *[]*string {
	_init_.Initialize()

	if err := validateFn_ChunklistParameters(value, chunkSize); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"chunklist",
		[]interface{}{value, chunkSize},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrhost.html cidrhost} calculates a full host IP address for a given host number within a given IP network address prefix.
// Experimental.
func Fn_Cidrhost(prefix *string, hostnum *float64) *string {
	_init_.Initialize()

	if err := validateFn_CidrhostParameters(prefix, hostnum); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrhost",
		[]interface{}{prefix, hostnum},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrnetmask.html cidrnetmask} converts an IPv4 address prefix given in CIDR notation into a subnet mask address.
// Experimental.
func Fn_Cidrnetmask(prefix *string) *string {
	_init_.Initialize()

	if err := validateFn_CidrnetmaskParameters(prefix); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrnetmask",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrsubnet.html cidrsubnet} calculates a subnet address within given IP network address prefix.
// Experimental.
func Fn_Cidrsubnet(prefix *string, newbits *float64, netnum *float64) *string {
	_init_.Initialize()

	if err := validateFn_CidrsubnetParameters(prefix, newbits, netnum); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrsubnet",
		[]interface{}{prefix, newbits, netnum},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/cidrsubnets.html cidrsubnets} calculates a sequence of consecutive IP address ranges within a particular CIDR prefix.
// Experimental.
func Fn_Cidrsubnets(prefix *string, newbits *[]*float64) *[]*string {
	_init_.Initialize()

	if err := validateFn_CidrsubnetsParameters(prefix, newbits); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"cidrsubnets",
		[]interface{}{prefix, newbits},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/coalesce.html coalesce} takes any number of arguments and returns the first one that isn't null or an empty string.
// Experimental.
func Fn_Coalesce(value *[]interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_CoalesceParameters(value); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"coalesce",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/coalescelist.html coalescelist} takes any number of list arguments and returns the first one that isn't empty.
// Experimental.
func Fn_Coalescelist(value *[]*[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_CoalescelistParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"coalescelist",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/compact.html compact} takes a list of strings and returns a new list with any empty string elements removed.
// Experimental.
func Fn_Compact(value *[]*string) *[]*string {
	_init_.Initialize()

	if err := validateFn_CompactParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"compact",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/concat.html concat} takes two or more lists and combines them into a single list.
// Experimental.
func Fn_Concat(value *[]*[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_ConcatParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"concat",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/contains.html contains} determines whether a given list or set contains a given single value as one of its elements.
// Experimental.
func Fn_Contains(list interface{}, value interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_ContainsParameters(list, value); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"contains",
		[]interface{}{list, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/csvdecode.html csvdecode} decodes a string containing CSV-formatted data and produces a list of maps representing that data.
// Experimental.
func Fn_Csvdecode(value *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_CsvdecodeParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"csvdecode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/dirname.html dirname} takes a string containing a filesystem path and removes the last portion from it.
// Experimental.
func Fn_Dirname(value *string) *string {
	_init_.Initialize()

	if err := validateFn_DirnameParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"dirname",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/distinct.html distinct} takes a list and returns a new list with any duplicate elements removed.
// Experimental.
func Fn_Distinct(list interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_DistinctParameters(list); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"distinct",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/element.html element} retrieves a single element from a list.
// Experimental.
func Fn_Element(list interface{}, index *float64) interface{} {
	_init_.Initialize()

	if err := validateFn_ElementParameters(list, index); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"element",
		[]interface{}{list, index},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/file.html file} takes a string containing a filesystem path and removes all except the last portion from it.
// Experimental.
func Fn_File(value *string) *string {
	_init_.Initialize()

	if err := validateFn_FileParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"file",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filebase64.html filebase64} reads the contents of a file at the given path and returns them as a base64-encoded string.
// Experimental.
func Fn_Filebase64(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filebase64Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filebase64",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filebase64sha256.html filebase64sha256} is a variant of base64sha256 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filebase64sha256(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filebase64sha256Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filebase64sha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filebase64sha512.html filebase64sha512} is a variant of base64sha512 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filebase64sha512(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filebase64sha512Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filebase64sha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/fileexists.html fileexists} determines whether a file exists at a given path.
// Experimental.
func Fn_Fileexists(value *string) IResolvable {
	_init_.Initialize()

	if err := validateFn_FileexistsParameters(value); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"fileexists",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filemd5.html filemd5} is a variant of md5 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filemd5(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filemd5Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filemd5",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/fileset.html fileset} enumerates a set of regular file names given a path and pattern.
// Experimental.
func Fn_Fileset(path *string, pattern *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_FilesetParameters(path, pattern); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"fileset",
		[]interface{}{path, pattern},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filesha1.html filesha1} is a variant of sha1 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filesha1(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filesha1Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filesha1",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filesha256.html filesha256} is a variant of sha256 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filesha256(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filesha256Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filesha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/filesha512.html filesha512} is a variant of sha512 that hashes the contents of a given file rather than a literal string.
// Experimental.
func Fn_Filesha512(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Filesha512Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"filesha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/flatten.html flatten} takes a list and replaces any elements that are lists with a flattened sequence of the list contents.
// Experimental.
func Fn_Flatten(list interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_FlattenParameters(list); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"flatten",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/floor.html floor} returns the closest whole number that is less than or equal to the given value, which may be a fraction.
// Experimental.
func Fn_Floor(value *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_FloorParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"floor",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/format.html format} produces a string by formatting a number of other values according to a specification string.
// Experimental.
func Fn_Format(spec *string, values *[]interface{}) *string {
	_init_.Initialize()

	if err := validateFn_FormatParameters(spec, values); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"format",
		[]interface{}{spec, values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/formatdate.html formatdate} converts a timestamp into a different time format.
// Experimental.
func Fn_Formatdate(spec *string, timestamp *string) *string {
	_init_.Initialize()

	if err := validateFn_FormatdateParameters(spec, timestamp); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"formatdate",
		[]interface{}{spec, timestamp},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/formatlist.html formatlist} produces a list of strings by formatting a number of other values according to a specification string.
// Experimental.
func Fn_Formatlist(spec *string, values *[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_FormatlistParameters(spec, values); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"formatlist",
		[]interface{}{spec, values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/indent.html indent} adds a given number of spaces to the beginnings of all but the first line in a given multi-line string.
// Experimental.
func Fn_Indent(indentation *float64, value *string) *string {
	_init_.Initialize()

	if err := validateFn_IndentParameters(indentation, value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"indent",
		[]interface{}{indentation, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/index.html index} finds the element index for a given value in a list.
// Experimental.
func Fn_Index(list interface{}, value interface{}) *float64 {
	_init_.Initialize()

	if err := validateFn_IndexParameters(list, value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"index",
		[]interface{}{list, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/join.html join} produces a string by concatenating together all elements of a given list of strings with the given delimiter.
// Experimental.
func Fn_Join(separator *string, value *[]*string) *string {
	_init_.Initialize()

	if err := validateFn_JoinParameters(separator, value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"join",
		[]interface{}{separator, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/jsondecode.html jsondecode} interprets a given string as JSON, returning a representation of the result of decoding that string.
// Experimental.
func Fn_Jsondecode(value *string) interface{} {
	_init_.Initialize()

	if err := validateFn_JsondecodeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"jsondecode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/jsonencode.html jsonencode} encodes a given value to a string using JSON syntax.
// Experimental.
func Fn_Jsonencode(value interface{}) *string {
	_init_.Initialize()

	if err := validateFn_JsonencodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"jsonencode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/keys.html keys} takes a map and returns a list containing the keys from that map.
// Experimental.
func Fn_Keys(map_ interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_KeysParameters(map_); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"keys",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/length.html length} determines the length of a given list, map, or string.
// Experimental.
func Fn_LengthOf(value interface{}) *float64 {
	_init_.Initialize()

	if err := validateFn_LengthOfParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"lengthOf",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/log.html log} returns the logarithm of a given number in a given base.
// Experimental.
func Fn_Log(value *float64, base *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_LogParameters(value, base); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"log",
		[]interface{}{value, base},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/lookup.html lookup} retrieves the value of a single element from a map, given its key. If the given key does not exist, the given default value is returned instead.
// Experimental.
func Fn_Lookup(value interface{}, key interface{}, defaultValue interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_LookupParameters(value, key, defaultValue); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"lookup",
		[]interface{}{value, key, defaultValue},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/lower.html lower} converts all cased letters in the given string to lowercase.
// Experimental.
func Fn_Lower(value *string) *string {
	_init_.Initialize()

	if err := validateFn_LowerParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"lower",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/matchkeys.html matchkeys} constructs a new list by taking a subset of elements from one list whose indexes match the corresponding indexes of values in another list.
// Experimental.
func Fn_Matchkeys(valuesList interface{}, keysList interface{}, searchSet interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_MatchkeysParameters(valuesList, keysList, searchSet); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"matchkeys",
		[]interface{}{valuesList, keysList, searchSet},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/max.html max} takes one or more numbers and returns the greatest number from the set.
// Experimental.
func Fn_Max(values *[]*float64) *float64 {
	_init_.Initialize()

	if err := validateFn_MaxParameters(values); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"max",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/md5.html md5} computes the MD5 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Md5(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Md5Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"md5",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/merge.html merge} takes an arbitrary number of maps or objects, and returns a single map or object that contains a merged set of elements from all arguments.
// Experimental.
func Fn_MergeLists(values *[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_MergeListsParameters(values); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"mergeLists",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/merge.html merge} takes an arbitrary number of maps or objects, and returns a single map or object that contains a merged set of elements from all arguments.
// Experimental.
func Fn_MergeMaps(values *[]interface{}) *map[string]*string {
	_init_.Initialize()

	if err := validateFn_MergeMapsParameters(values); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"mergeMaps",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/min.html min} takes one or more numbers and returns the smallest number from the set.
// Experimental.
func Fn_Min(values *[]*float64) *float64 {
	_init_.Initialize()

	if err := validateFn_MinParameters(values); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"min",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/nonsensitive.html nonsensitive} takes a sensitive value and returns a copy of that value with the sensitive marking removed, thereby exposing the sensitive value.
// Experimental.
func Fn_Nonsensitive(expression interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_NonsensitiveParameters(expression); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"nonsensitive",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/one.html one} takes a list, set, or tuple value with either zero or one elements.
// Experimental.
func Fn_One(list interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_OneParameters(list); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"one",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/parseint.html parseInt} parses the given string as a representation of an integer in the specified base and returns the resulting number. The base must be between 2 and 62 inclusive.
// Experimental.
func Fn_ParseInt(value *string, base *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_ParseIntParameters(value, base); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"parseInt",
		[]interface{}{value, base},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/pathexpand.html pathexpand} takes a string containing a filesystem path and removes the last portion from it.
// Experimental.
func Fn_Pathexpand(value *string) *string {
	_init_.Initialize()

	if err := validateFn_PathexpandParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"pathexpand",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/pow.html pow} calculates an exponent, by raising its first argument to the power of the second argument.
// Experimental.
func Fn_Pow(value *float64, power *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_PowParameters(value, power); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"pow",
		[]interface{}{value, power},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/range.html range} generates a list of numbers using a start value, a limit value, and a step value.
// Experimental.
func Fn_Range(start *float64, limit *float64, step *float64) *[]*string {
	_init_.Initialize()

	if err := validateFn_RangeParameters(start, limit); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"range",
		[]interface{}{start, limit, step},
		&returns,
	)

	return returns
}

// Use this function to wrap a string and escape it properly for the use in Terraform This is only needed in certain scenarios (e.g., if you have unescaped double quotes in the string).
// Experimental.
func Fn_RawString(str *string) *string {
	_init_.Initialize()

	if err := validateFn_RawStringParameters(str); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"rawString",
		[]interface{}{str},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/regex.html regex} applies a regular expression to a string and returns the matching substrings in pattern.
// Experimental.
func Fn_Regex(pattern *string, value *string) *string {
	_init_.Initialize()

	if err := validateFn_RegexParameters(pattern, value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"regex",
		[]interface{}{pattern, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/regexall.html regexall} applies a regular expression to a string and returns a list of all matches.
// Experimental.
func Fn_Regexall(pattern *string, value *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_RegexallParameters(pattern, value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"regexall",
		[]interface{}{pattern, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/replace.html replace} searches a given string for another given substring, and replaces each occurrence with a given replacement string.
// Experimental.
func Fn_Replace(value *string, substring *string, replacement *string) *string {
	_init_.Initialize()

	if err := validateFn_ReplaceParameters(value, substring, replacement); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"replace",
		[]interface{}{value, substring, replacement},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/reverse.html reverse} takes a sequence and produces a new sequence of the same length with all of the same elements as the given sequence but in reverse order.
// Experimental.
func Fn_Reverse(values interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_ReverseParameters(values); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"reverse",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/rsadecrypt.html rsadecrypt} decrypts an RSA-encrypted ciphertext, returning the corresponding cleartext.
// Experimental.
func Fn_Rsadecrypt(ciphertext *string, privatekey *string) *string {
	_init_.Initialize()

	if err := validateFn_RsadecryptParameters(ciphertext, privatekey); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"rsadecrypt",
		[]interface{}{ciphertext, privatekey},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sensitive.html sensitive} takes any value and returns a copy of it marked so that Terraform will treat it as sensitive, with the same meaning and behavior as for sensitive input variables.
// Experimental.
func Fn_Sensitive(expression interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_SensitiveParameters(expression); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sensitive",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/setintersection.html setintersection} function takes multiple sets and produces a single set containing only the elements that all of the given sets have in common.
// Experimental.
func Fn_Setintersection(values *[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_SetintersectionParameters(values); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setintersection",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/setproduct.html setproduct} function finds all of the possible combinations of elements from all of the given sets by computing the Cartesian product.
// Experimental.
func Fn_Setproduct(values *[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_SetproductParameters(values); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setproduct",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/slice.html setsubtract} function returns a new set containing the elements from the first set that are not present in the second set.
// Experimental.
func Fn_Setsubtract(minuend interface{}, subtrahend interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_SetsubtractParameters(minuend, subtrahend); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setsubtract",
		[]interface{}{minuend, subtrahend},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/setunion.html setunion} function takes multiple sets and produces a single set containing the elements from all of the given sets.
// Experimental.
func Fn_Setunion(values *[]interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_SetunionParameters(values); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"setunion",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sha1.html sha1} computes the SHA1 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Sha1(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Sha1Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sha1",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sha256.html sha256} computes the SHA256 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Sha256(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Sha256Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sha256",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sha512.html sha512} computes the SHA512 hash of a given string and encodes it with hexadecimal digits.
// Experimental.
func Fn_Sha512(value *string) *string {
	_init_.Initialize()

	if err := validateFn_Sha512Parameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sha512",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/signum.html signum} determines the sign of a number, returning a number between -1 and 1 to represent the sign.
// Experimental.
func Fn_Signum(value *float64) *float64 {
	_init_.Initialize()

	if err := validateFn_SignumParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"signum",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/slice.html slice} extracts some consecutive elements from within a list.
// Experimental.
func Fn_Slice(list interface{}, startindex *float64, endindex *float64) *[]*string {
	_init_.Initialize()

	if err := validateFn_SliceParameters(list, startindex, endindex); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"slice",
		[]interface{}{list, startindex, endindex},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sort.html sort} takes a list of strings and returns a new list with those strings sorted lexicographically.
// Experimental.
func Fn_Sort(list interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_SortParameters(list); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sort",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/split.html split} produces a list by dividing a given string at all occurrences of a given separator.
// Experimental.
func Fn_Split(seperator *string, value *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_SplitParameters(seperator, value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"split",
		[]interface{}{seperator, value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/strrev.html strrev} reverses the characters in a string.
// Experimental.
func Fn_Strrev(value *string) *string {
	_init_.Initialize()

	if err := validateFn_StrrevParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"strrev",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/substr.html substr} extracts a substring from a given string by offset and length.
// Experimental.
func Fn_Substr(value *string, offset *float64, length *float64) *string {
	_init_.Initialize()

	if err := validateFn_SubstrParameters(value, offset, length); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"substr",
		[]interface{}{value, offset, length},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/sum.html sum} takes a list or set of numbers and returns the sum of those numbers.
// Experimental.
func Fn_Sum(list interface{}) *float64 {
	_init_.Initialize()

	if err := validateFn_SumParameters(list); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"sum",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/templatefile.html templatefile} reads the file at the given path and renders its content as a template using a supplied set of template variables.
// Experimental.
func Fn_Templatefile(path *string, vars interface{}) *string {
	_init_.Initialize()

	if err := validateFn_TemplatefileParameters(path, vars); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"templatefile",
		[]interface{}{path, vars},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/textdecodebase64.html textdecodebase64} function decodes a string that was previously Base64-encoded, and then interprets the result as characters in a specified character encoding.
// Experimental.
func Fn_Textdecodebase64(value *string, encodingName *string) *string {
	_init_.Initialize()

	if err := validateFn_Textdecodebase64Parameters(value, encodingName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"textdecodebase64",
		[]interface{}{value, encodingName},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/textencodebase64.html textencodebase64}  encodes the unicode characters in a given string using a specified character encoding, returning the result base64 encoded because Terraform language strings are always sequences of unicode characters.
// Experimental.
func Fn_Textencodebase64(value *string, encodingName *string) *string {
	_init_.Initialize()

	if err := validateFn_Textencodebase64Parameters(value, encodingName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"textencodebase64",
		[]interface{}{value, encodingName},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/timeadd.html timeadd} adds a duration to a timestamp, returning a new timestamp.
// Experimental.
func Fn_Timeadd(timestamp *string, duration *string) *string {
	_init_.Initialize()

	if err := validateFn_TimeaddParameters(timestamp, duration); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"timeadd",
		[]interface{}{timestamp, duration},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/timestamp.html timestamp} returns a UTC timestamp string in RFC 3339 format.
// Experimental.
func Fn_Timestamp() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"timestamp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/title.html title} converts the first letter of each word in the given string to uppercase.
// Experimental.
func Fn_Title(value *string) *string {
	_init_.Initialize()

	if err := validateFn_TitleParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"title",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tobool.html tobool} converts its argument to a boolean value.
// Experimental.
func Fn_Tobool(expression interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_ToboolParameters(expression); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tobool",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tolist.html tolist} converts its argument to a list value.
// Experimental.
func Fn_Tolist(expression interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_TolistParameters(expression); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tolist",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tomap.html tomap} converts its argument to a map value.
// Experimental.
func Fn_Tomap(expression interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_TomapParameters(expression); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tomap",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tonumber.html tonumber} converts its argument to a number value.
// Experimental.
func Fn_Tonumber(expression interface{}) *float64 {
	_init_.Initialize()

	if err := validateFn_TonumberParameters(expression); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tonumber",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/toset.html toset} converts its argument to a set value.
// Experimental.
func Fn_Toset(expression interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_TosetParameters(expression); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"toset",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/tostring.html tostring} converts its argument to a string value.
// Experimental.
func Fn_Tostring(expression interface{}) *string {
	_init_.Initialize()

	if err := validateFn_TostringParameters(expression); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"tostring",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/transpose.html transpose} takes a map of lists of strings and swaps the keys and values to produce a new map of lists of strings.
// Experimental.
func Fn_Transpose(value interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_TransposeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"transpose",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trim.html trim} removes the specified characters from the start and end of the given string.
// Experimental.
func Fn_Trim(value *string, replacement *string) *string {
	_init_.Initialize()

	if err := validateFn_TrimParameters(value, replacement); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trim",
		[]interface{}{value, replacement},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trimprefix.html trimprefix} removes the specified prefix from the start of the given string.
// Experimental.
func Fn_Trimprefix(value *string, prefix *string) *string {
	_init_.Initialize()

	if err := validateFn_TrimprefixParameters(value, prefix); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trimprefix",
		[]interface{}{value, prefix},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trimspace.html trimspace} removes any space characters from the start and end of the given string.
// Experimental.
func Fn_Trimspace(value *string) *string {
	_init_.Initialize()

	if err := validateFn_TrimspaceParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trimspace",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/trimsuffix.html trimsuffix} removes the specified suffix from the end of the given string.
// Experimental.
func Fn_Trimsuffix(value *string, suffix *string) *string {
	_init_.Initialize()

	if err := validateFn_TrimsuffixParameters(value, suffix); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"trimsuffix",
		[]interface{}{value, suffix},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/try.html try} evaluates all of its argument expressions in turn and returns the result of the first one that does not produce any errors.
// Experimental.
func Fn_Try(expression *[]interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_TryParameters(expression); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"try",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/upper.html upper} converts all cased letters in the given string to uppercase.
// Experimental.
func Fn_Upper(value *string) *string {
	_init_.Initialize()

	if err := validateFn_UpperParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"upper",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/urlencode.html urlencode} applies URL encoding to a given string.
// Experimental.
func Fn_Urlencode(value *string) *string {
	_init_.Initialize()

	if err := validateFn_UrlencodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"urlencode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/uuid.html uuid} generates a unique identifier string.
// Experimental.
func Fn_Uuid() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"uuid",
		nil, // no parameters
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/uuidv5.html uuidv5} generates a unique identifier string.
// Experimental.
func Fn_Uuidv5(namespace *string, name *string) *string {
	_init_.Initialize()

	if err := validateFn_Uuidv5Parameters(namespace, name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"uuidv5",
		[]interface{}{namespace, name},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/values.html values} takes a map and returns a list containing the values of the elements in that map.
// Experimental.
func Fn_Values(value interface{}) *[]*string {
	_init_.Initialize()

	if err := validateFn_ValuesParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"values",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/yamldecode.html yamldecode} parses a string as a subset of YAML, and produces a representation of its value.
// Experimental.
func Fn_Yamldecode(value *string) interface{} {
	_init_.Initialize()

	if err := validateFn_YamldecodeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"yamldecode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/yamlencode.html yamlencode} encodes a given value to a string using JSON syntax.
// Experimental.
func Fn_Yamlencode(value interface{}) *string {
	_init_.Initialize()

	if err := validateFn_YamlencodeParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"yamlencode",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// {@link https://www.terraform.io/docs/language/functions/zipmap.html zipmap} constructs a map from a list of keys and a corresponding list of values.
// Experimental.
func Fn_Zipmap(keyslist interface{}, valueslist interface{}) interface{} {
	_init_.Initialize()

	if err := validateFn_ZipmapParameters(keyslist, valueslist); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Fn",
		"zipmap",
		[]interface{}{keyslist, valueslist},
		&returns,
	)

	return returns
}

