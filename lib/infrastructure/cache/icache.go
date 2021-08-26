package cache

import (
	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

const (
	MEMORY string = "InMemory"
	REDIS  string = "Redis"
	NONE   string = "None"
)

func NewCache() output.ICacheStorage {

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
		logs.GetLogs().WriteMessage(logs.ERROR, "no data storage selected", nil)
	}

	return res
}
