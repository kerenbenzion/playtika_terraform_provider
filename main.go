package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
        "github.com/rafi-barel/playtika-terraform-provider-drp/drpv4"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() *schema.Provider {
                        return drpv4.Provider()
                },
        })
}
