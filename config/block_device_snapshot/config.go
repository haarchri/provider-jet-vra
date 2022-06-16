package block_device_snapshot

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_block_device_snapshot", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "block_device_snapshot"
	})

	p.AddResourceConfigurator("vra_block_device", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "block_device"

		r.References["block_device"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vra/apis/blueprint/v1alpha1.block_device",
		}
	})
}