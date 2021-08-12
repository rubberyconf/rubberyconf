package rules

import (
	"container/list"
)

type RuleCountry struct {
}

func (me *RuleCountry) CheckRule(f FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var strBased StringArrayBased
	key := "country"
	currentValue := vars[key]
	return strBased.CheckRule(f.Country, currentValue, matches, total, key)
}
