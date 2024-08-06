package apilist

import (
	"Practice/common"
	"Practice/driver"
	"Practice/repository/repoimpl"
	"Practice/service"
	"github.com/gin-gonic/gin"
)

func GetItem() {
	MySQL := driver.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	callRepo := repoimpl.CallRepoImpl{MySQL: MySQL}
	callService := service.NewGetCallService(callRepo)

	r := gin.Default()
	r.GET("/v1/items", callService.GetValueField())
	r.Run()
}
