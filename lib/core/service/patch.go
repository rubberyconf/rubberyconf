package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"

	"github.com/imdario/mergo"
)

func (bb *ServiceFeature) PatchFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (int, error) {

	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	/*if !result {
		return NotResult, nil
	}*/

	status, featureDefOriginal, err := bb.GetFeatureFull(ctx, vars)

	if status != Success {
		return NoContent, err
	}

	if err := mergo.Merge(featureDefOriginal, ruberConf, mergo.WithOverride); err != nil {
		return Unknown, err
	}
	var featureSelected output.FeatureKeyValue
	featureSelected.Value = featureDefOriginal

	res := bb.updateCache(ctx, featureSelected)
	if !res {
		return Unknown, nil
	}

	bb.datasource.DeleteFeature(ctx, featureSelected)
	res = bb.datasource.CreateFeature(ctx, featureSelected)
	if !res {
		return Unknown, nil
	}
	return Success, nil

}
