package service

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
)

func (bb *ServiceFeature) CreateFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (int, error) {

	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	featureSelected, result := bb.datasource.EnableFeature(vars)

	if !result {
		return NotResult, nil
	}

	featureSelected.Value = &ruberConf

	res := bb.updateCache(ctx, featureSelected)
	if !res {
		return Unknown, nil
	}

	res = bb.datasource.CreateFeature(ctx, featureSelected)
	if !res {
		return Unknown, nil
	}
	return Success, nil

}
