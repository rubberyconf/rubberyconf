package datasource

import (
	"context"
	"sync"

	feature "github.com/rubberyconf/language/lib"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

type DataSourceInMemory struct {
	mtx      sync.RWMutex
	features map[string]feature.FeatureDefinition
}

func NewDataSourceInMemory() *DataSourceInMemory {

	inMemDataSource := new(DataSourceInMemory)
	inMemDataSource.features = make(map[string]feature.FeatureDefinition)
	return inMemDataSource
}
func (source *DataSourceInMemory) GetFeature(ctx context.Context, feature output.FeatureKeyValue) (bool, error) {
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

func (source *DataSourceInMemory) DeleteFeature(ctx context.Context, feature output.FeatureKeyValue) bool {
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

func (source *DataSourceInMemory) CreateFeature(ctx context.Context, feature output.FeatureKeyValue) bool {
	source.mtx.Lock()
	defer source.mtx.Unlock()

	source.features[feature.Key] = *feature.Value
	return true
}

func (source *DataSourceInMemory) EnableFeature(keys map[string]string) (output.FeatureKeyValue, bool) {
	return enableFeature(keys)
}

func (source *DataSourceInMemory) ReviewDependencies() {
}
