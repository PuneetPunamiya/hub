// Code generated by goa v3.2.2, DO NOT EDIT.
//
// admin HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"strings"

	admin "github.com/tektoncd/hub/api/gen/admin"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateAgentResponse returns an encoder for responses returned by the
// admin CreateAgent endpoint.
func EncodeCreateAgentResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*admin.CreateAgentResult)
		enc := encoder(ctx, w)
		body := NewCreateAgentResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateAgentRequest returns a decoder for requests sent to the admin
// CreateAgent endpoint.
func DecodeCreateAgentRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateAgentRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateAgentRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token string
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateAgentPayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeCreateAgentError returns an encoder for errors returned by the
// CreateAgent admin endpoint.
func EncodeCreateAgentError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-payload":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateAgentInvalidPayloadResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-payload")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "invalid-token":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateAgentInvalidTokenResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-token")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "invalid-scopes":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateAgentInvalidScopesResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "internal-error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateAgentInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal-error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}