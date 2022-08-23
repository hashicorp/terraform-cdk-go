package testingmatchers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdktf.testingMatchers.AssertionReturn",
		reflect.TypeOf((*AssertionReturn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "pass", GoGetter: "Pass"},
		},
		func() interface{} {
			return &jsiiProxy_AssertionReturn{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.testingMatchers.TerraformConstructor",
		reflect.TypeOf((*TerraformConstructor)(nil)).Elem(),
	)
}
