package metrics

import (
	"fmt"
	"testing"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

func TestUpdateMetricsRegister(t *testing.T) {
	var tests = []struct {
		createdAt, updatedAt time.Time
		feature              string
		counter              int64
		expectedCounter      int64
	}{
		{time.Now(), time.Now(), "feature1", 0, 1},
		{time.Now(), time.Now(), "feature1", 10, 11},
		{time.Now(), time.Now(), "feature2", 0, 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s, counter: %d, expectedCounter: %d", tt.feature, tt.counter, tt.expectedCounter)
		t.Run(testname, func(t *testing.T) {
			var metric MongoMetrics
			metric.Feature = tt.feature
			metric.Counter = tt.counter
			metric.CreatedAt = tt.createdAt
			metric.UpdatedAt = tt.updatedAt
			metric.Update()
			if metric.Counter != tt.expectedCounter {
				t.Errorf("got %d, want %d", metric.Counter, tt.expectedCounter)
			}
		})
	}
}

func TestUpdateMetrics(t *testing.T) {
	var tests = []struct {
		feature         string
		expectedCounter int64
	}{
		{"feature1", 1},
		{"feature1", 2},
		{"feature2", 1},
		{"feature1", 3},
	}
	config.NewConfiguration("../../config/local.yml")
	metricsService := CreateMetrics()

	for _, tt := range tests {
		testname := fmt.Sprintf("feature: %s, expectedCounter: %d", tt.feature, tt.expectedCounter)
		t.Run(testname, func(t *testing.T) {
			res, err := metricsService.Update(tt.feature)
			if err != nil {
				t.Fatalf("error running the test %s", err)
			} else {
				if res.Counter != tt.expectedCounter {
					t.Errorf("got %d, want %d", res.Counter, tt.expectedCounter)
				}
			}
		})
	}

	for _, tt := range tests {
		metricsService.Remove(tt.feature)
	}

}
