package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestInMemoryOptions(t *testing.T) {

	tests := getCommonScenarios()

	storage := NewDataStorageInMemory()

	for _, tt := range tests {
		testname := fmt.Sprintf("key: %s, value: %s, duration: %d", tt.key, tt.value, tt.duration)
		t.Run(testname, func(t *testing.T) {
			completed, err := storage.SetValue(tt.key, tt.value, tt.duration)
			if !completed || err != nil {
				t.Errorf(" error storing value key: %s", tt.key)
			}
			time.Sleep(500 * time.Millisecond)
			_, found, err := storage.GetValue(tt.key)
			if err != nil && found == tt.found {
				t.Errorf("got %t, want %t", found, tt.found)
			}
		})
	}
}
