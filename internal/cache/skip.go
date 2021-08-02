package cache

import (
	"sync"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type skip struct {
}

var (
	skipped     *skip
	onceSkipped sync.Once
)

func NewDataStorageSkip() *skip {

	onceSkipped.Do(func() {
		skipped = new(skip)
	})
	return skipped
}

func (nc *skip) GetValue(key string) (*feature.FeatureDefinition, bool, error) {
	found := false
	return nil, found, nil
}

func (nc *skip) DeleteValue(key string) (bool, error) {
	return true, nil
}

func (nc *skip) SetValue(key string, value *feature.FeatureDefinition, expiration time.Duration) (bool, error) {
	return true, nil
}
