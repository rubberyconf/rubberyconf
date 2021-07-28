package datasource

import (
	"log"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

type IDataSource interface {
	GetFeature(feature string) (interface{}, bool)
	DeleteFeature(feature string) bool
	CreateFeature(feature string, featureDescription interface{}) bool
}

const (
	GOGS     string = "Gogs"
	GITHUB   string = "GitHub"
	INMEMORY string = "InMemory"
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

	return res
}
