package cache

import (
	"sync"
	"time"
)

type itemInMemory struct {
	value    interface{}
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

func (nc *inMemoryClient) GetValue(key string) (interface{}, bool) {

	val, ok := nc.values[key]
	if !ok {
		return nil, true
	}
	currentTime := time.Now()
	diff := currentTime.Sub(val.initTime)

	if diff > val.ttl {
		return nil, true
	} else {
		return val.value, false
	}

}

func (nc *inMemoryClient) SetValue(key string, value interface{}, expiration time.Duration) bool {

	aux := itemInMemory{value: value, ttl: expiration, initTime: time.Now()}
	nc.values[key] = aux
	return true
}
func (nc *inMemoryClient) DeleteValue(key string) bool {

	delete(nc.values, key)
	return true
}
