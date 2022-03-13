package router

import (
	"net/http"

	h "github.com/aditya109/library-system/internal/handlers"
	"github.com/go-openapi/runtime/middleware"
)

// Route is container template for server routes
type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
	Handler         http.Handler
}

type routes []Route

var routeList = routes{

	// swagger:route GET / home welcome
	// Returns a welcome message
	// responses:
	// 	200: WelcomeResponse
	Route{
		"Welcome",
		"GET",
		"/",
		h.WelcomeHandler,
		nil,
	},

	// swagger:route GET /docs docs swaggerDocumentation
	// Returns swagger specification uunder OpenAPIv3 documeted APIs
	Route{
		"swaggerDocumentation",
		"GET",
		"/docs",
		nil,
		middleware.Redoc(middleware.RedocOpts{
			SpecURL: "/swagger.yaml",
		}, nil),
	},
	Route{
		"Swagger JSON",
		"GET",
		"/swagger.yaml",
		nil,
		http.FileServer(http.Dir("./api/swagger")),
	},

	// swagger:route GET /items items listItems
	// Returns a list of items, no query params required
	// responses:
	// 	200: GetItemsResponse
	Route{
		"GetItems",
		"GET",
		"/items",
		h.GetItemsHandler,
		nil,
	},
}
