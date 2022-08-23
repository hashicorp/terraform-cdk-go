// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type DataTerraformRemoteStateAzurermConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// Experimental.
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
	// (Required) The Name of the Storage Container within the Storage Account.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// (Required) The name of the Blob used to retrieve/store Terraform's State file inside the Storage Container.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// (Required) The Name of the Storage Account.
	// Experimental.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// access_key - (Optional) The Access Key used to access the Blob Storage Account.
	//
	// This can also be sourced from the ARM_ACCESS_KEY environment variable.
	// Experimental.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// (Optional) The Client ID of the Service Principal.
	//
	// This can also be sourced from the ARM_CLIENT_ID environment variable.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// (Optional) The Client Secret of the Service Principal.
	//
	// This can also be sourced from the ARM_CLIENT_SECRET environment variable.
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// (Optional) The Custom Endpoint for Azure Resource Manager. This can also be sourced from the ARM_ENDPOINT environment variable.
	//
	// NOTE: An endpoint should only be configured when using Azure Stack.
	// Experimental.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// (Optional) The Azure Environment which should be used.
	//
	// This can also be sourced from the ARM_ENVIRONMENT environment variable.
	//   Possible values are public, china, german, stack and usgovernment. Defaults to public.
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// (Optional) The path to a custom Managed Service Identity endpoint which is automatically determined if not specified.
	//
	// This can also be sourced from the ARM_MSI_ENDPOINT environment variable.
	// Experimental.
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// (Required) The Name of the Resource Group in which the Storage Account exists.
	// Experimental.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// (Optional) The SAS Token used to access the Blob Storage Account.
	//
	// This can also be sourced from the ARM_SAS_TOKEN environment variable.
	// Experimental.
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// (Optional) The Subscription ID in which the Storage Account exists.
	//
	// This can also be sourced from the ARM_SUBSCRIPTION_ID environment variable.
	// Experimental.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// (Optional) The Tenant ID in which the Subscription exists.
	//
	// This can also be sourced from the ARM_TENANT_ID environment variable.
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// (Optional) Should Managed Service Identity authentication be used?
	//
	// This can also be sourced from the ARM_USE_MSI environment variable.
	// Experimental.
	UseMsi *bool `field:"optional" json:"useMsi" yaml:"useMsi"`
}

