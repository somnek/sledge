package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Store struct {
	ctx context.Context
	rdb *redis.Client
}

func (s *Store) get(key string) string {
	rdb, ctx := s.rdb, s.ctx
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "Does not exist!"
	} else {
		return val
	}
}

func (s *Store) add(key string, val string) {
	err := s.rdb.Set(s.ctx, key, val, 0).Err()
	if err == redis.Nil {
		fmt.Println(err)
	}
}

func (s *Store) del(key string) {
	if err := s.rdb.Del(s.ctx, key).Err(); err != nil {
		panic(err)
	}
}

func (s *Store) exists(key string) bool {
	return s.rdb.Exists(s.ctx, key).Val() == 1
}

func (s *Store) getKeys() (ret []string) {
	iter := s.rdb.Scan(s.ctx, 0, "", 0).Iterator()
	for iter.Next(s.ctx) {
		key := iter.Val()
		ret = append(ret, key)
	}
	return
}
