package tools

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/internal/service"
)

func ProcessHTTPAnswer(result int, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/text; charset=UTF-8") //by default
	switch result {
	case service.NotResult:
		w.WriteHeader(http.StatusBadRequest)
		return
	case service.NoContent:
		w.WriteHeader(http.StatusNoContent)
		return
	case service.Unknown:
		w.WriteHeader(http.StatusInternalServerError)
		return
	case service.Success:
		w.WriteHeader(http.StatusOK)
		return
	}

}
