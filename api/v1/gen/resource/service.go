// Code generated by goa v3.3.1, DO NOT EDIT.
//
// resource service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package resource

import (
	"context"

	resourceviews "github.com/tektoncd/hub/api/v1/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// The resource service provides details about all kind of resources
type Service interface {
	// Find resources by a combination of name, kind,catalog and tags
	Query(context.Context, *QueryPayload) (res *Resources, err error)
	// List all resources sorted by rating and name
	List(context.Context, *ListPayload) (res *Resources, err error)
	// Find all versions of a resource by its id
	VersionsByID(context.Context, *VersionsByIDPayload) (res *ResourceVersions, err error)
	// Find resource using name of catalog & name, kind and version of resource
	ByCatalogKindNameVersion(context.Context, *ByCatalogKindNameVersionPayload) (res *ResourceVersion, err error)
	// Find a resource using its version's id
	ByVersionID(context.Context, *ByVersionIDPayload) (res *ResourceVersion, err error)
	// Find resources using name of catalog, resource name and kind of resource
	ByCatalogKindName(context.Context, *ByCatalogKindNamePayload) (res *Resource, err error)
	// Find a resource using it's id
	ByID(context.Context, *ByIDPayload) (res *Resource, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "resource"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [7]string{"Query", "List", "VersionsByID", "ByCatalogKindNameVersion", "ByVersionId", "ByCatalogKindName", "ById"}

// QueryPayload is the payload type of the resource service Query method.
type QueryPayload struct {
	// Name of resource
	Name string
	// Catalogs of resource to filter by
	Catalogs []string
	// Kinds of resource to filter by
	Kinds []string
	// Category associated with a resource to filter by
	Categories []string
	// Tags associated with a resource to filter by
	Tags []string
	// Maximum number of resources to be returned
	Limit uint
	// Strategy used to find matching resources
	Match string
}

// Resources is the result type of the resource service Query method.
type Resources struct {
	Data ResourceDataCollection
}

// ListPayload is the payload type of the resource service List method.
type ListPayload struct {
	// Maximum number of resources to be returned
	Limit uint
}

// VersionsByIDPayload is the payload type of the resource service VersionsByID
// method.
type VersionsByIDPayload struct {
	// ID of a resource
	ID uint
}

// ResourceVersions is the result type of the resource service VersionsByID
// method.
type ResourceVersions struct {
	Data *Versions
}

// ByCatalogKindNameVersionPayload is the payload type of the resource service
// ByCatalogKindNameVersion method.
type ByCatalogKindNameVersionPayload struct {
	// name of catalog
	Catalog string
	// kind of resource
	Kind string
	// name of resource
	Name string
	// version of resource
	Version string
}

// ResourceVersion is the result type of the resource service
// ByCatalogKindNameVersion method.
type ResourceVersion struct {
	Data *ResourceVersionData
}

// ByVersionIDPayload is the payload type of the resource service ByVersionId
// method.
type ByVersionIDPayload struct {
	// Version ID of a resource's version
	VersionID uint
}

// ByCatalogKindNamePayload is the payload type of the resource service
// ByCatalogKindName method.
type ByCatalogKindNamePayload struct {
	// name of catalog
	Catalog string
	// kind of resource
	Kind string
	// Name of resource
	Name string
	// To find resource compatible with a Tekton pipelines version, use this param
	Pipelinesversion *string
}

// Resource is the result type of the resource service ByCatalogKindName method.
type Resource struct {
	Data *ResourceData
}

// ByIDPayload is the payload type of the resource service ById method.
type ByIDPayload struct {
	// ID of a resource
	ID uint
}

type ResourceDataCollection []*ResourceData

// The resource type describes resource information.
type ResourceData struct {
	// ID is the unique id of the resource
	ID uint
	// Name of resource
	Name string
	// Type of catalog to which resource belongs
	Catalog *Catalog
	// Categories related to resource
	Categories []*Category
	// Kind of resource
	Kind string
	// Latest version of resource
	LatestVersion *ResourceVersionData
	// Tags related to resource
	Tags []*Tag
	// Rating of resource
	Rating float64
	// List of all versions of a resource
	Versions []*ResourceVersionData
}

type Catalog struct {
	// ID is the unique id of the catalog
	ID uint
	// Name of catalog
	Name string
	// Type of catalog
	Type string
	// URL of catalog
	URL string
}

type Category struct {
	// ID is the unique id of the category
	ID uint
	// Name of category
	Name string
}

// The Version result type describes resource's version information.
type ResourceVersionData struct {
	// ID is the unique id of resource's version
	ID uint
	// Version of resource
	Version string
	// Display name of version
	DisplayName string
	// Description of version
	Description string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string
	// Raw URL of resource's yaml file of the version
	RawURL string
	// Web URL of resource's yaml file of the version
	WebURL string
	// Timestamp when version was last updated
	UpdatedAt string
	// Resource to which the version belongs
	Resource *ResourceData
}

type Tag struct {
	// ID is the unique id of tag
	ID uint
	// Name of tag
	Name string
}

// The Versions type describes response for versions by resource id API.
type Versions struct {
	// Latest Version of resource
	Latest *ResourceVersionData
	// List of all versions of resource
	Versions []*ResourceVersionData
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal-error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInvalidKind builds a goa.ServiceError from an error.
func MakeInvalidKind(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid-kind",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewResources initializes result type Resources from viewed result type
// Resources.
func NewResources(vres *resourceviews.Resources) *Resources {
	return newResources(vres.Projected)
}

// NewViewedResources initializes viewed result type Resources from result type
// Resources using the given view.
func NewViewedResources(res *Resources, view string) *resourceviews.Resources {
	p := newResourcesView(res)
	return &resourceviews.Resources{Projected: p, View: "default"}
}

// NewResourceVersions initializes result type ResourceVersions from viewed
// result type ResourceVersions.
func NewResourceVersions(vres *resourceviews.ResourceVersions) *ResourceVersions {
	return newResourceVersions(vres.Projected)
}

// NewViewedResourceVersions initializes viewed result type ResourceVersions
// from result type ResourceVersions using the given view.
func NewViewedResourceVersions(res *ResourceVersions, view string) *resourceviews.ResourceVersions {
	p := newResourceVersionsView(res)
	return &resourceviews.ResourceVersions{Projected: p, View: "default"}
}

// NewResourceVersion initializes result type ResourceVersion from viewed
// result type ResourceVersion.
func NewResourceVersion(vres *resourceviews.ResourceVersion) *ResourceVersion {
	return newResourceVersion(vres.Projected)
}

// NewViewedResourceVersion initializes viewed result type ResourceVersion from
// result type ResourceVersion using the given view.
func NewViewedResourceVersion(res *ResourceVersion, view string) *resourceviews.ResourceVersion {
	p := newResourceVersionView(res)
	return &resourceviews.ResourceVersion{Projected: p, View: "default"}
}

// NewResource initializes result type Resource from viewed result type
// Resource.
func NewResource(vres *resourceviews.Resource) *Resource {
	return newResource(vres.Projected)
}

// NewViewedResource initializes viewed result type Resource from result type
// Resource using the given view.
func NewViewedResource(res *Resource, view string) *resourceviews.Resource {
	p := newResourceView(res)
	return &resourceviews.Resource{Projected: p, View: "default"}
}

// newResources converts projected type Resources to service type Resources.
func newResources(vres *resourceviews.ResourcesView) *Resources {
	res := &Resources{}
	if vres.Data != nil {
		res.Data = newResourceDataCollectionWithoutVersion(vres.Data)
	}
	return res
}

// newResourcesView projects result type Resources to projected type
// ResourcesView using the "default" view.
func newResourcesView(res *Resources) *resourceviews.ResourcesView {
	vres := &resourceviews.ResourcesView{}
	if res.Data != nil {
		vres.Data = newResourceDataCollectionViewWithoutVersion(res.Data)
	}
	return vres
}

// newResourceDataCollectionInfo converts projected type ResourceDataCollection
// to service type ResourceDataCollection.
func newResourceDataCollectionInfo(vres resourceviews.ResourceDataCollectionView) ResourceDataCollection {
	res := make(ResourceDataCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceDataInfo(n)
	}
	return res
}

// newResourceDataCollectionWithoutVersion converts projected type
// ResourceDataCollection to service type ResourceDataCollection.
func newResourceDataCollectionWithoutVersion(vres resourceviews.ResourceDataCollectionView) ResourceDataCollection {
	res := make(ResourceDataCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceDataWithoutVersion(n)
	}
	return res
}

// newResourceDataCollection converts projected type ResourceDataCollection to
// service type ResourceDataCollection.
func newResourceDataCollection(vres resourceviews.ResourceDataCollectionView) ResourceDataCollection {
	res := make(ResourceDataCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceData(n)
	}
	return res
}

// newResourceDataCollectionViewInfo projects result type
// ResourceDataCollection to projected type ResourceDataCollectionView using
// the "info" view.
func newResourceDataCollectionViewInfo(res ResourceDataCollection) resourceviews.ResourceDataCollectionView {
	vres := make(resourceviews.ResourceDataCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceDataViewInfo(n)
	}
	return vres
}

// newResourceDataCollectionViewWithoutVersion projects result type
// ResourceDataCollection to projected type ResourceDataCollectionView using
// the "withoutVersion" view.
func newResourceDataCollectionViewWithoutVersion(res ResourceDataCollection) resourceviews.ResourceDataCollectionView {
	vres := make(resourceviews.ResourceDataCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceDataViewWithoutVersion(n)
	}
	return vres
}

// newResourceDataCollectionView projects result type ResourceDataCollection to
// projected type ResourceDataCollectionView using the "default" view.
func newResourceDataCollectionView(res ResourceDataCollection) resourceviews.ResourceDataCollectionView {
	vres := make(resourceviews.ResourceDataCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceDataView(n)
	}
	return vres
}

// newResourceDataInfo converts projected type ResourceData to service type
// ResourceData.
func newResourceDataInfo(vres *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Catalog != nil {
		res.Catalog = newCatalogMin(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = newResourceVersionData(vres.LatestVersion)
	}
	return res
}

// newResourceDataWithoutVersion converts projected type ResourceData to
// service type ResourceData.
func newResourceDataWithoutVersion(vres *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Categories != nil {
		res.Categories = make([]*Category, len(vres.Categories))
		for i, val := range vres.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Catalog != nil {
		res.Catalog = newCatalogMin(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = newResourceVersionDataWithoutResource(vres.LatestVersion)
	}
	return res
}

// newResourceData converts projected type ResourceData to service type
// ResourceData.
func newResourceData(vres *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Categories != nil {
		res.Categories = make([]*Category, len(vres.Categories))
		for i, val := range vres.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Versions != nil {
		res.Versions = make([]*ResourceVersionData, len(vres.Versions))
		for i, val := range vres.Versions {
			res.Versions[i] = transformResourceviewsResourceVersionDataViewToResourceVersionData(val)
		}
	}
	if vres.Catalog != nil {
		res.Catalog = newCatalogMin(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = newResourceVersionDataWithoutResource(vres.LatestVersion)
	}
	return res
}

// newResourceDataViewInfo projects result type ResourceData to projected type
// ResourceDataView using the "info" view.
func newResourceDataViewInfo(res *ResourceData) *resourceviews.ResourceDataView {
	vres := &resourceviews.ResourceDataView{
		ID:     &res.ID,
		Name:   &res.Name,
		Kind:   &res.Kind,
		Rating: &res.Rating,
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Catalog != nil {
		vres.Catalog = newCatalogViewMin(res.Catalog)
	}
	return vres
}

// newResourceDataViewWithoutVersion projects result type ResourceData to
// projected type ResourceDataView using the "withoutVersion" view.
func newResourceDataViewWithoutVersion(res *ResourceData) *resourceviews.ResourceDataView {
	vres := &resourceviews.ResourceDataView{
		ID:     &res.ID,
		Name:   &res.Name,
		Kind:   &res.Kind,
		Rating: &res.Rating,
	}
	if res.Categories != nil {
		vres.Categories = make([]*resourceviews.CategoryView, len(res.Categories))
		for i, val := range res.Categories {
			vres.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Catalog != nil {
		vres.Catalog = newCatalogViewMin(res.Catalog)
	}
	if res.LatestVersion != nil {
		vres.LatestVersion = newResourceVersionDataViewWithoutResource(res.LatestVersion)
	}
	return vres
}

// newResourceDataView projects result type ResourceData to projected type
// ResourceDataView using the "default" view.
func newResourceDataView(res *ResourceData) *resourceviews.ResourceDataView {
	vres := &resourceviews.ResourceDataView{
		ID:     &res.ID,
		Name:   &res.Name,
		Kind:   &res.Kind,
		Rating: &res.Rating,
	}
	if res.Categories != nil {
		vres.Categories = make([]*resourceviews.CategoryView, len(res.Categories))
		for i, val := range res.Categories {
			vres.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Versions != nil {
		vres.Versions = make([]*resourceviews.ResourceVersionDataView, len(res.Versions))
		for i, val := range res.Versions {
			vres.Versions[i] = transformResourceVersionDataToResourceviewsResourceVersionDataView(val)
		}
	}
	if res.Catalog != nil {
		vres.Catalog = newCatalogViewMin(res.Catalog)
	}
	if res.LatestVersion != nil {
		vres.LatestVersion = newResourceVersionDataViewWithoutResource(res.LatestVersion)
	}
	return vres
}

// newCatalogMin converts projected type Catalog to service type Catalog.
func newCatalogMin(vres *resourceviews.CatalogView) *Catalog {
	res := &Catalog{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	return res
}

// newCatalog converts projected type Catalog to service type Catalog.
func newCatalog(vres *resourceviews.CatalogView) *Catalog {
	res := &Catalog{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.URL != nil {
		res.URL = *vres.URL
	}
	return res
}

// newCatalogViewMin projects result type Catalog to projected type CatalogView
// using the "min" view.
func newCatalogViewMin(res *Catalog) *resourceviews.CatalogView {
	vres := &resourceviews.CatalogView{
		ID:   &res.ID,
		Name: &res.Name,
		Type: &res.Type,
	}
	return vres
}

// newCatalogView projects result type Catalog to projected type CatalogView
// using the "default" view.
func newCatalogView(res *Catalog) *resourceviews.CatalogView {
	vres := &resourceviews.CatalogView{
		ID:   &res.ID,
		Name: &res.Name,
		Type: &res.Type,
		URL:  &res.URL,
	}
	return vres
}

// newResourceVersionDataTiny converts projected type ResourceVersionData to
// service type ResourceVersionData.
func newResourceVersionDataTiny(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.Resource != nil {
		res.Resource = newResourceData(vres.Resource)
	}
	return res
}

// newResourceVersionDataMin converts projected type ResourceVersionData to
// service type ResourceVersionData.
func newResourceVersionDataMin(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.Resource != nil {
		res.Resource = newResourceData(vres.Resource)
	}
	return res
}

// newResourceVersionDataWithoutResource converts projected type
// ResourceVersionData to service type ResourceVersionData.
func newResourceVersionDataWithoutResource(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.MinPipelinesVersion != nil {
		res.MinPipelinesVersion = *vres.MinPipelinesVersion
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Resource != nil {
		res.Resource = newResourceData(vres.Resource)
	}
	return res
}

// newResourceVersionData converts projected type ResourceVersionData to
// service type ResourceVersionData.
func newResourceVersionData(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.MinPipelinesVersion != nil {
		res.MinPipelinesVersion = *vres.MinPipelinesVersion
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Resource != nil {
		res.Resource = newResourceDataInfo(vres.Resource)
	}
	return res
}

// newResourceVersionDataViewTiny projects result type ResourceVersionData to
// projected type ResourceVersionDataView using the "tiny" view.
func newResourceVersionDataViewTiny(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:      &res.ID,
		Version: &res.Version,
	}
	return vres
}

// newResourceVersionDataViewMin projects result type ResourceVersionData to
// projected type ResourceVersionDataView using the "min" view.
func newResourceVersionDataViewMin(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:      &res.ID,
		Version: &res.Version,
		RawURL:  &res.RawURL,
		WebURL:  &res.WebURL,
	}
	return vres
}

// newResourceVersionDataViewWithoutResource projects result type
// ResourceVersionData to projected type ResourceVersionDataView using the
// "withoutResource" view.
func newResourceVersionDataViewWithoutResource(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:                  &res.ID,
		Version:             &res.Version,
		DisplayName:         &res.DisplayName,
		Description:         &res.Description,
		MinPipelinesVersion: &res.MinPipelinesVersion,
		RawURL:              &res.RawURL,
		WebURL:              &res.WebURL,
		UpdatedAt:           &res.UpdatedAt,
	}
	return vres
}

// newResourceVersionDataView projects result type ResourceVersionData to
// projected type ResourceVersionDataView using the "default" view.
func newResourceVersionDataView(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:                  &res.ID,
		Version:             &res.Version,
		DisplayName:         &res.DisplayName,
		Description:         &res.Description,
		MinPipelinesVersion: &res.MinPipelinesVersion,
		RawURL:              &res.RawURL,
		WebURL:              &res.WebURL,
		UpdatedAt:           &res.UpdatedAt,
	}
	if res.Resource != nil {
		vres.Resource = newResourceDataViewInfo(res.Resource)
	}
	return vres
}

// newResourceVersions converts projected type ResourceVersions to service type
// ResourceVersions.
func newResourceVersions(vres *resourceviews.ResourceVersionsView) *ResourceVersions {
	res := &ResourceVersions{}
	if vres.Data != nil {
		res.Data = newVersions(vres.Data)
	}
	return res
}

// newResourceVersionsView projects result type ResourceVersions to projected
// type ResourceVersionsView using the "default" view.
func newResourceVersionsView(res *ResourceVersions) *resourceviews.ResourceVersionsView {
	vres := &resourceviews.ResourceVersionsView{}
	if res.Data != nil {
		vres.Data = newVersionsView(res.Data)
	}
	return vres
}

// newVersions converts projected type Versions to service type Versions.
func newVersions(vres *resourceviews.VersionsView) *Versions {
	res := &Versions{}
	if vres.Versions != nil {
		res.Versions = make([]*ResourceVersionData, len(vres.Versions))
		for i, val := range vres.Versions {
			res.Versions[i] = transformResourceviewsResourceVersionDataViewToResourceVersionData(val)
		}
	}
	if vres.Latest != nil {
		res.Latest = newResourceVersionDataMin(vres.Latest)
	}
	return res
}

// newVersionsView projects result type Versions to projected type VersionsView
// using the "default" view.
func newVersionsView(res *Versions) *resourceviews.VersionsView {
	vres := &resourceviews.VersionsView{}
	if res.Versions != nil {
		vres.Versions = make([]*resourceviews.ResourceVersionDataView, len(res.Versions))
		for i, val := range res.Versions {
			vres.Versions[i] = transformResourceVersionDataToResourceviewsResourceVersionDataView(val)
		}
	}
	if res.Latest != nil {
		vres.Latest = newResourceVersionDataViewMin(res.Latest)
	}
	return vres
}

// newResourceVersion converts projected type ResourceVersion to service type
// ResourceVersion.
func newResourceVersion(vres *resourceviews.ResourceVersionView) *ResourceVersion {
	res := &ResourceVersion{}
	if vres.Data != nil {
		res.Data = newResourceVersionData(vres.Data)
	}
	return res
}

// newResourceVersionView projects result type ResourceVersion to projected
// type ResourceVersionView using the "default" view.
func newResourceVersionView(res *ResourceVersion) *resourceviews.ResourceVersionView {
	vres := &resourceviews.ResourceVersionView{}
	if res.Data != nil {
		vres.Data = newResourceVersionDataView(res.Data)
	}
	return vres
}

// newResource converts projected type Resource to service type Resource.
func newResource(vres *resourceviews.ResourceView) *Resource {
	res := &Resource{}
	if vres.Data != nil {
		res.Data = newResourceData(vres.Data)
	}
	return res
}

// newResourceView projects result type Resource to projected type ResourceView
// using the "default" view.
func newResourceView(res *Resource) *resourceviews.ResourceView {
	vres := &resourceviews.ResourceView{}
	if res.Data != nil {
		vres.Data = newResourceDataView(res.Data)
	}
	return vres
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

// transformResourceviewsCategoryViewToCategory builds a value of type
// *Category from a value of type *resourceviews.CategoryView.
func transformResourceviewsCategoryViewToCategory(v *resourceviews.CategoryView) *Category {
	if v == nil {
		return nil
	}
	res := &Category{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformResourceviewsResourceVersionDataViewToResourceVersionData builds a
// value of type *ResourceVersionData from a value of type
// *resourceviews.ResourceVersionDataView.
func transformResourceviewsResourceVersionDataViewToResourceVersionData(v *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	if v == nil {
		return nil
	}
	res := &ResourceVersionData{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
	}
	if v.Resource != nil {
		res.Resource = transformResourceviewsResourceDataViewToResourceData(v.Resource)
	}

	return res
}

// transformResourceviewsResourceDataViewToResourceData builds a value of type
// *ResourceData from a value of type *resourceviews.ResourceDataView.
func transformResourceviewsResourceDataViewToResourceData(v *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if v.ID != nil {
		res.ID = *v.ID
	}
	if v.Name != nil {
		res.Name = *v.Name
	}
	if v.Kind != nil {
		res.Kind = *v.Kind
	}
	if v.Rating != nil {
		res.Rating = *v.Rating
	}
	if v.Categories != nil {
		res.Categories = make([]*Category, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*Tag, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if v.Versions != nil {
		res.Versions = make([]*ResourceVersionData, len(v.Versions))
		for i, val := range v.Versions {
			res.Versions[i] = transformResourceviewsResourceVersionDataViewToResourceVersionData(val)
		}
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

// transformCategoryToResourceviewsCategoryView builds a value of type
// *resourceviews.CategoryView from a value of type *Category.
func transformCategoryToResourceviewsCategoryView(v *Category) *resourceviews.CategoryView {
	res := &resourceviews.CategoryView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformResourceVersionDataToResourceviewsResourceVersionDataView builds a
// value of type *resourceviews.ResourceVersionDataView from a value of type
// *ResourceVersionData.
func transformResourceVersionDataToResourceviewsResourceVersionDataView(v *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	res := &resourceviews.ResourceVersionDataView{
		ID:                  &v.ID,
		Version:             &v.Version,
		DisplayName:         &v.DisplayName,
		Description:         &v.Description,
		MinPipelinesVersion: &v.MinPipelinesVersion,
		RawURL:              &v.RawURL,
		WebURL:              &v.WebURL,
		UpdatedAt:           &v.UpdatedAt,
	}
	if v.Resource != nil {
		res.Resource = transformResourceDataToResourceviewsResourceDataView(v.Resource)
	}

	return res
}

// transformResourceDataToResourceviewsResourceDataView builds a value of type
// *resourceviews.ResourceDataView from a value of type *ResourceData.
func transformResourceDataToResourceviewsResourceDataView(v *ResourceData) *resourceviews.ResourceDataView {
	res := &resourceviews.ResourceDataView{
		ID:     &v.ID,
		Name:   &v.Name,
		Kind:   &v.Kind,
		Rating: &v.Rating,
	}
	if v.Categories != nil {
		res.Categories = make([]*resourceviews.CategoryView, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*resourceviews.TagView, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if v.Versions != nil {
		res.Versions = make([]*resourceviews.ResourceVersionDataView, len(v.Versions))
		for i, val := range v.Versions {
			res.Versions[i] = transformResourceVersionDataToResourceviewsResourceVersionDataView(val)
		}
	}

	return res
}

// transformResourceviewsCatalogViewToCatalog builds a value of type *Catalog
// from a value of type *resourceviews.CatalogView.
func transformResourceviewsCatalogViewToCatalog(v *resourceviews.CatalogView) *Catalog {
	res := &Catalog{
		ID:   *v.ID,
		Name: *v.Name,
		Type: *v.Type,
		URL:  *v.URL,
	}

	return res
}

// transformCatalogToResourceviewsCatalogView builds a value of type
// *resourceviews.CatalogView from a value of type *Catalog.
func transformCatalogToResourceviewsCatalogView(v *Catalog) *resourceviews.CatalogView {
	res := &resourceviews.CatalogView{
		ID:   &v.ID,
		Name: &v.Name,
		Type: &v.Type,
		URL:  &v.URL,
	}

	return res
}
