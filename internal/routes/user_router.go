package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/controllers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/middlewares"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userController := services.UserService{}
	userHandler := controllers.UserHandler{UserService: &userController}
	userRouter.GET("/", middlewares.JWTAuthMiddleware(), userHandler.GetCurrentUserController)
	userRouter.POST("/create/", userHandler.CreateUserController)
	userRouter.GET("/email-verify/", userHandler.VerifyEmailController)
	userRouter.POST("/login", userHandler.LoginController)
	userRouter.GET("/generate-otp/", userHandler.GenerateOTP)
	userRouter.POST("/verify-otp/", userHandler.VerifyOTP)
}
