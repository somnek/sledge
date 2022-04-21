package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func connect() *Store {
	connection := redis.NewClient(&redis.Options{
		Addr: "localhost:6666",
		DB:   0,
	})

	rdb := Store{
		ctx: context.Background(),
		rdb: connection,
	}

	return &rdb
}
