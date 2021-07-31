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

func (source *DataSourceGithub) GetFeature(feature *Feature) bool {
	log.Panicf("error github not implemented yet")
	return false
}

func (source *DataSourceGithub) DeleteFeature(feature Feature) bool {
	log.Panicf(errorMessage)
	return false
}

func (source *DataSourceGithub) CreateFeature(feature Feature) bool {

	log.Panicf(errorMessage)
	return false
}

func (source *DataSourceGithub) EnableFeature(keys map[string]string) (Feature, bool) {
	return gitEnableFeature(keys)
}
