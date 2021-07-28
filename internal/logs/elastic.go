package logs

import (
	"context"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/rubberyconf/rubberyconf/internal/config"
)

type ElasticLog struct {
	es    *elasticsearch.Client
	index string
}

var (
	elasticLogging *ElasticLog
	elasticLogOnce sync.Once
)

func NewElasticLog() *ElasticLog {

	elasticLogOnce.Do(func() {
		elasticLogging = new(ElasticLog)
		conf := config.GetConfiguration()

		cfg := elasticsearch.Config{
			Addresses: []string{
				conf.Elastic.Url,
			},
		}
		var err error
		elasticLogging.es, err = elasticsearch.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		elasticLogging.index = conf.Elastic.Logs.Index

	})
	return elasticLogging
}

func (eg *ElasticLog) WriteMessage(message string) {

	req := esapi.IndexRequest{
		Index: eg.index,
		//DocumentID: strconv.Itoa(i + 1),
		Body:    strings.NewReader(message),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), eg.es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

}
