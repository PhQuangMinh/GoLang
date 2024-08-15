package consumer

import (
	"PhoneCall/common"
	models "PhoneCall/model"
	"PhoneCall/repository"
	"PhoneCall/service/connection"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func popQueueResult(rabbit *repository.RabbitMQ, callRepo *repository.CallRepoImpl, nameQueue string) {
	messages, err := rabbit.Pop(nameQueue)
	if err != nil {
		fmt.Println(err)
		return
	}
	cha := make(chan bool)

	go func() {
		for d := range messages {
			var cal models.Call
			er := json.Unmarshal(d.Body, &cal)
			if er != nil {
				fmt.Println(er)
				continue
			}
			cal.CallResult = "Success"
			updateDatabase(callRepo, cal)
		}
		cha <- true
	}()
	<-cha
	fmt.Println("Thanh cong")
}

func updateDatabase(callRepo *repository.CallRepoImpl, data models.Call) {
	r := gin.Default()
	//r.PUT("/v1/items", callRepo.UpdateCall(data))
	r.Run()
}

func UpdateResult(nameQueue string) {
	MySQL := connection.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	callRepo := repository.CallRepoImpl{MySQL: MySQL}

	channel, connect, err := connection.ConnectRabbit("QueueSolve", common.LinkRabbit)
	if err != nil {
		fmt.Println(err)
		return
	}
	rabbit := repository.NewRabbitMQ(channel, connect)
	defer rabbit.Connect.Close()
	defer rabbit.Channel.Close()

	popQueueResult(rabbit, &callRepo, nameQueue)
}
