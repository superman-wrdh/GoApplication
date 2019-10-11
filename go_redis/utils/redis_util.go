package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis"
)

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		//Addr:     "192.168.199.199:6379",
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	return client
	// Output: PONG <nil>
}

type RedisCache struct {
	client *redis.Client
}

func (c *RedisCache) Put(key string, obj interface{}) bool {
	var buf bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buf) // Will write to network.
	// Encode (send) some values.
	err := enc.Encode(obj)
	if err != nil {
		return false
	}
	e := c.client.Set(key, buf.Bytes(), 0).Err()
	if e != nil {
		return false
	}
	return true
}
func (c *RedisCache) Get(key string, dst interface{}) (interface{}, error) {
	data, err := c.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	err = dec.Decode(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (c *RedisCache) Flush() bool {
	return false
}
func (c *RedisCache) IsExits() bool {
	return false
}
func (c *RedisCache) Expire() bool {
	return false
}

type User struct {
	Name string
	Age  int
}

func main() {

	client := GetClient()
	redisCache := RedisCache{client: client}
	var u = &User{"superman", 22}
	key := "man"
	redisCache.Put(key, u)
	user, err := redisCache.Get(key, &User{Age: 1})
	if err != nil {
		fmt.Print(user)
	}
	fmt.Println("--------------")
	if o, ok := user.(*User); ok {
		fmt.Println(o.Age)
		fmt.Println(o.Name)
	}
	fmt.Println("--------------")
	o := user.(*User)
	fmt.Println(o)
	fmt.Println("--------------")
	fmt.Println(u)
}
