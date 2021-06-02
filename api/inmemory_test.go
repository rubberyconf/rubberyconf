package api

import (
	"fmt"
	"testing"
	"time"
)

func TestInMemoryOptions(t *testing.T) {

	var tests = []struct {
		key, value string
		duration   time.Duration
		want       string
	}{
		{"key1", "hello1", 1 * time.Second, "hello1"},  // retrieve value with corret TTL
		{"key2", "hello2", 400 * time.Millisecond, ""}, //retrieve value with TTL completed
	}

	storage := NewDataStorageInMemory()

	for _, tt := range tests {
		testname := fmt.Sprintf("key: %s, value: %s, duration: %d", tt.key, tt.value, tt.duration)
		t.Run(testname, func(t *testing.T) {
			storage.SetValue(tt.key, tt.value, tt.duration)
			time.Sleep(500 * time.Millisecond)
			res, err := storage.GetValue(tt.key)
			if !err && res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}
}
