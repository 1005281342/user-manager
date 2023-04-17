package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/1005281342/user-manager/db"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.8.42:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return &RedisCache{client: client}
}

func (c *RedisCache) Get(key string, value interface{}) error {
	data, err := c.client.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return db.ErrNoResult
	} else if err != nil {
		return err
	}
	return json.Unmarshal(data, value)
}

func (c *RedisCache) Set(key string, value interface{}, duration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(context.Background(), key, data, duration).Err()
}

func (c *RedisCache) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}
