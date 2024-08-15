package main

import (
	"PhoneCall/handlers/middlewares"
	"PhoneCall/repository"
	"PhoneCall/service/callservice"
	"PhoneCall/service/connection"
	"PhoneCall/service/redisservice"
	"PhoneCall/service/userservice"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	nameDatabase := os.Getenv("NAME_DATABASE")
	port := os.Getenv("PORT")
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	MySQL := connection.ConnectDB(user, password, port, nameDatabase)
	userRepo := repository.NewUserRepoImpl(MySQL)
	client := connection.ConnectRedis()
	redisService := redisservice.NewRedisService(client)
	userService := userservice.NewUserService(userRepo, *redisService)

	callRepo := repository.NewCallRepoImpl(MySQL)
	callService := callservice.NewCallService(callRepo)

	routerUser := router.Group("/v2")
	{
		calls := routerUser.Group("/calls", middlewares.AuthMiddleware())
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
			users.POST("/logout", userService.Logout())
			users.GET("/:id", middlewares.AuthMiddleware(), userService.GetUserById)
			users.GET("", middlewares.AuthMiddleware(), userService.GetUsers)
			users.PUT("/:id", middlewares.AuthMiddleware(), userService.UpdateUserInfoById)
			users.PUT("/password/:id", middlewares.AuthMiddleware(), userService.UpdateUserPasswordInfoById)
			users.DELETE("/:id", middlewares.AuthMiddleware(), userService.DeleteUserById)
		}
	}
	router.Run()
}
