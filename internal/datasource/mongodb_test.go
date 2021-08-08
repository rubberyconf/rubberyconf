package datasource

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/feature"
)

func TestDataSourceMongoDB_Creates(t *testing.T) {

	feDef := feature.FeatureDefinition{}
	feDef.Name = "name"
	feDef.Meta.Description = "test description"
	feDef.Default.Value.Data = "value"
	feDef.Default.Value.Type = "string"
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

	datasource := NewDataSourceMongoDB()

	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s ", tt.feature)
		t.Run(testname, func(t *testing.T) {
			fea := Feature{Key: tt.feature, Value: &tt.featureDef}
			res := datasource.CreateFeature(fea)
			if !res {
				t.Errorf("error in %s creation", tt.feature)
			}
		})
	}
	for _, tt := range tests {
		fea := Feature{Key: tt.feature, Value: &tt.featureDef}
		res := datasource.DeleteFeature(fea)
		if !res {
			t.Errorf("error in %s deletion", tt.feature)
		}
	}
}
