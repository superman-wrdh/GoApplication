package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func InitClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.199.199:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
	// Output: PONG <nil>
}

func TestClient() {
	client := InitClient()
	//key string, value interface{}
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	TestClient()
}
