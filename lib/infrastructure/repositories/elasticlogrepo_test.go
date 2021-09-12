package repositories

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
)

func TestElasticLogRepository_Creates(t *testing.T) {

	var tests = []struct {
		label    string
		level    string
		sentence string
	}{
		{"log1", "debug", "test write this sentence in elasticIndex"},
	}
	conf := config.GetConfiguration()
	if conf == nil {
		path, _ := os.Getwd()
		config.NewConfiguration(filepath.Join(path, "../../../config/local.yml"))
	}
	ctx := context.TODO()
	//initElastic(ctx)
	repo := NewElasticRepository()

	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s ", tt.label)
		t.Run(testname, func(t *testing.T) {
			res, err := repo.WriteMessage(ctx, tt.level, tt.sentence, nil)
			if !res || err != nil {
				t.Errorf("error in %s creation", tt.label)
			}
		})
	}
}

func initElastic(ctx context.Context) (testcontainers.Container, string) {

	conf := config.GetConfiguration()
	urlSplit := strings.Split(conf.Elastic.Url, ":")
	esPort := urlSplit[2]
	elastic, err := startEsContainer(esPort, "9300")
	if err != nil {
		log.Fatalf("Could not start ES container: " + err.Error())
	}
	ip, err := elastic.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get host where the container is exposed: " + err.Error())
	}
	port, err := elastic.MappedPort(ctx, nat.Port(esPort))
	if err != nil {
		log.Fatalf("Could not retrive the mapped port: " + err.Error())
	}
	baseUrl := fmt.Sprintf("http://%s:%s", ip, port.Port())
	resp := createLogsIndexMapping(baseUrl)
	log.Printf("status code: %d", resp.StatusCode)
	return elastic, baseUrl
}

func createLogsIndexMapping(endpoint string) *http.Response {
	mapping := `
{"mappings" : {
  "properties" : {
    "level" : {
      "type" : "text"
    },
	"message" : {
      "type" : "text"
    },
    "id" : {
      "type" : "keyword"
    }
  }
}}`
	conf := config.GetConfiguration()
	return createIndexMapping(endpoint, conf.Elastic.Logs.Index, mapping)
}

func createIndexMapping(endpoint, indexName string, mappingsJson string) *http.Response {
	req, err := http.NewRequest(http.MethodPut, endpoint+"/"+indexName, bytes.NewBuffer([]byte(mappingsJson)))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalf("Could not create a mapping: " + indexName)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not create index mapping")
	}
	defer resp.Body.Close()

	return resp
}

func startEsContainer(restPort string, nodesPort string) (testcontainers.Container, error) {
	ctx := context.Background()

	rp := fmt.Sprintf("%s:%s/tcp", restPort, restPort)
	np := fmt.Sprintf("%s:%s/tcp", nodesPort, nodesPort)

	reqes5 := testcontainers.ContainerRequest{
		Image:        "elasticsearch:7.11.0",
		Name:         "es7-mock",
		Env:          map[string]string{"discovery.type": "single-node"},
		ExposedPorts: []string{rp, np},
		WaitingFor:   wait.ForLog("started"),
	}
	elastic, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: reqes5,
		Started:          true,
	})

	return elastic, err
}
