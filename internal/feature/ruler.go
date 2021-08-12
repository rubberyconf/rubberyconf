package feature

import (
	"container/list"
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/feature/rules"
)

type BasicRuleMethods interface {
	CheckRule(f rules.FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool)
}

type RulerMaster struct {
	rules *list.List
}

var (
	onceRuleMaster sync.Once
	rulesEnabled   *RulerMaster
)

func NewRuleMaster() *RulerMaster {

	onceRuleMaster.Do(func() {

		rulesEnabled = new(RulerMaster)
		rulesEnabled.rules = list.New()
		rulesEnabled.RegisterRules()
	})
	return rulesEnabled
}

func (me RulerMaster) RegisterRules() {

	me.rules.PushBack(new(rules.RuleEnvironment))
	me.rules.PushBack(new(rules.RuleTimer))
	me.rules.PushBack(new(rules.RuleVersion))
	me.rules.PushBack(new(rules.RuleQueryString))
	me.rules.PushBack(new(rules.RuleHeader))

}

func (me RulerMaster) CheckRules(f rules.FeatureRule, vars map[string]string) (int, *list.List) {
	total := 0
	matches := list.New()

	for e := me.rules.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		aux := e.Value.(BasicRuleMethods)
		aux.CheckRule(f, vars, matches, &total)
	}
	return total, matches

}
