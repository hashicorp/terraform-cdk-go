//go:build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPostProcessor) validatePostProcessParameters(input interface{}, context IResolveContext) error {
	return nil
}

