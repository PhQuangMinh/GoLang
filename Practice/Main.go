package main

import (
	"Practice/controller/database"
	"Practice/controller/rabbitmq"
	"Practice/model"
	"time"
)

func main() {
	data := model.Call{
		PhoneNumber:       "12345678",
		Metadata:          "{\"Age\": 10, \"name\": \"Minh\"}",
		CallResult:        "INIT",
		CallTime:          time.Now(),
		ReceiveResultTime: time.Now(),
		CallAnsweredTime:  time.Now(),
		CallEndedTime:     time.Now(),
	}
	go database.PostData(data)
	go rabbitmq.Push(data, "QueueSolve")
	time.Sleep(2 * time.Second)
	rabbitmq.GetQueue("QueueSolve", 3*time.Second)
	time.Sleep(10 * time.Second)
	rabbitmq.GetQueueResult("QueueResult")
	time.Sleep(2 * time.Second)
}
