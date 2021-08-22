package output

import (
	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
)

type FeatureKeyValue struct {
	Key   string
	Value *feature.FeatureDefinition
}
