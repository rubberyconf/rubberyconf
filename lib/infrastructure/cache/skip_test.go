package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

func TestSkipCache(t *testing.T) {

	var tests = []testCase{
		{"key1", nil, 1 * time.Second, nil, false},
	}

	storage := NewDataStorageSkip()

	ctx := context.Background()
	for _, tt := range tests {
		testname := fmt.Sprintf("key: %s, duration: %d", tt.key, tt.duration)
		t.Run(testname, func(t *testing.T) {
			feature := output.FeatureKeyValue{Key: tt.key, Value: tt.value}
			found, err := storage.SetValue(ctx, feature, tt.duration)
			if err == nil && found == false {
				t.Errorf("error setting key %s", tt.key)
			}
			time.Sleep(100 * time.Millisecond)
			_, found, err = storage.GetValue(ctx, tt.key)
			if err != nil && found == tt.found {
				t.Errorf("got %t, want %t", found, tt.found)
			}
		})
	}
}
