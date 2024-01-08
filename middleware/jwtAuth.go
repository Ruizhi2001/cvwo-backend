package middleware

import (
	"fmt"
	"net/http"

	"github.com/Ruizhi2001/cvwo-backend/helper"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		fmt.Println("JWT Authenticated")
		context.Next()
	}
}
