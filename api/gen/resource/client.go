// Code generated by goa v3.1.3, DO NOT EDIT.
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
	AllResourcesEndpoint goa.Endpoint
}

// NewClient initializes a "resource" service client given the endpoints.
func NewClient(allResources goa.Endpoint) *Client {
	return &Client{
		AllResourcesEndpoint: allResources,
	}
}

// AllResources calls the "AllResources" endpoint of the "resource" service.
func (c *Client) AllResources(ctx context.Context, p *AllResourcesPayload) (res ResourceCollection, err error) {
	var ires interface{}
	ires, err = c.AllResourcesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(ResourceCollection), nil
}
