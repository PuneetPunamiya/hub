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
	"context"

	"github.com/tektoncd/hub/api/gen/status"
	"github.com/tektoncd/hub/api/pkg/app"
)

// status service implementation.
type service struct {
	app.Service
}

// New returns the status service implementation.
func New(api app.BaseConfig) status.Service {
	return &service{api.Service("status")}
}

// Return status 'ok' when the server has started successfully
func (s *service) Status(ctx context.Context) (res []*status.Server, err error) {

	service := []*status.Services{}
	service = append(service, &status.Services{Name: "api", Status: "ok"})

	db := s.DB(ctx).DB()

	if err := db.Ping(); err != nil {
		service = append(service, &status.Services{Name: "db", Status: err.Error()})
		res = append(res, &status.Server{Services: service})
		return res, nil
	}

	service = append(service, &status.Services{Name: "db", Status: "ok"})

	res = append(res, &status.Server{Services: service})
	return res, nil
}
