// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Stores the state in the Consul KV store at a given path. This backend supports state locking.
//
// Read more about this backend in the Terraform docs:
// https://developer.hashicorp.com/terraform/language/settings/backends/consul
// Experimental.
type ConsulBackendConfig struct {
	// (Required) Access token.
	// Experimental.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// (Required) Path in the Consul KV store.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// (Optional) DNS name and port of your Consul endpoint specified in the format dnsname:port.
	//
	// Defaults to the local agent HTTP listener.
	// Experimental.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// (Optional) A path to a PEM-encoded certificate authority used to verify the remote agent's certificate.
	// Experimental.
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// (Optional) A path to a PEM-encoded certificate provided to the remote agent;
	//
	// requires use of key_file.
	// Experimental.
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// (Optional) The datacenter to use.
	//
	// Defaults to that of the agent.
	// Experimental.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// (Optional) true to compress the state data using gzip, or false (the default) to leave it uncompressed.
	// Experimental.
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// (Optional) HTTP Basic Authentication credentials to be used when communicating with Consul, in the format of either user or user:pass.
	// Experimental.
	HttpAuth *string `field:"optional" json:"httpAuth" yaml:"httpAuth"`
	// (Optional) A path to a PEM-encoded private key, required if cert_file is specified.
	// Experimental.
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// (Optional) false to disable locking.
	//
	// This defaults to true, but will require session permissions with Consul and
	// at least kv write permissions on $path/.lock to perform locking.
	// Experimental.
	Lock *bool `field:"optional" json:"lock" yaml:"lock"`
	// (Optional) Specifies what protocol to use when talking to the given address,either http or https.
	//
	// SSL support can also be triggered by setting then environment variable CONSUL_HTTP_SSL to true.
	// Experimental.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

