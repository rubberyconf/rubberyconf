package service

import "context"

func (bb *ServiceFeature) DeleteFeature(ctx context.Context, vars map[string]string) (int, error) {

	//_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	//if !result {
	//	return NotResult, nil
	//}
	featureSelected := bb.datasource.EnableFeature(vars)
	//res, err :=
	bb.cache.DeleteValue(ctx, featureSelected.Key)
	//if res {
	//	return Unknown, err
	//}
	res := bb.datasource.DeleteFeature(ctx, featureSelected)
	if res {
		return Unknown, nil
	}

	return Success, nil

}
