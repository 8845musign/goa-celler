package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Title("The virtual wine cellar")
	Description("A smple goa service")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("bottle", func() {
	BasePath("/bottles")
	DefaultMedia(BottleMedia)

	Action("show", func() {
		Description("Get bottle by id")
		Routing(GET("/:bottleID"))
		Params(func() {
			Param("bottleID", Integer, "Bottle ID")
		})

		Response(OK)
		Response(NotFound)
	})
})

var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	Description("A bottle of wine")

	Attributes(func() {
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Attribute("id", "href", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})
