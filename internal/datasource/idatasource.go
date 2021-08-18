package datasource

import (
	"context"
	"log"
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type Feature struct {
	Key   string
	Value *feature.FeatureDefinition
}

type IDataSource interface {
	EnableFeature(aux map[string]string) (Feature, bool)
	GetFeature(ctx context.Context, feature *Feature) (bool, error)
	DeleteFeature(ctx context.Context, feature Feature) bool
	CreateFeature(ctx context.Context, feature Feature) bool
	reviewDependencies()
}

const (
	GOGS       string = "Gogs"
	MONGODB    string = "Mongodb"
	GITHUB     string = "GitHub"
	INMEMORY   string = "InMemory"
	keyFeature string = "feature"
)

var (
	dataSourceEnabled *IDataSource
	onceDataSource    sync.Once
)

func NewDataSourceSource(ctx context.Context) *IDataSource {

	onceDataSource.Do(func() {
		conf := config.GetConfiguration()
		var res *IDataSource
		typeSource := conf.Api.Source
		if typeSource == GOGS {
			res = NewDataSourceGogs(ctx)
		} else if typeSource == INMEMORY {
			res = NewDataSourceInMemory(ctx)
		} else if typeSource == GITHUB {
			res = NewDataSourceGithub(ctx)
		} else {
			log.Fatal("no data source selected")
		}
		res.reviewDependencies()
		dataSourceEnabled = res
	})

	return dataSourceEnabled
}
