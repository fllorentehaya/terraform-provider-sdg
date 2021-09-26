package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		// Add the resources supported by this provider to this map.
		ResourcesMap: map[string]*schema.Resource{
			"sdg_batch_template": resourceBatchTemplate(),
		},
	}
}
