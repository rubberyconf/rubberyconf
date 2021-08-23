package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"

	"github.com/matishsiao/goInfo"
)

type ElasticLogRepository struct {
}

type elasticDocs struct {
	Level     string
	Message   string
	Metainfo  interface{}
	TimeStamp time.Time
	OsInfo    *goInfo.GoInfoObject
}

func NewElasticRepository() *ElasticLogRepository {
	aux := new(ElasticLogRepository)
	return aux
}

func (metric *ElasticLogRepository) timeOut() time.Duration {
	timeout, err := time.ParseDuration(config.GetConfiguration().Elastic.TimeOut)
	if err != nil {
		timeout = 1 * time.Second
	}
	return timeout
}

func (me *ElasticLogRepository) connect() *elasticsearch.Client {

	conf := config.GetConfiguration()

	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.Elastic.Url,
		},
	}
	var err error
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return client

}

func jsonStruct(doc elasticDocs) string {

	// Create struct instance of the Elasticsearch fields struct object
	gi := goInfo.GetInfo()
	docStruct := &elasticDocs{
		Level:     doc.Level,
		Message:   doc.Message,
		TimeStamp: time.Now(), //.Format("2020-12-01 13:00:00") ,
		OsInfo:    gi,
	}
	docStruct.Metainfo = doc.Metainfo

	//fmt.Println("\ndocStruct:", docStruct)
	//fmt.Println("docStruct TYPE:", reflect.TypeOf(docStruct))

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return string(err.Error())
	}
	return string(b)
}

func (me *ElasticLogRepository) WriteMessage(ctx context.Context, level string, message string, metainfo interface{}) (bool, error) {

	conf := config.GetConfiguration()
	client := me.connect()

	doc := elasticDocs{Level: level, Message: message, Metainfo: metainfo}

	docStr := jsonStruct(doc)

	req := esapi.IndexRequest{
		Index:   conf.Elastic.Logs.Index,
		Body:    strings.NewReader(docStr),
		Refresh: "true",
	}
	ctxElastic, cancel := context.WithTimeout(ctx, me.timeOut())
	res, err := req.Do(ctxElastic, client)
	if err != nil && res.StatusCode == 201 {
		log.Fatalf("Error getting response from Elastic error: %v", err)
		cancel()
		return false, err
	}
	defer res.Body.Close()
	cancel()

	if res.IsError() {
		log.Printf("%s ERROR indexing document in Elastic message= %s", res.Status(), message)
		return false, nil
	}
	return true, nil

}
