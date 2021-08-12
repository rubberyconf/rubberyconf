package rules

import "container/list"

type QueryParam struct {
	Key   string   `yaml:"key"`
	Value []string `yaml:"value"`
}

type RuleQueryString struct {
}

func (me *RuleQueryString) CheckRule(r FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {

	if r.QueryString.Key != "" {
		ok := me.evaluate(r.QueryString, vars)
		if ok {
			matches.PushBack("querystring")
			*total++
		}
		return ok, false
	}
	return false, true
}
func (me *RuleQueryString) evaluate(q QueryParam, vars map[string]string) bool {

	value, ok := vars[q.Key]
	if !ok {
		return false
	}

	for _, candidate := range q.Value {
		if candidate == value {
			return true
		}

	}

	return false
}
