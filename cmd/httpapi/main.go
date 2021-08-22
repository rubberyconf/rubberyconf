package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rubberyconf/rubberyconf/lib/application/httpapi"
	"github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/service"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/cache"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/datasource"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/repositories"
)

func loadConfiguration() *configuration.Config {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "local"
	}

	pathToConfigFile := filepath.Join(path, fmt.Sprintf("../../config/%s.yml", environment))

	conf := configuration.NewConfiguration(pathToConfigFile)
	b, _ := json.MarshalIndent(conf, "", "   ")
	logs.GetLogs().WriteMessage("debug", fmt.Sprintf("Configuration loaded:\n%s\nEnvironment: %s ", string(b), environment), nil)
	return conf
}

func main() {

	loadConfiguration()

	repository := repositories.NewMetricsRepository()
	datasource := datasource.NewDataSourceSource()
	cache := cache.NewCache()

	logs:= []{"",""}

	service1 := service.NewServiceFeature(repository, datasource, cache)
	service1.SetLogs(logs)

	server := httpapi.NewHTTPServer(service1)

	server.Start()

}
