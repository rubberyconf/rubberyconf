package business

import (
	"time"

	"github.com/rubberyconf/rubberyconf/internal/feature"
)

func (bb Business) CreateFeature(vars map[string]string, ruberConf feature.FeatureDefinition) (int, error) {

	conf, cacheValue, source, featureSelected, result := preRequisites(vars)

	if !result {
		return NotResult, nil
	}

	//ruberConf := feature.FeatureDefinition{}
	//ruberConf.LoadFromJsonBinary(b)

	featureSelected.Value = &ruberConf

	timeInText := conf.Api.DefaultTTL
	if ruberConf.Default.TTL != "" {
		timeInText = ruberConf.Default.TTL
	}
	u, _ := time.ParseDuration(timeInText)
	res, _ := cacheValue.SetValue(featureSelected.Key, featureSelected.Value, time.Duration(u.Seconds()))
	if !res {
		return Unknown, nil
	}

	res = source.CreateFeature(featureSelected)
	if !res {
		return Unknown, nil
	}
	//go metrics.GetMetrics().Update(featureSelected.Key)
	return Success, nil

}
