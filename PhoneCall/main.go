package main

import (
	"PhoneCall/common"
	"PhoneCall/handlers/middlewares"
	"PhoneCall/repository"
	"PhoneCall/service/connection"
	"PhoneCall/service/redisservice"
	"PhoneCall/service/userservice"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	nameDatabase := os.Getenv("NAME_DATABASE")
	port := os.Getenv("PORT")
	fmt.Println(user, password, nameDatabase, port)
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	MySQL := connection.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	userRepo := repository.NewUserRepoImpl(MySQL)
	client := connection.ConnectRedis()
	redisService := redisservice.NewRedisService(client)
	userService := userservice.NewUserService(userRepo, *redisService)
	routerUser := router.Group("/v2")
	{
		users := routerUser.Group("/users")
		{
			users.POST("/register", userService.Signup())
			users.POST("/login", userService.Login())
			users.POST("/logout", userService.Logout())
			users.GET("/:id", middlewares.AuthMiddleware(), userService.GetUserById)
			users.GET("", middlewares.AuthMiddleware(), userService.GetUsers)
			users.PUT("/:id", middlewares.AuthMiddleware(), userService.UpdateUserById)
			users.DELETE("/:id", middlewares.AuthMiddleware(), userService.DeleteUserById)
		}
	}
	router.Run()
}
