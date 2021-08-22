package httpapi

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/lib/application/httpapi/handlers"
	handlersConf "github.com/rubberyconf/rubberyconf/lib/application/httpapi/handlers/conf"
	handlersFeature "github.com/rubberyconf/rubberyconf/lib/application/httpapi/handlers/feature"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (me *HTTPServer) getRoutes() Routes {

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
			"/conf/{feature:[a-zA-Z0-9_]}",
			handlersConf.ConfigurationGET(me.service),
		},
		Route{
			"configuration",
			"POST",
			"/conf/{feature:[a-zA-Z0-9_]}",
			handlersConf.ConfigurationPOST(me.service),
		},
		Route{
			"configuration",
			"DELETE",
			"/conf/{feature:[a-zA-Z0-9_]}",
			handlersConf.ConfigurationDELETE(me.service),
		},
		Route{
			"configuration",
			"PATCH",
			"/conf/{feature:[a-zA-Z0-9_]}",
			handlersConf.ConfigurationPATCH(me.service),
		},
		Route{
			"feature",
			"GET",
			"/feature/{feature:[a-zA-Z0-9_]}",
			handlersFeature.FeatureGET(me.service),
		},
	}

	return routes

}
