package helpers

import (
	"PhoneCall/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateTokens(id int64, email string, firstName string, lastName string, userType string) (signedToken string, signedRefreshToken string, err error) {
	//Tạo một access token trong 24h -> không hoạt động 24h, tự out -> phải đăng nhập lại
	fmt.Println(id)
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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

// UpdateTokens Update lại các token mới
func UpdateTokens(id int64, signedToken string, signedRefreshToken string,
	UserRepo repository.UserRepo) error {
	UpdateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updates := map[string]interface{}{
		"token":         signedToken,
		"refresh_token": signedRefreshToken,
		"updated_at":    UpdateAt,
	}
	err := UserRepo.UpdateValueFields(id, updates)
	if err != nil {
		return err
	}
	return nil
}

func ValidateToken(tokenString string, c *gin.Context) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errorValidate": err.Error(),
		})
		return false
	}

	claims, check := token.Claims.(jwt.MapClaims)
	if !check {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return false
	}

	if int64(claims["exp"].(float64)) < time.Now().Local().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Expired token",
		})
		return false
	}
	return true
}

func ExtractIdUser(ctx *gin.Context) (int64, error) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		ctx.Abort()
		return -1, nil
	}
	tokenString = tokenString[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	fmt.Println(token)
	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	//fmt.Printf("%T", claims["Id"])
	fmt.Println(claims["Id"])
	if ok && token.Valid {
		id := int64(claims["Id"].(float64))
		return id, nil
	}
	return -1, nil
}
