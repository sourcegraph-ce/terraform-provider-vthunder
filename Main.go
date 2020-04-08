package main

import (
	"github.com/terraform-providers/terraform-provider-vthunder/vthunder"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vthunder.Provider})
}
