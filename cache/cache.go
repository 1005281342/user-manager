package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Connect() error {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetClient() *redis.Client {
	return client
}
