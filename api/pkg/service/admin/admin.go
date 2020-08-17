package admin

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"

	"github.com/tektoncd/hub/api/gen/admin"
	"github.com/tektoncd/hub/api/pkg/app"
)

type service struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
	jwtKey string
}

// New returns the admin service implementation.
func New(api app.Config) admin.Service {
	return &service{api.Logger(), api.DB(), api.JWTSigningKey()}
}

// JWTAuth implements the authorization logic for service "admin" for the "jwt"
// security scheme.
func (s *service) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return ctx, fmt.Errorf("not implemented")
}

// Create a agent user with required scopes
func (s *service) CreateAgent(ctx context.Context, p *admin.CreateAgentPayload) (res *admin.CreateAgentResult, err error) {
	s.logger.Info("admin.CreateAgent")
	return
}
