package datasource

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
)

func TestDataSourceMongoDB_Creates(t *testing.T) {

	feDef := feature.FeatureDefinition{}
	feDef.Name = "name"
	feDef.Meta.Description = "test description"
	feDef.Default.Value.Data = "value"
	feDef.Default.Value.Type = feature.ValueText
	feDef.Default.TTL = "3s"

	var tests = []struct {
		feature    string
		featureDef feature.FeatureDefinition
	}{
		{"feature1", feDef},
	}
	conf := config.GetConfiguration()
	if conf == nil {
		path, _ := os.Getwd()
		config.NewConfiguration(filepath.Join(path, "../../config/local.yml"))
	}
	ctx := context.TODO()
	datasource := NewDataSourceMongoDB()

	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s ", tt.feature)
		t.Run(testname, func(t *testing.T) {
			fea := output.FeatureKeyValue{Key: tt.feature, Value: &tt.featureDef}
			res := datasource.CreateFeature(ctx, fea)
			if !res {
				t.Errorf("error in %s creation", tt.feature)
			}
		})
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s ", tt.feature)
		t.Run(testname, func(t *testing.T) {
			fea := output.FeatureKeyValue{Key: tt.feature, Value: nil} // &tt.featureDef}
			res, err := datasource.GetFeature(ctx, fea)
			if !res && err != nil && fea.Value == nil {
				t.Errorf("error in %s get", tt.feature)
			}
		})
	}
	for _, tt := range tests {
		fea := output.FeatureKeyValue{Key: tt.feature, Value: &tt.featureDef}
		res := datasource.DeleteFeature(ctx, fea)
		if !res {
			t.Errorf("error in %s deletion", tt.feature)
		}
	}
}
