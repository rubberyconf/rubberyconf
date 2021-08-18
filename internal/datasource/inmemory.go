package datasource

import (
	"context"
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type DataSourceInMemory struct {
	mtx      sync.RWMutex
	features map[string]feature.FeatureDefinition
}

var (
	inMemDataSource *DataSourceInMemory
	onceInMemory    sync.Once
)

func NewDataSourceInMemory(ctx context.Context) *DataSourceInMemory {

	onceInMemory.Do(func() {
		inMemDataSource = new(DataSourceInMemory)
		inMemDataSource.features = make(map[string]feature.FeatureDefinition)
	})
	return inMemDataSource
}
func (source *DataSourceInMemory) GetFeature(ctx context.Context, feature *Feature) (bool, error) {
	var found bool
	source.mtx.Lock()
	defer source.mtx.Unlock()
	//var content feature.FeatureDefinition
	if content, found := source.features[feature.Key]; found {
		feature.Value = &content
	} else {
		feature.Value = nil
	}
	return found, nil
}

func (source *DataSourceInMemory) DeleteFeature(ctx context.Context, feature Feature) bool {
	source.mtx.Lock()
	defer source.mtx.Unlock()

	_, ok := source.features[feature.Key]
	if ok {
		delete(source.features, feature.Key)
		return true
	} else {
		return false
	}
}

func (source *DataSourceInMemory) CreateFeature(ctx context.Context, feature Feature) bool {
	source.mtx.Lock()
	defer source.mtx.Unlock()

	source.features[feature.Key] = *feature.Value
	return true
}

func (source *DataSourceInMemory) EnableFeature(keys map[string]string) (Feature, bool) {
	return enableFeature(keys)
}

func (source *DataSourceInMemory) reviewDependencies() {
}
