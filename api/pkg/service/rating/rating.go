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
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"

	"github.com/tektoncd/hub/api/gen/rating"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/token"
)

type service struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
	jwtKey string
}

var (
	invalidTokenError  = rating.MakeInvalidToken(fmt.Errorf("invalid user token"))
	invalidScopesError = rating.MakeInvalidScopes(fmt.Errorf("user not authorized"))
	fetchError         = rating.MakeInternalError(fmt.Errorf("failed to fetch rating"))
	updateError        = rating.MakeInternalError(fmt.Errorf("failed to update rating"))
)

type contextKey string

var (
	contextKeyUserID = contextKey("user-id")
)

// New returns the rating service implementation.
func New(api app.Config) rating.Service {
	return &service{api.Logger(), api.DB(), api.JWTSigningKey()}
}

// JWTAuth implements the authorization logic for service "rating" for the
// "jwt" security scheme.
func (s *service) JWTAuth(ctx context.Context, jwt string, scheme *security.JWTScheme) (context.Context, error) {

	claims, err := token.Verify(jwt, s.jwtKey)
	if err != nil {
		return ctx, invalidTokenError
	}

	err = token.ValidateScopes(claims, scheme)
	if err != nil {
		return ctx, invalidScopesError
	}

	// Gets user id and passes in context to API Method
	userID, ok := claims["id"].(float64)
	if !ok {
		return ctx, invalidTokenError
	}

	ctx = context.WithValue(ctx, contextKeyUserID, uint(userID))

	return ctx, nil
}

// Find user's rating for a resource
func (s *service) Get(ctx context.Context, p *rating.GetPayload) (res *rating.GetResult, err error) {

	userID := ctx.Value(contextKeyUserID).(uint)

	q := s.db.Where("user_id = ? AND resource_id = ?", userID, p.ID)

	r := &model.UserResourceRating{}
	if err := q.Find(&r).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &rating.GetResult{Rating: 0}, nil
		}
		s.logger.Error(err)
		return nil, fetchError
	}

	res = &rating.GetResult{
		Rating: r.Rating,
	}

	return res, nil
}

// Update user's rating for a resource
func (s *service) Update(ctx context.Context, p *rating.UpdatePayload) (res *rating.UpdateResult, err error) {

	userID := ctx.Value(contextKeyUserID).(uint)

	// update user's rating for the resource
	q := s.db.Where("user_id = ? AND resource_id = ?", userID, p.ID)

	rat := &model.UserResourceRating{
		UserID:     userID,
		ResourceID: p.ID,
		Rating:     p.Rating,
	}
	if err := q.Assign(&model.UserResourceRating{Rating: p.Rating}).
		FirstOrCreate(rat).Error; err != nil {
		return nil, updateError
	}

	// evaluates average rating of the resource
	q = s.db.Model(&model.UserResourceRating{}).Where("resource_id = ?", p.ID).Select("avg(rating)")

	var avg float64
	q.Row().Scan(&avg)

	// updates resource average's rating in resource table
	q = s.db.Model(&model.Resource{}).Where("id = ?", p.ID)

	avg = math.Round(avg*10) / 10
	if err := q.Update("rating", avg).
		Error; err != nil {
		return nil, updateError
	}

	res = &rating.UpdateResult{
		AvgRating: avg,
	}

	return res, nil
}
