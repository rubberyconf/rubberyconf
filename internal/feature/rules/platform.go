package rules

import (
	"container/list"
)

type RulePlatform struct {
}

func (me *RulePlatform) CheckRule(f FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var strBased StringArrayBased
	key := "platform"
	currentValue := vars[key]
	return strBased.CheckRule(f.Platform, currentValue, matches, total, key)
}
