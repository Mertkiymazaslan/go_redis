package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Redis Demo")

	//bad practice to hardcode the connection details but it is just a small demo script!
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

}
