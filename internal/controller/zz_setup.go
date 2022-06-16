/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

<<<<<<< HEAD

	accountaws "github.com/crossplane-contrib/provider-jet-vra/internal/controller/aws/accountaws"
	accountazure "github.com/crossplane-contrib/provider-jet-vra/internal/controller/azure/accountazure"

	blueprint "github.com/crossplane-contrib/provider-jet-vra/internal/controller/blueprint/blueprint"
	version "github.com/crossplane-contrib/provider-jet-vra/internal/controller/blueprint/version"
	itementitlement "github.com/crossplane-contrib/provider-jet-vra/internal/controller/catalogitementitlement/itementitlement"
	sourceblueprint "github.com/crossplane-contrib/provider-jet-vra/internal/controller/catalogsourceblueprint/sourceblueprint"
	sourceentitlement "github.com/crossplane-contrib/provider-jet-vra/internal/controller/catalogsourceentitlement/sourceentitlement"

	source "github.com/crossplane-contrib/provider-jet-vra/internal/controller/contentsource/source"
	deployment "github.com/crossplane-contrib/provider-jet-vra/internal/controller/deployment/deployment"

	accountgcp "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccountgcp/accountgcp"
	accountnsxt "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccountnsxt/accountnsxt"
	source "github.com/crossplane-contrib/provider-jet-vra/internal/controller/contentsource/source"
	deployment "github.com/crossplane-contrib/provider-jet-vra/internal/controller/deployment/deployment"

	profile "github.com/crossplane-contrib/provider-jet-vra/internal/controller/flavorprofile/profile"
	balancer "github.com/crossplane-contrib/provider-jet-vra/internal/controller/loadbalancer/balancer"

	machine "github.com/crossplane-contrib/provider-jet-vra/internal/controller/machine/machine"
	iprange "github.com/crossplane-contrib/provider-jet-vra/internal/controller/network/iprange"
	network "github.com/crossplane-contrib/provider-jet-vra/internal/controller/network/network"
	profile "github.com/crossplane-contrib/provider-jet-vra/internal/controller/networkprofile/profile"


	project "github.com/crossplane-contrib/provider-jet-vra/internal/controller/project/project"
	providerconfig "github.com/crossplane-contrib/provider-jet-vra/internal/controller/providerconfig"
	zone "github.com/crossplane-contrib/provider-jet-vra/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{

		accountaws.Setup,
		accountazure.Setup,

		blueprint.Setup,
		version.Setup,
		itementitlement.Setup,
		sourceblueprint.Setup,
		sourceentitlement.Setup,

		source.Setup,
		deployment.Setup,

		accountgcp.Setup,
		accountnsxt.Setup,
		source.Setup,
		deployment.Setup,

		profile.Setup,
		profile.Setup,
		balancer.Setup,

		machine.Setup,
		iprange.Setup,
		network.Setup,
		profile.Setup,


		project.Setup,
		providerconfig.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}