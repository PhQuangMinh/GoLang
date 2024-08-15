package apilist

import (
	"PhoneCall/common"
	"PhoneCall/repository"
	"PhoneCall/service/callservice"
	"PhoneCall/service/connection"
	"github.com/gin-gonic/gin"
)

func GetItem() {
	MySQL := connection.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	callRepo := repository.CallRepoImpl{MySQL: MySQL}
	callService := callservice.NewGetCallService(&callRepo)

	r := gin.Default()
	r.GET("/v1/items", callService.GetValueField())
	r.Run()
}
