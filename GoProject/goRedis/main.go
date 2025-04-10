package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func main() {
	ctx := context.Background()
	rdb.Set(ctx, "goredistestkey", "goredistestvalue", 10*time.Second)
	err := rdb.Set(ctx, "goredistestkey", "goredistestvalue", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "goredistestkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(("goredistestkey"), val)

	result, err := rdb.Do(ctx, "get", "goredistestkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("do gordistestkey", result.(string))
}
