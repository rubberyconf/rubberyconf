package tools

import (
	"net/http"

	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func ProcessHTTPAnswer(result input.ServiceResult, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/text; charset=UTF-8") //by default
	switch result {
	case input.NotResult:
		w.WriteHeader(http.StatusBadRequest)
		return
	case input.NoContent:
		w.WriteHeader(http.StatusNoContent)
		return
	case input.Unknown:
		w.WriteHeader(http.StatusInternalServerError)
		return
	case input.Success:
		w.WriteHeader(http.StatusOK)
		return
	}

}
