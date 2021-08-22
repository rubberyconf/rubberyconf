package logs

import (
	"encoding/json"
	"log"
	"sync"
)

type ConsoleLog struct {
}

var (
	consoleLogging *ConsoleLog
	consoleLogOnce sync.Once
)

func NewConsoleLog() *ConsoleLog {

	consoleLogOnce.Do(func() {
		consoleLogging = new(ConsoleLog)
	})
	return consoleLogging
}

func (lg *ConsoleLog) WriteMessage(level string, message string, metainfo interface{}) {
	if metainfo == nil {
		log.Printf("%s - %s", level, message)
	} else {
		bytearr, _ := json.Marshal(metainfo)
		log.Printf("%s - %s \n %v \n", level, message, string(bytearr))
	}
}
