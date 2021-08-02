package api

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/api/handlers"
	handlersConf "github.com/rubberyconf/rubberyconf/api/handlers/conf"
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
		handlersConf.ConfigurationGET,
	},
	Route{
		"configuration",
		"GET",
		"/conf/{feature}",
		handlersConf.ConfigurationGET,
	},
	Route{
		"configuration",
		"POST",
		"/conf/{feature}",
		handlersConf.ConfigurationPOST,
	},
	Route{
		"configuration",
		"DELETE",
		"/conf/{feature}",
		handlersConf.ConfigurationDELETE,
	},
}
