package api

import (
	"log"
	"time"
)

type IDataStorage interface {
	SetValue(key string, value interface{}, timeout time.Duration)
	GetValue(key string) (interface{}, bool)
}

func SelectStorage(conf *Config) IDataStorage {

	var res IDataStorage
	typeStorage := conf.Api.Type
	if typeStorage == "inMemory" {
		res = NewDataStorageInMemory()
	} else if typeStorage == "redis" {
		res = NewDataStorageRedis()
	} else {
		log.Fatal("no data storage selected")
	}

	return res
}
