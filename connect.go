package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func connect() *redis.Client {
	opts := &redis.Options{
		Addr:     "localhost:6666",
		Password: "",
		DB:       0,
	}
	return redis.NewClient(opts)
}

func Ping(ctx context.Context) {
	rdb := connect()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
}
