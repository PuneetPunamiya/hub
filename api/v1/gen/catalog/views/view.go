// Code generated by goa v3.18.0, DO NOT EDIT.
//
// catalog views
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

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
	// Provider of catalog
	Provider *string
}

var (
	// CatalogMap is a map indexing the attribute names of Catalog by view name.
	CatalogMap = map[string][]string{
		"min": {
			"id",
			"name",
			"type",
		},
		"default": {
			"id",
			"name",
			"type",
			"url",
			"provider",
		},
	}
)

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
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []any{"official", "community"}))
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
	if result.Provider == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provider", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "official" || *result.Type == "community") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []any{"official", "community"}))
		}
	}
	return
}
