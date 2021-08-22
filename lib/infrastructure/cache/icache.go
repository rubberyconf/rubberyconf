package cache

import (
	"log"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

const (
	MEMORY string = "InMemory"
	REDIS  string = "Redis"
	NONE   string = "None"
)

func NewCache() *output.ICacheStorage {

	conf := config.GetConfiguration()
	var res output.ICacheStorage
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

	return &res
}
