package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ConfigureRouter provides a route-configured *mux.Router object
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeList {
		var handler http.Handler
		if route.HandlerFunction != nil {
			handler = ConfigureHandler(route.HandlerFunction, route)
		} else {
			handler = route.Handler
		}
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

// ConfigureHandler enhances a handler function with ServeHTTP
func ConfigureHandler(inner http.Handler, route Route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)
	})
}
