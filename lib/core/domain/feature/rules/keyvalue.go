package rules

import "container/list"

type RuleKeyValue struct {
}

func (me *RuleKeyValue) CheckRule(r KeyValue, vars map[string]string, matches *list.List, total *int, label string) (bool, bool) {

	if r.Key != "" {
		ok := me.evaluate(r, vars)
		if ok {
			matches.PushBack(label)
			*total++
		}
		return ok, false
	}
	return false, true
}
func (me *RuleKeyValue) evaluate(q KeyValue, vars map[string]string) bool {

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
