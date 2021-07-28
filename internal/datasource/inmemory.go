package datasource

import (
	//"context"

	"sync"
)

type DataSourceInMemory struct {
	features map[string]interface{}
}

var (
	inMemDataSource *DataSourceInMemory
	onceInMemory    sync.Once
)

func NewDataSourceInMemory() *DataSourceInMemory {

	onceInMemory.Do(func() {
		inMemDataSource = new(DataSourceInMemory)
		inMemDataSource.features = make(map[string]interface{})
	})
	return inMemDataSource
}
func (source *DataSourceInMemory) GetFeature(feature string) (interface{}, bool) {
	aux, ok := source.features[feature]
	if ok {
		return aux, true
	} else {
		return nil, false
	}
}

func (source *DataSourceInMemory) DeleteFeature(feature string) bool {
	_, ok := source.features[feature]
	if ok {
		delete(source.features, feature)
		return true
	} else {
		return false
	}
}

func (source *DataSourceInMemory) CreateFeature(feature string, featureDescription interface{}) bool {
	source.features[feature] = featureDescription
	return true
}
