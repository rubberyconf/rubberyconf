package api

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/api/handlers"
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
		handlers.Index,
	},
	Route{
		"configuration",
		"GET",
		"/conf/{feature}/{branch}",
		handlers.ConfigurationGET,
	},
	Route{
		"configuration",
		"GET",
		"/conf/{feature}",
		handlers.ConfigurationGET,
	},
	Route{
		"configuration",
		"POST",
		"/conf/{feature}",
		handlers.ConfigurationPOST,
	},
	Route{
		"configuration",
		"DELETE",
		"/conf/{feature}",
		handlers.ConfigurationDELETE,
	},
}
