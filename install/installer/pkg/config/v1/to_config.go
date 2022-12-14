// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package config

import (
	"fmt"

	"github.com/gitpod-io/gitpod/installer/pkg/config"
)

func (v version) MapToConfig(rawCfg interface{}) (config.Config, error) {
	cfg := rawCfg.(*Config)

	fmt.Println(cfg)

	return config.Config{
		VersionedCfg: cfg,

		Kind:       string(cfg.Kind),
		Domain:     cfg.Domain,
		Repository: cfg.Repository,
	}, nil
}
