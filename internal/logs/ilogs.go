package logs

import (
	"log"
	"os"
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/stringarr"
)

type iLogs interface {
	WriteMessage(level string, message string, metainfo interface{})
}

const (
	CONSOLE string = "Console"
	ELASTIC string = "Elastic"

	INFO  string = "warning"
	DEBUG string = "debug"
	ERROR string = "error"
)

type Logs struct {
	logs map[string]iLogs
}

var (
	allLogs  *Logs
	logsOnce sync.Once
)

func GetLogs() *Logs {

	logsOnce.Do(func() {
		conf := config.GetConfiguration()
		reviewDependencies(conf)
		allLogs = new(Logs)
		allLogs.logs = make(map[string]iLogs)

		for _, log := range conf.Api.Logs {
			switch log {
			case CONSOLE:
				allLogs.logs[CONSOLE] = NewConsoleLog()
			case ELASTIC:
				allLogs.logs[ELASTIC] = NewElasticLog()
			}
		}

	})
	return allLogs
}

func (logs *Logs) WriteMessage(level string, message string, metainfo interface{}) {

	for _, lg := range logs.logs {
		lg.WriteMessage(level, message, metainfo)
	}
}

func reviewDependencies(conf *config.Config) {

	if stringarr.Include(conf.Api.Logs, ELASTIC) && conf.Elastic.Url == "" {
		log.Fatalf("elastic server dependency enabled but not configured, check config yml file.")
		os.Exit(2)
	}

}
