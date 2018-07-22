package redis

import (
	"github.com/go-redis/redis"
	"log"
	"fmt"
)

type Redis struct {
	Addr string
	Rediscli *redis.Client
}

func (r *Redis) Get(key string) string {
	val, err := r.Rediscli.Get(key).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

func (r *Redis) Set(key, val string) {
	err := r.Rediscli.Set(key, val, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func NewRedis(addr string) *Redis {
	if "" == addr {
		log.Panic("redis addres is nil")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,//"localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return &Redis{addr, client}
}