package service

import (
	"context"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/cache"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/datasource"
)

type Service struct {
}

const (
	NotResult = iota
	NoContent = iota
	Unknown   = iota
	Success   = iota
)

func preRequisites(ctx context.Context, vars map[string]string) (*config.Config, cache.IDataStorage, *datasource.IDataSource, datasource.Feature, bool) {
	conf := config.GetConfiguration()
	cacheValue := cache.SelectCache(conf)
	source := datasource.NewDataSourceSource(ctx)

	feature, result := source.EnableFeature(vars)

	return conf, cacheValue, source, feature, result
}

func updateCache(ctx context.Context, featureSelected datasource.Feature, cacheValue cache.IDataStorage) bool {
	conf := config.GetConfiguration()
	timeInText := conf.Api.DefaultTTL
	if featureSelected.Value.Default.TTL != "" {
		timeInText = featureSelected.Value.Default.TTL
	}
	u, _ := time.ParseDuration(timeInText)
	res, _ := cacheValue.SetValue(ctx, featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	if !res {
		return false
	}
	return true
}
