package output

import (
	"context"
	"time"

	feature "github.com/rubberyconf/language/lib"
)

type ICacheStorage interface {
	SetValue(ctx context.Context, f FeatureKeyValue, timeout time.Duration) (bool, error)
	GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error)
	DeleteValue(ctx context.Context, key string) (bool, error)
}
