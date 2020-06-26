package design

import (
	. "goa.design/goa/v3/dsl"
)

var Resource = ResultType("application/vnd.hub.resource", func() {

	Description("The resource service provides all resources information.")
	TypeName("Resource")

	Attribute("id", UInt, "ID is the unique id of the resource", func() {
		Example("id", 1)
	})

	Attribute("name", String, "Name of the resource", func() {
		Example("name", "golang")
	})

	Attribute("displayName", String, "Display name of the resource", func() {
		Example("displayName", "golang")
	})

	Attribute("catalog", Catalog, "Type of catalog where resource belongs")

	Attribute("type", String, "Type of resource", func() {
		Example("type", "task")
	})

	Attribute("description", String, "Description of resource", func() {
		Example("descripiton", "This Task is Golang task to build Go projects")
	})

	Attribute("latest_version", String, "Latest version o resource", func() {
		Example("latest_version", "0.1")
	})

	Attribute("tags", ArrayOf(ResourceTag), "Tags related to resources")

	Attribute("rating", UInt, "Rating of resource", func() {
		Example("rating", 5)
	})

	Attribute("last_updated_at", String, "Date when resource was last updated", func() {
		Example("last_updated_at", "2020-03-26T14:09:08.19599+05:30")
	})

	Attribute("versions", ArrayOf(Versions), "Version of resource")

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("displayName")
		Attribute("catalog")
		Attribute("type")
		Attribute("description")
		Attribute("latest_version")
		Attribute("tags")
		Attribute("rating")
		Attribute("last_updated_at")
		Attribute("versions")
	})

	View("extended", func() {
		Attribute("id")
		Attribute("name")
		Attribute("displayName")
		Attribute("catalog")
		Attribute("type")
		Attribute("description")
		Attribute("latest_version")
		Attribute("tags")
		Attribute("rating")
		Attribute("last_updated_at")
	})
	Required("id", "name", "displayName", "catalog", "type", "description", "latest_version", "tags", "rating", "last_updated_at", "versions")

})

var Versions = Type("Versions", func() {
	Attribute("versionId", UInt, "Version ID of the resource to be fetched", func() {
		Example("versionId", 1)
	})

	Attribute("version", String, "Version of the resource to be fetched", func() {
		Example("versionId", "0,1")
	})
	Required("versionId", "version")

})

var Catalog = Type("Catalog", func() {
	Attribute("id", UInt, "ID is the unique id of the catalog", func() {
		Example("id", 1)
	})
	Attribute("type", String, "Type of catalog", func() {
		Example("type", "Official")
	})
	Required("id", "type")
})

var Category = Type("category", func() {
	Attribute("id", UInt, "unique id of category", func() {
		Example("id", 1)
	})
	Attribute("name", String, "name of category", func() {
		Example("name", "Image-build")
	})
	Attribute("tags", ArrayOf(ResourceTag), "list of tag associated with category")
	Required("id", "name", "tags")
})
var ResourceTag = Type("Tag", func() {
	Attribute("id", UInt, "Id is the unique id of tags", func() {
		Example("id", 1)
	})
	Attribute("name", String, "name of tag", func() {
		Example("name", "cli")
	})
	Required("id", "name")
})
