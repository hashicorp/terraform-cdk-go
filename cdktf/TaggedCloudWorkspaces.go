// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"
)

// A set of Terraform Cloud workspace tags.
//
// You will be able to use this working directory with any workspaces that have all of the specified tags, and can use the terraform workspace commands to switch between them or create new workspaces. New workspaces will automatically have the specified tags. This option conflicts with name.
// Experimental.
type TaggedCloudWorkspaces interface {
	CloudWorkspace
	// Experimental.
	Tags() *[]*string
	// Experimental.
	SetTags(val *[]*string)
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TaggedCloudWorkspaces
type jsiiProxy_TaggedCloudWorkspaces struct {
	jsiiProxy_CloudWorkspace
}

func (j *jsiiProxy_TaggedCloudWorkspaces) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}


// Experimental.
func NewTaggedCloudWorkspaces(tags *[]*string) TaggedCloudWorkspaces {
	_init_.Initialize()

	if err := validateNewTaggedCloudWorkspacesParameters(tags); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaggedCloudWorkspaces{}

	_jsii_.Create(
		"cdktf.TaggedCloudWorkspaces",
		[]interface{}{tags},
		&j,
	)

	return &j
}

// Experimental.
func NewTaggedCloudWorkspaces_Override(t TaggedCloudWorkspaces, tags *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TaggedCloudWorkspaces",
		[]interface{}{tags},
		t,
	)
}

func (j *jsiiProxy_TaggedCloudWorkspaces)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (t *jsiiProxy_TaggedCloudWorkspaces) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

