package datasource

import (
	//"context"

	"context"
	"log"
	"os"
	"sync"
	"time"

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

func NewDataSourceGithub(ctx context.Context) *DataSourceGithub {

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
func (source *DataSourceGithub) timeOut() time.Duration {
	timeout, err := time.ParseDuration(config.GetConfiguration().GitServer.TimeOut)
	if err != nil {
		timeout = 1 * time.Second
	}
	return timeout
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

func (source *DataSourceGithub) GetFeature(ctx context.Context, feature *Feature) (bool, error) {

	conf := config.GetConfiguration()

	ctxGitHub, cancel := context.WithTimeout(ctx, source.timeOut())
	fc, _, _, err := source.client.Repositories.GetContents(
		ctxGitHub,
		conf.GitServer.Username,
		conf.GitServer.Repo,
		feature.Key,
		nil)
	cancel()

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

func (source *DataSourceGithub) DeleteFeature(ctx context.Context, feature Feature) bool {
	log.Panicf(errorMessage)
	return false
}

func (source *DataSourceGithub) CreateFeature(ctx context.Context, feature Feature) bool {

	conf := config.GetConfiguration()

	//fileContent := []byte("This is the content of my file\nand the 2nd line of it")
	out, err := yaml.Marshal(feature.Value)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "error marshaling yaml", err)
		return false
	}

	branch := "" // TODO: getting from querystring!!!! Where is it?
	if branch == "" {
		branch = conf.GitServer.BranchDefault
	}

	fileName := feature.Key + ".yml"
	// Note: the file needs to be absent from the repository as you are not
	// specifying a SHA reference here.
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("rubberyconf commit"),
		Content: out,
		Branch:  github.String(branch),
		Committer: &github.CommitAuthor{
			Name:  github.String("rubberyconf on behalf of " + conf.GitServer.Username),
			Email: github.String(conf.GitServer.Email)},
	}
	ctxGitHub, cancel := context.WithTimeout(ctx, source.timeOut())
	_, _, err = source.client.Repositories.CreateFile(ctxGitHub, conf.GitServer.Organization, conf.GitServer.Repo, fileName, opts)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "impossible create feature in github", err)
		cancel()
		return false
	}
	cancel()

	return false
}

func (source *DataSourceGithub) EnableFeature(keys map[string]string) (Feature, bool) {
	return gitEnableFeature(keys)
}
func (source *DataSourceGithub) reviewDependencies() {
	reviewDependencies()
	conf := config.GetConfiguration()
	if conf.Api.Source == GOGS {
		if conf.GitServer.ApiToken == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not apitoken configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Username == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not username configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Email == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not email configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Organization == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not email configured, check config yml file.", nil)
			os.Exit(2)
		}
	}
}
