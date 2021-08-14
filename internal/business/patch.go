package business

import (
	"github.com/rubberyconf/rubberyconf/internal/feature"

	"github.com/imdario/mergo"
)

func (bb Business) PatchFeature(vars map[string]string, ruberConf feature.FeatureDefinition) (int, error) {

	_, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		return NotResult, nil
	}

	status, featureDefOriginal, err := bb.GetFeatureFull(vars)

	if status != Success {
		return NoContent, err
	}

	if err := mergo.Merge(featureDefOriginal, ruberConf, mergo.WithOverride); err != nil {
		return Unknown, err
	}

	featureSelected.Value = featureDefOriginal

	res := updateCache(featureSelected, cacheValue)
	if !res {
		return Unknown, nil
	}

	source.DeleteFeature(featureSelected)
	res = source.CreateFeature(featureSelected)
	if !res {
		return Unknown, nil
	}
	return Success, nil

}
