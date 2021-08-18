package conf

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/internal/service"
)

func ConfigurationDELETE(w http.ResponseWriter, r *http.Request) {

	var logic service.Service

	vars := mux.Vars(r)

	result, _ := logic.DeleteFeature(vars)

	tools.ProcessHTTPAnswer(result, w)

}
