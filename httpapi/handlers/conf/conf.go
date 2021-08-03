package handlers

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/internal/business"
)

func processHTTPAnswer(result int, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/text; charset=UTF-8") //by default
	switch result {
	case business.NotResult:
		w.WriteHeader(http.StatusBadRequest)
		return
	case business.NoContent:
		w.WriteHeader(http.StatusNoContent)
		return
	case business.Unknown:
		w.WriteHeader(http.StatusInternalServerError)
		return
	case business.Success:
		w.WriteHeader(http.StatusOK)
		return
	}

}
