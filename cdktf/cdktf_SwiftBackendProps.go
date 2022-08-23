// Cloud Development Kit for Terraform
package cdktf


// Experimental.
type SwiftBackendProps struct {
	// Experimental.
	Container *string `field:"required" json:"container" yaml:"container"`
	// Experimental.
	ApplicationCredentialId *string `field:"optional" json:"applicationCredentialId" yaml:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `field:"optional" json:"applicationCredentialName" yaml:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `field:"optional" json:"applicationCredentialSecret" yaml:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `field:"optional" json:"archiveContainer" yaml:"archiveContainer"`
	// Experimental.
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// Experimental.
	CacertFile *string `field:"optional" json:"cacertFile" yaml:"cacertFile"`
	// Experimental.
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// Experimental.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// Experimental.
	DefaultDomain *string `field:"optional" json:"defaultDomain" yaml:"defaultDomain"`
	// Experimental.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Experimental.
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
	// Experimental.
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	ProjectDomainId *string `field:"optional" json:"projectDomainId" yaml:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `field:"optional" json:"projectDomainName" yaml:"projectDomainName"`
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Experimental.
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Experimental.
	TenantName *string `field:"optional" json:"tenantName" yaml:"tenantName"`
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Experimental.
	UserDomainId *string `field:"optional" json:"userDomainId" yaml:"userDomainId"`
	// Experimental.
	UserDomainName *string `field:"optional" json:"userDomainName" yaml:"userDomainName"`
	// Experimental.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Experimental.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

