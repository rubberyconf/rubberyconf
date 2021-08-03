package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/internal/business"
)

func ConfigurationPOST(w http.ResponseWriter, r *http.Request) {

	var logic business.Business
	vars := mux.Vars(r)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, _ := logic.CreateFeature(vars, b)

	processHTTPAnswer(result, w)

}
