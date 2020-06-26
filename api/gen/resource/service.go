// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// The resource service provides all resources information
type Service interface {
	// Get all Resources
	AllResources(context.Context, *AllResourcesPayload) (res ResourceCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "resource"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"AllResources"}

// AllResourcesPayload is the payload type of the resource service AllResources
// method.
type AllResourcesPayload struct {
	// Number of resources
	Limit uint
}

// ResourceCollection is the result type of the resource service AllResources
// method.
type ResourceCollection []*Resource

// The resource service provides all resources information.
type Resource struct {
	// ID is the unique id of the resource
	ID uint
	// Name of the resource
	Name string
	// Display name of the resource
	DisplayName string
	// Type of catalog where resource belongs
	Catalog *Catalog
	// Type of resource
	Type string
	// Description of resource
	Description string
	// Latest version o resource
	LatestVersion string
	// Tags related to resources
	Tags []*Tag
	// Rating of resource
	Rating uint
	// Date when resource was last updated
	LastUpdatedAt string
	// Version of resource
	Versions []*Versions
}

type Catalog struct {
	// ID is the unique id of the catalog
	ID uint
	// Type of catalog
	Type string
}

type Tag struct {
	// Id is the unique id of tags
	ID uint
	// name of tag
	Name string
}

type Versions struct {
	// Version ID of the resource to be fetched
	VersionID uint
	// Version of the resource to be fetched
	Version string
}

// MakeDbDown builds a goa.ServiceError from an error.
func MakeDbDown(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "db-down",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewResourceCollection initializes result type ResourceCollection from viewed
// result type ResourceCollection.
func NewResourceCollection(vres resourceviews.ResourceCollection) ResourceCollection {
	var res ResourceCollection
	switch vres.View {
	case "default", "":
		res = newResourceCollection(vres.Projected)
	case "extended":
		res = newResourceCollectionExtended(vres.Projected)
	}
	return res
}

// NewViewedResourceCollection initializes viewed result type
// ResourceCollection from result type ResourceCollection using the given view.
func NewViewedResourceCollection(res ResourceCollection, view string) resourceviews.ResourceCollection {
	var vres resourceviews.ResourceCollection
	switch view {
	case "default", "":
		p := newResourceCollectionView(res)
		vres = resourceviews.ResourceCollection{Projected: p, View: "default"}
	case "extended":
		p := newResourceCollectionViewExtended(res)
		vres = resourceviews.ResourceCollection{Projected: p, View: "extended"}
	}
	return vres
}

// newResourceCollection converts projected type ResourceCollection to service
// type ResourceCollection.
func newResourceCollection(vres resourceviews.ResourceCollectionView) ResourceCollection {
	res := make(ResourceCollection, len(vres))
	for i, n := range vres {
		res[i] = newResource(n)
	}
	return res
}

// newResourceCollectionExtended converts projected type ResourceCollection to
// service type ResourceCollection.
func newResourceCollectionExtended(vres resourceviews.ResourceCollectionView) ResourceCollection {
	res := make(ResourceCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceExtended(n)
	}
	return res
}

// newResourceCollectionView projects result type ResourceCollection to
// projected type ResourceCollectionView using the "default" view.
func newResourceCollectionView(res ResourceCollection) resourceviews.ResourceCollectionView {
	vres := make(resourceviews.ResourceCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceView(n)
	}
	return vres
}

// newResourceCollectionViewExtended projects result type ResourceCollection to
// projected type ResourceCollectionView using the "extended" view.
func newResourceCollectionViewExtended(res ResourceCollection) resourceviews.ResourceCollectionView {
	vres := make(resourceviews.ResourceCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceViewExtended(n)
	}
	return vres
}

// newResource converts projected type Resource to service type Resource.
func newResource(vres *resourceviews.ResourceView) *Resource {
	res := &Resource{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = *vres.LatestVersion
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.LastUpdatedAt != nil {
		res.LastUpdatedAt = *vres.LastUpdatedAt
	}
	if vres.Catalog != nil {
		res.Catalog = transformResourceviewsCatalogViewToCatalog(vres.Catalog)
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Versions != nil {
		res.Versions = make([]*Versions, len(vres.Versions))
		for i, val := range vres.Versions {
			res.Versions[i] = transformResourceviewsVersionsViewToVersions(val)
		}
	}
	return res
}

// newResourceExtended converts projected type Resource to service type
// Resource.
func newResourceExtended(vres *resourceviews.ResourceView) *Resource {
	res := &Resource{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = *vres.LatestVersion
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.LastUpdatedAt != nil {
		res.LastUpdatedAt = *vres.LastUpdatedAt
	}
	if vres.Catalog != nil {
		res.Catalog = transformResourceviewsCatalogViewToCatalog(vres.Catalog)
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	return res
}

// newResourceView projects result type Resource to projected type ResourceView
// using the "default" view.
func newResourceView(res *Resource) *resourceviews.ResourceView {
	vres := &resourceviews.ResourceView{
		ID:            &res.ID,
		Name:          &res.Name,
		DisplayName:   &res.DisplayName,
		Type:          &res.Type,
		Description:   &res.Description,
		LatestVersion: &res.LatestVersion,
		Rating:        &res.Rating,
		LastUpdatedAt: &res.LastUpdatedAt,
	}
	if res.Catalog != nil {
		vres.Catalog = transformCatalogToResourceviewsCatalogView(res.Catalog)
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Versions != nil {
		vres.Versions = make([]*resourceviews.VersionsView, len(res.Versions))
		for i, val := range res.Versions {
			vres.Versions[i] = transformVersionsToResourceviewsVersionsView(val)
		}
	}
	return vres
}

// newResourceViewExtended projects result type Resource to projected type
// ResourceView using the "extended" view.
func newResourceViewExtended(res *Resource) *resourceviews.ResourceView {
	vres := &resourceviews.ResourceView{
		ID:            &res.ID,
		Name:          &res.Name,
		DisplayName:   &res.DisplayName,
		Type:          &res.Type,
		Description:   &res.Description,
		LatestVersion: &res.LatestVersion,
		Rating:        &res.Rating,
		LastUpdatedAt: &res.LastUpdatedAt,
	}
	if res.Catalog != nil {
		vres.Catalog = transformCatalogToResourceviewsCatalogView(res.Catalog)
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	return vres
}

// transformResourceviewsCatalogViewToCatalog builds a value of type *Catalog
// from a value of type *resourceviews.CatalogView.
func transformResourceviewsCatalogViewToCatalog(v *resourceviews.CatalogView) *Catalog {
	if v == nil {
		return nil
	}
	res := &Catalog{
		ID:   *v.ID,
		Type: *v.Type,
	}

	return res
}

// transformResourceviewsTagViewToTag builds a value of type *Tag from a value
// of type *resourceviews.TagView.
func transformResourceviewsTagViewToTag(v *resourceviews.TagView) *Tag {
	if v == nil {
		return nil
	}
	res := &Tag{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformResourceviewsVersionsViewToVersions builds a value of type
// *Versions from a value of type *resourceviews.VersionsView.
func transformResourceviewsVersionsViewToVersions(v *resourceviews.VersionsView) *Versions {
	if v == nil {
		return nil
	}
	res := &Versions{
		VersionID: *v.VersionID,
		Version:   *v.Version,
	}

	return res
}

// transformCatalogToResourceviewsCatalogView builds a value of type
// *resourceviews.CatalogView from a value of type *Catalog.
func transformCatalogToResourceviewsCatalogView(v *Catalog) *resourceviews.CatalogView {
	res := &resourceviews.CatalogView{
		ID:   &v.ID,
		Type: &v.Type,
	}

	return res
}

// transformTagToResourceviewsTagView builds a value of type
// *resourceviews.TagView from a value of type *Tag.
func transformTagToResourceviewsTagView(v *Tag) *resourceviews.TagView {
	res := &resourceviews.TagView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformVersionsToResourceviewsVersionsView builds a value of type
// *resourceviews.VersionsView from a value of type *Versions.
func transformVersionsToResourceviewsVersionsView(v *Versions) *resourceviews.VersionsView {
	res := &resourceviews.VersionsView{
		VersionID: &v.VersionID,
		Version:   &v.Version,
	}

	return res
}
