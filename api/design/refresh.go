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

package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("refresh", func() {
	Description("Catalog refresh")

	Method("CatalogRefresh", func() {
		Description("Dummy api for cron job testing")
		Security(JWTAuth, func() {
			Scope("catalog:refresh")
		})

		Payload(func() {
			Token("token", String, "JWT of an agent")
			Required("token")
		})
		HTTP(func() {
			PUT("/catalog/refresh")
			Header("token:Authorization")

			Response(StatusOK)
		})
	})
})
