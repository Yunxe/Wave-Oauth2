package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var RDB *redis.Client

func RedisConnection(ctx context.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       1,                // use default DB
	})
	RDB = rdb

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}
