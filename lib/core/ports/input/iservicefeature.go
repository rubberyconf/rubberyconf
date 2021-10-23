package input

import (
	"context"

	feature "github.com/rubberyconf/language/lib"
)

type ServiceResult int

const (
	NotResult ServiceResult = iota
	NoContent ServiceResult = iota
	Unknown   ServiceResult = iota
	Success   ServiceResult = iota
)

type IServiceFeature interface {
	CreateFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (ServiceResult, error)
	DeleteFeature(ctx context.Context, vars map[string]string) (ServiceResult, error)
	GetFeatureOnlyValue(ctx context.Context, vars map[string]string) (ServiceResult, interface{}, string, error)
	GetFeatureFull(ctx context.Context, vars map[string]string) (ServiceResult, *feature.FeatureDefinition, error)
	PatchFeature(ctx context.Context, vars map[string]string, ruberConf feature.FeatureDefinition) (ServiceResult, error)
}
