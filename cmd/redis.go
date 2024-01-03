package cmd

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

// Create a new redis client.
func NewClient(url string) (*Rdb, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	c := &Rdb{redis.NewClient(opts)}
	c.Ping()
	return c, nil
}

// cmd: PING
func (r *Rdb) Ping() {
	_, err := r.client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
}

// cmd: EXISTS
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

// cmd: SET
func (r *Rdb) Set(ctx context.Context, key string, val interface{}) error {
	return r.client.Set(ctx, key, val, 0).Err()
}

// cmd: GET
func (r *Rdb) Get(ctx context.Context, key string) (string, error) {
	if !r.Exists(ctx, key) {
		return "", fmt.Errorf("key not found")
	}
	return r.client.Get(ctx, key).Result()
}

// cmd: DEL
func (r *Rdb) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// cmd: KEYS
func (r *Rdb) Keys(ctx context.Context, pattern string) ([]string, error) {
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// cmd: TYPE
func (r *Rdb) Type(ctx context.Context, key string) (string, error) {
	kind, err := r.client.Type(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return kind, err
}

// cmd: HKEYS
func (r *Rdb) HKeys(ctx context.Context, key string) ([]string, error) {
	fields, err := r.client.HKeys(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return fields, nil
}

// cmd: HGET
func (r *Rdb) HGet(ctx context.Context, key, field string) (string, error) {
	val, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// cmd: SMEMBERS
func (r *Rdb) SMembers(ctx context.Context, key string) ([]string, error) {
	val, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// cmd: LRANGE
func (r *Rdb) LRange(ctx context.Context, key string) ([]string, error) {
	val, err := r.client.LRange(ctx, key, 0, -1).Result() // assuming getting all element
	if err != nil {
		return []string{}, nil
	}
	return val, nil
}

// ValFromHash returns a map of field-value pairs from a hash type.
func (r *Rdb) ValFromHash(ctx context.Context, key string) (map[string]string, error) {
	m := make(map[string]string)

	fields, err := r.HKeys(ctx, key)
	if err != nil {
		return nil, err
	}

	for _, f := range fields {
		val, err := r.HGet(ctx, key, f)
		if err != nil {
			return nil, err
		}

		m[f] = val
	}

	return m, nil
}

// ExtractVal returns the value of a key based on its type.
// string -> string
// hash   -> map[string]string
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

	case "list":
		val, err = r.LRange(ctx, key)
		if err != nil {
			return nil, err
		}

	case "set":
		val, err = r.SMembers(ctx, key)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

// GetRecords returns a slice of Record structs.
// Get all records in database and convert them to Record structs.
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

func (r *Rdb) Close() error {
	return r.client.Close()
}
