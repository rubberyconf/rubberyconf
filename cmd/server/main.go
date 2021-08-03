package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rubberyconf/rubberyconf/api"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/logs"
)

func loadConfiguration(path string) *config.Config {

	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "local"
	}

	conf := config.NewConfiguration(filepath.Join(path, fmt.Sprintf("../../config/%s.yml", environment)))
	b, _ := json.MarshalIndent(conf, "", "   ")
	logs.GetLogs().WriteMessage("debug", fmt.Sprintf("Configuration loaded:\n%s\nEnvironment: %s ", string(b), environment), nil)
	return conf
}

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	conf := loadConfiguration(path)

	router := api.NewRouter()
	datasource.SelectSource()

	logs.GetLogs().WriteMessage("info", fmt.Sprintf("rubberyconf api started at port: %s", conf.Api.Port), nil)

	log.Fatal(http.ListenAndServe(":"+conf.Api.Port, router))

}
