package middlewares

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		// Check if the token matches the predefined token
		if token == os.Getenv("TOKEN") {
			ctx.Next() // Continue with the request
			return
		}
		// Validate Token
		err := services.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "User not authorized!")
			ctx.Abort()
			return
		}
		// Before request
		ctx.Next()
		//	After request logic
	}
}
