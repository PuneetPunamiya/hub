// Code generated by goa v3.11.1, DO NOT EDIT.
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
// Query may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
func (c *Client) Query(ctx context.Context, p *QueryPayload) (res *Resources, err error) {
	var ires interface{}
	ires, err = c.QueryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resources), nil
}

// List calls the "List" endpoint of the "resource" service.
// List may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
func (c *Client) List(ctx context.Context, p *ListPayload) (res *Resources, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resources), nil
}

// VersionsByID calls the "VersionsByID" endpoint of the "resource" service.
// VersionsByID may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
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
// ByCatalogKindNameVersion may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
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
// ByCatalogKindNameVersionReadme may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
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
// ByCatalogKindNameVersionYaml may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
func (c *Client) ByCatalogKindNameVersionYaml(ctx context.Context, p *ByCatalogKindNameVersionYamlPayload) (res *ResourceVersionYaml, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameVersionYamlEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersionYaml), nil
}

// ByVersionID calls the "ByVersionId" endpoint of the "resource" service.
// ByVersionID may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
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
// ByCatalogKindName may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
func (c *Client) ByCatalogKindName(ctx context.Context, p *ByCatalogKindNamePayload) (res *Resource, err error) {
	var ires interface{}
	ires, err = c.ByCatalogKindNameEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resource), nil
}

// ByID calls the "ById" endpoint of the "resource" service.
// ByID may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
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
// GetRawYamlByCatalogKindNameVersion may return the following errors:
//	- "internal-error" (type *goa.ServiceError): Internal Server Error
//	- "not-found" (type *goa.ServiceError): Resource Not Found Error
//	- "invalid-kind" (type *goa.ServiceError): Invalid Resource Kind
//	- error: internal error
func (c *Client) GetRawYamlByCatalogKindNameVersion(ctx context.Context, p *GetRawYamlByCatalogKindNameVersionPayload) (resp io.ReadCloser, err error) {
	var ires interface{}
	ires, err = c.GetRawYamlByCatalogKindNameVersionEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*GetRawYamlByCatalogKindNameVersionResponseData)
	return o.Body, nil
}
