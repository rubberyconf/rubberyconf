package conf

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	feature "github.com/rubberyconf/language/lib"
	"github.com/rubberyconf/rubberyconf/lib/application/httpapi/handlers/tools"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func ConfigurationPOST(service input.IServiceFeature) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ruberConf := feature.FeatureDefinition{}
		ruberConf.LoadFromJsonBinary(b)

		result, _ := service.CreateFeature(r.Context(), vars, ruberConf)

		tools.ProcessHTTPAnswer(result, w)

	}
}
