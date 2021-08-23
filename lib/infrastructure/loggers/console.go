package loggers

import (
	"encoding/json"
	"log"

	"github.com/rubberyconf/rubberyconf/lib/core/logs"
)

type ConsoleLog struct {
}

const (
	CONSOLE string = "Console"
)

func NewConsoleLog() *logs.ILogs {

	var llg logs.ILogs
	consoleLogging := new(ConsoleLog)
	llg = consoleLogging
	return &llg
}

func (lg *ConsoleLog) WriteMessage(level string, message string, metainfo interface{}) {
	if metainfo == nil {
		log.Printf("%s - %s", level, message)
	} else {
		bytearr, _ := json.Marshal(metainfo)
		log.Printf("%s - %s \n %v \n", level, message, string(bytearr))
	}
}
