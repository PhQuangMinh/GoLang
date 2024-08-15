package userservice

import (
	"PhoneCall/controller/helpers"
	"PhoneCall/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (userService *UserService) GetUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		handlers.LogErr(err.Error())
		return
	}
	//Kiểm tra xem nếu là user đúng id chưa, nếu không thì ko đc truy cập
	err = helpers.MatchUserTypeId(c, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := userService.UserRepo.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (userService *UserService) GetUsers(c *gin.Context) {
	err := helpers.CheckUserType(c, "ADMIN")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	users, err := userService.UserRepo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
