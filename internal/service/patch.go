package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/internal/feature"

	"github.com/imdario/mergo"
)

func (bb Service) PatchFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (int, error) {

	_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	if !result {
		return NotResult, nil
	}

	status, featureDefOriginal, err := bb.GetFeatureFull(ctx, vars)

	if status != Success {
		return NoContent, err
	}

	if err := mergo.Merge(featureDefOriginal, ruberConf, mergo.WithOverride); err != nil {
		return Unknown, err
	}

	featureSelected.Value = featureDefOriginal

	res := updateCache(ctx, featureSelected, cacheValue)
	if !res {
		return Unknown, nil
	}

	source.DeleteFeature(ctx, featureSelected)
	res = source.CreateFeature(ctx, featureSelected)
	if !res {
		return Unknown, nil
	}
	return Success, nil

}
