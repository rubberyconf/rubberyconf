package cache

import (
	"context"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

type SkipCache struct {
}

func NewDataStorageSkip() *SkipCache {

	skipped := new(SkipCache)
	return skipped
}

func (nc *SkipCache) GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error) {
	found := false
	return nil, found, nil
}

func (nc *SkipCache) DeleteValue(ctx context.Context, key string) (bool, error) {
	return true, nil
}

func (nc *SkipCache) SetValue(ctx context.Context, feature output.FeatureKeyValue, expiration time.Duration) (bool, error) {
	return true, nil
}
