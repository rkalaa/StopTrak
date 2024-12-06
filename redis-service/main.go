package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis Test")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // remove hardcode
		Password: "",               //remove hardcode pass
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

}
