package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("resource", func() {
	Description("The resource service provides all resources information")

	Error("db-down", ErrorResult)

	// Method to get all resources
	Method("AllResources", func() {
		Description("Get all Resources")
		Payload(func() {
			Attribute("limit", UInt, "Number of resources", func() {
				Default(100)
			})
		})

		Result(CollectionOf(Resource), func() { View("extended") })

		HTTP(func() {
			GET("/resources")
			Param("limit")
			Response(StatusOK)
			Response("db-down", StatusInternalServerError)
		})
	})
})
