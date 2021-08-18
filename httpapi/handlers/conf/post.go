package conf

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/service"
)

func ConfigurationPOST(w http.ResponseWriter, r *http.Request) {

	var logic service.Service
	vars := mux.Vars(r)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ruberConf := feature.FeatureDefinition{}
	ruberConf.LoadFromJsonBinary(b)

	result, _ := logic.CreateFeature(vars, ruberConf)

	tools.ProcessHTTPAnswer(result, w)

}
