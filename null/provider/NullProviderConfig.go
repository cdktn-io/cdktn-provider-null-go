// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type NullProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/null/3.3.1/docs#alias NullProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

