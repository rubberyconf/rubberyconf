package logs

import (
	"log"
	"os"
	"sync"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/stringarr"
)

type iLogs interface {
	WriteMessage(level string, message string, metainfo interface{})
}

const (
	CONSOLE string = "Console"
	ELASTIC string = "Elastic"

	INFO    string = "info"
	DEBUG   string = "debug"
	ERROR   string = "error"
	WARNING string = "warning"
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

	if checkLevel(level) {
		for _, lg := range logs.logs {
			lg.WriteMessage(level, message, metainfo)
		}
	}
}

func reviewDependencies(conf *config.Config) {

	if stringarr.Include(conf.Api.Logs, ELASTIC) && conf.Elastic.Url == "" {
		log.Fatalf("elastic server dependency enabled but not configured, check config yml file.")
		os.Exit(2)
	}

}

func checkLevel(level string) bool {

	conf := config.GetConfiguration()
	levelThreshold := conf.Api.Options.LogLevel
	switch levelThreshold {
	case DEBUG:
		return true
	case INFO:
		if level == ERROR ||
			level == INFO ||
			level == WARNING {
			return true
		}
	case WARNING:
		if level == ERROR ||
			level == WARNING {
			return true
		}
	case ERROR:
		if level == ERROR {
			return true
		}
	}
	return false

}
