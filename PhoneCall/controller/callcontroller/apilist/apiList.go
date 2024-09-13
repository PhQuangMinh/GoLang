package apilist

import (
	"PhoneCall/common"
	"PhoneCall/repository"
	"PhoneCall/service/callservice"
	"PhoneCall/service/connection"
	"github.com/gin-gonic/gin"
)

func GetList() {
	r := gin.Default()
	MySQL := connection.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	callRepo := repository.CallRepoImpl{MySQL: MySQL}
	callService := callservice.NewGetCallService(&callRepo)
	r.GET("/v1/items", callService.GetCallsTime())
	r.Run()
}
