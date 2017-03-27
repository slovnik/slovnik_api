package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct defines the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type is a shorthand for array of Route
type Routes []Route

// NewRouter creates new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"translate",
		"GET",
		"/api/translate/{word}",
		translate,
	},
	Route{
		"search",
		"GET",
		"/api/search/{term}",
		search,
	},
}
