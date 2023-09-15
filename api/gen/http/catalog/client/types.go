// Code generated by goa v3.13.0, DO NOT EDIT.
//
// catalog HTTP client types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	catalog "github.com/tektoncd/hub/api/gen/catalog"
	catalogviews "github.com/tektoncd/hub/api/gen/catalog/views"
	goa "goa.design/goa/v3/pkg"
)

// RefreshResponseBody is the type of the "catalog" service "Refresh" endpoint
// HTTP response body.
type RefreshResponseBody struct {
	// id of the job
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the catalog
	CatalogName *string `form:"catalogName,omitempty" json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// status of the job
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// RefreshAllResponseBody is the type of the "catalog" service "RefreshAll"
// endpoint HTTP response body.
type RefreshAllResponseBody []*JobResponse

// CatalogErrorResponseBody is the type of the "catalog" service "CatalogError"
// endpoint HTTP response body.
type CatalogErrorResponseBody struct {
	// Catalog errors
	Data []*CatalogErrorsResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// RefreshNotFoundResponseBody is the type of the "catalog" service "Refresh"
// endpoint HTTP response body for the "not-found" error.
type RefreshNotFoundResponseBody struct {
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

// RefreshInternalErrorResponseBody is the type of the "catalog" service
// "Refresh" endpoint HTTP response body for the "internal-error" error.
type RefreshInternalErrorResponseBody struct {
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

// RefreshAllInternalErrorResponseBody is the type of the "catalog" service
// "RefreshAll" endpoint HTTP response body for the "internal-error" error.
type RefreshAllInternalErrorResponseBody struct {
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

// CatalogErrorInternalErrorResponseBody is the type of the "catalog" service
// "CatalogError" endpoint HTTP response body for the "internal-error" error.
type CatalogErrorInternalErrorResponseBody struct {
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

// JobResponse is used to define fields on response body types.
type JobResponse struct {
	// id of the job
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the catalog
	CatalogName *string `form:"catalogName,omitempty" json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// status of the job
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// CatalogErrorsResponseBody is used to define fields on response body types.
type CatalogErrorsResponseBody struct {
	// Catalog Errror type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Catalog Error message
	Errors []string `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
}

// NewRefreshJobOK builds a "catalog" service "Refresh" endpoint result from a
// HTTP "OK" response.
func NewRefreshJobOK(body *RefreshResponseBody) *catalogviews.JobView {
	v := &catalogviews.JobView{
		ID:          body.ID,
		CatalogName: body.CatalogName,
		Status:      body.Status,
	}

	return v
}

// NewRefreshNotFound builds a catalog service Refresh endpoint not-found error.
func NewRefreshNotFound(body *RefreshNotFoundResponseBody) *goa.ServiceError {
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

// NewRefreshInternalError builds a catalog service Refresh endpoint
// internal-error error.
func NewRefreshInternalError(body *RefreshInternalErrorResponseBody) *goa.ServiceError {
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

// NewRefreshAllJobOK builds a "catalog" service "RefreshAll" endpoint result
// from a HTTP "OK" response.
func NewRefreshAllJobOK(body []*JobResponse) []*catalog.Job {
	v := make([]*catalog.Job, len(body))
	for i, val := range body {
		v[i] = unmarshalJobResponseToCatalogJob(val)
	}

	return v
}

// NewRefreshAllInternalError builds a catalog service RefreshAll endpoint
// internal-error error.
func NewRefreshAllInternalError(body *RefreshAllInternalErrorResponseBody) *goa.ServiceError {
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

// NewCatalogErrorResultOK builds a "catalog" service "CatalogError" endpoint
// result from a HTTP "OK" response.
func NewCatalogErrorResultOK(body *CatalogErrorResponseBody) *catalog.CatalogErrorResult {
	v := &catalog.CatalogErrorResult{}
	v.Data = make([]*catalog.CatalogErrors, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalCatalogErrorsResponseBodyToCatalogCatalogErrors(val)
	}

	return v
}

// NewCatalogErrorInternalError builds a catalog service CatalogError endpoint
// internal-error error.
func NewCatalogErrorInternalError(body *CatalogErrorInternalErrorResponseBody) *goa.ServiceError {
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

// ValidateCatalogErrorResponseBody runs the validations defined on
// CatalogErrorResponseBody
func ValidateCatalogErrorResponseBody(body *CatalogErrorResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateCatalogErrorsResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateRefreshNotFoundResponseBody runs the validations defined on
// Refresh_not-found_Response_Body
func ValidateRefreshNotFoundResponseBody(body *RefreshNotFoundResponseBody) (err error) {
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

// ValidateRefreshInternalErrorResponseBody runs the validations defined on
// Refresh_internal-error_Response_Body
func ValidateRefreshInternalErrorResponseBody(body *RefreshInternalErrorResponseBody) (err error) {
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

// ValidateRefreshAllInternalErrorResponseBody runs the validations defined on
// RefreshAll_internal-error_Response_Body
func ValidateRefreshAllInternalErrorResponseBody(body *RefreshAllInternalErrorResponseBody) (err error) {
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

// ValidateCatalogErrorInternalErrorResponseBody runs the validations defined
// on CatalogError_internal-error_Response_Body
func ValidateCatalogErrorInternalErrorResponseBody(body *CatalogErrorInternalErrorResponseBody) (err error) {
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

// ValidateJobResponse runs the validations defined on JobResponse
func ValidateJobResponse(body *JobResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.CatalogName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("catalogName", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	return
}

// ValidateCatalogErrorsResponseBody runs the validations defined on
// CatalogErrorsResponseBody
func ValidateCatalogErrorsResponseBody(body *CatalogErrorsResponseBody) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Errors == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errors", "body"))
	}
	return
}
