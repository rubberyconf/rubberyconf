package cache

import (
	"context"
	"sync"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

type itemInMemory struct {
	value    *feature.FeatureDefinition
	ttl      time.Duration
	initTime time.Time
}

type inProcessCache struct {
	mtx    sync.RWMutex
	values map[string]itemInMemory
}

func NewDataStorageInMemory() *inProcessCache {

	inMemClient := new(inProcessCache)
	inMemClient.values = make(map[string]itemInMemory)
	return inMemClient
}

func (nc *inProcessCache) GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error) {
	nc.mtx.Lock()
	defer nc.mtx.Unlock()

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

func (nc *inProcessCache) SetValue(ctx context.Context, feature output.FeatureKeyValue, expiration time.Duration) (bool, error) {

	nc.mtx.Lock()
	defer nc.mtx.Unlock()
	aux := itemInMemory{value: feature.Value, ttl: expiration, initTime: time.Now()}
	nc.values[feature.Key] = aux
	return true, nil
}
func (nc *inProcessCache) DeleteValue(ctx context.Context, key string) (bool, error) {

	nc.mtx.Lock()
	defer nc.mtx.Unlock()
	delete(nc.values, key)
	return true, nil
}
