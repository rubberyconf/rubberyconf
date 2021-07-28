package logs

import (
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

type iLogs interface {
	WriteMessage(message string)
}

const (
	CONSOLE string = "Console"
	ELASTIC string = "Elastic"
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

func (logs *Logs) WriteMessage(message string) {

	for _, lg := range logs.logs {
		lg.WriteMessage(message)
	}
}
