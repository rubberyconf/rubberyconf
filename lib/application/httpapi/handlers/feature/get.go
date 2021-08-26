package feature

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/lib/application/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func FeatureGET(service *input.IServiceFeature) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		result, content, typeContent, err := service.GetFeatureOnlyValue(r.Context(), vars)
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
}
