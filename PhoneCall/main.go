package main

import (
	"PhoneCall/common"
	"PhoneCall/handlers"
	"PhoneCall/handlers/middlewares"
	"PhoneCall/repository"
	"PhoneCall/service/callservice"
	"PhoneCall/service/connection"
	"PhoneCall/service/redisservice"
	"PhoneCall/service/userservice"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	handlers.InitLogging()
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	MySQL := connection.ConnectDB(common.USER, common.PASSWORD, common.PORT, common.NAME_DATABASE)
	userRepo := repository.NewUserRepoImpl(MySQL)
	client := connection.ConnectRedis()
	redis := redisservice.NewRedisService(client)
	userService := userservice.NewUserService(userRepo, *redis)

	callRepo := repository.NewCallRepoImpl(MySQL)
	callService := callservice.NewCallService(callRepo)

	routerUser := router.Group("/v2")
	{
		test := routerUser.Group("/test")
		{
			test.POST("", HELLO())
		}
		calls := routerUser.Group("/calls", middlewares.AuthMiddleware(redis))
		{
			calls.GET("", callService.GetCallsTime())
			calls.POST("", callService.CreateNewCall())
			calls.PUT("/:id", callService.UpdateCall())
			calls.DELETE("/:id", callService.DeleteCall())
		}
		users := routerUser.Group("/users")
		{
			users.POST("/register", userService.Signup())
			users.POST("/login", userService.Login())
			users.POST("/logout", middlewares.AuthMiddleware(redis), userService.Logout())
			users.GET("/:id", middlewares.AuthMiddleware(redis), userService.GetUserById)
			users.GET("", middlewares.AuthMiddleware(redis), userService.GetUsers)
			users.PUT("/:id", middlewares.AuthMiddleware(redis), userService.UpdateUserInfoById)
			users.PUT("/password/:id", middlewares.AuthMiddleware(redis), userService.UpdateUserPasswordInfoById)
			users.DELETE("/:id", middlewares.AuthMiddleware(redis), userService.DeleteUserById)
		}
	}
	router.Run()
}

func HELLO() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "HELLO")
	}
}
