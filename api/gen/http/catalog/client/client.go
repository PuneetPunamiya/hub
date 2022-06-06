// Code generated by goa v3.7.6, DO NOT EDIT.
//
// catalog client HTTP transport
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the catalog service endpoint HTTP clients.
type Client struct {
	// Refresh Doer is the HTTP client used to make requests to the Refresh
	// endpoint.
	RefreshDoer goahttp.Doer

	// RefreshAll Doer is the HTTP client used to make requests to the RefreshAll
	// endpoint.
	RefreshAllDoer goahttp.Doer

	// CatalogError Doer is the HTTP client used to make requests to the
	// CatalogError endpoint.
	CatalogErrorDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the catalog service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RefreshDoer:         doer,
		RefreshAllDoer:      doer,
		CatalogErrorDoer:    doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Refresh returns an endpoint that makes HTTP requests to the catalog service
// Refresh server.
func (c *Client) Refresh() goa.Endpoint {
	var (
		encodeRequest  = EncodeRefreshRequest(c.encoder)
		decodeResponse = DecodeRefreshResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRefreshRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RefreshDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("catalog", "Refresh", err)
		}
		return decodeResponse(resp)
	}
}

// RefreshAll returns an endpoint that makes HTTP requests to the catalog
// service RefreshAll server.
func (c *Client) RefreshAll() goa.Endpoint {
	var (
		encodeRequest  = EncodeRefreshAllRequest(c.encoder)
		decodeResponse = DecodeRefreshAllResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRefreshAllRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RefreshAllDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("catalog", "RefreshAll", err)
		}
		return decodeResponse(resp)
	}
}

// CatalogError returns an endpoint that makes HTTP requests to the catalog
// service CatalogError server.
func (c *Client) CatalogError() goa.Endpoint {
	var (
		encodeRequest  = EncodeCatalogErrorRequest(c.encoder)
		decodeResponse = DecodeCatalogErrorResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCatalogErrorRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CatalogErrorDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("catalog", "CatalogError", err)
		}
		return decodeResponse(resp)
	}
}
