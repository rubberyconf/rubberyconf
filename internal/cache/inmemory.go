package cache

import (
	"context"
	"sync"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/feature"
)

type itemInMemory struct {
	value    *feature.FeatureDefinition
	ttl      time.Duration
	initTime time.Time
}

type inMemoryClient struct {
	values map[string]itemInMemory
}

var (
	inMemClient  *inMemoryClient
	onceInMemory sync.Once
)

func NewDataStorageInMemory() *inMemoryClient {

	onceInMemory.Do(func() {
		inMemClient = new(inMemoryClient)
		inMemClient.values = make(map[string]itemInMemory)
	})
	return inMemClient
}

func (nc *inMemoryClient) GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error) {
	var found bool
	var content itemInMemory

	if content, found = nc.values[key]; !found {
		return nil, false, nil
	}
	currentTime := time.Now()
	diff := currentTime.Sub(content.initTime)

	if diff > content.ttl {
		return nil, false, nil
	} else {
		return content.value, true, nil
	}

}

func (nc *inMemoryClient) SetValue(ctx context.Context, key string, value *feature.FeatureDefinition, expiration time.Duration) (bool, error) {

	aux := itemInMemory{value: value, ttl: expiration, initTime: time.Now()}
	nc.values[key] = aux
	return true, nil
}
func (nc *inMemoryClient) DeleteValue(ctx context.Context, key string) (bool, error) {

	delete(nc.values, key)
	return true, nil
}
