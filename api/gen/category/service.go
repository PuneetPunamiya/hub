// Code generated by goa v3.8.1, DO NOT EDIT.
//
// category service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package category

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The category service provides details about category
type Service interface {
	// List all categories along with their tags sorted by name
	List(context.Context) (res *ListResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "category"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"list"}

type Category struct {
	// ID is the unique id of the category
	ID uint
	// Name of category
	Name string
}

// ListResult is the result type of the category service list method.
type ListResult struct {
	Data []*Category
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal-error", false, false, false)
}
