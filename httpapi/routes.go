package httpapi

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers"
	handlersConf "github.com/rubberyconf/rubberyconf/httpapi/handlers/conf"
	handlersFeature "github.com/rubberyconf/rubberyconf/httpapi/handlers/feature"
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
	Route{
		"feature",
		"GET",
		"/feature/{feature}",
		handlersFeature.FeatureGET,
	},
}
