package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func RedisConnect() {
	// Redis connect
	client := redis.NewClient(&redis.Options{
		Addr:     "microservice-redis:6379",
		Password: "",
		DB:       0,
	})

	// Redis Close
	defer func(client *redis.Client) {
		_ = client.Close()
	}(client)

}

func (r *Redis) RedisWrite() error {
	return r.client.Set(context.Background(), "example-key", "example-value", 0).Err()
}

func (r *Redis) RedisRead() (string, error) {
	value, err := r.client.Get(context.Background(), "example-key").Result()
	return value, err
}
