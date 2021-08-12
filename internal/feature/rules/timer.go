package rules

import (
	"container/list"
	"time"
)

type FeatureTimer struct {
	TriggerTime string `yaml:"triggerTime"`
}
type RuleTimer struct {
}

func (me *RuleTimer) CheckRule(r FeatureRule, vars map[string]string, matches *list.List, total *int) (bool, bool) {
	if len(r.FeatureTm.TriggerTime) > 0 {
		ok := me.evaluate(r.FeatureTm)
		if ok {
			matches.PushBack("timer")
			*total++
		}
		return ok, false
	}
	return false, true

}

func (me *RuleTimer) evaluate(t FeatureTimer) bool {

	inTimeSpan := func(check time.Time) bool {
		start, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
		today := time.Now()

		res := check.After(start) && check.Before(today)
		return res
	}

	const layout = "Jan 2, 2006 at 3:04pm (MST)"

	tm, err := time.Parse(layout, t.TriggerTime)
	if err != nil {
		//logs.GetLogs().WriteMessage("error", fmt.Sprintf("error parsing triggerTime: %s", t.TriggerTime), nil)
		return false
	}

	res := inTimeSpan(tm)
	return res

}
