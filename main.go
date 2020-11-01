package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/scastria/terraform-provider-adops/adops"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: adops.Provider,
	})
}
