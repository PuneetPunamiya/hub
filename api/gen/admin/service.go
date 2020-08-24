// Code generated by goa v3.2.2, DO NOT EDIT.
//
// admin service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Admin service
type Service interface {
	// Create an agent user with required scopes
	CreateAgent(context.Context, *CreateAgentPayload) (res *CreateAgentResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "admin"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"CreateAgent"}

// CreateAgentPayload is the payload type of the admin service CreateAgent
// method.
type CreateAgentPayload struct {
	// Name of Agent
	Name string
	// Scopes required for Agent
	Scopes []string
	// User JWT
	Token string
}

// CreateAgentResult is the result type of the admin service CreateAgent method.
type CreateAgentResult struct {
	// Agent JWT
	Token string
}

// MakeInvalidPayload builds a goa.ServiceError from an error.
func MakeInvalidPayload(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid-payload",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInvalidToken builds a goa.ServiceError from an error.
func MakeInvalidToken(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid-token",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInvalidScopes builds a goa.ServiceError from an error.
func MakeInvalidScopes(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid-scopes",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal-error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}