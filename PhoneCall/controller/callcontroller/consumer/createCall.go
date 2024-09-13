package consumer

import (
	"PhoneCall/common"
	error2 "PhoneCall/handlers"
	"PhoneCall/repository"
	"PhoneCall/service/connection"
	"github.com/gin-gonic/gin"
)

func CreateNewCall(nameQueue string) {
	//MySQL := connection.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	//callRepo := repository.CallRepoImpl{MySQL: MySQL}

	channel, connect, err := connection.ConnectRabbit("QueueSolve", common.LinkRabbit)
	if err != nil {
		error2.LogErr(err.Error())
		return
	}
	rabbit := repository.NewRabbitMQ(channel, connect)
	defer rabbit.Connect.Close()
	defer rabbit.Channel.Close()

	r := gin.Default()
	//r.POST("/v1/items/", callRepo.Post(rabbit, nameQueue))
	r.Run()
}
