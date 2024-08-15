package userservice

import (
	"PhoneCall/handlers"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

func (userService *UserService) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User not found",
			})
			return
		}
		//Khi logout thì xóa token đi
		err := userService.RedisService.Client.Del(context.Background(), "access_token_"+id).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			handlers.LogErr(err.Error())
			return
		}

		c.JSON(http.StatusOK, "logout successfully")
	}
}
