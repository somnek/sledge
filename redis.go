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

func (r *Rdb) Keys(ctx context.Context, pattern string) ([]string, error) {
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *Rdb) Close() error {
	return r.client.Close()
}

func (r *Rdb) Type(ctx context.Context, key string) (string, error) {
	kind, err := r.client.Type(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return kind, err
}

func (r *Rdb) HKeys(ctx context.Context, key string) ([]string, error) {
	fields, err := r.client.HKeys(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (r *Rdb) Hget(ctx context.Context, key, field string) (string, error) {
	val, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *Rdb) ValFromHash(ctx context.Context, key string) (map[string]string, error) {
	m := make(map[string]string)

	fields, err := r.HKeys(ctx, key)
	if err != nil {
		return nil, err
	}

	for _, f := range fields {
		val, err := r.Hget(ctx, key, f)
		if err != nil {
			return nil, err
		}

		m[f] = val
	}

	return m, nil
}

func (r *Rdb) ExtractVal(ctx context.Context, key, kind string) (interface{}, error) {
	var val interface{}
	var err error

	switch kind {
	case "string":
		val, err = r.Get(ctx, key)
		if err != nil {
			return nil, err
		}

	case "hash":
		val, err = r.ValFromHash(ctx, key)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (r *Rdb) GetRecords(ctx context.Context, pattern string) ([]Record, error) {

	keys, err := r.Keys(ctx, pattern)
	if err != nil {
		return nil, err
	}

	records := make([]Record, len(keys))

	for i, key := range keys {
		kind, err := r.Type(ctx, key)
		if err != nil {
			return nil, err
		}

		val, err := r.ExtractVal(ctx, key, kind)
		if err != nil {
			return nil, err
		}

		r := Record{
			key:  key,
			val:  val,
			kind: kind,
		}

		records[i] = r
	}

	return records, nil
}
