package api

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"configuration",
		"GET",
		"/conf/{feature}/{branch}",
		Configuration,
	},
	Route{
		"configuration",
		"GET",
		"/conf/{feature}",
		Configuration,
	},
}
