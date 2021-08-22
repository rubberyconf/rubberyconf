package output

import (
	"context"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
)

type ICacheStorage interface {
	SetValue(ctx context.Context, feature FeatureKeyValue, timeout time.Duration) (bool, error)
	GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error)
	DeleteValue(ctx context.Context, key string) (bool, error)
}
