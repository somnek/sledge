package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func Ping(db int) string {
	rdb := connect(db)
	err := rdb.rdb.Ping(context.Background()).Err()
	if err != nil {
		return "PENG"
	}
	return "PONG"
}

func connect(db int) *Store {
	connection := redis.NewClient(&redis.Options{
		Addr: "localhost:6666",
		DB:   db,
	},
	)
	rdb := Store{
		ctx: context.Background(),
		rdb: connection, // bad naming
	}
	return &rdb
}
