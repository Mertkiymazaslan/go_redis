package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Redis Demo")

	ctx := context.Background()
	client, err := initializeRedis(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize the redis client: %s", err.Error())
		return
	}

	//basic set and get string example
	{
		err = client.Set(ctx, "name", "Mert", 0).Err()
		if err != nil {
			fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
			return
		}

		val, err := client.Get(ctx, "name").Result()
		if err != nil {
			fmt.Printf("Failed to get value from the redis instance: %s", err.Error())
			return
		}
		fmt.Println("name", val)
	}

	//store Person struct in redis and read it using uuid (more production-like)
	{
		type Person struct {
			ID         string
			Name       string `json:"name"`
			Age        int    `json:"age"`
			Occupation string `json:"occupation"`
		}

		mertID := uuid.NewString()
		jsonString, err := json.Marshal(Person{
			ID:         mertID,
			Name:       "Mert",
			Age:        22,
			Occupation: "Software Engineer",
		})
		if err != nil {
			fmt.Printf("failed to marshal: %s\n", err.Error())
			return
		}

		err = client.Set(ctx, mertID, jsonString, 0).Err()
		if err != nil {
			fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
			return
		}

		val, err := client.Get(ctx, mertID).Result()
		if err != nil {
			fmt.Printf("Failed to get value from the redis instance: %s", err.Error())
			return
		}
		var mert Person
		if err = json.Unmarshal([]byte(val), &mert); err != nil {
			fmt.Printf("failed to marshal: %s\n", err.Error())
			return
		}
		fmt.Printf("name: %s, age: %d, job: %s ", mert.Name, mert.Age, mert.Occupation)
	}
}

func initializeRedis(ctx context.Context) (*redis.Client, error) {
	//bad practice to hardcode the connection details but it is just a small demo script!
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(ping)

	return client, nil
}
