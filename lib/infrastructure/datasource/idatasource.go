package datasource

import (
	"log"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

/*
type IDataSource interface {
	EnableFeature(aux map[string]string) (output.FeatureKeyValue, bool)
	GetFeature(ctx context.Context, feature *output.FeatureKeyValue) (bool, error)
	DeleteFeature(ctx context.Context, feature output.FeatureKeyValue) bool
	CreateFeature(ctx context.Context, feature output.FeatureKeyValue) bool
	reviewDependencies()
}*/

const (
	GOGS       string = "Gogs"
	MONGODB    string = "Mongodb"
	GITHUB     string = "GitHub"
	INMEMORY   string = "InMemory"
	keyFeature string = "feature"
)

func NewDataSourceSource() *output.IDataSource {

	conf := config.GetConfiguration()
	var res output.IDataSource
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
	res.ReviewDependencies()
	dataSourceEnabled := &res

	return dataSourceEnabled
}
