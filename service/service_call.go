package service

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/gta4roy/goRedisDBClient/model"
)

func CreateNewPerson(redisClient *redis.Client, name string, age int, salary float32, phone string) (bool, string) {
	var personID string

	personID = uuid.New().String()
	personObj := model.NewPerson()
	personObj.Age = age
	personObj.Name = name
	personObj.Phone = phone
	personObj.Salary = salary

	pJson, err := json.Marshal(personObj)

	if err != nil {
		fmt.Println("Error in converting to json object")
		return false, ""
	}

	status := redisClient.Set(personID, pJson, 0)

	if status.Err() != nil {
		fmt.Println("Exception in Saving ", status.Err())
		return false, ""
	}

	return true, personID
}

func GetPerson(redisClient *redis.Client, key string) (bool, *model.Person) {

	status := redisClient.Get(key)
	if status.Err() != nil {
		fmt.Println("Error in retrieving ", status.Err())
		return false, nil
	}
	value, _ := status.Result()
	fmt.Println("Json Value :", value)

	var personObj model.Person

	_err := json.Unmarshal([]byte(value), &personObj)

	if _err != nil {
		fmt.Println("Error in Parsing the json ", _err)
		return false, nil
	}
	return true, &personObj
}

func DeleteRecord(redisClient *redis.Client, key string) bool {

	status := redisClient.Del(key)

	if status.Err() != nil {
		fmt.Println("Error in deleting the record ", status.Err())
		return false
	}

	return true
}

func GetSetRecord(redisClient *redis.Client, key string, data *model.Person) bool {

	pJson, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error in converting to json object")
		return false
	}
	status := redisClient.GetSet(key, pJson)

	if status.Err() != nil {
		fmt.Println("Error in Getting the key ")
		return false
	}

	return true
}

func GetAll(redisClient *redis.Client) []model.Person {

	allKeys := redisClient.Keys("*")

	keySet, err := allKeys.Result()

	if err != nil {
		fmt.Println("Error While fetching Keys ", err.Error())
	}

	var personsResultSet []model.Person

	for _, key := range keySet {

		fmt.Println("Key ", key)

		valueSet := redisClient.Get(key)

		if valueSet.Err() != nil {
			fmt.Println("Error While Fetching Result  ", err.Error())
			break
		}

		personObject, _ := valueSet.Result()

		var personObj model.Person

		_err := json.Unmarshal([]byte(personObject), &personObj)

		if _err != nil {
			fmt.Println("Error in Parsing the json ", _err)
		}

		personsResultSet = append(personsResultSet, personObj)

	}
	return personsResultSet
}
