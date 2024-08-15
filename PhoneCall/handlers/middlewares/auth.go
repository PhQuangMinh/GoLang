package middlewares

import (
	"PhoneCall/controller/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"ErrorAuth": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}
		ok := helpers.ValidateToken(tokenString, ctx)
		if !ok {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
