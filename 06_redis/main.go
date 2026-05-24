package main

import (
	"context"
	"fmt"

	redis1 "github.com/go-redis/redis"
	redis2 "github.com/redis/go-redis/v9"
)

// redis的golang客户端有2种
// github.com/go-redis/redis, 参数不带上下文
// github.com/redis/go-redis/v9, 参数必须带上下文

func demo1() {
	// 参数不带上下文的用法
	rdb := redis1.NewClient(&redis1.Options{
		Addr:     "192.168.31.136:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("could not connect to Redis")
		return
	}
	fmt.Println(pong)

	val, _ := rdb.Get("name").Result()
	fmt.Println(val)

	err = rdb.Set("new_key", "new_val", 0).Err()
	if err != nil {
		fmt.Println("can not set key")
	}
}

func demo2() {
	var ctx = context.Background()
	rdb := redis2.NewClient(&redis2.Options{
		Addr:     "192.168.31.136:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("could not connect to Redis")
	}
	fmt.Println(pong)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis2.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func main() {
	demo2()
}
