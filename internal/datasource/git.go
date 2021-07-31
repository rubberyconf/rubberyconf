package datasource

import "strings"

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
