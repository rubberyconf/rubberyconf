package metrics

import (
	"fmt"
	"testing"
	"time"

	//config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
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
			var metric output.Metrics
			metric.Feature = tt.feature
			metric.Counter = tt.counter
			metric.CreatedAt = tt.createdAt
			metric.UpdatedAt = tt.updatedAt
			UpdateValue(&metric)
			if metric.Counter != tt.expectedCounter {
				t.Errorf("got %d, want %d", metric.Counter, tt.expectedCounter)
			}
		})
	}
}
