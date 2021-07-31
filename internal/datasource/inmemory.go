package datasource

import (
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
func (source *DataSourceInMemory) GetFeature(feature *Feature) bool {
	var ok bool
	feature.Value, ok = source.features[feature.Key]
	return ok
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
	source.features[feature.Key] = feature.Value
	return true
}

func (source *DataSourceInMemory) EnableFeature(keys map[string]string) (Feature, bool) {
	fe1 := Feature{Key: "", Value: nil}
	key := keys[keyFeature]
	if key == "" {
		return fe1, false
	}
	fe1.Key = key
	return fe1, true
}
