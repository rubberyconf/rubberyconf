package conf

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func ConfigurationPATCH(service *input.IServiceFeature) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ruberConf := feature.FeatureDefinition{}
		ruberConf.LoadFromJsonBinary(b)

		result, _ := service.PatchFeature(vars, ruberConf)

		tools.ProcessHTTPAnswer(result, w)
	}
}
