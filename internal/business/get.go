package business

import (
	"time"

	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func (bb Business) GetFeature(vars map[string]string) (int, interface{}, string) {
	conf, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		return NotResult, nil, ""
	}

	updateCache := false
	val, found, _ := cacheValue.GetValue(featureSelected.Key)
	if !found {
		found, err := source.GetFeature(&featureSelected)

		if err == nil && !found {
			return NoContent, nil, ""
		}
		if err != nil {
			return Unknown, nil, ""
		}
		updateCache = true
	} else {
		featureSelected.Value = val
	}

	if updateCache {
		timeInText := conf.Api.DefaultTTL
		if featureSelected.Value.Default.TTL != "" {
			timeInText = featureSelected.Value.Default.TTL
		}
		u, _ := time.ParseDuration(timeInText)
		cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	}

	finalresult, err := featureSelected.Value.GetFinalValue(vars)
	finaltype := featureSelected.Value.Default.Value.Type
	if err != nil {
		return Unknown, nil, ""
	} else {
		go metrics.GetMetrics().Update(featureSelected.Key)
		return Success, finalresult, finaltype
	}

}
