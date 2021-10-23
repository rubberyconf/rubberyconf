package service

import (
	"context"

	feature "github.com/rubberyconf/language/lib"
	inputPort "github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func (bb *ServiceFeature) CreateFeature(
	ctx context.Context,
	vars map[string]string,
	ruberConf feature.FeatureDefinition) (inputPort.ServiceResult, error) {

	featureSelected, result := bb.datasource.EnableFeature(vars)

	if !result {
		return inputPort.NotResult, nil
	}

	featureSelected.Value = &ruberConf

	res := bb.updateCache(ctx, featureSelected)
	if !res {
		return inputPort.Unknown, nil
	}

	res = bb.datasource.CreateFeature(ctx, featureSelected)
	if !res {
		return inputPort.Unknown, nil
	}
	return inputPort.Success, nil

}
