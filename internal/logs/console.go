package logs

import (
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

func (lg *ConsoleLog) WriteMessage(message string) {
	log.Printf("%s", message)
}
