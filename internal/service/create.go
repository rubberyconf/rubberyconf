package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/internal/feature"
)

func (bb Service) CreateFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (int, error) {

	_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	if !result {
		return NotResult, nil
	}

	featureSelected.Value = &ruberConf

	res := updateCache(ctx, featureSelected, cacheValue)
	if !res {
		return Unknown, nil
	}

	res = source.CreateFeature(ctx, featureSelected)
	if !res {
		return Unknown, nil
	}
	return Success, nil

}
