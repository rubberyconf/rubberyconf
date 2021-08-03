package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/internal/business"
)

func ConfigurationGET(w http.ResponseWriter, r *http.Request) {

	var logic business.Business

	vars := mux.Vars(r)
	result, content, typeContent := logic.GetFeature(vars)

	processHTTPAnswer(result, w)

	if result == business.Success {
		if typeContent == "json" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		}
		w.Write([]byte(fmt.Sprintf("%v", content)))
	}

}
