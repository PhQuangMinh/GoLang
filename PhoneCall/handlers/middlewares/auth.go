package middlewares

import (
	"PhoneCall/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var jwtKey = []byte("123456")

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"ErrorAuth": http.StatusUnauthorized,
			})
			context.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		ok := helpers.ValidateToken(tokenString, context)
		if !ok {
			context.Abort()
			return
		}
		context.Next()
	}
}
