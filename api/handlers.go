package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to RubberyConfg.io!")
}

func ConfigurationGet(w http.ResponseWriter, r *http.Request) {

	conf := GetConfiguration()
	storage := SelectStorage(conf)

	vars := mux.Vars(r)
	feature := vars["feature"]
	if feature == "" {
		log.Printf("no feature specified")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateCache := false
	val, err := storage.GetValue(feature)
	if err {
		branch := vars["branch"]
		if branch == "" {
			branch = "master"
		}
		log.Printf("feature: %s in branch: %s requested...", feature, branch)

		val = getValueFromGitRepo(feature+".yml", branch)
		updateCache = true
	}
	ruberConf := RubberyConfig{}
	ruberConf.load(val)

	if updateCache {
		storage.SetValue(feature, val, time.Duration(ruberConf.TimeToLive.Value)*time.Second)
	}

	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", ruberConf.Value.(interface{}))))

}
