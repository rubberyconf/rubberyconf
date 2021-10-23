package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	feature "github.com/rubberyconf/language/lib"
	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"
)

const (
	DEFAULTDB int = 0
)

type RedisCache struct {
}

func NewDataStorageRedis() *RedisCache {

	redisClnt := new(RedisCache)
	return redisClnt
}

func (aux *RedisCache) connect() *redis.Client {
	conf := config.GetConfiguration()
	rbd := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Url,
		Password: conf.Redis.Password,
		DB:       DEFAULTDB,
	})
	return rbd
}
func (aux *RedisCache) disconnect(client *redis.Client) {
	err := client.Close()
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "error closing redis client", err)
	}

}
func (aux *RedisCache) timeOut() time.Duration {
	timeout, err := time.ParseDuration(config.GetConfiguration().Redis.TimeOut)
	if err != nil {
		timeout = 1 * time.Second
	}
	return timeout
}
func (aux *RedisCache) GetValue(ctx context.Context, key string) (*feature.FeatureDefinition, bool, error) {

	rbd := aux.connect()
	defer aux.disconnect(rbd)
	ctxRedis, cancel := context.WithTimeout(ctx, aux.timeOut())
	val, err := rbd.Get(ctxRedis, key).Result()
	cancel()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("error geting value from redis key: %s", key), err)
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

func (aux *RedisCache) SetValue(ctx context.Context, feature output.FeatureKeyValue, expiration time.Duration) (bool, error) {

	svalue, err := feature.Value.ToString()
	if err != nil {
		return false, err
	}
	rbd := aux.connect()
	defer aux.disconnect(rbd)
	ctxRedis, cancel := context.WithTimeout(ctx, aux.timeOut())
	err = rbd.Set(ctxRedis, feature.Key, svalue, expiration).Err()
	cancel()
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("error seting value to redis key: %s", feature.Key), err)
		return false, err
	} else {
		return true, nil
	}
}

func (aux *RedisCache) DeleteValue(ctx context.Context, key string) (bool, error) {

	rbd := aux.connect()
	defer aux.disconnect(rbd)
	ctxRedis, cancel := context.WithTimeout(ctx, aux.timeOut())
	err := rbd.Del(ctxRedis, key).Err()
	cancel()
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("error seting value to redis key: %s", key), err)
		return false, err
	} else {
		return true, nil
	}
}
