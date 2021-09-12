package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/rubberyconf/rubberyconf/lib/application/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/lib/application/grpcapi/servers"
	"github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/service"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/cache"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/datasource"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/observability/loggers"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/repositories"

	"google.golang.org/grpc"
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

	loadLogs(conf)

	logs.GetLogs().WriteMessage("debug", fmt.Sprintf("Configuration loaded:\n%s\nEnvironment: %s ", string(b), environment), nil)
	return conf
}

func loadLogs(conf *configuration.Config) {

	for _, log := range conf.Api.Logs {
		switch log {
		case loggers.CONSOLE:
			logs.GetLogs().AddLog(loggers.CONSOLE, loggers.NewConsoleLog())
		case loggers.ELASTIC:
			logs.GetLogs().AddLog(loggers.ELASTIC, loggers.NewElasticLog())
		}
	}

}

func main() {

	conf := loadConfiguration()

	repository := repositories.NewMetricsRepository()
	datasource := datasource.NewDataSourceSource()
	cache := cache.NewCache()

	service := service.NewServiceFeature(repository, datasource, cache)

	address := "0.0.0.0:" + conf.Api.Port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	confServer := new(servers.ConfServer)
	confServer.SetService(service)
	grpcapipb.RegisterRubberyConfServiceServer(s, confServer)
	featureServer := new(servers.FeatureServer)
	featureServer.SetService(service)
	grpcapipb.RegisterRubberyFeatureServiceServer(s, featureServer)

	s.Serve(lis)
}
