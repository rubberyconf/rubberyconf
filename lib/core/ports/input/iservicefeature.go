package input

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
)

type IServiceFeature interface {
	CreateFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (int, error)
	DeleteFeature(ctx context.Context, vars map[string]string) (int, error)
	GetFeatureOnlyValue(ctx context.Context, vars map[string]string) (int, interface{}, string, error)
	GetFeatureFull(ctx context.Context, vars map[string]string) (int, *feature.FeatureDefinition, error)
	PatchFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (int, error)
}
