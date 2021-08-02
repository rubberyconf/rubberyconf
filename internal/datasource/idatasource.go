package datasource

import (
	"log"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type Feature struct {
	Key   string
	Value *feature.FeatureDefinition
}

type IDataSource interface {
	EnableFeature(aux map[string]string) (Feature, bool)
	GetFeature(feature *Feature) (bool, error)
	DeleteFeature(feature Feature) bool
	CreateFeature(feature Feature) bool
	reviewDependencies(conf *config.Config)
}

const (
	GOGS       string = "Gogs"
	MONGODB    string = "Mongodb"
	GITHUB     string = "GitHub"
	INMEMORY   string = "InMemory"
	keyFeature string = "feature"
)

func SelectSource() IDataSource {

	conf := config.GetConfiguration()
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
	res.reviewDependencies(conf)

	return res
}
