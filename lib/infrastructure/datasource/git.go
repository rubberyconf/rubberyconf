package datasource

import (
	"os"
	"strings"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

const (
	errorMessage = "error, it doesn't implemented yet, use git client in your source"
	gitbranch    = "branch"
)

func gitEnableFeature(keys map[string]string) (output.FeatureKeyValue, bool) {
	fe1 := output.FeatureKeyValue{Key: "", Value: nil}
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
		if conf.GitServer.BranchDefault == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not default branch configured, check config yml file.", nil)
			os.Exit(2)
		}
		if conf.GitServer.Repo == "" {
			logs.GetLogs().WriteMessage(logs.ERROR, "git server dependency enabled but not default branch configured, check config yml file.", nil)
			os.Exit(2)
		}
	}
}
