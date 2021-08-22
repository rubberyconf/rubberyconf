package datasource

import "github.com/rubberyconf/rubberyconf/lib/core/ports/output"

func enableFeature(keys map[string]string) (output.FeatureKeyValue, bool) {
	fe1 := output.FeatureKeyValue{Key: "", Value: nil}
	key := keys[keyFeature]
	if key == "" {
		return fe1, false
	}
	fe1.Key = key
	return fe1, true
}
