// Code generated by goa v3.4.0, DO NOT EDIT.
//
// resource client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "resource" service client.
type Client struct {
	ListEndpoint                           goa.Endpoint
	VersionsByIDEndpoint                   goa.Endpoint
	ByCatalogKindNameVersionEndpoint       goa.Endpoint
	ByCatalogKindNameVersionReadmeEndpoint goa.Endpoint
	ByCatalogKindNameVersionYamlEndpoint   goa.Endpoint
	ByVersionIDEndpoint                    goa.Endpoint
	ByCatalogKindNameEndpoint              goa.Endpoint
}

// NewClient initializes a "resource" service client given the endpoints.
func NewClient(list, versionsByID, byCatalogKindNameVersion, byCatalogKindNameVersionReadme, byCatalogKindNameVersionYaml, byVersionID, byCatalogKindName goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:                           list,
		VersionsByIDEndpoint:                   versionsByID,
		ByCatalogKindNameVersionEndpoint:       byCatalogKindNameVersion,
		ByCatalogKindNameVersionReadmeEndpoint: byCatalogKindNameVersionReadme,
		ByCatalogKindNameVersionYamlEndpoint:   byCatalogKindNameVersionYaml,
		ByVersionIDEndpoint:                    byVersionID,
		ByCatalogKindNameEndpoint:              byCatalogKindName,
	}
}

// List calls the "List" endpoint of the "resource" service.
func (c *Client) List(ctx context.Context) (res *Resources, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Resources), nil
}

// VersionsByID calls the "VersionsByID" endpoint of the "resource" service.
func (c *Client) VersionsByID(ctx context.Context, p *VersionsByIDPayload) (res *VersionsByIDResult, err error) {
	var ires interface{}
	ires, err = c.VersionsByIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*VersionsByIDResult), nil
}

// ByCatalogKindNameVersion calls the "ByCatalogKindNameVersion" endpoint of
// the "resource" service.
func (c *Client) ByCatalogKindNameVersion(ctx context.Context, p *ByCatalogKindNameVersionPayload) (res *ByCatalogKindNameVersionResult, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ByCatalogKindNameVersionResult), nil
}

// ByCatalogKindNameVersionReadme calls the "ByCatalogKindNameVersionReadme"
// endpoint of the "resource" service.
func (c *Client) ByCatalogKindNameVersionReadme(ctx context.Context, p *ByCatalogKindNameVersionReadmePayload) (res *ByCatalogKindNameVersionReadmeResult, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionReadmeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ByCatalogKindNameVersionReadmeResult), nil
}

// ByCatalogKindNameVersionYaml calls the "ByCatalogKindNameVersionYaml"
// endpoint of the "resource" service.
func (c *Client) ByCatalogKindNameVersionYaml(ctx context.Context, p *ByCatalogKindNameVersionYamlPayload) (res *ByCatalogKindNameVersionYamlResult, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionYamlEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ByCatalogKindNameVersionYamlResult), nil
}

// ByVersionID calls the "ByVersionId" endpoint of the "resource" service.
func (c *Client) ByVersionID(ctx context.Context, p *ByVersionIDPayload) (res *ByVersionIDResult, err error) {
	var ires interface{}
	ires, err = c.ByVersionIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ByVersionIDResult), nil
}

// ByCatalogKindName calls the "ByCatalogKindName" endpoint of the "resource"
// service.
func (c *Client) ByCatalogKindName(ctx context.Context, p *ByCatalogKindNamePayload) (res *ByCatalogKindNameResult, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ByCatalogKindNameResult), nil
}
