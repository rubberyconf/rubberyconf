package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigurationDELETE(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, _ := cacheValue.DeleteValue(featureSelected.Key)
	if res {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res = source.DeleteFeature(featureSelected)
	if res {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
