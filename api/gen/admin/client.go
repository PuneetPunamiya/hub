// Code generated by goa v3.8.1, DO NOT EDIT.
//
// admin client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "admin" service client.
type Client struct {
	UpdateAgentEndpoint   goa.Endpoint
	RefreshConfigEndpoint goa.Endpoint
}

// NewClient initializes a "admin" service client given the endpoints.
func NewClient(updateAgent, refreshConfig goa.Endpoint) *Client {
	return &Client{
		UpdateAgentEndpoint:   updateAgent,
		RefreshConfigEndpoint: refreshConfig,
	}
}

// UpdateAgent calls the "UpdateAgent" endpoint of the "admin" service.
func (c *Client) UpdateAgent(ctx context.Context, p *UpdateAgentPayload) (res *UpdateAgentResult, err error) {
	var ires interface{}
	ires, err = c.UpdateAgentEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateAgentResult), nil
}

// RefreshConfig calls the "RefreshConfig" endpoint of the "admin" service.
func (c *Client) RefreshConfig(ctx context.Context, p *RefreshConfigPayload) (res *RefreshConfigResult, err error) {
	var ires interface{}
	ires, err = c.RefreshConfigEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*RefreshConfigResult), nil
}
