package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/internal/cache"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/configurations"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func Configuration(w http.ResponseWriter, r *http.Request) {

	conf := config.GetConfiguration()
	storage := cache.SelectStorage(conf)

	vars := mux.Vars(r)
	featureSelected := vars["feature"]
	if featureSelected == "" {
		log.Printf("no feature specified")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateCache := false
	val, err := storage.GetValue(featureSelected)
	if err {
		source := datasource.SelectSource()
		if conf.Api.Source == datasource.INMEMORY {
			val, _ = source.GetFeature(featureSelected)
		} else {
			branch := vars["branch"]
			if branch == "" {
				branch = "master"
			}
			log.Printf("feature: %s in branch: %s requested...", featureSelected, branch)

			partialUrl := strings.Join([]string{branch, "/", featureSelected + ".yml"}, "")
			val, _ = source.GetFeature(partialUrl)
		}
		updateCache = true
	}
	ruberConf := feature.RubberyConfig{}
	ruberConf.Load(val)
	configurations.ParseConfiguration(&ruberConf) //TODO

	if updateCache {
		u, _ := time.ParseDuration(ruberConf.Default.TTL)
		storage.SetValue(featureSelected, val, time.Duration(u.Seconds()))
	}

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", ruberConf.Default.Value.Data.(interface{}))))

	metrics.GetMetrics().Update(featureSelected)
}
