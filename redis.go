package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Rdb struct {
	client *redis.Client
}

func NewClient(addr string, db int) (*Rdb, error) {
	opts := &redis.Options{
		Addr:     addr,
		Password: "",
		DB:       db,
	}
	c := &Rdb{redis.NewClient(opts)}
	c.Ping()
	return c, nil
}

func (r *Rdb) Ping() {
	_, err := r.client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}
}

func (r *Rdb) Exists(ctx context.Context, key string) bool {
	exist, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}

	if exist == 1 {
		return true
	}
	return false
}

func (r *Rdb) Set(ctx context.Context, key string, val interface{}) error {
	return r.client.Set(ctx, key, val, 0).Err()
}

func (r *Rdb) Get(ctx context.Context, key string) (string, error) {
	if !r.Exists(ctx, key) {
		return "", fmt.Errorf("key not found")
	}
	return r.client.Get(ctx, key).Result()
}

func (r *Rdb) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Rdb) Close() error {
	return r.client.Close()
}
