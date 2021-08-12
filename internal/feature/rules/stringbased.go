package rules

import (
	"container/list"

	str "github.com/rubberyconf/rubberyconf/internal/stringarr"
)

type StringArrayBased struct {
}

func (me *StringArrayBased) CheckRule(potentialValues []string, currentValue string, matches *list.List, total *int, label string) (bool, bool) {

	if len(potentialValues) > 0 {
		ok := me.evaluate(potentialValues, currentValue)
		if ok {
			matches.PushBack(label)
			*total++
		}
		return ok, false
	}
	return false, true
}

func (me *StringArrayBased) evaluate(potentialValues []string, currentValue string) bool {
	return str.Include(potentialValues, currentValue)
}
