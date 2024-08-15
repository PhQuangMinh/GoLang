package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type SignedDetails struct {
	Id        int64
	Email     string
	FirstName string
	LastName  string
	UserType  string
	jwt.StandardClaims
}

var secretKey = os.Getenv("SECRET_KEY")

func GenerateTokens(id int64, email string, firstName string, lastName string, userType string) (signedToken string, signedRefreshToken string, err error) {
	//Tạo một access token trong 24h -> không hoạt động 24h, tự out -> phải đăng nhập lại
	claims := &SignedDetails{
		Id:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	//Tạo refresh token, khi còn hạn -> cung cấp access token mới giúp duy trì đăng nhập
	//, hết hạn -> đăng nhập lại
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(200)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func ValidateToken(tokenString string, c *gin.Context) bool {
	//Giai ma token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errorValidate": err.Error(),
		})
		return false
	}

	claims, check := token.Claims.(*SignedDetails)
	if !check {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return false
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Expired token",
		})
		return false
	}
	c.Set("email", claims.Email)
	c.Set("first_name", claims.FirstName)
	c.Set("id", strconv.FormatInt(claims.Id, 10))
	c.Set("last_name", claims.LastName)
	c.Set("user_type", claims.UserType)
	return true
}
