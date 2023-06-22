package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/lisaycliu/todoist-rest-go"
	"github.com/lisaycliu/terraform-provider-todoist/to"

)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return todoist.Provider()
		},
	})
}