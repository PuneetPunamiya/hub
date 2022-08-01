// Code generated by goa v3.7.14, DO NOT EDIT.
//
// resource client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package resource

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "resource" service client.
type Client struct {
	QueryEndpoint                              goa.Endpoint
	ListEndpoint                               goa.Endpoint
	VersionsByIDEndpoint                       goa.Endpoint
	ByCatalogKindNameVersionEndpoint           goa.Endpoint
	ByCatalogKindNameVersionReadmeEndpoint     goa.Endpoint
	ByCatalogKindNameVersionYamlEndpoint       goa.Endpoint
	ByVersionIDEndpoint                        goa.Endpoint
	ByCatalogKindNameEndpoint                  goa.Endpoint
	ByIDEndpoint                               goa.Endpoint
	GetRawYamlByCatalogKindNameVersionEndpoint goa.Endpoint
}

// NewClient initializes a "resource" service client given the endpoints.
func NewClient(query, list, versionsByID, byCatalogKindNameVersion, byCatalogKindNameVersionReadme, byCatalogKindNameVersionYaml, byVersionID, byCatalogKindName, byID, getRawYamlByCatalogKindNameVersion goa.Endpoint) *Client {
	return &Client{
		QueryEndpoint:                              query,
		ListEndpoint:                               list,
		VersionsByIDEndpoint:                       versionsByID,
		ByCatalogKindNameVersionEndpoint:           byCatalogKindNameVersion,
		ByCatalogKindNameVersionReadmeEndpoint:     byCatalogKindNameVersionReadme,
		ByCatalogKindNameVersionYamlEndpoint:       byCatalogKindNameVersionYaml,
		ByVersionIDEndpoint:                        byVersionID,
		ByCatalogKindNameEndpoint:                  byCatalogKindName,
		ByIDEndpoint:                               byID,
		GetRawYamlByCatalogKindNameVersionEndpoint: getRawYamlByCatalogKindNameVersion,
	}
}

// Query calls the "Query" endpoint of the "resource" service.
func (c *Client) Query(ctx context.Context, p *QueryPayload) (res *Resources, err error) {
	var ires interface{}
	ires, err = c.QueryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resources), nil
}

// List calls the "List" endpoint of the "resource" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res *Resources, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resources), nil
}

// VersionsByID calls the "VersionsByID" endpoint of the "resource" service.
func (c *Client) VersionsByID(ctx context.Context, p *VersionsByIDPayload) (res *ResourceVersions, err error) {
	var ires interface{}
	ires, err = c.VersionsByIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersions), nil
}

// ByCatalogKindNameVersion calls the "ByCatalogKindNameVersion" endpoint of
// the "resource" service.
func (c *Client) ByCatalogKindNameVersion(ctx context.Context, p *ByCatalogKindNameVersionPayload) (res *ResourceVersion, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersion), nil
}

// ByCatalogKindNameVersionReadme calls the "ByCatalogKindNameVersionReadme"
// endpoint of the "resource" service.
func (c *Client) ByCatalogKindNameVersionReadme(ctx context.Context, p *ByCatalogKindNameVersionReadmePayload) (res *ResourceVersionReadme, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionReadmeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersionReadme), nil
}

// ByCatalogKindNameVersionYaml calls the "ByCatalogKindNameVersionYaml"
// endpoint of the "resource" service.
func (c *Client) ByCatalogKindNameVersionYaml(ctx context.Context, p *ByCatalogKindNameVersionYamlPayload) (res *ResourceVersionYaml, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionYamlEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersionYaml), nil
}

// ByVersionID calls the "ByVersionId" endpoint of the "resource" service.
func (c *Client) ByVersionID(ctx context.Context, p *ByVersionIDPayload) (res *ResourceVersion, err error) {
	var ires interface{}
	ires, err = c.ByVersionIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersion), nil
}

// ByCatalogKindName calls the "ByCatalogKindName" endpoint of the "resource"
// service.
func (c *Client) ByCatalogKindName(ctx context.Context, p *ByCatalogKindNamePayload) (res *Resource, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resource), nil
}

// ByID calls the "ById" endpoint of the "resource" service.
func (c *Client) ByID(ctx context.Context, p *ByIDPayload) (res *Resource, err error) {
	var ires interface{}
	ires, err = c.ByIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resource), nil
}

// GetRawYamlByCatalogKindNameVersion calls the
// "GetRawYamlByCatalogKindNameVersion" endpoint of the "resource" service.
func (c *Client) GetRawYamlByCatalogKindNameVersion(ctx context.Context, p *GetRawYamlByCatalogKindNameVersionPayload) (resp io.ReadCloser, err error) {
	var ires interface{}
	ires, err = c.GetRawYamlByCatalogKindNameVersionEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*GetRawYamlByCatalogKindNameVersionResponseData)
	return o.Body, nil
}
