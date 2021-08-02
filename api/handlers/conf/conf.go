package handlers

import (
	"github.com/rubberyconf/rubberyconf/internal/cache"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
)

func preRequisites(vars map[string]string) (*config.Config, cache.IDataStorage, datasource.IDataSource, datasource.Feature, bool) {
	conf := config.GetConfiguration()
	cacheValue := cache.SelectCache(conf)
	source := datasource.SelectSource()

	feature, result := source.EnableFeature(vars)

	return conf, cacheValue, source, feature, result
}

/*
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
*/
