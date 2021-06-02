package api

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestInRedisOptions(t *testing.T) {

	var tests = []struct {
		key, value string
		duration   time.Duration
		want       string
	}{
		{"key1", "hello1", 1 * time.Second, "hello1"},  // retrieve value with corret TTL
		{"key2", "hello2", 100 * time.Millisecond, ""}, //retrieve value with TTL completed
	}

	conf := GetConfiguration()
	if conf == nil {
		path, _ := os.Getwd()
		NewConfiguration(filepath.Join(path, "../config/local.yml"))
	}

	storage := NewDataStorageRedis()

	for _, tt := range tests {
		testname := fmt.Sprintf("key: %s, value: %s, duration: %d", tt.key, tt.value, tt.duration)
		t.Run(testname, func(t *testing.T) {
			storage.SetValue(tt.key, tt.value, tt.duration)
			time.Sleep(500 * time.Millisecond)
			res, err := storage.GetValue(tt.key)
			if !err && res != tt.want {
				t.Errorf("got '%s', want '%s'", res, tt.want)
			}
		})
	}
}
