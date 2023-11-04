// Code generated by goa v3.14.0, DO NOT EDIT.
//
// resource HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	resource "github.com/tektoncd/hub/api/gen/resource"
	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
)

// ListResponseBody is the type of the "resource" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Data ResourceDataResponseBodyWithoutVersionCollection `form:"data" json:"data" xml:"data"`
}

// ResourceDataResponseBodyWithoutVersionCollection is used to define fields on
// response body types.
type ResourceDataResponseBodyWithoutVersionCollection []*ResourceDataResponseBodyWithoutVersion

// ResourceDataResponseBodyWithoutVersion is used to define fields on response
// body types.
type ResourceDataResponseBodyWithoutVersion struct {
	// ID is the unique id of the resource
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of resource
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog to which resource belongs
	Catalog *CatalogResponseBodyMin `form:"catalog" json:"catalog" xml:"catalog"`
	// Categories related to resource
	Categories []*CategoryResponseBody `form:"categories" json:"categories" xml:"categories"`
	// Kind of resource
	Kind string `form:"kind" json:"kind" xml:"kind"`
	// Url path of the resource in hub
	HubURLPath string `form:"hubURLPath" json:"hubURLPath" xml:"hubURLPath"`
	// Path of the api to get the raw yaml of resource from hub apiserver
	HubRawURLPath string `form:"hubRawURLPath" json:"hubRawURLPath" xml:"hubRawURLPath"`
	// Latest version of resource
	LatestVersion *ResourceVersionDataResponseBodyWithoutResource `form:"latestVersion" json:"latestVersion" xml:"latestVersion"`
	// Tags related to resource
	Tags []*TagResponseBody `form:"tags" json:"tags" xml:"tags"`
	// Platforms related to resource
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
	// Rating of resource
	Rating float64 `form:"rating" json:"rating" xml:"rating"`
}

// CatalogResponseBodyMin is used to define fields on response body types.
type CatalogResponseBodyMin struct {
	// ID is the unique id of the catalog
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of catalog
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog
	Type string `form:"type" json:"type" xml:"type"`
}

// CategoryResponseBody is used to define fields on response body types.
type CategoryResponseBody struct {
	// ID is the unique id of the category
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of category
	Name string `form:"name" json:"name" xml:"name"`
}

// ResourceVersionDataResponseBodyWithoutResource is used to define fields on
// response body types.
type ResourceVersionDataResponseBodyWithoutResource struct {
	// ID is the unique id of resource's version
	ID uint `form:"id" json:"id" xml:"id"`
	// Version of resource
	Version string `form:"version" json:"version" xml:"version"`
	// Display name of version
	DisplayName string `form:"displayName" json:"displayName" xml:"displayName"`
	// Deprecation status of a version
	Deprecated *bool `form:"deprecated,omitempty" json:"deprecated,omitempty" xml:"deprecated,omitempty"`
	// Description of version
	Description string `form:"description" json:"description" xml:"description"`
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string `form:"minPipelinesVersion" json:"minPipelinesVersion" xml:"minPipelinesVersion"`
	// Raw URL of resource's yaml file of the version
	RawURL string `form:"rawURL" json:"rawURL" xml:"rawURL"`
	// Web URL of resource's yaml file of the version
	WebURL string `form:"webURL" json:"webURL" xml:"webURL"`
	// Path of the api to get the raw yaml of resource from hub apiserver
	HubRawURLPath string `form:"hubRawURLPath" json:"hubRawURLPath" xml:"hubRawURLPath"`
	// Url path of the resource in hub
	HubURLPath string `form:"hubURLPath" json:"hubURLPath" xml:"hubURLPath"`
	// Timestamp when version was last updated
	UpdatedAt string `form:"updatedAt" json:"updatedAt" xml:"updatedAt"`
	// Platforms related to resource version
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
}

// PlatformResponseBody is used to define fields on response body types.
type PlatformResponseBody struct {
	// ID is the unique id of platform
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of platform
	Name string `form:"name" json:"name" xml:"name"`
}

// TagResponseBody is used to define fields on response body types.
type TagResponseBody struct {
	// ID is the unique id of tag
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of tag
	Name string `form:"name" json:"name" xml:"name"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "resource" service.
func NewListResponseBody(res *resourceviews.ResourcesView) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Data != nil {
		body.Data = make([]*ResourceDataResponseBodyWithoutVersion, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalResourceviewsResourceDataViewToResourceDataResponseBodyWithoutVersion(val)
		}
	} else {
		body.Data = []*ResourceDataResponseBodyWithoutVersion{}
	}
	return body
}

// NewQueryPayload builds a resource service Query endpoint payload.
func NewQueryPayload(name string, catalogs []string, categories []string, kinds []string, tags []string, platforms []string, limit uint, match string) *resource.QueryPayload {
	v := &resource.QueryPayload{}
	v.Name = name
	v.Catalogs = catalogs
	v.Categories = categories
	v.Kinds = kinds
	v.Tags = tags
	v.Platforms = platforms
	v.Limit = limit
	v.Match = match

	return v
}

// NewVersionsByIDPayload builds a resource service VersionsByID endpoint
// payload.
func NewVersionsByIDPayload(id uint) *resource.VersionsByIDPayload {
	v := &resource.VersionsByIDPayload{}
	v.ID = id

	return v
}

// NewByCatalogKindNameVersionPayload builds a resource service
// ByCatalogKindNameVersion endpoint payload.
func NewByCatalogKindNameVersionPayload(catalog string, kind string, name string, version string) *resource.ByCatalogKindNameVersionPayload {
	v := &resource.ByCatalogKindNameVersionPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v
}

// NewByVersionIDPayload builds a resource service ByVersionId endpoint payload.
func NewByVersionIDPayload(versionID uint) *resource.ByVersionIDPayload {
	v := &resource.ByVersionIDPayload{}
	v.VersionID = versionID

	return v
}

// NewByCatalogKindNamePayload builds a resource service ByCatalogKindName
// endpoint payload.
func NewByCatalogKindNamePayload(catalog string, kind string, name string, pipelinesversion *string) *resource.ByCatalogKindNamePayload {
	v := &resource.ByCatalogKindNamePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Pipelinesversion = pipelinesversion

	return v
}

// NewByIDPayload builds a resource service ById endpoint payload.
func NewByIDPayload(id uint) *resource.ByIDPayload {
	v := &resource.ByIDPayload{}
	v.ID = id

	return v
}
