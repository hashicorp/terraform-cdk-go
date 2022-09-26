//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func validateResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewResourceParameters(scope constructs.Construct, id *string) error {
	return nil
}

