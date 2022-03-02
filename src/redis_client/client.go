package redis_client

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClientInstance *redisClient

type redisClient struct {
	c *redis.Client
}

func GetRedisClientManager() *redisClient {
	return redisClientInstance
}

func init() {
	if redisClientInstance == nil {
		redisClientInstance = &redisClient{}
	}
}

func NewClient(addr string, password string) *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	redisClientInstance.c = rdb

	return redisClientInstance
}

func Close() error {
	return redisClientInstance.c.Close()
}

func Ping() (string, error) {
	err := redisClientInstance.c.Ping(ctx).Err()
	if err != nil {
		return "", err
	} else {
		return redisClientInstance.c.Ping(ctx).Val(), nil
	}
}

func Set(key, val string) error {
	err := redisClientInstance.c.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string) (string, error) {
	val, err := redisClientInstance.c.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func Incr(key string) (int64, error) {
	current, err := redisClientInstance.c.Incr(ctx, key).Result()
	if err != nil {
		return current, err
	}

	return current, nil
}
