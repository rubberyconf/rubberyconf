package datasource

import (
	//"context"

	"log"
	"sync"
)

// TODO: to be implemented
type DataSourceGithub struct {
	//Url string
}

var (
	githubDataSource *DataSourceGithub
	onceGitHub       sync.Once
)

func NewDataSourceGithub() *DataSourceGithub {

	onceGitHub.Do(func() {
		//conf := config.GetConfiguration()
		githubDataSource = new(DataSourceGithub)
		//githubDataSource.Url = strings.Join([]string{conf.GitServer.Url, "/raw/"}, "")
	})
	return githubDataSource
}

func (source *DataSourceGithub) GetFeature(partialUrl string) (interface{}, bool) {
	log.Panicf("error github not implemented yet")
	return nil, false
}

func (source *DataSourceGithub) DeleteFeature(feature string) bool {
	log.Panicf("error github not implemented yet")
	return false
}

func (source *DataSourceGithub) CreateFeature(feature string, featureDescription interface{}) bool {

	log.Panicf("error github not implemented yet")
	return false
}
