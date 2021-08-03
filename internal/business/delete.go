package business

func (bb Business) DeleteFeature(vars map[string]string) (int, error) {

	_, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		return NotResult, nil
	}

	//res, err :=
	cacheValue.DeleteValue(featureSelected.Key)
	//if res {
	//	return Unknown, err
	//}
	res := source.DeleteFeature(featureSelected)
	if res {
		return Unknown, nil
	}

	return Success, nil

}
