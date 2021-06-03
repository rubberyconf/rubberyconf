package datastorage

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rubberyconf/rubberyconf/internal/config"
)

type redisClient struct {
	rbd *redis.Client
}

var (
	ctx       = context.Background()
	redisClnt *redisClient
	onceRedis sync.Once
)

func NewDataStorageRedis() *redisClient {

	onceRedis.Do(func() {

		conf := config.GetConfiguration()

		if conf == nil {
			log.Fatalf("No configuration loaded")
		}

		redisClnt = new(redisClient)
		redisClnt.rbd = redis.NewClient(&redis.Options{
			Addr:     conf.Cache.Url,      // "localhost:6379",
			Password: conf.Cache.Password, // no password set
			DB:       0,                   // use default DB
		})
	})
	return redisClnt
}

func (aux *redisClient) GetValue(key string) (interface{}, bool) {

	val, err := aux.rbd.Get(ctx, key).Result()
	if err != nil {
		return "", true
	}
	return val, false

}

func (aux *redisClient) SetValue(key string, value interface{}, expiration time.Duration) {

	err := aux.rbd.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Fatalln(err)
	}
}
