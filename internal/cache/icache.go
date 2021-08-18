package cache

import (
	"context"
	"log"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type IDataStorage interface {
	SetValue(ctx context.Context, key string, value *feature.FeatureDefinition, timeout time.Duration) (bool, error)
	GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error)
	DeleteValue(ctx context.Context, key string) (bool, error)
}

const (
	MEMORY string = "InMemory"
	REDIS  string = "Redis"
	NONE   string = "None"
)

func SelectCache(conf *config.Config) IDataStorage {

	var res IDataStorage
	typeStorage := conf.Api.Cache
	if typeStorage == MEMORY {
		res = NewDataStorageInMemory()
	} else if typeStorage == REDIS {
		res = NewDataStorageRedis()
	} else if typeStorage == NONE {
		res = NewDataStorageSkip()
	} else {
		log.Fatal("no data storage selected")
	}

	return res
}
