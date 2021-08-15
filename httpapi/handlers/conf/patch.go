package conf

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/internal/business"
	"github.com/rubberyconf/rubberyconf/internal/feature"
)

func ConfigurationPATCH(w http.ResponseWriter, r *http.Request) {

	var logic business.Business
	vars := mux.Vars(r)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ruberConf := feature.FeatureDefinition{}
	ruberConf.LoadFromJsonBinary(b)

	result, _ := logic.PatchFeature(vars, ruberConf)

	tools.ProcessHTTPAnswer(result, w)

}