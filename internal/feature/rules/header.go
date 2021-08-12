package rules

import "container/list"

type RuleHeader struct {
}

func (me *RuleHeader) CheckRule(r FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var kk RuleKeyValue
	return kk.CheckRule(r.Header, vars, matches, total, "header")
}
