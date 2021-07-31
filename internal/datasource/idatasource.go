package datasource

import (
	"log"
	"os"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

type Feature struct {
	Key   string
	Value interface{}
}

type IDataSource interface {
	EnableFeature(aux map[string]string) (Feature, bool)
	GetFeature(feature *Feature) bool
	DeleteFeature(feature Feature) bool
	CreateFeature(feature Feature) bool
}

const (
	GOGS       string = "Gogs"
	GITHUB     string = "GitHub"
	INMEMORY   string = "InMemory"
	keyFeature string = "feature"
)

func SelectSource() IDataSource {

	conf := config.GetConfiguration()
	reviewDependencies(conf)
	var res IDataSource
	typeSource := conf.Api.Source
	if typeSource == GOGS {
		res = NewDataSourceGogs()
	} else if typeSource == INMEMORY {
		res = NewDataSourceInMemory()
	} else if typeSource == GITHUB {
		res = NewDataSourceGithub()
	} else {
		log.Fatal("no data source selected")
	}

	return res
}

func reviewDependencies(conf *config.Config) {
	if (conf.Api.Source == GOGS || conf.Api.Source == GITHUB) &&
		conf.GitServer.Url == "" {
		log.Fatalf("git server dependency enabled but not configured, check config yml file.")
		os.Exit(2)
	}
}
