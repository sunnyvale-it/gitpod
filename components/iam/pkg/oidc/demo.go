// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package oidc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	db "github.com/gitpod-io/gitpod/components/gitpod-db/go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// The demo config is used to setup a OIDC client with Google.
//
// This is a temporary way to boot the OIDC client service with a single
// configuration, e.g. mounted as secret into a preview environment.
//
// ‼️ This demo config will be removed once the configuration is read from DB.
type DemoConfig struct {
	Issuer       string `json:"issuer"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	RedirectURL  string `json:"redirectURL"`
}

func readDemoConfigFromFile(path string) (*DemoConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read test config: %w", err)
	}

	var config DemoConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func LoadDemoConfigIntoDB(dbConn *gorm.DB, cipherSet *db.CipherSet, path string) (*db.OIDCClientConfig, error) {
	testConfig, err := readDemoConfigFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read test config: %w", err)
	}

	data := db.OIDCSpec{
		ClientId:     testConfig.ClientID,
		ClientSecret: testConfig.ClientSecret,
	}

	encrypted, err := db.EncryptJSON(cipherSet, data)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data: %w", err)
	}

	entry, err := db.CreateOIDCClientConfig(context.Background(), dbConn, db.OIDCClientConfig{
		ID:     uuid.New(),
		Issuer: testConfig.Issuer,
		Data:   encrypted,
	})
	return &entry, err
}
