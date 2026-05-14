// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/null/3.3.0/docs/resources/resource#triggers Resource#triggers}
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

