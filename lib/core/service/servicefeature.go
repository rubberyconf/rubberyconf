package service

import (
	"context"
	"time"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	inputPort "github.com/rubberyconf/rubberyconf/lib/core/ports/input"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

type ServiceFeature struct {
	repository output.IMetricsRepository
	datasource output.IDataSource
	cache      output.ICacheStorage
}

func NewServiceFeature(
	repository output.IMetricsRepository,
	datasource output.IDataSource,
	cache output.ICacheStorage) inputPort.IServiceFeature {

	service := new(ServiceFeature)
	service.repository = repository
	service.datasource = datasource
	service.cache = cache

	return service
}

/*
func preRequisites(ctx context.Context, vars map[string]string) (*config.Config, output.ICacheStorage, *output.IDataSource, output.FeatureKeyValue, bool) {
	//conf := config.GetConfiguration()
	//cacheValue := cache.SelectCache(conf)
	//source := datasource.NewDataSourceSource(ctx)

	feature, result := source.EnableFeature(vars)

	return conf, cacheValue, source, feature, result
}
*/

func (service *ServiceFeature) updateCache(ctx context.Context, featureSelected output.FeatureKeyValue) bool {
	conf := config.GetConfiguration()
	timeInText := conf.Api.DefaultTTL
	if featureSelected.Value.Default.TTL != "" {
		timeInText = featureSelected.Value.Default.TTL
	}
	u, _ := time.ParseDuration(timeInText)
	res, _ := service.cache.SetValue(ctx, featureSelected, time.Duration(u.Seconds()))
	if !res {
		return false
	}
	return true
}
