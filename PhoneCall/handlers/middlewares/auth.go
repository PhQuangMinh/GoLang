package middlewares

import (
	"PhoneCall/controller/helpers"
	"PhoneCall/service/redisservice"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(redisService *redisservice.RedisService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"ErrorAuth": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]
		ok := helpers.ValidateToken(tokenString, ctx, redisService)
		if !ok {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
