package repositories

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

func TestMongoMatricsRepository_Creates(t *testing.T) {

	mets := output.Metrics{}
	mets.Feature = "featureTest" + time.Now().String()
	mets.Counter = 0
	mets.CreatedAt = time.Now()
	mets.UpdatedAt = time.Now()

	var tests = []struct {
		label   string
		metrics output.Metrics
	}{
		{mets.Feature, mets},
	}
	conf := config.GetConfiguration()
	if conf == nil {
		path, _ := os.Getwd()
		config.NewConfiguration(filepath.Join(path, "../../../config/local.yml"))
	}
	ctx := context.TODO()
	repo := NewMetricsRepository()

	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s ", tt.label)
		t.Run(testname, func(t *testing.T) {
			res, _ := repo.Store(ctx, &tt.metrics)
			if !res {
				t.Errorf("error in %s creation", tt.label)
			}
		})
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s ", tt.label)
		t.Run(testname, func(t *testing.T) {
			_, err := repo.Fetch(ctx, tt.label)
			if err != nil {
				t.Errorf("error in %s get", tt.label)
			}

		})
	}
	for _, tt := range tests {
		res, _ := repo.Remove(ctx, tt.label)
		if !res {
			t.Errorf("error in %s deletion", tt.label)
		}
	}
}
