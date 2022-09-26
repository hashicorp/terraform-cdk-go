//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Cloud Development Kit for Terraform
package cdktf

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LazyBase) validateAddPostProcessorParameters(postProcessor IPostProcessor) error {
	return nil
}

func (l *jsiiProxy_LazyBase) validateResolveParameters(context IResolveContext) error {
	return nil
}

func (l *jsiiProxy_LazyBase) validateResolveLazyParameters(context IResolveContext) error {
	return nil
}

