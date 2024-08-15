package userservice

import (
	"PhoneCall/controller/helpers"
	"PhoneCall/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var validate = validator.New()

func (userService *UserService) UpdateUserInfoById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user *model.UserUpdate
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Kiểm tra xem có đúng là user đó không
	checkUser := helpers.MatchUserTypeId(c, id)
	//Kiểm tra xem là admin hay không
	checkAdmin := helpers.CheckUserType(c, "ADMIN")
	if checkUser != nil && checkAdmin != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized to access this resource",
		})
		return
	}

	//Kiểm tra xem nó đúng cái field yêu cầu hay không
	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validateErr.Error(),
		})
		return
	}

	user, err = userService.UserRepo.UpdateUser(user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (userService *UserService) UpdateUserPasswordInfoById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := struct {
		Password string `json:"password" gorm:"column:password" validate:"required,min=8,max=1000"`
	}{}

	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Kiểm tra xem nó đúng cái field yêu cầu hay không
	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validateErr.Error(),
		})
		return
	}

	//Kiểm tra xem có đúng là user đó không
	checkUser := helpers.MatchUserTypeId(c, id)
	//Kiểm tra xem là admin hay không
	checkAdmin := helpers.CheckUserType(c, "ADMIN")
	if checkUser != nil && checkAdmin != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized to access this resource",
		})
		return
	}

	mp := map[string]interface{}{
		"password": userService.HashPassword(user.Password),
	}
	err = userService.UserRepo.UpdateValueFields(id, mp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "updated password user successfully")
}
