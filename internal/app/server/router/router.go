package router

import (
	"net/http"

	"library-server/configs"
	"library-server/internal/app/server/router/routes"

	"github.com/gorilla/mux"
)

func GetRouter(prefix string, config configs.Config) *mux.Router {
	router := mux.NewRouter()
	routes := routes.GetRoutes(prefix, config)

	for i := 0; i < len(routes); i++ {
		var route = routes[i]
		router.HandleFunc(route.Path, route.HandlerFunction).Methods(route.MethodTypes...)
	}

	// router.HandleFunc("/api/v1/books", GetBooksHandler)
	// router.HandleFunc("/api/v1/book/{id}", GetBookHandler)
	// router.HandleFunc("/api/v1/books", AddBooksHandler)
	// router.HandleFunc("/api/v1/book", AddBookHandler)
	http.Handle("/", router)
	return router
}
