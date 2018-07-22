package config

import "github.com/telo_demo/redis"

type Output struct {
	Name string `name`
	Type string `type`
}

type Outputs struct {
	OutputsArray []Output `outputs`
}

var Client = redis.NewRedis("redis:6379")