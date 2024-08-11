package service

import (
	"PhoneCall/handlers/errorpk"
	"PhoneCall/helpers"
	models "PhoneCall/models"
	"PhoneCall/repository"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validate "github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func NewUserService(UserRepo repository.UserRepo) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (userService *UserService) HashPassword(password string) string {
	by, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {

		log.Panic(err)
	}
	return string(by)
}

func (userService *UserService) VerifyPassword(userPassword string, foundUserPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	if err != nil {
		return false
	}
	return true
}

func (userService *UserService) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			errorpk.LogErr(err.Error())
			return
		}

		//Kiểm tra xem thông tin đki có gồm những thông tin bắt buộc hay không
		validationErr := validate.New().Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			errorpk.LogErr(validationErr.Error())
			return
		}

		//Kiểm tra email xem có tồn tại hay không
		_, err := userService.UserRepo.VerifyValueField("email", *user.Email)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				errorpk.LogErr(err.Error())
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email already exists"})
			errorpk.LogErr(err.Error())
			return
		}

		password := userService.HashPassword(*user.Password)
		user.Password = &password

		_, err = userService.UserRepo.VerifyValueField("phone_number", *user.PhoneNumber)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				errorpk.LogErr(err.Error())
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"errorpk": "Phone number already exists"})
			errorpk.LogErr(err.Error())
			return
		}

		user.CreateAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdateAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		token, refreshToken, _ := helpers.GenerateTokens(user.Id, *user.Email, *user.FirstName, *user.LastName, *user.UserType)
		user.Token = &token
		user.RefreshToken = &refreshToken
		userPost, err := userService.UserRepo.PostUser(&user)

		if err != nil {
			fmt.Println("User was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			errorpk.LogErr(err.Error())
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"user": userPost,
		})
	}
}

func (userService *UserService) LoginHTML() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		foundUser, err := userService.UserRepo.VerifyValueField("email", email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Email does not exist"})
			} else {
				c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": err.Error()})
			}
			errorpk.LogErr(err.Error())
			return
		}

		isValidPassword := userService.VerifyPassword(password, *foundUser.Password)
		if isValidPassword == false {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"error": "Email or password is incorrect",
			})
			errorpk.LogErr(err.Error())
			return
		}
		// Khi người dùng đăng nhập lại thì gen ra token mới vì token cũ có thể hết hạn
		token, refreshToken, _ := helpers.GenerateTokens(foundUser.Id, *foundUser.Email, *foundUser.FirstName,
			*foundUser.LastName, *foundUser.UserType)
		fmt.Println(foundUser.Id, token, refreshToken, userService.UserRepo)
		err = helpers.UpdateTokens(foundUser.Id, token, refreshToken, userService.UserRepo)
		if err != nil {
			errorpk.LogErr(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, "You have successfully logged in :)")
		c.HTML(http.StatusOK, "login.html", gin.H{"error": "Check username and password"})
		c.JSON(http.StatusOK, foundUser)
	}
}

func (userService *UserService) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUserInfo := struct {
			Email    *string `json:"email" validate:"required"`
			Password *string `json:"password" validate:"required"`
		}{}
		if err := c.ShouldBind(&loginUserInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			errorpk.LogErr(err.Error())
			return
		}
		foundUser, err := userService.UserRepo.VerifyValueField("email", *loginUserInfo.Email)
		fmt.Println(foundUser)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Email does not exist"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			errorpk.LogErr(err.Error())
			return
		}

		isValidPassword := userService.VerifyPassword(*loginUserInfo.Password, *foundUser.Password)
		if isValidPassword == false {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Email or password is incorrect",
			})
			errorpk.LogErr(err.Error())
			return
		}
		// Khi người dùng đăng nhập lại thì gen ra token mới vì token cũ có thể hết hạn
		token, refreshToken, _ := helpers.GenerateTokens(foundUser.Id, *foundUser.Email, *foundUser.FirstName,
			*foundUser.LastName, *foundUser.UserType)
		fmt.Println(foundUser.Id, token, refreshToken, userService.UserRepo)
		err = helpers.UpdateTokens(foundUser.Id, token, refreshToken, userService.UserRepo)
		if err != nil {
			errorpk.LogErr(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

func (userService *UserService) Logout() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := helpers.ExtractIdUser(context)
		fmt.Println(id, err)
		if id == -1 || err != nil {
			if id == -1 {
				context.JSON(http.StatusBadRequest, gin.H{
					"error": "User not found",
				})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
		err = helpers.UpdateTokens(id+1, "", "", userService.UserRepo)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			errorpk.LogErr(err.Error())
			return
		}
		context.JSON(http.StatusOK, "logout successfully")
	}
}

func (userService *UserService) GetUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error1": err.Error(),
		})
		errorpk.LogErr(err.Error())
		return
	}
	user, err := userService.UserRepo.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error2": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (userService *UserService) GetUsers(c *gin.Context) {
	users, err := userService.UserRepo.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (userService *UserService) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user *models.User
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = userService.UserRepo.UpdateUser(user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (userService *UserService) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = userService.UserRepo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "Deleted successful")
}
