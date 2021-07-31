package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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

type elasticDocs struct {
	Level    string
	Message  string
	Metainfo interface{}
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

func jsonStruct(doc elasticDocs) string {

	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &elasticDocs{
		Level:   doc.Level,
		Message: doc.Message,
	}
	docStruct.Metainfo = doc.Metainfo

	fmt.Println("\ndocStruct:", docStruct)
	fmt.Println("docStruct TYPE:", reflect.TypeOf(docStruct))

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return string(err.Error())
	}
	return string(b)
}

func (eg *ElasticLog) WriteMessage(level string, message string, metainfo interface{}) {

	doc := elasticDocs{Level: level, Message: message, Metainfo: metainfo}

	docStr := jsonStruct(doc)

	req := esapi.IndexRequest{
		Index:   eg.index,
		Body:    strings.NewReader(docStr),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), eg.es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("%s ERROR indexing document in Elastic message= %s", res.Status(), message)
	}

}
