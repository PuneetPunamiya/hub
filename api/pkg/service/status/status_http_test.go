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

package status

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"
	"gotest.tools/golden"

	"github.com/tektoncd/hub/api/gen/http/status/server"
	"github.com/tektoncd/hub/api/gen/status"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestOk_http(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	checker := goahttpcheck.New()
	checker.Mount(
		server.NewStatusHandler,
		server.MountStatusHandler,
		status.NewStatusEndpoint(New(tc)),
	)

	checker.Test(t, http.MethodGet, "/").Check().
		HasStatus(http.StatusOK).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		res, err := testutils.FormatJSON(b)
		assert.NoError(t, err)

		golden.Assert(t, res, fmt.Sprintf("%s.golden", t.Name()))
	})
}

func TestNotOk_http(t *testing.T) {
	tc := testutils.Setup(t)
	// testutils.LoadFixtures(t, tc.FixturePath())

	tc.DB().Close()

	checker := goahttpcheck.New()
	checker.Mount(
		server.NewStatusHandler,
		server.MountStatusHandler,
		status.NewStatusEndpoint(New(tc)),
	)

	checker.Test(t, http.MethodGet, "/").Check().
		HasStatus(http.StatusOK).Cb(func(r *http.Response) {
		b, readErr := ioutil.ReadAll(r.Body)
		assert.NoError(t, readErr)
		defer r.Body.Close()

		res, err := testutils.FormatJSON(b)
		assert.NoError(t, err)

		golden.Assert(t, res, fmt.Sprintf("%s.golden", t.Name()))
	})
}
