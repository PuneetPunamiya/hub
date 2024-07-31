// Code generated by goa v3.18.0, DO NOT EDIT.
//
// rating HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	rating "github.com/tektoncd/hub/api/gen/rating"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "rating" service "Update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// User rating for resource
	Rating *uint `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// GetResponseBody is the type of the "rating" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// User rating for resource
	Rating int `form:"rating" json:"rating" xml:"rating"`
}

// GetNotFoundResponseBody is the type of the "rating" service "Get" endpoint
// HTTP response body for the "not-found" error.
type GetNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetInternalErrorResponseBody is the type of the "rating" service "Get"
// endpoint HTTP response body for the "internal-error" error.
type GetInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetInvalidTokenResponseBody is the type of the "rating" service "Get"
// endpoint HTTP response body for the "invalid-token" error.
type GetInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetInvalidScopesResponseBody is the type of the "rating" service "Get"
// endpoint HTTP response body for the "invalid-scopes" error.
type GetInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotFoundResponseBody is the type of the "rating" service "Update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInternalErrorResponseBody is the type of the "rating" service "Update"
// endpoint HTTP response body for the "internal-error" error.
type UpdateInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInvalidTokenResponseBody is the type of the "rating" service "Update"
// endpoint HTTP response body for the "invalid-token" error.
type UpdateInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInvalidScopesResponseBody is the type of the "rating" service "Update"
// endpoint HTTP response body for the "invalid-scopes" error.
type UpdateInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "Get" endpoint of the "rating" service.
func NewGetResponseBody(res *rating.GetResult) *GetResponseBody {
	body := &GetResponseBody{
		Rating: res.Rating,
	}
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "Get" endpoint of the "rating" service.
func NewGetNotFoundResponseBody(res *goa.ServiceError) *GetNotFoundResponseBody {
	body := &GetNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInternalErrorResponseBody builds the HTTP response body from the
// result of the "Get" endpoint of the "rating" service.
func NewGetInternalErrorResponseBody(res *goa.ServiceError) *GetInternalErrorResponseBody {
	body := &GetInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInvalidTokenResponseBody builds the HTTP response body from the result
// of the "Get" endpoint of the "rating" service.
func NewGetInvalidTokenResponseBody(res *goa.ServiceError) *GetInvalidTokenResponseBody {
	body := &GetInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInvalidScopesResponseBody builds the HTTP response body from the
// result of the "Get" endpoint of the "rating" service.
func NewGetInvalidScopesResponseBody(res *goa.ServiceError) *GetInvalidScopesResponseBody {
	body := &GetInvalidScopesResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "Update" endpoint of the "rating" service.
func NewUpdateNotFoundResponseBody(res *goa.ServiceError) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInternalErrorResponseBody builds the HTTP response body from the
// result of the "Update" endpoint of the "rating" service.
func NewUpdateInternalErrorResponseBody(res *goa.ServiceError) *UpdateInternalErrorResponseBody {
	body := &UpdateInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInvalidTokenResponseBody builds the HTTP response body from the
// result of the "Update" endpoint of the "rating" service.
func NewUpdateInvalidTokenResponseBody(res *goa.ServiceError) *UpdateInvalidTokenResponseBody {
	body := &UpdateInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInvalidScopesResponseBody builds the HTTP response body from the
// result of the "Update" endpoint of the "rating" service.
func NewUpdateInvalidScopesResponseBody(res *goa.ServiceError) *UpdateInvalidScopesResponseBody {
	body := &UpdateInvalidScopesResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetPayload builds a rating service Get endpoint payload.
func NewGetPayload(id uint, token string) *rating.GetPayload {
	v := &rating.GetPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewUpdatePayload builds a rating service Update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id uint, token string) *rating.UpdatePayload {
	v := &rating.UpdatePayload{
		Rating: *body.Rating,
	}
	v.ID = id
	v.Token = token

	return v
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "body"))
	}
	if body.Rating != nil {
		if *body.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 0, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}
