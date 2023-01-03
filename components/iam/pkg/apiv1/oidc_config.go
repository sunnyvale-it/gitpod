// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package apiv1

import (
	"context"

	goidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/gitpod-io/gitpod/common-go/log"
	db "github.com/gitpod-io/gitpod/components/gitpod-db/go"
	v1 "github.com/gitpod-io/gitpod/components/iam-api/go/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func NewOIDCClientConfigService(dbConn *gorm.DB, cipher db.Cipher) *OIDCClientConfigService {
	return &OIDCClientConfigService{
		dbConn: dbConn,
		cipher: cipher,
	}
}

type OIDCClientConfigService struct {
	dbConn *gorm.DB
	cipher db.Cipher

	v1.UnimplementedOIDCServiceServer
}

func (s *OIDCClientConfigService) CreateClientConfig(ctx context.Context, req *v1.CreateClientConfigRequest) (*v1.CreateClientConfigResponse, error) {
	config := req.GetConfig()

	data, err := db.EncryptJSON[db.OIDCSpec](db.OIDCSpec{
		OAuth2: db.OAuth2Config{
			ClientID:     config.GetOauth2Config().GetClientId(),
			ClientSecret: config.GetOauth2Config().GetClientSecret(),
			RedirectURL:  config.GetOauth2Config().GetAuthorizationEndpoint(),
			// TODO: Extract scopes from request
			Scopes: []string{goidc.ScopeOpenID, "profile", "email"},
		},
		Verifier: db.VerifierConfig{
			ClientID: config.GetOauth2Config().GetClientId(),
		},
	})
	if err != nil {
		log.Log.WithError(err).Error("Failed to encrypt oidc client config.")
		return nil, status.Errorf(codes.Internal, "Failed to store OIDC client config.")
	}

	created, err := db.CreateOIDCCLientConfig(ctx, s.dbConn, db.OIDCClientConfig{
		ID:     uuid.New(),
		Issuer: config.GetOidcConfig().GetIssuer(),
		Data:   data,
	})
	if err != nil {
		log.Log.WithError(err).Error("Failed to store oidc client config in the database.")
		return nil, status.Errorf(codes.Internal, "Failed to store OIDC client config.")
	}

	return &v1.CreateClientConfigResponse{
		Config: &v1.OIDCClientConfig{
			Id: created.ID.String(),
			// TODO: Populate remainder of fields
		},
	}, nil

	return nil, status.Errorf(codes.Unimplemented, "method CreateClientConfig not implemented")
}
