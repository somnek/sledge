package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func snakeBite(key string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config files: %s", err)
	}
	// type asssertion
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func connect() *Store {
	password := snakeBite("password")
	connection := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       1,
		Password: password,
	})
	// connection := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6666",
	// 	DB:   0,
	// })

	rdb := Store{
		ctx: context.Background(),
		rdb: connection,
	}

	return &rdb
}
