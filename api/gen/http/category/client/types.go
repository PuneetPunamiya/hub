// Code generated by goa v3.1.3, DO NOT EDIT.
//
// category HTTP client types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	category "github.com/tektoncd/hub/api/gen/category"
	goa "goa.design/goa/v3/pkg"
)

// AllResponseBody is the type of the "category" service "All" endpoint HTTP
// response body.
type AllResponseBody []*CategoryResponse

// AllInternalErrorResponseBody is the type of the "category" service "All"
// endpoint HTTP response body for the "internal-error" error.
type AllInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CategoryResponse is used to define fields on response body types.
type CategoryResponse struct {
	// unique id of category
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of category
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// list of tag associated with category
	Tags []*TagResponse `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
}

// TagResponse is used to define fields on response body types.
type TagResponse struct {
	// Id is the unique id of tags
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of tag
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewAllCategoryOK builds a "category" service "All" endpoint result from a
// HTTP "OK" response.
func NewAllCategoryOK(body []*CategoryResponse) []*category.Category {
	v := make([]*category.Category, len(body))
	for i, val := range body {
		v[i] = unmarshalCategoryResponseToCategoryCategory(val)
	}
	return v
}

// NewAllInternalError builds a category service All endpoint internal-error
// error.
func NewAllInternalError(body *AllInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateAllInternalErrorResponseBody runs the validations defined on
// All_internal-error_Response_Body
func ValidateAllInternalErrorResponseBody(body *AllInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCategoryResponse runs the validations defined on categoryResponse
func ValidateCategoryResponse(body *CategoryResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
	}
	for _, e := range body.Tags {
		if e != nil {
			if err2 := ValidateTagResponse(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateTagResponse runs the validations defined on TagResponse
func ValidateTagResponse(body *TagResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
