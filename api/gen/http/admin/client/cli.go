// Code generated by goa v3.2.2, DO NOT EDIT.
//
// admin HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"encoding/json"
	"fmt"

	admin "github.com/tektoncd/hub/api/gen/admin"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateAgentPayload builds the payload for the admin CreateAgent
// endpoint from CLI flags.
func BuildCreateAgentPayload(adminCreateAgentBody string, adminCreateAgentToken string) (*admin.CreateAgentPayload, error) {
	var err error
	var body CreateAgentRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateAgentBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"name\": \"Voluptas beatae id qui.\",\n      \"scopes\": [\n         \"Possimus tempora omnis et nihil aut.\",\n         \"Quidem magni.\",\n         \"Nemo sint est omnis.\"\n      ]\n   }'")
		}
		if body.Scopes == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("scopes", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminCreateAgentToken
	}
	v := &admin.CreateAgentPayload{
		Name: body.Name,
	}
	if body.Scopes != nil {
		v.Scopes = make([]string, len(body.Scopes))
		for i, val := range body.Scopes {
			v.Scopes[i] = val
		}
	}
	v.Token = token

	return v, nil
}
