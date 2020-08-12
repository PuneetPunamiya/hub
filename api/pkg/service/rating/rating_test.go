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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tektoncd/hub/api/gen/rating"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestGet(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	ratingSvc := New(tc)
	ctx := context.WithValue(context.Background(), contextKeyUserID, uint(11))
	payload := &rating.GetPayload{ID: 1, Token: "token"}
	rat, err := ratingSvc.Get(ctx, payload)
	assert.NoError(t, err)
	assert.Equal(t, uint(5), rat.Rating)
}

func TestGet_RatingNotFound(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	ratingSvc := New(tc)
	ctx := context.WithValue(context.Background(), contextKeyUserID, uint(11))
	payload := &rating.GetPayload{ID: 111, Token: "token"}
	rat, err := ratingSvc.Get(ctx, payload)
	assert.NoError(t, err)
	assert.Equal(t, uint(0), rat.Rating)
}

func TestUpdate(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	ratingSvc := New(tc)
	ctx := context.WithValue(context.Background(), contextKeyUserID, uint(11))
	payload := &rating.UpdatePayload{ID: 1, Rating: 3, Token: "token"}
	rat, err := ratingSvc.Update(ctx, payload)
	assert.NoError(t, err)
	assert.Equal(t, float64(4), rat.AvgRating)
}
