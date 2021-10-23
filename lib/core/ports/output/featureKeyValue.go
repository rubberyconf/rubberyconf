package output

import (
	feature "github.com/rubberyconf/language/lib"
)

type FeatureKeyValue struct {
	Key   string
	Value *feature.FeatureDefinition
}
