// Code generated by goa v3.8.3, DO NOT EDIT.
//
// catalog service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package catalog

import (
	"context"

	catalogviews "github.com/tektoncd/hub/api/gen/catalog/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The Catalog Service exposes endpoints to interact with catalogs
type Service interface {
	// Refresh a Catalog by it's name
	Refresh(context.Context, *RefreshPayload) (res *Job, err error)
	// Refresh all catalogs
	RefreshAll(context.Context, *RefreshAllPayload) (res []*Job, err error)
	// List all errors occurred refreshing a catalog
	CatalogError(context.Context, *CatalogErrorPayload) (res *CatalogErrorResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "catalog"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"Refresh", "RefreshAll", "CatalogError"}

// CatalogErrorPayload is the payload type of the catalog service CatalogError
// method.
type CatalogErrorPayload struct {
	// Name of catalog
	CatalogName string
	// JWT
	Token string
}

// CatalogErrorResult is the result type of the catalog service CatalogError
// method.
type CatalogErrorResult struct {
	// Catalog errors
	Data []*CatalogErrors
}

// CatalogErrors define the errors that occurred during catalog refresh
type CatalogErrors struct {
	// Catalog Errror type
	Type string
	// Catalog Error message
	Errors []string
}

// Job is the result type of the catalog service Refresh method.
type Job struct {
	// id of the job
	ID uint
	// Name of the catalog
	CatalogName string
	// status of the job
	Status string
}

// RefreshAllPayload is the payload type of the catalog service RefreshAll
// method.
type RefreshAllPayload struct {
	// JWT
	Token string
}

// RefreshPayload is the payload type of the catalog service Refresh method.
type RefreshPayload struct {
	// Name of catalog
	CatalogName string
	// JWT
	Token string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal-error", false, false, false)
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not-found", false, false, false)
}

// NewJob initializes result type Job from viewed result type Job.
func NewJob(vres *catalogviews.Job) *Job {
	return newJob(vres.Projected)
}

// NewViewedJob initializes viewed result type Job from result type Job using
// the given view.
func NewViewedJob(res *Job, view string) *catalogviews.Job {
	p := newJobView(res)
	return &catalogviews.Job{Projected: p, View: "default"}
}

// newJob converts projected type Job to service type Job.
func newJob(vres *catalogviews.JobView) *Job {
	res := &Job{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.CatalogName != nil {
		res.CatalogName = *vres.CatalogName
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	return res
}

// newJobView projects result type Job to projected type JobView using the
// "default" view.
func newJobView(res *Job) *catalogviews.JobView {
	vres := &catalogviews.JobView{
		ID:          &res.ID,
		CatalogName: &res.CatalogName,
		Status:      &res.Status,
	}
	return vres
}
