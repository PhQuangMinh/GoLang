package main

import (
	"PhoneCall/common"
	"PhoneCall/driver"
	"PhoneCall/handlers/middlewares"
	"PhoneCall/repository/repoimpl"
	"PhoneCall/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")
	router.Use(middlewares.CORSMiddleware())
	MySQL := driver.ConnectDB(common.User, common.Password, common.Port, common.NameDB)
	userRepoImpl := repoimpl.NewUserRepoImpl(MySQL)
	userService := service.NewUserService(userRepoImpl)
	routerUser := router.Group("/v2")
	{
		users := routerUser.Group("/users")
		{
			users.POST("/register", userService.Signup())
			users.POST("/login", userService.Login())
			users.POST("/logout", userService.Logout())
			users.GET("/:id", middlewares.AuthMiddleware(), userService.GetUserById)
			users.GET("", middlewares.AuthMiddleware(), userService.GetUsers)
			users.PUT("/:id", middlewares.AuthMiddleware(), userService.UpdateUser)
			users.DELETE("/:id", middlewares.AuthMiddleware(), userService.DeleteUser)
		}
	}
	router.Run()
}


