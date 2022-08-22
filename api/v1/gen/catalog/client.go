// Code generated by goa v3.8.3, DO NOT EDIT.
//
// catalog client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package catalog

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "catalog" service client.
type Client struct {
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "catalog" service client given the endpoints.
func NewClient(list goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
	}
}

// List calls the "List" endpoint of the "catalog" service.
func (c *Client) List(ctx context.Context) (res *ListResult, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}
