package business

import (
	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func (bb Business) GetFeatureOnlyValue(vars map[string]string) (int, interface{}, string, error) {

	status, featureSelected, err := bb.getFeature(vars)
	if status == Success {
		finalresult, err := featureSelected.Value.GetFinalValue(vars)
		finaltype := featureSelected.Value.Default.Value.Type
		if err != nil {
			return Unknown, nil, "", err
		} else {
			return Success, finalresult, finaltype, err
		}
	} else {
		return status, nil, "", err
	}
}
func (bb Business) GetFeatureFull(vars map[string]string) (int, *feature.FeatureDefinition, error) {

	status, featureSelected, err := bb.getFeature(vars)
	if status == Success {
		return status, featureSelected.Value, nil
	} else {
		return status, nil, err
	}
}

func (bb Business) getFeature(vars map[string]string) (int, datasource.Feature, error) {
	_, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		return NotResult, featureSelected, nil
	}

	updateCacheFlag := false
	val, found, _ := cacheValue.GetValue(featureSelected.Key)
	if !found {
		found, err := source.GetFeature(&featureSelected)

		if err == nil && !found {
			return NoContent, featureSelected, nil
		}
		if err != nil {
			return Unknown, featureSelected, err
		}
		updateCacheFlag = true
	} else {
		featureSelected.Value = val
	}

	if updateCacheFlag {
		updateCache(featureSelected, cacheValue)
		/*timeInText := conf.Api.DefaultTTL
		if featureSelected.Value.Default.TTL != "" {
			timeInText = featureSelected.Value.Default.TTL
		}
		u, _ := time.ParseDuration(timeInText)
		cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
		*/
	}

	go metrics.GetMetrics().Update(featureSelected.Key)

	return Success, featureSelected, nil

}
