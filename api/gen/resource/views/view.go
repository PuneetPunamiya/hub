// Code generated by goa v3.3.1, DO NOT EDIT.
//
// resource views
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Resources is the viewed result type that is projected based on a view.
type Resources struct {
	// Type to project
	Projected *ResourcesView
	// View to render
	View string
}

// ResourceVersions is the viewed result type that is projected based on a view.
type ResourceVersions struct {
	// Type to project
	Projected *ResourceVersionsView
	// View to render
	View string
}

// ResourceVersion is the viewed result type that is projected based on a view.
type ResourceVersion struct {
	// Type to project
	Projected *ResourceVersionView
	// View to render
	View string
}

// Resource is the viewed result type that is projected based on a view.
type Resource struct {
	// Type to project
	Projected *ResourceView
	// View to render
	View string
}

// ResourcesView is a type that runs validations on a projected type.
type ResourcesView struct {
	Data ResourceDataCollectionView
}

// ResourceDataCollectionView is a type that runs validations on a projected
// type.
type ResourceDataCollectionView []*ResourceDataView

// ResourceDataView is a type that runs validations on a projected type.
type ResourceDataView struct {
	// ID is the unique id of the resource
	ID *uint
	// Name of resource
	Name *string
	// Type of catalog to which resource belongs
	Catalog *CatalogView
	// Categories related to resource
	Categories []*CategoryView
	// Kind of resource
	Kind *string
	// Latest version of resource
	LatestVersion *ResourceVersionDataView
	// Tags related to resource
	Tags []*TagView
	// Rating of resource
	Rating *float64
	// List of all versions of a resource
	Versions []*ResourceVersionDataView
}

// CatalogView is a type that runs validations on a projected type.
type CatalogView struct {
	// ID is the unique id of the catalog
	ID *uint
	// Name of catalog
	Name *string
	// Type of catalog
	Type *string
	// URL of catalog
	URL *string
}

// CategoryView is a type that runs validations on a projected type.
type CategoryView struct {
	// ID is the unique id of the category
	ID *uint
	// Name of category
	Name *string
}

// ResourceVersionDataView is a type that runs validations on a projected type.
type ResourceVersionDataView struct {
	// ID is the unique id of resource's version
	ID *uint
	// Version of resource
	Version *string
	// Display name of version
	DisplayName *string
	// Description of version
	Description *string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion *string
	// Raw URL of resource's yaml file of the version
	RawURL *string
	// Web URL of resource's yaml file of the version
	WebURL *string
	// Timestamp when version was last updated
	UpdatedAt *string
	// Resource to which the version belongs
	Resource *ResourceDataView
}

// TagView is a type that runs validations on a projected type.
type TagView struct {
	// ID is the unique id of tag
	ID *uint
	// Name of tag
	Name *string
}

// ResourceVersionsView is a type that runs validations on a projected type.
type ResourceVersionsView struct {
	Data *VersionsView
}

// VersionsView is a type that runs validations on a projected type.
type VersionsView struct {
	// Latest Version of resource
	Latest *ResourceVersionDataView
	// List of all versions of resource
	Versions []*ResourceVersionDataView
}

// ResourceVersionView is a type that runs validations on a projected type.
type ResourceVersionView struct {
	Data *ResourceVersionDataView
}

// ResourceView is a type that runs validations on a projected type.
type ResourceView struct {
	Data *ResourceDataView
}

var (
	// ResourcesMap is a map of attribute names in result type Resources indexed by
	// view name.
	ResourcesMap = map[string][]string{
		"default": []string{
			"data",
		},
	}
	// ResourceVersionsMap is a map of attribute names in result type
	// ResourceVersions indexed by view name.
	ResourceVersionsMap = map[string][]string{
		"default": []string{
			"data",
		},
	}
	// ResourceVersionMap is a map of attribute names in result type
	// ResourceVersion indexed by view name.
	ResourceVersionMap = map[string][]string{
		"default": []string{
			"data",
		},
	}
	// ResourceMap is a map of attribute names in result type Resource indexed by
	// view name.
	ResourceMap = map[string][]string{
		"default": []string{
			"data",
		},
	}
	// ResourceDataCollectionMap is a map of attribute names in result type
	// ResourceDataCollection indexed by view name.
	ResourceDataCollectionMap = map[string][]string{
		"info": []string{
			"id",
			"name",
			"catalog",
			"kind",
			"tags",
			"rating",
		},
		"withoutVersion": []string{
			"id",
			"name",
			"catalog",
			"categories",
			"kind",
			"latestVersion",
			"tags",
			"rating",
		},
		"default": []string{
			"id",
			"name",
			"catalog",
			"categories",
			"kind",
			"latestVersion",
			"tags",
			"rating",
			"versions",
		},
	}
	// ResourceDataMap is a map of attribute names in result type ResourceData
	// indexed by view name.
	ResourceDataMap = map[string][]string{
		"info": []string{
			"id",
			"name",
			"catalog",
			"kind",
			"tags",
			"rating",
		},
		"withoutVersion": []string{
			"id",
			"name",
			"catalog",
			"categories",
			"kind",
			"latestVersion",
			"tags",
			"rating",
		},
		"default": []string{
			"id",
			"name",
			"catalog",
			"categories",
			"kind",
			"latestVersion",
			"tags",
			"rating",
			"versions",
		},
	}
	// CatalogMap is a map of attribute names in result type Catalog indexed by
	// view name.
	CatalogMap = map[string][]string{
		"min": []string{
			"id",
			"name",
			"type",
		},
		"default": []string{
			"id",
			"name",
			"type",
			"url",
		},
	}
	// ResourceVersionDataMap is a map of attribute names in result type
	// ResourceVersionData indexed by view name.
	ResourceVersionDataMap = map[string][]string{
		"tiny": []string{
			"id",
			"version",
		},
		"min": []string{
			"id",
			"version",
			"rawURL",
			"webURL",
		},
		"withoutResource": []string{
			"id",
			"version",
			"displayName",
			"description",
			"minPipelinesVersion",
			"rawURL",
			"webURL",
			"updatedAt",
		},
		"default": []string{
			"id",
			"version",
			"displayName",
			"description",
			"minPipelinesVersion",
			"rawURL",
			"webURL",
			"updatedAt",
			"resource",
		},
	}
	// VersionsMap is a map of attribute names in result type Versions indexed by
	// view name.
	VersionsMap = map[string][]string{
		"default": []string{
			"latest",
			"versions",
		},
	}
)

// ValidateResources runs the validations defined on the viewed result type
// Resources.
func ValidateResources(result *Resources) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResourcesView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateResourceVersions runs the validations defined on the viewed result
// type ResourceVersions.
func ValidateResourceVersions(result *ResourceVersions) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResourceVersionsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateResourceVersion runs the validations defined on the viewed result
// type ResourceVersion.
func ValidateResourceVersion(result *ResourceVersion) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResourceVersionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateResource runs the validations defined on the viewed result type
// Resource.
func ValidateResource(result *Resource) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResourceView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateResourcesView runs the validations defined on ResourcesView using
// the "default" view.
func ValidateResourcesView(result *ResourcesView) (err error) {

	if result.Data != nil {
		if err2 := ValidateResourceDataCollectionViewWithoutVersion(result.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceDataCollectionViewInfo runs the validations defined on
// ResourceDataCollectionView using the "info" view.
func ValidateResourceDataCollectionViewInfo(result ResourceDataCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateResourceDataViewInfo(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceDataCollectionViewWithoutVersion runs the validations
// defined on ResourceDataCollectionView using the "withoutVersion" view.
func ValidateResourceDataCollectionViewWithoutVersion(result ResourceDataCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateResourceDataViewWithoutVersion(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceDataCollectionView runs the validations defined on
// ResourceDataCollectionView using the "default" view.
func ValidateResourceDataCollectionView(result ResourceDataCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateResourceDataView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceDataViewInfo runs the validations defined on
// ResourceDataView using the "info" view.
func ValidateResourceDataViewInfo(result *ResourceDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "result"))
	}
	if result.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "result"))
	}
	for _, e := range result.Tags {
		if e != nil {
			if err2 := ValidateTagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Catalog != nil {
		if err2 := ValidateCatalogViewMin(result.Catalog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceDataViewWithoutVersion runs the validations defined on
// ResourceDataView using the "withoutVersion" view.
func ValidateResourceDataViewWithoutVersion(result *ResourceDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Categories == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("categories", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "result"))
	}
	if result.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "result"))
	}
	for _, e := range result.Categories {
		if e != nil {
			if err2 := ValidateCategoryView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Tags {
		if e != nil {
			if err2 := ValidateTagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Catalog != nil {
		if err2 := ValidateCatalogViewMin(result.Catalog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.LatestVersion != nil {
		if err2 := ValidateResourceVersionDataViewWithoutResource(result.LatestVersion); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceDataView runs the validations defined on ResourceDataView
// using the "default" view.
func ValidateResourceDataView(result *ResourceDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Categories == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("categories", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "result"))
	}
	if result.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "result"))
	}
	if result.Versions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("versions", "result"))
	}
	for _, e := range result.Categories {
		if e != nil {
			if err2 := ValidateCategoryView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Tags {
		if e != nil {
			if err2 := ValidateTagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Versions {
		if e != nil {
			if err2 := ValidateResourceVersionDataView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Catalog != nil {
		if err2 := ValidateCatalogViewMin(result.Catalog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.LatestVersion != nil {
		if err2 := ValidateResourceVersionDataViewWithoutResource(result.LatestVersion); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCatalogViewMin runs the validations defined on CatalogView using the
// "min" view.
func ValidateCatalogViewMin(result *CatalogView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "official" || *result.Type == "community") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"official", "community"}))
		}
	}
	return
}

// ValidateCatalogView runs the validations defined on CatalogView using the
// "default" view.
func ValidateCatalogView(result *CatalogView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "official" || *result.Type == "community") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"official", "community"}))
		}
	}
	return
}

// ValidateCategoryView runs the validations defined on CategoryView.
func ValidateCategoryView(result *CategoryView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateResourceVersionDataViewTiny runs the validations defined on
// ResourceVersionDataView using the "tiny" view.
func ValidateResourceVersionDataViewTiny(result *ResourceVersionDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	return
}

// ValidateResourceVersionDataViewMin runs the validations defined on
// ResourceVersionDataView using the "min" view.
func ValidateResourceVersionDataViewMin(result *ResourceVersionDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	if result.RawURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rawURL", "result"))
	}
	if result.WebURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("webURL", "result"))
	}
	if result.RawURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.rawURL", *result.RawURL, goa.FormatURI))
	}
	if result.WebURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.webURL", *result.WebURL, goa.FormatURI))
	}
	return
}

// ValidateResourceVersionDataViewWithoutResource runs the validations defined
// on ResourceVersionDataView using the "withoutResource" view.
func ValidateResourceVersionDataViewWithoutResource(result *ResourceVersionDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	if result.DisplayName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("displayName", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.MinPipelinesVersion == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("minPipelinesVersion", "result"))
	}
	if result.RawURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rawURL", "result"))
	}
	if result.WebURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("webURL", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updatedAt", "result"))
	}
	if result.RawURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.rawURL", *result.RawURL, goa.FormatURI))
	}
	if result.WebURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.webURL", *result.WebURL, goa.FormatURI))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updatedAt", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}

// ValidateResourceVersionDataView runs the validations defined on
// ResourceVersionDataView using the "default" view.
func ValidateResourceVersionDataView(result *ResourceVersionDataView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	if result.DisplayName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("displayName", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.MinPipelinesVersion == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("minPipelinesVersion", "result"))
	}
	if result.RawURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rawURL", "result"))
	}
	if result.WebURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("webURL", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updatedAt", "result"))
	}
	if result.RawURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.rawURL", *result.RawURL, goa.FormatURI))
	}
	if result.WebURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.webURL", *result.WebURL, goa.FormatURI))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updatedAt", *result.UpdatedAt, goa.FormatDateTime))
	}
	if result.Resource != nil {
		if err2 := ValidateResourceDataViewInfo(result.Resource); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTagView runs the validations defined on TagView.
func ValidateTagView(result *TagView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateResourceVersionsView runs the validations defined on
// ResourceVersionsView using the "default" view.
func ValidateResourceVersionsView(result *ResourceVersionsView) (err error) {

	if result.Data != nil {
		if err2 := ValidateVersionsView(result.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVersionsView runs the validations defined on VersionsView using the
// "default" view.
func ValidateVersionsView(result *VersionsView) (err error) {
	if result.Versions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("versions", "result"))
	}
	for _, e := range result.Versions {
		if e != nil {
			if err2 := ValidateResourceVersionDataView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Latest != nil {
		if err2 := ValidateResourceVersionDataViewMin(result.Latest); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceVersionView runs the validations defined on
// ResourceVersionView using the "default" view.
func ValidateResourceVersionView(result *ResourceVersionView) (err error) {

	if result.Data != nil {
		if err2 := ValidateResourceVersionDataView(result.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceView runs the validations defined on ResourceView using the
// "default" view.
func ValidateResourceView(result *ResourceView) (err error) {

	if result.Data != nil {
		if err2 := ValidateResourceDataView(result.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
