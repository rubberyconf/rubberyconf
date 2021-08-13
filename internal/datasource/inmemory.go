package datasource

import (
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type DataSourceInMemory struct {
	features map[string]feature.FeatureDefinition
}

var (
	inMemDataSource *DataSourceInMemory
	onceInMemory    sync.Once
)

func NewDataSourceInMemory() *DataSourceInMemory {

	onceInMemory.Do(func() {
		inMemDataSource = new(DataSourceInMemory)
		inMemDataSource.features = make(map[string]feature.FeatureDefinition)
	})
	return inMemDataSource
}
func (source *DataSourceInMemory) GetFeature(feature *Feature) (bool, error) {
	var found bool
	//var content feature.FeatureDefinition
	if content, found := source.features[feature.Key]; found {
		feature.Value = &content
	} else {
		feature.Value = nil
	}
	return found, nil
}

func (source *DataSourceInMemory) DeleteFeature(feature Feature) bool {
	_, ok := source.features[feature.Key]
	if ok {
		delete(source.features, feature.Key)
		return true
	} else {
		return false
	}
}

func (source *DataSourceInMemory) CreateFeature(feature Feature) bool {
	source.features[feature.Key] = *feature.Value
	return true
}

func (source *DataSourceInMemory) EnableFeature(keys map[string]string) (Feature, bool) {
	return enableFeature(keys)
}

func (source *DataSourceInMemory) reviewDependencies() {
}
