package cache

import (
	"sync"
	"time"
)

type skip struct {
}

var (
	skipped     *skip
	onceSkipped sync.Once
)

func NewDataStorageSkip() *skip {

	onceSkipped.Do(func() {
		skipped = new(skip)
	})
	return skipped
}

func (nc *skip) GetValue(key string) (interface{}, bool) {
	return nil, true
}

func (nc *skip) DeleteValue(key string) bool {
	return true
}

func (nc *skip) SetValue(key string, value interface{}, expiration time.Duration) bool {
	return true
}
