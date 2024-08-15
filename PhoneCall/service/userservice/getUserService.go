package userservice

import (
	"PhoneCall/controller/helpers"
	"PhoneCall/handlers"
	"PhoneCall/model"
	"fmt"
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
	//Kieemrr tra xem nếu là admin mới truy cập tất cả user
	err := helpers.CheckUserType(c, "ADMIN")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var paging model.Paging
	if err := c.ShouldBind(&paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	paging.Process()

	numberOfUsers, err := userService.UserRepo.GetNumberOfUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	users, err := userService.UserRepo.GetUsers(&paging)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(numberOfUsers)
	c.JSON(http.StatusOK, gin.H{
		"users":      users,
		"pagination": numberOfUsers,
	})
}
