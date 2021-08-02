package cache

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/logs"
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
			Addr:     conf.Redis.Url,      // "localhost:6379",
			Password: conf.Redis.Password, // no password set
			DB:       0,                   // use default DB
		})
	})
	return redisClnt
}

func (aux *redisClient) GetValue(key string) (*feature.FeatureDefinition, bool, error) {

	val, err := aux.rbd.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("error geting value from redis key: %s", key), err)
		return nil, false, err
	}
	var feat *feature.FeatureDefinition
	feat = new(feature.FeatureDefinition)
	err = feat.LoadFromString(val)
	if err != nil {
		return nil, false, err
	}
	return feat, true, nil

}

func (aux *redisClient) SetValue(key string, value *feature.FeatureDefinition, expiration time.Duration) (bool, error) {

	svalue, err := value.ToString()
	if err != nil {
		return false, err
	}
	err = aux.rbd.Set(ctx, key, svalue, expiration).Err()
	if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("error seting value to redis key: %s", key), err)
		return false, err
	} else {
		return true, nil
	}
}

func (aux *redisClient) DeleteValue(key string) (bool, error) {

	err := aux.rbd.Del(ctx, key).Err()
	if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("error seting value to redis key: %s", key), err)
		return false, err
	} else {
		return true, nil
	}
}
