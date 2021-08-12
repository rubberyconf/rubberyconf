package rules

import (
	"container/list"
	"os"

	str "github.com/rubberyconf/rubberyconf/internal/stringarr"
)

type RuleEnvironment struct {
}

func (me *RuleEnvironment) CheckRule(f FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	if len(f.Environment) > 0 {
		ok := me.evaluateEnv(f.Environment)
		if ok {
			matches.PushBack("environment")
			*total++
		}
		return ok, false
	}
	return false, true
}

func (me *RuleEnvironment) evaluateEnv(envs []string) bool {
	currentEnv := os.Getenv("ENV")
	return str.Include(envs, currentEnv)
}
