package cache

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"time"
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_host"),
		Password: beego.AppConfig.String("redis_auth"), // no password set
		DB:       8,                                    // use default DB
	})
	_ = client.Set("key", "value", time.Second*1000).Err()
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_host"),
		Password: beego.AppConfig.String("redis_auth"), // no password set
		DB:       8,                                    // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
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
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
