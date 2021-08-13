package business

import (
	"time"

	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func (bb Business) GetFeatureOnlyValue(vars map[string]string) (int, interface{}, string) {

	status, featureSelected := bb.getFeature(vars)
	if status == Success {
		finalresult, err := featureSelected.Value.GetFinalValue(vars)
		finaltype := featureSelected.Value.Default.Value.Type
		if err != nil {
			return Unknown, nil, ""
		} else {
			return Success, finalresult, finaltype
		}
	} else {
		return status, nil, ""
	}
}
func (bb Business) GetFeatureFull(vars map[string]string) (int, *feature.FeatureDefinition) {

	status, featureSelected := bb.getFeature(vars)
	if status == Success {
		return status, featureSelected.Value
	} else {
		return status, nil
	}
}

func (bb Business) getFeature(vars map[string]string) (int, datasource.Feature) {
	conf, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		return NotResult, featureSelected
	}

	updateCache := false
	val, found, _ := cacheValue.GetValue(featureSelected.Key)
	if !found {
		found, err := source.GetFeature(&featureSelected)

		if err == nil && !found {
			return NoContent, featureSelected
		}
		if err != nil {
			return Unknown, featureSelected
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

	go metrics.GetMetrics().Update(featureSelected.Key)

	return Success, featureSelected

	/*finalresult, err := featureSelected.Value.GetFinalValue(vars)
	finaltype := featureSelected.Value.Default.Value.Type
	if err != nil {
		return Unknown, nil, ""
	} else {
		return Success, finalresult, finaltype
	}*/

}
