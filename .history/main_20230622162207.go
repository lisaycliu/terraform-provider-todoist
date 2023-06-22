package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/lisaycliu/todoist-rest-go"
	"github.com/lisaycliu/terraform-provider-todoist.git"

)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return todoist.Provider()
		},
	})
}