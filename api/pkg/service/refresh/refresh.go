package refresh

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	refresh "github.com/tektoncd/hub/api/gen/refresh"
	"github.com/tektoncd/hub/api/pkg/app"
	"goa.design/goa/v3/security"
)

// refresh service example implementation.
// The example methods log the requests and return zero values.
type service struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

// NewRefresh returns the refresh service implementation.
func New(api app.BaseConfig) refresh.Service {
	return &service{api.Logger(), api.DB()}
}

// JWTAuth implements the authorization logic for service "refresh" for the
// "jwt" security scheme.
func (s *service) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	s.logger.Info("Successfully called")
	s.logger.Info("Token ", token)
	return ctx, fmt.Errorf("not implemented")
}

// Dummy api for cron job testing
func (s *service) CatalogRefresh(ctx context.Context, p *refresh.CatalogRefreshPayload) (err error) {
	s.logger.Info("refresh.CatalogRefresh")
	return
}
