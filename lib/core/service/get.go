package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/metrics"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

func (bb *ServiceFeature) GetFeatureOnlyValue(ctx context.Context, vars map[string]string) (int, interface{}, string, error) {

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
func (bb *ServiceFeature) GetFeatureFull(ctx context.Context, vars map[string]string) (int, *feature.FeatureDefinition, error) {

	status, featureSelected, err := bb.getFeature(ctx, vars)
	if status == Success {
		return status, featureSelected.Value, nil
	} else {
		return status, nil, err
	}
}

func (bb *ServiceFeature) getFeature(ctx context.Context, vars map[string]string) (int, output.FeatureKeyValue, error) {
	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	//if !result {
	//	return NotResult, featureSelected, nil
	//}
	featureSelected := bb.datasource.EnableFeature(vars)

	updateCacheFlag := false
	val, found, _ := bb.cache.GetValue(ctx, featureSelected.Key)
	if !found {
		found, err := bb.datasource.GetFeature(ctx, &featureSelected)

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
		bb.updateCache(ctx, featureSelected)
	}

	go metrics.NewMetrics().Update(ctx, featureSelected.Key)

	return Success, featureSelected, nil

}
