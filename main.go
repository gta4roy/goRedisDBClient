package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	redisconnection "github.com/gta4roy/goRedisDBClient/redis_connection"
)

func main() {

	var redisClient redis.Client
	fmt.Println("Redis DB Client Main Module ")
	redisconnection.ConnectToRedisDB(&redisClient)

}
