package datasource

import (
	//"context"

	"context"
	"log"
	"os"
	"sync"

	"github.com/google/go-github/v38/github"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/logs"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

// TODO: to be implemented
type DataSourceGithub struct {
	//Url string
	client *github.Client
}

var (
	githubDataSource *DataSourceGithub
	onceGitHub       sync.Once
)

func NewDataSourceGithub() *DataSourceGithub {

	onceGitHub.Do(func() {
		conf := config.GetConfiguration()
		githubDataSource = new(DataSourceGithub)

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: conf.GitServer.ApiToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		githubDataSource.client = github.NewClient(tc)

		_, _, err := githubDataSource.client.Repositories.List(ctx, "", nil)
		githubDataSource.checkErrors(err)

	})
	return githubDataSource
}

func (source *DataSourceGithub) checkErrors(err error) bool {
	if _, ok := err.(*github.RateLimitError); ok {
		logs.GetLogs().WriteMessage("error", "error getting access to github, rate limiting reached", err)
		return true
	} else if _, ok := err.(*github.AcceptedError); ok {
		logs.GetLogs().WriteMessage("error", "error on github side (scheduled), check api token", err)
		return true
	} else if err != nil {
		logs.GetLogs().WriteMessage("error", "error getting access to github, check api token", err)
		return true
	}
	return false
}

func (source *DataSourceGithub) GetFeature(feature *Feature) (bool, error) {

	conf := config.GetConfiguration()
	//repo, response, err := source.client.Repositories.Get(context.Background(), conf.GitServer.Username, conf.GitServer.Repo)
	/*if ok := source.checkErrors(err); ok {
		return false, err
	}*/

	fc, _, _, err := source.client.Repositories.GetContents(
		context.Background(),
		conf.GitServer.Username,
		conf.GitServer.Repo,
		feature.Key,
		nil)

	if ok := source.checkErrors(err); ok {
		return false, err
	}

	strConten := fc.Content

	err = yaml.Unmarshal([]byte(*strConten), feature.Value)
	if err != nil {
		return false, err
	}

	return true, nil
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
func (source *DataSourceGithub) reviewDependencies() {
	reviewDependencies()
	conf := config.GetConfiguration()
	if conf.Api.Source == GOGS {
		if conf.GitServer.Username == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not username configured, check config yml file.", nil)
			os.Exit(2)
		}
	}
}
