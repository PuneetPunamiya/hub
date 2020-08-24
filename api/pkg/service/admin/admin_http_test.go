// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"

	"github.com/tektoncd/hub/api/gen/admin"
	"github.com/tektoncd/hub/api/gen/http/admin/server"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

// Token for the user with github name "foo-bar" and github login "foo"
// It has a scope "agent:create" along with default scope
const validTokenWithAgentCreateScope = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJpZCI6MTEsImxvZ2luIjoiZm9vIiwibmFtZSI6ImZvby1iYXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiLCJhZ2VudDpjcmVhdGUiXX0." +
	"DcaSiYNyiRpTDVmr7rJs3P-B0RgoDfby5si2UAkcnIM"

// Token for the agent with name "agent-007" with scopes ["test:read"]
const agentToken007 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJpZCI6MTAwMDEsIm5hbWUiOiJhZ2VudC0wMDciLCJzY29wZXMiOlsidGVzdDpyZWFkIl19." +
	"5kgbZ1nsrr8mvrDx1Wa-NgMc9W14B3LVsxXJ_wt99DY"

// Token for the agent with name "agent-007" with scopes ["test:read","agent:create"]
const agentToken007Updated = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJpZCI6MzEsIm5hbWUiOiJhZ2VudC0wMDEiLCJzY29wZXMiOlsidGVzdDpyZWFkIiwiYWdlbnQ6Y3JlYXRlIl19." +
	"QbZNJcsb0okpk7bnyWKPXnoR-Qpk6ru12Ysmpkm_rLY"

func CreateAgentChecker(tc *testutils.TestConfig) *goahttpcheck.APIChecker {
	checker := goahttpcheck.New()
	checker.Mount(server.NewCreateAgentHandler,
		server.MountCreateAgentHandler,
		admin.NewCreateAgentEndpoint(New(tc), New(tc).(admin.Auther).JWTAuth))
	return checker
}

func TestCreateAgent_Http(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	data := []byte(`{"name": "agent-007","scopes": ["test:read"]}`)

	CreateAgentChecker(tc).Test(t, http.MethodPut, "/system/user/agent").
		WithHeader("Authorization", validTokenWithAgentCreateScope).WithBody(data).
		Check().
		HasStatus(200).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, agentToken007, jsonMap["token"])
	})
}

func TestCreateAgent_Http_NormalUserExistWithName(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	data := []byte(`{"name": "foo-bar","scopes": ["test:read"]}`)

	CreateAgentChecker(tc).Test(t, http.MethodPut, "/system/user/agent").
		WithHeader("Authorization", validTokenWithAgentCreateScope).WithBody(data).
		Check().
		HasStatus(400).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, "invalid-payload", jsonMap["name"])
		assert.Equal(t, "normal user exists with name: foo-bar", jsonMap["message"])
	})
}

func TestCreateAgent_Http_InvalidScopeCase(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	data := []byte(`{"name": "agent-001","scopes": ["invalid:scope"]}`)

	CreateAgentChecker(tc).Test(t, http.MethodPut, "/system/user/agent").
		WithHeader("Authorization", validTokenWithAgentCreateScope).WithBody(data).
		Check().
		HasStatus(400).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, "invalid-payload", jsonMap["name"])
	})
}

func TestCreateAgent_Http_UpdateCase(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	data := []byte(`{"name": "agent-001","scopes": ["test:read","agent:create"]}`)

	CreateAgentChecker(tc).Test(t, http.MethodPut, "/system/user/agent").
		WithHeader("Authorization", validTokenWithAgentCreateScope).WithBody(data).
		Check().
		HasStatus(200).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, agentToken007Updated, jsonMap["token"])
	})
}
