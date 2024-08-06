package consumer

import (
	"Practice/common"
	"Practice/driver"
	"Practice/repository/repoimpl"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateNewCall(nameQueue string) {
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

	r := gin.Default()
	r.POST("/v1/items/", callRepo.Post(rabbit, nameQueue))
	r.Run()
}
