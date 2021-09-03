package main

import (
	"context"
	"fmt"
	"os"

	"dronetest/str"

	"github.com/go-redis/redis/v8"
)

func main() {
	var key string

	if len(os.Args) > 1 {
		key = os.Args[1]
	} else {
		key = "test"
	}

	options, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(options)

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("redis not accessibe: %v\n", err)
		return
	}

	value, err := rdb.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		fmt.Printf("failed to read redis value for key %q: %v\n", key, err)
		return
	} else if err == redis.Nil {
		fmt.Printf("Key %q has no value.\n", key)
	} else {
		fmt.Printf("Value for key %q is %q.\n", key, value)
		fmt.Printf("Reversed value for key %q is %q.\n", key, str.Reverse(value))
	}
}