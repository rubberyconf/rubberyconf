package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/metrics"
	inputPort "github.com/rubberyconf/rubberyconf/lib/core/ports/input"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

func (bb *ServiceFeature) GetFeatureOnlyValue(
	ctx context.Context,
	vars map[string]string) (inputPort.ServiceResult, interface{}, string, error) {

	status, featureSelected, err := bb.getFeature(ctx, vars)
	if status == inputPort.Success {
		finalresult, err := featureSelected.Value.GetFinalValue(vars)
		finaltype := featureSelected.Value.Default.Value.Type
		if err != nil {
			return inputPort.Unknown, nil, "", err
		} else {
			return inputPort.Success, finalresult, finaltype.String(), err
		}
	} else {
		return status, nil, "", err
	}
}
func (bb *ServiceFeature) GetFeatureFull(ctx context.Context, vars map[string]string) (inputPort.ServiceResult, *feature.FeatureDefinition, error) {

	status, featureSelected, err := bb.getFeature(ctx, vars)
	if status == inputPort.Success {
		return status, featureSelected.Value, nil
	} else {
		return status, nil, err
	}
}

func (bb *ServiceFeature) getFeature(ctx context.Context, vars map[string]string) (inputPort.ServiceResult, output.FeatureKeyValue, error) {
	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	//if !result {
	//	return NotResult, featureSelected, nil
	//}
	featureSelected, _ := bb.datasource.EnableFeature(vars)

	updateCacheFlag := false
	val, found, _ := bb.cache.GetValue(ctx, featureSelected.Key)
	if !found {
		found, err := bb.datasource.GetFeature(ctx, featureSelected)

		if err == nil && !found {
			return inputPort.NoContent, featureSelected, nil
		}
		if err != nil {
			return inputPort.Unknown, featureSelected, err
		}
		updateCacheFlag = true
	} else {
		featureSelected.Value = val
	}

	if updateCacheFlag {
		bb.updateCache(ctx, featureSelected)
	}

	go metrics.Update(ctx, featureSelected.Key, bb.repository)

	return inputPort.Success, featureSelected, nil

}
