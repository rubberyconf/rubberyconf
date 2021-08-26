package datasource

import (
	//"context"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

type DataSourceGogs struct {
	Url string
}

var (
	gogsDataSource *DataSourceGogs
	onceGogs       sync.Once
	//errorMessage  = "error gogs not implemented yet, use git client in your source"
)

func NewDataSourceGogs() *DataSourceGogs {

	onceGogs.Do(func() {
		conf := config.GetConfiguration()
		gogsDataSource = new(DataSourceGogs)
		gogsDataSource.Url = strings.Join([]string{conf.GitServer.Url, "/raw/"}, "")
	})
	return gogsDataSource
}

func (source *DataSourceGogs) GetFeature(ctx context.Context, feat output.FeatureKeyValue) (bool, error) {

	client := &http.Client{}
	finalURL := source.Url + feat.Key

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("imossible reach this host: %s", finalURL), err)
		feat.Value = nil
		return false, err
	}

	//req.Header.Add("Accept", "application/vnd.github.v3.raw")
	//req.Header.Add("authorization", "token 929f19719c9c9aac8c37c3a3766ebfce211cf5a9")
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("error reaching repo %s", finalURL)
		feat.Value = nil
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("error processing answer")
	}
	sb := string(body)
	//featureDef := stringToFeatureDefinition(sb)
	feat.Value = new(feature.FeatureDefinition)
	feat.Value.LoadFromString(sb)
	return true, nil

}

func (source *DataSourceGogs) DeleteFeature(ctx context.Context, feature output.FeatureKeyValue) bool {
	log.Panicf(errorMessage)
	return false
}

func (source *DataSourceGogs) CreateFeature(ctx context.Context, feature output.FeatureKeyValue) bool {
	log.Panicf(errorMessage)
	return false
}

func (source *DataSourceGogs) EnableFeature(keys map[string]string) (output.FeatureKeyValue, bool) {
	return gitEnableFeature(keys)
}

func (source *DataSourceGogs) ReviewDependencies() {
	reviewDependencies()
	conf := config.GetConfiguration()
	if conf.Api.Source == GOGS {
		if conf.GitServer.Url == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not url configured, check config yml file.", nil)
			os.Exit(2)
		}
	}
}
