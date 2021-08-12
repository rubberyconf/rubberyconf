package rules

import (
	"container/list"
)

type RuleUserId struct {
}

func (me *RuleUserId) CheckRule(f FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var strBased StringArrayBased
	key := "userid"
	currentValue := vars[key]
	return strBased.CheckRule(f.UserId, currentValue, matches, total, key)
}
