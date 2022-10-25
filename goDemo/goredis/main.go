package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	RedisConn()
}
func RedisConn() {
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.4:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic("连接失败！")
	}
	fmt.Println("连接成功", pong)

	err = client.Set("name", "zhang", 0).Err()
	if err != nil {
		panic(err)
	}
	var val string
	val, err = client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name的值为:", val)

}
