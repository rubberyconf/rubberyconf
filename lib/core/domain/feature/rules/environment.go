package rules

import (
	"container/list"
	"os"
)

type RuleEnvironment struct {
}

func (me *RuleEnvironment) CheckRule(f FeatureRule, _ map[string]string, matches *list.List, total *int) (bool, bool) {

	var strBased StringArrayBased

	currentEnv := os.Getenv("ENV")
	return strBased.CheckRule(f.Environment, currentEnv, matches, total, "environment")
}
