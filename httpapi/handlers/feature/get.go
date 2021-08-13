package feature

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/internal/business"
)

func FeatureGET(w http.ResponseWriter, r *http.Request) {

	var logic business.Business

	vars := mux.Vars(r)
	result, content, typeContent := logic.GetFeatureOnlyValue(vars)

	tools.ProcessHTTPAnswer(result, w)

	if result == business.Success {
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
