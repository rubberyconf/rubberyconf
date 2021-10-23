package service

import (
	"context"

	feature "github.com/rubberyconf/language/lib"
	inputPort "github.com/rubberyconf/rubberyconf/lib/core/ports/input"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"

	"github.com/imdario/mergo"
)

func (bb *ServiceFeature) PatchFeature(
	ctx context.Context, vars map[string]string,
	ruberConf feature.FeatureDefinition) (inputPort.ServiceResult, error) {

	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	/*if !result {
		return NotResult, nil
	}*/

	status, featureDefOriginal, err := bb.GetFeatureFull(ctx, vars)

	if status != inputPort.Success {
		return inputPort.NoContent, err
	}

	if err := mergo.Merge(featureDefOriginal, ruberConf, mergo.WithOverride); err != nil {
		return inputPort.Unknown, err
	}
	var featureSelected output.FeatureKeyValue
	featureSelected.Value = featureDefOriginal

	res := bb.updateCache(ctx, featureSelected)
	if !res {
		return inputPort.Unknown, nil
	}

	bb.datasource.DeleteFeature(ctx, featureSelected)
	res = bb.datasource.CreateFeature(ctx, featureSelected)
	if !res {
		return inputPort.Unknown, nil
	}
	return inputPort.Success, nil

}
