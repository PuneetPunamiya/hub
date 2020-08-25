// Code generated by goa v3.2.2, DO NOT EDIT.
//
// refresh service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package refresh

import (
	"context"

	"goa.design/goa/v3/security"
)

// Catalog refresh
type Service interface {
	// Dummy api for cron job testing
	CatalogRefresh(context.Context, *CatalogRefreshPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "refresh"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"CatalogRefresh"}

// CatalogRefreshPayload is the payload type of the refresh service
// CatalogRefresh method.
type CatalogRefreshPayload struct {
	// JWT of an agent
	Token string
}
