//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_TaggedCloudWorkspaces) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func validateNewTaggedCloudWorkspacesParameters(tags *[]*string) error {
	return nil
}

