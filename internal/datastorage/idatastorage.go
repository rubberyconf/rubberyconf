package datastorage

import (
	"log"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

type IDataStorage interface {
	SetValue(key string, value interface{}, timeout time.Duration)
	GetValue(key string) (interface{}, bool)
}

const (
	MEMORY string = "inMemory"
	REDIS  string = "redis"
)

func SelectStorage(conf *config.Config) IDataStorage {

	var res IDataStorage
	typeStorage := conf.Api.Type
	if typeStorage == MEMORY {
		res = NewDataStorageInMemory()
	} else if typeStorage == REDIS {
		res = NewDataStorageRedis()
	} else {
		log.Fatal("no data storage selected")
	}

	return res
}
