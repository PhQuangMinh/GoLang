package userservice

import (
	"PhoneCall/controller/helpers"
	"PhoneCall/handlers"
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
		handlers.LogDebug("id")
		return
	}

	user := struct {
		OldPassword string `json:"old_password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required"`
	}{}

	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		handlers.LogDebug("SHOULD BIND")
		return
	}

	//Kiểm tra xem nó đúng cái field yêu cầu hay không
	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validateErr.Error(),
		})
		handlers.LogDebug("Validate Err")
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
		handlers.LogDebug("unauthorized to access this resource")
		return
	}

	foundUser, err := userService.UserRepo.GetFullInfoUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		handlers.LogDebug("Find User Err")
		return
	}

	isValidPassword := userService.VerifyPassword(user.OldPassword, foundUser.Password)
	if isValidPassword == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password is incorrect",
		})
		handlers.LogErr("password is incorrect")
		return
	}

	mp := map[string]interface{}{
		"password": userService.HashPassword(user.NewPassword),
	}
	err = userService.UserRepo.UpdateValueFields(id, mp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		handlers.LogDebug("Update User Err")
		return
	}
	c.JSON(http.StatusOK, "updated password user successfully")
	handlers.LogInfo("updated password user successfully")
}
