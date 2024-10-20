// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package licensemanager

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the licensemanager group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_licensemanager_association", func(r *config.Resource) {
		r.References["license_configuration_arn"] = config.Reference{
			Type:      "LicenseConfiguration",
			Extractor: common.PathARNExtractor,
		}
	})
}
