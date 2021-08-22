package output

import (
	"context"
)

type IDataSource interface {
	EnableFeature(aux map[string]string) (FeatureKeyValue, bool)
	GetFeature(ctx context.Context, feature FeatureKeyValue) (bool, error)
	DeleteFeature(ctx context.Context, feature FeatureKeyValue) bool
	CreateFeature(ctx context.Context, feature FeatureKeyValue) bool
	ReviewDependencies()
}
