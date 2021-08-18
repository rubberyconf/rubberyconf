package service

import "context"

func (bb Service) DeleteFeature(ctx context.Context, vars map[string]string) (int, error) {

	_, cacheValue, source, featureSelected, result := preRequisites(ctx, vars)

	if !result {
		return NotResult, nil
	}

	//res, err :=
	cacheValue.DeleteValue(ctx, featureSelected.Key)
	//if res {
	//	return Unknown, err
	//}
	res := source.DeleteFeature(ctx, featureSelected)
	if res {
		return Unknown, nil
	}

	return Success, nil

}
