package conf

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func ConfigurationDELETE(service *input.IServiceFeature) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		result, _ := service.DeleteFeature(vars)

		tools.ProcessHTTPAnswer(result, w)
	}
}
