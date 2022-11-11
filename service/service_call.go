package service

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gta4roy/goRedisDBClient/model"
	redisconnection "github.com/gta4roy/goRedisDBClient/redis_connection"
)

func CreateNewPerson(name string, age int, salary float32, phone string) string {
	var personID string

	personObj := model.NewPerson()
	personObj.Age = age
	personObj.Name = name
	personObj.Phone = phone
	personObj.Salary = salary

	var redisClient redis.Client
	if redisconnection.ConnectToRedisDB(&redisClient) {
		fmt.Println("Redis is connected to Data Saving")

		redisClient.Set("key", *personObj)
	}
	return personID
}
