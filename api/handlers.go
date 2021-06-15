package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/datastorage"
	"github.com/rubberyconf/rubberyconf/internal/feature"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to RubberyConf.io!")
}

func Configuration(w http.ResponseWriter, r *http.Request) {

	conf := config.GetConfiguration()
	storage := datastorage.SelectStorage(conf)

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
		branch := vars["branch"]
		if branch == "" {
			branch = "master"
		}
		log.Printf("feature: %s in branch: %s requested...", featureSelected, branch)

		source := datasource.NewDataSourceGogs()

		//source.Url = strings.Join([]string{conf.GitServer.Url, "/raw/", branch, "/", featureSelected + ".yml"}, "")
		partialUrl := strings.Join([]string{branch, "/", featureSelected + ".yml"}, "")
		val, err = source.GetFeature(partialUrl)
		if err {
			return
		}
		updateCache = true
	}
	ruberConf := feature.RubberyConfig{}
	ruberConf.Load(val)

	if updateCache {
		u, _ := time.ParseDuration(ruberConf.Default.TTL)
		storage.SetValue(featureSelected, val, time.Duration(u.Seconds()))
	}

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", ruberConf.Default.Value.Data.(interface{}))))

}
