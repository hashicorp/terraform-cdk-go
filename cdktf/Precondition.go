// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Cloud Development Kit for Terraform
package cdktf


// Terraform checks a precondition before evaluating the object it is associated with.
// Experimental.
type Precondition struct {
	// This is a boolean expression that should return true if the intended assumption or guarantee is fulfilled or false if it does not.
	// Experimental.
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// This contains the text that Terraform will include as part of error messages when it detects an unmet condition.
	// Experimental.
	ErrorMessage *string `field:"required" json:"errorMessage" yaml:"errorMessage"`
}

