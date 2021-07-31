package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"github.com/rubberyconf/rubberyconf/internal/cache"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func preRequisites(vars map[string]string) (*config.Config, cache.IDataStorage, datasource.IDataSource, datasource.Feature, bool) {
	conf := config.GetConfiguration()
	cacheValue := cache.SelectCache(conf)
	source := datasource.SelectSource()

	feature, result := source.EnableFeature(vars)

	//featureSelected := vars["feature"]
	//result := true
	//if !result {
	//	log.Printf("no feature specified")
	//	result = false
	//}
	return conf, cacheValue, source, feature, result
}

func ConfigurationGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	conf, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateCache := false
	val, err := cacheValue.GetValue(featureSelected.Key)
	if err {
		err = source.GetFeature(&featureSelected)

		if val == nil && !err {
			w.Header().Set("Content-Type", "application/text; charset=UTF-8")
			w.WriteHeader(http.StatusNoContent)
			return

		}
		if val == nil && err {
			w.Header().Set("Content-Type", "application/text; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		updateCache = true
	}

	ruberConf := feature.RubberyConfig{}
	ruberConf.LoadFromYaml(val)

	if updateCache {
		timeInText := conf.Api.DefaultTTL
		if ruberConf.Default.TTL != "" {
			timeInText = ruberConf.Default.TTL
		}
		u, _ := time.ParseDuration(timeInText)
		cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	}

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", ruberConf.Default.Value.Data.(interface{}))))

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
	ruberConf := feature.RubberyConfig{}
	ruberConf.LoadFromJsonBinary(b)

	bYaml, err := yaml.Marshal(ruberConf)

	featureSelected.Value = string(bYaml)

	timeInText := conf.Api.DefaultTTL
	if ruberConf.Default.TTL != "" {
		timeInText = ruberConf.Default.TTL
	}
	u, _ := time.ParseDuration(timeInText)
	errCache := cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	if !errCache {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errFea := source.CreateFeature(featureSelected)
	if !errFea {
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

	err := cacheValue.DeleteValue(featureSelected.Key)
	if err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = source.DeleteFeature(featureSelected)
	if err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
