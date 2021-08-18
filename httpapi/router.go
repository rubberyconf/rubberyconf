package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() http.Handler { ////*mux.Router { //

	router := mux.NewRouter().StrictSlash(true)

	router.Use(newServerMiddleware())

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	/*	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost},
		AllowCredentials: true,
	})*/

	handler := cors.AllowAll().Handler(router)

	//handler := c.Handler(router)

	return handler
}
