package main

import (
	"context"
	"fmt"
	"log"
)

/*
get
add
del
exists
keys
*/

func get(ctx context.Context, key string) string {
	rdb := connect()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func add(ctx context.Context, key string, val string) {
	rdb := connect()
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func del(ctx context.Context, key string) {
	rdb := connect()
	if err := rdb.Del(ctx, key).Err(); err != nil {
		log.Fatal(err)
	}
}

func keys(ctx context.Context) []string {
	rdb := connect()

	var cur uint64
	var keys []string
	var max int64 = 5

	for {
		var err error
		keys, cur, err = rdb.Scan(ctx, cur, "*", max).Result()
		if err != nil {
			log.Fatal(err)
			break
		}

		if cur == 0 {
			break
		}
	}
	return keys
}

func exists(ctx context.Context, key string) bool {
	rdb := connect()
	return rdb.Exists(ctx, key).Val() == 1
}

func show(ctx context.Context) {
	keys := keys(ctx)
	for _, k := range keys {
		v := get(ctx, k)
		fmt.Println(k, v)
	}
}
