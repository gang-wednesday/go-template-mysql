package redisutil

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var windowTime int64 = 1
var numberOfRequests int64 = 100

func GetClient() *redis.Client {
	host := "localhost"
	if os.Getenv("ENVIRONMENT_NAME") == "docker" {
		host = "redis"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func SetCounter(ctx context.Context, rdb *redis.Client, key string, value int64) error {
	duration := time.Minute.Nanoseconds() * windowTime
	err := rdb.Set(ctx, key, value, time.Duration(duration)).Err()
	return err
}

func GetCounter(ctx context.Context, rdb *redis.Client, key string) (int64, error) {
	counter, err := rdb.Get(ctx, key).Int64()
	if err != nil {
		if err == redis.Nil {
			err = SetCounter(ctx, rdb, key, 1)
			if err != nil {
				return 0, err
			}
		}
	}
	if counter > numberOfRequests {
		return counter, errors.New("request limit breached!!!")
	}
	return counter, nil

}
