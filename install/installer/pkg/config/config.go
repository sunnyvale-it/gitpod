// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package config

// This is the config used internally. The config versions (eg, v1) is the
// public contract whereas this is the internal contract.
// There are no stability guarantees in the internal configuration
type Config struct {
	// This represents the original config
	VersionedCfg interface{}

	Kind       string
	Domain     string
	Repository string
}
