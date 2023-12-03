package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/controllers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userRouter.GET("/", middlewares.JWTAuthMiddleware(), controllers.GetCurrentUserController)
	userRouter.POST("/create/", controllers.CreateUserController)
	userRouter.GET("/email-verify/", controllers.VerifyEmailController)
	userRouter.POST("/login", controllers.LoginController)
	userRouter.GET("/generate-otp/", controllers.GenerateOTP)
	userRouter.POST("/verify-otp/", controllers.VerifyOTP)
}
