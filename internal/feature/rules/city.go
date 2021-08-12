package rules

import (
	"container/list"
)

type RuleCity struct {
}

func (me *RuleCity) CheckRule(f FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var strBased StringArrayBased
	key := "city"
	currentValue := vars[key]
	return strBased.CheckRule(f.City, currentValue, matches, total, key)
}
