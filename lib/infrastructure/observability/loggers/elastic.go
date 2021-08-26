package loggers

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/infrastructure/repositories"
)

type ElasticLog struct {
	repository *repositories.ElasticLogRepository
}

const (
	ELASTIC string = "Elastic"
)

func NewElasticLog() *logs.ILogs {

	var llg logs.ILogs
	consoleLogging := new(ElasticLog)
	consoleLogging.repository = repositories.NewElasticRepository()
	llg = consoleLogging
	return &llg
}

func (lg *ElasticLog) WriteMessage(level logs.LogTypeMessage, message string, metainfo interface{}) {
	lg.repository.WriteMessage(context.Background(), string(level), message, metainfo)
}

/*
func reviewDependencies(conf *config.Config) {

	if stringarr.Include(conf.Api.Logs, ELASTIC) && conf.Elastic.Url == "" {
		log.Fatalf("elastic server dependency enabled but not configured, check config yml file.")
		os.Exit(2)
	}
}
*/
