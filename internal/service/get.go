package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/internal/datasource"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/metrics"
)

func (bb Service) GetFeatureOnlyValue(ctx context.Context, vars map[string]string) (int, interface{}, string, error) {

	status, featureSelected, err := bb.getFeature(ctx, vars)
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
func (bb Service) GetFeatureFull(ctx context.Context, vars map[string]string) (int, *feature.FeatureDefinition, error) {

	status, featureSelected, err := bb.getFeature(ctx, vars)
	if status == Success {
		return status, featureSelected.Value, nil
	} else {
		return status, nil, err
	}
}

func (bb Service) getFeature(ctx context.Context, vars map[string]string) (int, datasource.Feature, error) {
	_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	if !result {
		return NotResult, featureSelected, nil
	}

	updateCacheFlag := false
	val, found, _ := cacheValue.GetValue(ctx, featureSelected.Key)
	if !found {
		found, err := source.GetFeature(ctx, &featureSelected)

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
		updateCache(ctx, featureSelected, cacheValue)
	}

	go metrics.GetMetrics().Update(ctx, featureSelected.Key)

	return Success, featureSelected, nil

}
