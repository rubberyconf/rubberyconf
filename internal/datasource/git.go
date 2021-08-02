package datasource

import (
	"os"
	"strings"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/logs"
)

const (
	errorMessage = "error, it doesn't implemented yet, use git client in your source"
	gitbranch    = "branch"
	gitmaster    = "master"
)

func gitEnableFeature(keys map[string]string) (Feature, bool) {
	fe1 := Feature{Key: "", Value: nil}
	feature := keys[keyFeature]
	if feature == "" {
		return fe1, false
	}
	branch := keys[gitbranch]
	if branch == "" {
		branch = gitmaster
	}
	fe1.Key = strings.Join([]string{branch, "/", feature + ".yml"}, "")
	return fe1, true
}

func reviewDependencies(conf *config.Config) {
	if (conf.Api.Source == GOGS || conf.Api.Source == GITHUB) &&
		conf.GitServer.Url == "" {
		logs.GetLogs().WriteMessage("error", "git server dependency enabled but not configured, check config yml file.", nil)
		os.Exit(2)
	}
}
