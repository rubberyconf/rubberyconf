package cache

import (
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
)

type testCase struct {
	key      string
	value    *feature.FeatureDefinition
	duration time.Duration
	want     *feature.FeatureDefinition
	found    bool
}

func getCommonScenarios() []testCase {

	var feDef feature.FeatureDefinition

	feDef.Name = "name"
	feDef.Meta.Description = "test description"
	feDef.Default.Value.Data = "value"
	feDef.Default.Value.Type = "string"
	feDef.Default.TTL = "3s"

	var tests = []testCase{
		{"key1", &feDef, 1 * time.Second, &feDef, true},      // retrieve value with corret TTL
		{"key2", &feDef, 400 * time.Millisecond, nil, false}, //retrieve value with TTL completed
	}

	return tests
}
