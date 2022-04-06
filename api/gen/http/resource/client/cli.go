// Code generated by goa v3.4.0, DO NOT EDIT.
//
// resource HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"fmt"
	"strconv"

	resource "github.com/tektoncd/hub/api/gen/resource"
	goa "goa.design/goa/v3/pkg"
)

// BuildVersionsByIDPayload builds the payload for the resource VersionsByID
// endpoint from CLI flags.
func BuildVersionsByIDPayload(resourceVersionsByIDID string) (*resource.VersionsByIDPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(resourceVersionsByIDID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &resource.VersionsByIDPayload{}
	v.ID = id

	return v, nil
}

// BuildByCatalogKindNameVersionPayload builds the payload for the resource
// ByCatalogKindNameVersion endpoint from CLI flags.
func BuildByCatalogKindNameVersionPayload(resourceByCatalogKindNameVersionCatalog string, resourceByCatalogKindNameVersionKind string, resourceByCatalogKindNameVersionName string, resourceByCatalogKindNameVersionVersion string) (*resource.ByCatalogKindNameVersionPayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameVersionCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameVersionKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []interface{}{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameVersionName
	}
	var version string
	{
		version = resourceByCatalogKindNameVersionVersion
	}
	v := &resource.ByCatalogKindNameVersionPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildByCatalogKindNameVersionReadmePayload builds the payload for the
// resource ByCatalogKindNameVersionReadme endpoint from CLI flags.
func BuildByCatalogKindNameVersionReadmePayload(resourceByCatalogKindNameVersionReadmeCatalog string, resourceByCatalogKindNameVersionReadmeKind string, resourceByCatalogKindNameVersionReadmeName string, resourceByCatalogKindNameVersionReadmeVersion string) (*resource.ByCatalogKindNameVersionReadmePayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameVersionReadmeCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameVersionReadmeKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []interface{}{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameVersionReadmeName
	}
	var version string
	{
		version = resourceByCatalogKindNameVersionReadmeVersion
	}
	v := &resource.ByCatalogKindNameVersionReadmePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildByCatalogKindNameVersionYamlPayload builds the payload for the resource
// ByCatalogKindNameVersionYaml endpoint from CLI flags.
func BuildByCatalogKindNameVersionYamlPayload(resourceByCatalogKindNameVersionYamlCatalog string, resourceByCatalogKindNameVersionYamlKind string, resourceByCatalogKindNameVersionYamlName string, resourceByCatalogKindNameVersionYamlVersion string) (*resource.ByCatalogKindNameVersionYamlPayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameVersionYamlCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameVersionYamlKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []interface{}{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameVersionYamlName
	}
	var version string
	{
		version = resourceByCatalogKindNameVersionYamlVersion
	}
	v := &resource.ByCatalogKindNameVersionYamlPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildByVersionIDPayload builds the payload for the resource ByVersionId
// endpoint from CLI flags.
func BuildByVersionIDPayload(resourceByVersionIDVersionID string) (*resource.ByVersionIDPayload, error) {
	var err error
	var versionID uint
	{
		var v uint64
		v, err = strconv.ParseUint(resourceByVersionIDVersionID, 10, 64)
		versionID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for versionID, must be UINT")
		}
	}
	v := &resource.ByVersionIDPayload{}
	v.VersionID = versionID

	return v, nil
}

// BuildByCatalogKindNamePayload builds the payload for the resource
// ByCatalogKindName endpoint from CLI flags.
func BuildByCatalogKindNamePayload(resourceByCatalogKindNameCatalog string, resourceByCatalogKindNameKind string, resourceByCatalogKindNameName string, resourceByCatalogKindNamePipelinesversion string) (*resource.ByCatalogKindNamePayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []interface{}{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameName
	}
	var pipelinesversion *string
	{
		if resourceByCatalogKindNamePipelinesversion != "" {
			pipelinesversion = &resourceByCatalogKindNamePipelinesversion
			if pipelinesversion != nil {
				err = goa.MergeErrors(err, goa.ValidatePattern("pipelinesversion", *pipelinesversion, "^\\d+(?:\\.\\d+){0,2}$"))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &resource.ByCatalogKindNamePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Pipelinesversion = pipelinesversion

	return v, nil
}
