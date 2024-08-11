package consumer

import (
	"PhoneCall/common"
	"PhoneCall/driver"
	models "PhoneCall/models"
	"PhoneCall/repository/repoimpl"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func popQueueResult(rabbit *repoimpl.RabbitMQ, callRepo *repoimpl.CallRepoImpl, nameQueue string) {
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

func updateDatabase(callRepo *repoimpl.CallRepoImpl, data models.Call) {
	r := gin.Default()
	r.PUT("/v1/items", callRepo.Update(data))
	r.Run()
}

func UpdateResult(nameQueue string) {
	MySQL := driver.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	callRepo := repoimpl.CallRepoImpl{MySQL: MySQL}

	channel, connect, err := driver.ConnectRabbit("QueueSolve", common.LinkRabbit)
	if err != nil {
		fmt.Println(err)
		return
	}
	rabbit := repoimpl.NewRabbitMQ(channel, connect)
	defer rabbit.Connect.Close()
	defer rabbit.Channel.Close()

	popQueueResult(rabbit, &callRepo, nameQueue)
}
