// Code generated by goa v3.4.0, DO NOT EDIT.
//
// resource HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	resource "github.com/tektoncd/hub/api/gen/resource"
	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "resource" service "List" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListResourcePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("resource", "List", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the resource
// List endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusMovedPermanently:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("resource", "List", err)
			}
			p := NewListResourcesMovedPermanently(&body)
			view := "default"
			vres := &resourceviews.Resources{Projected: p, View: view}
			if err = resourceviews.ValidateResources(vres); err != nil {
				return nil, goahttp.ErrValidationError("resource", "List", err)
			}
			res := resource.NewResources(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("resource", "List", resp.StatusCode, string(body))
		}
	}
}

// BuildVersionsByIDRequest instantiates a HTTP request object with method and
// path set to call the "resource" service "VersionsByID" endpoint
func (c *Client) BuildVersionsByIDRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint
	)
	{
		p, ok := v.(*resource.VersionsByIDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("resource", "VersionsByID", "*resource.VersionsByIDPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VersionsByIDResourcePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("resource", "VersionsByID", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeVersionsByIDResponse returns a decoder for responses returned by the
// resource VersionsByID endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeVersionsByIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				location string
				err      error
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("location", "header"))
			}
			location = locationRaw
			err = goa.MergeErrors(err, goa.ValidateFormat("location", location, goa.FormatURI))

			if err != nil {
				return nil, goahttp.ErrValidationError("resource", "VersionsByID", err)
			}
			res := NewVersionsByIDResultFound(location)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("resource", "VersionsByID", resp.StatusCode, string(body))
		}
	}
}

// BuildByCatalogKindNameVersionRequest instantiates a HTTP request object with
// method and path set to call the "resource" service
// "ByCatalogKindNameVersion" endpoint
func (c *Client) BuildByCatalogKindNameVersionRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		catalog string
		kind    string
		name    string
		version string
	)
	{
		p, ok := v.(*resource.ByCatalogKindNameVersionPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("resource", "ByCatalogKindNameVersion", "*resource.ByCatalogKindNameVersionPayload", v)
		}
		catalog = p.Catalog
		kind = p.Kind
		name = p.Name
		version = p.Version
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ByCatalogKindNameVersionResourcePath(catalog, kind, name, version)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("resource", "ByCatalogKindNameVersion", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeByCatalogKindNameVersionResponse returns a decoder for responses
// returned by the resource ByCatalogKindNameVersion endpoint. restoreBody
// controls whether the response body should be restored after having been read.
func DecodeByCatalogKindNameVersionResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				location string
				err      error
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("location", "header"))
			}
			location = locationRaw
			err = goa.MergeErrors(err, goa.ValidateFormat("location", location, goa.FormatURI))

			if err != nil {
				return nil, goahttp.ErrValidationError("resource", "ByCatalogKindNameVersion", err)
			}
			res := NewByCatalogKindNameVersionResultFound(location)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("resource", "ByCatalogKindNameVersion", resp.StatusCode, string(body))
		}
	}
}

// BuildByCatalogKindNameVersionReadmeRequest instantiates a HTTP request
// object with method and path set to call the "resource" service
// "ByCatalogKindNameVersionReadme" endpoint
func (c *Client) BuildByCatalogKindNameVersionReadmeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		catalog string
		kind    string
		name    string
		version string
	)
	{
		p, ok := v.(*resource.ByCatalogKindNameVersionReadmePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("resource", "ByCatalogKindNameVersionReadme", "*resource.ByCatalogKindNameVersionReadmePayload", v)
		}
		catalog = p.Catalog
		kind = p.Kind
		name = p.Name
		version = p.Version
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ByCatalogKindNameVersionReadmeResourcePath(catalog, kind, name, version)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("resource", "ByCatalogKindNameVersionReadme", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeByCatalogKindNameVersionReadmeResponse returns a decoder for responses
// returned by the resource ByCatalogKindNameVersionReadme endpoint.
// restoreBody controls whether the response body should be restored after
// having been read.
func DecodeByCatalogKindNameVersionReadmeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				location string
				err      error
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("location", "header"))
			}
			location = locationRaw
			err = goa.MergeErrors(err, goa.ValidateFormat("location", location, goa.FormatURI))

			if err != nil {
				return nil, goahttp.ErrValidationError("resource", "ByCatalogKindNameVersionReadme", err)
			}
			res := NewByCatalogKindNameVersionReadmeResultFound(location)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("resource", "ByCatalogKindNameVersionReadme", resp.StatusCode, string(body))
		}
	}
}

// unmarshalResourceDataResponseBodyToResourceviewsResourceDataView builds a
// value of type *resourceviews.ResourceDataView from a value of type
// *ResourceDataResponseBody.
func unmarshalResourceDataResponseBodyToResourceviewsResourceDataView(v *ResourceDataResponseBody) *resourceviews.ResourceDataView {
	res := &resourceviews.ResourceDataView{
		ID:         v.ID,
		Name:       v.Name,
		Kind:       v.Kind,
		HubURLPath: v.HubURLPath,
		Rating:     v.Rating,
	}
	res.Catalog = unmarshalCatalogResponseBodyToResourceviewsCatalogView(v.Catalog)
	res.Categories = make([]*resourceviews.CategoryView, len(v.Categories))
	for i, val := range v.Categories {
		res.Categories[i] = unmarshalCategoryResponseBodyToResourceviewsCategoryView(val)
	}
	res.LatestVersion = unmarshalResourceVersionDataResponseBodyToResourceviewsResourceVersionDataView(v.LatestVersion)
	res.Tags = make([]*resourceviews.TagView, len(v.Tags))
	for i, val := range v.Tags {
		res.Tags[i] = unmarshalTagResponseBodyToResourceviewsTagView(val)
	}
	res.Platforms = make([]*resourceviews.PlatformView, len(v.Platforms))
	for i, val := range v.Platforms {
		res.Platforms[i] = unmarshalPlatformResponseBodyToResourceviewsPlatformView(val)
	}
	res.Versions = make([]*resourceviews.ResourceVersionDataView, len(v.Versions))
	for i, val := range v.Versions {
		res.Versions[i] = unmarshalResourceVersionDataResponseBodyToResourceviewsResourceVersionDataView(val)
	}

	return res
}

// unmarshalCatalogResponseBodyToResourceviewsCatalogView builds a value of
// type *resourceviews.CatalogView from a value of type *CatalogResponseBody.
func unmarshalCatalogResponseBodyToResourceviewsCatalogView(v *CatalogResponseBody) *resourceviews.CatalogView {
	res := &resourceviews.CatalogView{
		ID:       v.ID,
		Name:     v.Name,
		Type:     v.Type,
		URL:      v.URL,
		Provider: v.Provider,
	}

	return res
}

// unmarshalCategoryResponseBodyToResourceviewsCategoryView builds a value of
// type *resourceviews.CategoryView from a value of type *CategoryResponseBody.
func unmarshalCategoryResponseBodyToResourceviewsCategoryView(v *CategoryResponseBody) *resourceviews.CategoryView {
	res := &resourceviews.CategoryView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalResourceVersionDataResponseBodyToResourceviewsResourceVersionDataView
// builds a value of type *resourceviews.ResourceVersionDataView from a value
// of type *ResourceVersionDataResponseBody.
func unmarshalResourceVersionDataResponseBodyToResourceviewsResourceVersionDataView(v *ResourceVersionDataResponseBody) *resourceviews.ResourceVersionDataView {
	res := &resourceviews.ResourceVersionDataView{
		ID:                  v.ID,
		Version:             v.Version,
		DisplayName:         v.DisplayName,
		Deprecated:          v.Deprecated,
		Description:         v.Description,
		MinPipelinesVersion: v.MinPipelinesVersion,
		RawURL:              v.RawURL,
		WebURL:              v.WebURL,
		UpdatedAt:           v.UpdatedAt,
		HubURLPath:          v.HubURLPath,
	}
	res.Platforms = make([]*resourceviews.PlatformView, len(v.Platforms))
	for i, val := range v.Platforms {
		res.Platforms[i] = unmarshalPlatformResponseBodyToResourceviewsPlatformView(val)
	}
	res.Resource = unmarshalResourceDataResponseBodyToResourceviewsResourceDataView(v.Resource)

	return res
}

// unmarshalPlatformResponseBodyToResourceviewsPlatformView builds a value of
// type *resourceviews.PlatformView from a value of type *PlatformResponseBody.
func unmarshalPlatformResponseBodyToResourceviewsPlatformView(v *PlatformResponseBody) *resourceviews.PlatformView {
	res := &resourceviews.PlatformView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalTagResponseBodyToResourceviewsTagView builds a value of type
// *resourceviews.TagView from a value of type *TagResponseBody.
func unmarshalTagResponseBodyToResourceviewsTagView(v *TagResponseBody) *resourceviews.TagView {
	res := &resourceviews.TagView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}
