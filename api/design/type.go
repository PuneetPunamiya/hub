package design

import (
	. "goa.design/goa/v3/dsl"
)

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
