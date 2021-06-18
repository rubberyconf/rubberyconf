package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

func writeTraceLog(message string) {
	log.Printf("%s", message)
}
func writeTraceElastic(message string) {

	conf := config.GetConfiguration()

	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.Elastic.Url,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	req := esapi.IndexRequest{
		Index: conf.Elastic.Index,
		//DocumentID: strconv.Itoa(i + 1),
		Body:    strings.NewReader(message),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		message := fmt.Sprintf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))

		writeTraceLog(message)
		writeTraceElastic(message)

	})
}
