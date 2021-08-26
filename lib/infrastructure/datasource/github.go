package datasource

import (
	"context"
	"log"
	"os"
	"time"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

type DataSourceGithub struct {
}

func NewDataSourceGithub() *DataSourceGithub {
	return new(DataSourceGithub)
}

func (source *DataSourceGithub) connect(ctx context.Context) *github.Client {
	conf := config.GetConfiguration()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: conf.GitServer.ApiToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, _, err := client.Repositories.List(ctx, "", nil)
	source.checkErrors(err)

	return client
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
		logs.GetLogs().WriteMessage(logs.ERROR, "error getting access to github, rate limiting reached", err)
		return true
	} else if _, ok := err.(*github.AcceptedError); ok {
		logs.GetLogs().WriteMessage(logs.ERROR, "error on github side (scheduled), check api token", err)
		return true
	} else if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "error getting access to github, check api token", err)
		return true
	}
	return false
}

func (source *DataSourceGithub) GetFeature(ctx context.Context, feature output.FeatureKeyValue) (bool, error) {

	conf := config.GetConfiguration()

	ctxGitHub, cancel := context.WithTimeout(ctx, source.timeOut())
	client := source.connect(ctxGitHub)

	fc, _, _, err := client.Repositories.GetContents(
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

func (source *DataSourceGithub) DeleteFeature(ctx context.Context, feature output.FeatureKeyValue) bool {
	log.Panicf(errorMessage)
	return false
}

func (source *DataSourceGithub) CreateFeature(ctx context.Context, feature output.FeatureKeyValue) bool {

	conf := config.GetConfiguration()

	//fileContent := []byte("This is the content of my file\nand the 2nd line of it")
	out, err := yaml.Marshal(feature.Value)
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "error marshaling yaml", err)
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
	client := source.connect(ctxGitHub)
	_, _, err = client.Repositories.CreateFile(ctxGitHub, conf.GitServer.Organization, conf.GitServer.Repo, fileName, opts)
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "impossible create feature in github", err)
		cancel()
		return false
	}
	cancel()

	return false
}

func (source *DataSourceGithub) EnableFeature(keys map[string]string) (output.FeatureKeyValue, bool) {
	return gitEnableFeature(keys)
}
func (source *DataSourceGithub) ReviewDependencies() {
	reviewDependencies()
	conf := config.GetConfiguration()
	if conf.Api.Source == GOGS {
		if conf.GitServer.ApiToken == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not apitoken configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Username == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not username configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Email == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not email configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Organization == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not email configured, check config yml file.", nil)
			os.Exit(2)
		}
	}
}
