package logs

import (
	"sync"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
)

type ILogs interface {
	WriteMessage(level string, message string, metainfo interface{})
}

const (
	INFO    string = "info"
	DEBUG   string = "debug"
	ERROR   string = "error"
	WARNING string = "warning"
)

type Logs struct {
	logs map[string]*ILogs
}

var (
	allLogs  *Logs
	logsOnce sync.Once
)

func GetLogs() *Logs {

	logsOnce.Do(func() {
		allLogs = new(Logs)
		allLogs.logs = make(map[string]*ILogs)

	})
	return allLogs
}

func (logs *Logs) WriteMessage(level string, message string, metainfo interface{}) {

	if logs.checkLevel(level) {
		for _, lg := range logs.logs {
			(*lg).WriteMessage(level, message, metainfo)
		}
	}
}

func (logs *Logs) AddLog(key string, lg *ILogs) {
	logs.logs[key] = lg
}

func (logs *Logs) checkLevel(level string) bool {

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
