package redisconnection

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectToRedisDB() (bool, *redis.Client) {

	isConnected := false
	fmt.Println("Connecting to Local redis server ")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("Unable to connect to Redis ", err)
	} else {
		isConnected = true
		fmt.Println("Connected to redis ", pong)
	}
	return isConnected, redisClient
}
