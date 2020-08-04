// Code generated by goa v3.2.2, DO NOT EDIT.
//
// auth HTTP client encoders and decoders
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

	auth "github.com/tektoncd/hub/api/gen/auth"
	goahttp "goa.design/goa/v3/http"
)

// BuildAuthenticateRequest instantiates a HTTP request object with method and
// path set to call the "auth" service "Authenticate" endpoint
func (c *Client) BuildAuthenticateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthenticateAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("auth", "Authenticate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAuthenticateRequest returns an encoder for requests sent to the auth
// Authenticate server.
func EncodeAuthenticateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*auth.AuthenticatePayload)
		if !ok {
			return goahttp.ErrInvalidType("auth", "Authenticate", "*auth.AuthenticatePayload", v)
		}
		values := req.URL.Query()
		values.Add("code", p.Code)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeAuthenticateResponse returns a decoder for responses returned by the
// auth Authenticate endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeAuthenticateResponse may return the following errors:
//	- "invalid-code" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "invalid-token" (type *goa.ServiceError): http.StatusUnauthorized
//	- "invalid-scopes" (type *goa.ServiceError): http.StatusForbidden
//	- error: internal error
func DecodeAuthenticateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body AuthenticateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "Authenticate", err)
			}
			err = ValidateAuthenticateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "Authenticate", err)
			}
			res := NewAuthenticateResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body AuthenticateInvalidCodeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "Authenticate", err)
			}
			err = ValidateAuthenticateInvalidCodeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "Authenticate", err)
			}
			return nil, NewAuthenticateInvalidCode(&body)
		case http.StatusInternalServerError:
			var (
				body AuthenticateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "Authenticate", err)
			}
			err = ValidateAuthenticateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "Authenticate", err)
			}
			return nil, NewAuthenticateInternalError(&body)
		case http.StatusUnauthorized:
			var (
				body AuthenticateInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "Authenticate", err)
			}
			err = ValidateAuthenticateInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "Authenticate", err)
			}
			return nil, NewAuthenticateInvalidToken(&body)
		case http.StatusForbidden:
			var (
				body AuthenticateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "Authenticate", err)
			}
			err = ValidateAuthenticateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("auth", "Authenticate", err)
			}
			return nil, NewAuthenticateInvalidScopes(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("auth", "Authenticate", resp.StatusCode, string(body))
		}
	}
}
