// Code generated by goa v3.3.1, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	RefreshAccessToken goa.Endpoint
	NewRefreshToken    goa.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		RefreshAccessToken: NewRefreshAccessTokenEndpoint(s, a.JWTAuth),
		NewRefreshToken:    NewNewRefreshTokenEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.RefreshAccessToken = m(e.RefreshAccessToken)
	e.NewRefreshToken = m(e.NewRefreshToken)
}

// NewRefreshAccessTokenEndpoint returns an endpoint function that calls the
// method "RefreshAccessToken" of service "user".
func NewRefreshAccessTokenEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RefreshAccessTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"rating:read", "rating:write", "agent:create", "catalog:refresh", "config:refresh", "refresh:token"},
			RequiredScopes: []string{"refresh:token"},
		}
		ctx, err = authJWTFn(ctx, p.RefreshToken, &sc)
		if err != nil {
			return nil, err
		}
		return s.RefreshAccessToken(ctx, p)
	}
}

// NewNewRefreshTokenEndpoint returns an endpoint function that calls the
// method "NewRefreshToken" of service "user".
func NewNewRefreshTokenEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*NewRefreshTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"rating:read", "rating:write", "agent:create", "catalog:refresh", "config:refresh", "refresh:token"},
			RequiredScopes: []string{"refresh:token"},
		}
		ctx, err = authJWTFn(ctx, p.RefreshToken, &sc)
		if err != nil {
			return nil, err
		}
		return s.NewRefreshToken(ctx, p)
	}
}