package service

import (
	"context"

	inputPort "github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

func (bb *ServiceFeature) DeleteFeature(ctx context.Context, vars map[string]string) (inputPort.ServiceResult, error) {

	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	//if !result {
	//	return NotResult, nil
	//}
	featureSelected, _ := bb.datasource.EnableFeature(vars)
	//res, err :=
	bb.cache.DeleteValue(ctx, featureSelected.Key)
	//if res {
	//	return Unknown, err
	//}
	res := bb.datasource.DeleteFeature(ctx, featureSelected)
	if res {
		return inputPort.Unknown, nil
	}

	return inputPort.Success, nil

}
