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
)

func gitEnableFeature(keys map[string]string) (Feature, bool) {
	fe1 := Feature{Key: "", Value: nil}
	conf := config.GetConfiguration()
	feature := keys[keyFeature]
	if feature == "" {
		return fe1, false
	}
	branch := keys[gitbranch]
	if branch == "" {
		branch = conf.GitServer.BranchDefault
	}
	fe1.Key = strings.Join([]string{branch, "/", feature + ".yml"}, "")
	return fe1, true
}

func reviewDependencies() {

	conf := config.GetConfiguration()
	if conf.Api.Source == GOGS || conf.Api.Source == GITHUB {
		if conf.GitServer.Url == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not url configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.BranchDefault == "" {
			logs.GetLogs().WriteMessage("error", "git server dependency enabled but not default branch configured, check config yml file.", nil)
			os.Exit(2)
		}
	}
}
