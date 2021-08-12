package rules

import (
	"container/list"
)

type RuleUserGroup struct {
}

func (me *RuleUserGroup) CheckRule(f FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var strBased StringArrayBased
	key := "usergroup"
	currentValue := vars[key]
	return strBased.CheckRule(f.UserGroup, currentValue, matches, total, key)
}
