// Code generated by goa v3.2.2, DO NOT EDIT.
//
// admin endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "admin" service endpoints.
type Endpoints struct {
	CreateAgent goa.Endpoint
}

// NewEndpoints wraps the methods of the "admin" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		CreateAgent: NewCreateAgentEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "admin" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.CreateAgent = m(e.CreateAgent)
}

// NewCreateAgentEndpoint returns an endpoint function that calls the method
// "CreateAgent" of service "admin".
func NewCreateAgentEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateAgentPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write", "agent:create"},
			RequiredScopes: []string{"agent:create"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.CreateAgent(ctx, p)
	}
}