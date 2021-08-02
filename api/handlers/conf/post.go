package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func ConfigurationPOST(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	conf, cacheValue, source, featureSelected, result := preRequisites(vars)
	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	ruberConf := feature.FeatureDefinition{}
	ruberConf.LoadFromJsonBinary(b)

	featureSelected.Value = &ruberConf

	timeInText := conf.Api.DefaultTTL
	if ruberConf.Default.TTL != "" {
		timeInText = ruberConf.Default.TTL
	}
	u, _ := time.ParseDuration(timeInText)
	res, _ := cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	if !res {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res = source.CreateFeature(featureSelected)
	if !res {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	metrics.GetMetrics().Update(featureSelected.Key)

}
