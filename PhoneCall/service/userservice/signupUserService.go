package userservice

import (
	"PhoneCall/handlers"
	"PhoneCall/model"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func (userService *UserService) HashPassword(password string) string {
	by, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(by)
}

func (userService *UserService) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userSignUp model.User
		if err := c.ShouldBind(&userSignUp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			handlers.LogErr(err.Error())
			return
		}

		//Kiểm tra xem thông tin đki có gồm những thông tin bắt buộc hay không
		validationErr := validate.Struct(userSignUp)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			handlers.LogErr(validationErr.Error())
			return
		}

		//Kiểm tra email xem có tồn tại hay không
		_, err := userService.UserRepo.VerifyValueField("email", userSignUp.Email)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				handlers.LogErr(err.Error())
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email already exists"})
			handlers.LogErr(err.Error())
			return
		}
		//Tạm thời kiểm tra 2 loại ADMIN, USER sau này có thể thêm các loại khác tránh bị đăng kí sai kiểu user
		if userSignUp.UserType != "ADMIN" && userSignUp.UserType != "USER" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UserType is invalid"})
			return
		}

		password := userService.HashPassword(userSignUp.Password)
		userSignUp.Password = password

		_, err = userService.UserRepo.VerifyValueField("phone_number", userSignUp.PhoneNumber)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				handlers.LogErr(err.Error())
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already exists"})
			handlers.LogErr(err.Error())
			return
		}

		_, err = userService.UserRepo.CreateNewUser(&userSignUp)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			handlers.LogErr(err.Error())
			return
		}

		c.JSON(http.StatusCreated, "Created Successfully")
	}
}
