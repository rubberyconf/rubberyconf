package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func ConfigurationGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	conf, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateCache := false
	val, found, _ := cacheValue.GetValue(featureSelected.Key)
	if !found {
		found, err := source.GetFeature(&featureSelected)

		if err == nil && !found {
			w.Header().Set("Content-Type", "application/text; charset=UTF-8")
			w.WriteHeader(http.StatusNoContent)
			return

		}
		if err != nil {
			w.Header().Set("Content-Type", "application/text; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		updateCache = true
	} else {
		featureSelected.Value = val
	}

	if updateCache {
		timeInText := conf.Api.DefaultTTL
		if featureSelected.Value.Default.TTL != "" {
			timeInText = featureSelected.Value.Default.TTL
		}
		u, _ := time.ParseDuration(timeInText)
		cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	}

	finalresult, err := featureSelected.Value.GetFinalValue()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/text; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%v", finalresult)))
	}

	metrics.GetMetrics().Update(featureSelected.Key)
}
