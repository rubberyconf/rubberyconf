package cache

import (
	"log"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

type IDataStorage interface {
	SetValue(key string, value interface{}, timeout time.Duration) bool
	GetValue(key string) (interface{}, bool)
}

const (
	MEMORY string = "InMemory"
	REDIS  string = "Redis"
	NONE   string = "None"
)

func SelectStorage(conf *config.Config) IDataStorage {

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
