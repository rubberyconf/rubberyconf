package business

import (
	"github.com/rubberyconf/rubberyconf/internal/cache"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
)

type Business struct {
}

/*type iBusiness interface {
	CreateFeature(vars map[string]string, b []byte) (int, error)
	DeleteFeature(vars map[string]string) (int, error)
	GetFeature(vars map[string]string) (int, interface{}, string)
}*/

const (
	NotResult = iota
	NoContent = iota
	Unknown   = iota
	Success   = iota
)

func preRequisites(vars map[string]string) (*config.Config, cache.IDataStorage, datasource.IDataSource, datasource.Feature, bool) {
	conf := config.GetConfiguration()
	cacheValue := cache.SelectCache(conf)
	source := datasource.SelectSource()

	feature, result := source.EnableFeature(vars)

	return conf, cacheValue, source, feature, result
}
