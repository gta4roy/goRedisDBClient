package main

import (
	"fmt"

	"github.com/gta4roy/goRedisDBClient/model"
	redisconnection "github.com/gta4roy/goRedisDBClient/redis_connection"
	"github.com/gta4roy/goRedisDBClient/service"
)

func main() {

	fmt.Println("Redis DB Client Main Module ")
	isConnected, redisClient := redisconnection.ConnectToRedisDB()

	if isConnected {
		isSaved, id := service.CreateNewPerson(redisClient, "Ahijit", 34, 12000.0, "78297212297")

		if isSaved {
			fmt.Println(id)
			_, person := service.GetPerson(redisClient, id)

			fmt.Println("Original Person Details ....")
			fmt.Println("name ", person.Name)

			newPerson := model.NewPersonWithArgs("Rahul", 23, "898287286", 2300.8)

			service.GetSetRecord(redisClient, id, newPerson)

			_, person = service.GetPerson(redisClient, id)

			fmt.Println("Modified Person Details ....")
			fmt.Println("name ", person.Name)

			personArray := service.GetAll(redisClient)

			for _, personObj := range personArray {
				fmt.Println("name ", personObj.Name, " age ", personObj.Age, " Salary ", personObj.Salary)
			}

			isDeleted := service.DeleteRecord(redisClient, id)

			if isDeleted {
				fmt.Println("Record is deleted ")
			}
		}

	}

}
