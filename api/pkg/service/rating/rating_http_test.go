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

package rating

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"

	"github.com/tektoncd/hub/api/gen/http/rating/server"
	"github.com/tektoncd/hub/api/gen/rating"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

// Token for the user with github name "foo-bar" and github login "foo"
const validToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJpZCI6MTEsImxvZ2luIjoiZm9vIiwibmFtZSI6ImZvby1iYXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0." +
	"picev_yRsVjEzmkpQdErSQGB4mLfmb8U47kgu7olWHA"

// Token with Invalid Scopes
const tokenWithInvalidScope = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJpZCI6MTIzLCJsb2dpbiI6ImFiYyIsIm5hbWUiOiJJbnZhbGlkIFNjb3BlIn0." +
	"9I0u6eJi5t-ocw2rDq-um97FjOz_Gfi8Lt33kwjskb8"

func GetChecker(tc *testutils.TestConfig) *goahttpcheck.APIChecker {
	checker := goahttpcheck.New()
	checker.Mount(server.NewGetHandler,
		server.MountGetHandler,
		rating.NewGetEndpoint(New(tc), New(tc).(rating.Auther).JWTAuth))
	return checker
}

func TestGet_Http_InvalidToken(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	GetChecker(tc).Test(t, http.MethodGet, "/resource/1/rating").
		WithHeader("Authorization", "invalidToken").Check().
		HasStatus(401).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, "invalid-token", jsonMap["name"])
	})
}

func TestGet_Http_InvalidScopes(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	GetChecker(tc).Test(t, http.MethodGet, "/resource/1/rating").
		WithHeader("Authorization", tokenWithInvalidScope).Check().
		HasStatus(403).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, "invalid-scopes", jsonMap["name"])
	})
}

func TestGet_Http(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	GetChecker(tc).Test(t, http.MethodGet, "/resource/1/rating").
		WithHeader("Authorization", validToken).Check().
		HasStatus(200).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, float64(5), jsonMap["rating"])
	})
}

func TestGet_Http_RatingNotFound(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	GetChecker(tc).Test(t, http.MethodGet, "/resource/11/rating").
		WithHeader("Authorization", validToken).Check().
		HasStatus(200).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, float64(0), jsonMap["rating"])
	})
}

func UpdateChecker(tc *testutils.TestConfig) *goahttpcheck.APIChecker {
	checker := goahttpcheck.New()
	checker.Mount(server.NewUpdateHandler,
		server.MountUpdateHandler,
		rating.NewUpdateEndpoint(New(tc), New(tc).(rating.Auther).JWTAuth))
	return checker
}

func TestUpdate_Http(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	data := []byte(`{"rating": 2}`)

	UpdateChecker(tc).Test(t, http.MethodPut, "/resource/1/rating").
		WithHeader("Authorization", validToken).WithBody(data).
		Check().
		HasStatus(200).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		var jsonMap map[string]interface{}
		marshallErr := json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, marshallErr)

		assert.Equal(t, float64(3.5), jsonMap["avgRating"])
	})
}
