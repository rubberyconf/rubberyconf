package conf

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/internal/service"
)

func ConfigurationGET(w http.ResponseWriter, r *http.Request) {

	var logic service.Service

	vars := mux.Vars(r)
	result, content, err := logic.GetFeatureFull(r.Context(), vars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	tools.ProcessHTTPAnswer(result, w)

	if result == service.Success {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		bytes, err := json.Marshal(content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(bytes)
		}
	}

}
