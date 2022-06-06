// Code generated by goa v3.7.6, DO NOT EDIT.
//
// rating client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package rating

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "rating" service client.
type Client struct {
	GetEndpoint    goa.Endpoint
	UpdateEndpoint goa.Endpoint
}

// NewClient initializes a "rating" service client given the endpoints.
func NewClient(get, update goa.Endpoint) *Client {
	return &Client{
		GetEndpoint:    get,
		UpdateEndpoint: update,
	}
}

// Get calls the "Get" endpoint of the "rating" service.
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *GetResult, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetResult), nil
}

// Update calls the "Update" endpoint of the "rating" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}
