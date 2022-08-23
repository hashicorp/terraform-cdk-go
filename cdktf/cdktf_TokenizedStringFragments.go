// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// Fragments of a concatenated string containing stringified Tokens.
// Experimental.
type TokenizedStringFragments interface {
	// Returns the first token.
	// Experimental.
	FirstToken() IResolvable
	// Returns the first value.
	// Experimental.
	FirstValue() interface{}
	// Return all intrinsic fragments from this string.
	// Experimental.
	Intrinsic() *[]IResolvable
	// Returns the number of fragments.
	// Experimental.
	Length() *float64
	// Return all literals from this string.
	// Experimental.
	Literals() *[]IResolvable
	// Return all Tokens from this string.
	// Experimental.
	Tokens() *[]IResolvable
	// Adds an intrinsic fragment.
	// Experimental.
	AddIntrinsic(value interface{})
	// Adds a literal fragment.
	// Experimental.
	AddLiteral(lit interface{})
	// Adds a token fragment.
	// Experimental.
	AddToken(token IResolvable)
	// Combine the string fragments using the given joiner.
	//
	// If there are any.
	// Experimental.
	Join(concat IFragmentConcatenator) interface{}
	// Apply a transformation function to all tokens in the string.
	// Experimental.
	MapTokens(mapper ITokenMapper) TokenizedStringFragments
}

// The jsii proxy struct for TokenizedStringFragments
type jsiiProxy_TokenizedStringFragments struct {
	_ byte // padding
}

func (j *jsiiProxy_TokenizedStringFragments) FirstToken() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"firstToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) FirstValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Intrinsic() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"intrinsic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Literals() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"literals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Tokens() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


// Experimental.
func NewTokenizedStringFragments() TokenizedStringFragments {
	_init_.Initialize()

	j := jsiiProxy_TokenizedStringFragments{}

	_jsii_.Create(
		"cdktf.TokenizedStringFragments",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTokenizedStringFragments_Override(t TokenizedStringFragments) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TokenizedStringFragments",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddIntrinsic(value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addIntrinsic",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddLiteral(lit interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addLiteral",
		[]interface{}{lit},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) AddToken(token IResolvable) {
	_jsii_.InvokeVoid(
		t,
		"addToken",
		[]interface{}{token},
	)
}

func (t *jsiiProxy_TokenizedStringFragments) Join(concat IFragmentConcatenator) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"join",
		[]interface{}{concat},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenizedStringFragments) MapTokens(mapper ITokenMapper) TokenizedStringFragments {
	var returns TokenizedStringFragments

	_jsii_.Invoke(
		t,
		"mapTokens",
		[]interface{}{mapper},
		&returns,
	)

	return returns
}

