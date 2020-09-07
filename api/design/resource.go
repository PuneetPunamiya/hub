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

var _ = Service("resource", func() {
	Description("The resource service provides details about all kind of resources")

	Error("internal-error", ErrorResult, "Internal Server Error")
	Error("not-found", ErrorResult, "Resource Not Found Error")

	Method("Query", func() {
		Description("Find resources by a combination of name, kind")
		Payload(func() {
			Attribute("name", String, "Name of resource", func() {
				Default("")
			})
			Attribute("kinds", ArrayOf(String), "Type of resource")
			Attribute("tags", ArrayOf(String), "Tags")
			Attribute("limit", UInt, "Maximum number of resources to be returned", func() {
				Default(100)
			})
			Attribute("exact", Boolean, "Find the exact resource", func() {
				Default(false)
			})
		})
		Result(CollectionOf(Resource), func() {
			View("withoutVersion")
		})

		HTTP(func() {
			GET("/query")
			Param("name")
			Param("kinds")
			Param("tags")
			Param("limit")
			Param("exact")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
			Response("not-found", StatusNotFound)
		})
	})

	Method("List", func() {
		Description("List all resources sorted by rating and name")
		Payload(func() {
			Attribute("limit", UInt, "Maximum number of resources to be returned", func() {
				Default(100)
			})
		})
		Result(CollectionOf(Resource), func() {
			View("withoutVersion")
		})

		HTTP(func() {
			GET("/resources")
			Param("limit")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
		})
	})

	Method("VersionsByID", func() {
		Description("Find all versions of a resource by its id")
		Payload(func() {
			Attribute("id", UInt, "ID of a resource")
			Required("id")
		})
		Result(Versions)

		HTTP(func() {
			GET("/resource/{id}/versions")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
			Response("not-found", StatusNotFound)
		})
	})

	Method("ByKindNameVersion", func() {
		Description("Find resource using name, kind and version of resource")
		Payload(func() {
			Attribute("kind", String, "kind of resource", func() {
				Enum("task", "pipeline")
			})
			Attribute("name", String, "name of resource")
			Attribute("version", String, "version of resource")

			Required("kind", "name", "version")
		})
		Result(ResVersion, func() {
			View("default")
		})

		HTTP(func() {
			GET("/resource/{kind}/{name}/{version}")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
			Response("not-found", StatusNotFound)
		})
	})

	Method("ByVersionId", func() {
		Description("Find a resource using its version's id")
		Payload(func() {
			Attribute("versionID", UInt, "Version ID of a resource's version")
			Required("versionID")
		})
		Result(ResVersion, func() {
			View("default")
		})

		HTTP(func() {
			GET("/resource/version/{versionID}")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
			Response("not-found", StatusNotFound)
		})
	})

	Method("ByKindName", func() {
		Description("Find resources using name and kind")
		Payload(func() {
			Attribute("kind", String, "kind of resource", func() {
				Enum("task", "pipeline")
			})
			Attribute("name", String, "Name of resource")
			Required("kind", "name")
		})
		Result(CollectionOf(Resource), func() {
			View("withoutVersion")
		})

		HTTP(func() {
			GET("/resource/{kind}/{name}")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
			Response("not-found", StatusNotFound)
		})
	})

	Method("ById", func() {
		Description("Find a resource using it's id")
		Payload(func() {
			Attribute("id", UInt, "ID of a resource")
			Required("id")
		})
		Result(Resource, func() {
			View("default")
		})

		HTTP(func() {
			GET("/resource/{id}")

			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
			Response("not-found", StatusNotFound)
		})
	})

})
