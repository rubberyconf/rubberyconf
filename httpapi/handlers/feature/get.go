package feature

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/internal/service"
)

func FeatureGET(w http.ResponseWriter, r *http.Request) {

	var logic service.Service

	vars := mux.Vars(r)
	result, content, typeContent, err := logic.GetFeatureOnlyValue(r.Context(), vars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tools.ProcessHTTPAnswer(result, w)

	if result == service.Success {
		if typeContent == "json" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			bytes, err := json.Marshal(content)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write(bytes)
			}
		} else {
			w.Write([]byte(fmt.Sprintf("%v", content)))
		}
	}

}
