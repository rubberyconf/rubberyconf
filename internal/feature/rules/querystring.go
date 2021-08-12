package rules

import "container/list"

type RuleQueryString struct {
}

func (me *RuleQueryString) CheckRule(r FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	var kk RuleKeyValue
	return kk.CheckRule(r.QueryString, vars, matches, total, "querystring")
}
