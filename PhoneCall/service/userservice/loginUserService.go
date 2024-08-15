package userservice

import (
	"PhoneCall/controller/helpers"
	"PhoneCall/handlers"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func (userService *UserService) VerifyPassword(userPassword string, foundUserPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	if err != nil {
		return false
	}
	return true
}

func (userService *UserService) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUserInfo := struct {
			Email    *string `json:"email" validate:"required"`
			Password *string `json:"password" validate:"required"`
		}{}
		if err := c.ShouldBind(&loginUserInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			handlers.LogErr(err.Error())
			return
		}
		foundUser, err := userService.UserRepo.VerifyValueField("email", *loginUserInfo.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email does not exist"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			handlers.LogErr(err.Error())
			return
		}

		isValidPassword := userService.VerifyPassword(*loginUserInfo.Password, foundUser.Password)
		if isValidPassword == false {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email or password is incorrect",
			})
			handlers.LogErr(err.Error())
			return
		}
		// Khi người dùng đăng nhập lại thì gen ra token mới vì token cũ có thể hết hạn
		token, refreshToken, _ := helpers.GenerateTokens(foundUser.Id, foundUser.Email, *foundUser.FirstName,
			*foundUser.LastName, foundUser.UserType)
		if err != nil {
			handlers.LogErr(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = userService.RedisService.Client.Set(context.Background(),
			"access_token_"+strconv.FormatInt(foundUser.Id, 10), token, time.Hour*24).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		err = userService.RedisService.Client.Set(context.Background(),
			"refresh_token_"+strconv.FormatInt(foundUser.Id, 10), refreshToken, time.Hour*200).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"access_token":  token,
			"refresh_token": refreshToken,
		})
	}
}
