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

// func snakeBite(key string) string {
// 	// a file name config.env
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatalf("Error while reading config files: %s", err)
// 	}
// 	// type asssertion
// 	value, ok := viper.Get(key).(string)
// 	if !ok {
// 		log.Fatalf("Invalid type assertion")
// 	}
// 	return value
// }

// password := snakeBite("password")
// connection := redis.NewClient(&redis.Options{
// 	Addr:     "localhost:6379",
// 	DB:       db,
// 	Password: password,
// })
