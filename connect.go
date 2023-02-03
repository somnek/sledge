package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

const (
	TIMEOUT = 10000 // 10 second
)

func connect() *redis.Client {
	opts := &redis.Options{
		Addr:     "localhost:6666",
		Password: "",
		DB:       0,
	}
	return redis.NewClient(opts)
}

func Ping(ctx context.Context) string {
	rdb := connect()
	pong, err := rdb.WithTimeout(TIMEOUT).Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return pong
}
