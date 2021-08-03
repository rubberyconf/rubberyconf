package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rubberyconf/rubberyconf/internal/business"
)

func ConfigurationDELETE(w http.ResponseWriter, r *http.Request) {

	var logic business.Business

	vars := mux.Vars(r)

	result, _ := logic.DeleteFeature(vars)

	processHTTPAnswer(result, w)

}
