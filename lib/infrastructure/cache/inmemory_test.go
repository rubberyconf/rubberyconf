package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

func TestInMemoryOptions(t *testing.T) {

	tests := getCommonScenarios()

	storage := NewDataStorageInMemory()
	ctx := context.Background()
	for _, tt := range tests {
		testname := fmt.Sprintf("key: %s, value: %s, duration: %d", tt.key, tt.value, tt.duration)
		t.Run(testname, func(t *testing.T) {
			feature := output.FeatureKeyValue{Key: tt.key, Value: tt.value}
			completed, err := storage.SetValue(ctx, feature, tt.duration)
			if !completed || err != nil {
				t.Errorf(" error storing value key: %s", tt.key)
			}
			time.Sleep(500 * time.Millisecond)
			_, found, err := storage.GetValue(ctx, tt.key)
			if err != nil && found == tt.found {
				t.Errorf("got %t, want %t", found, tt.found)
			}
		})
	}
}
