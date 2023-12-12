package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/controllers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/middlewares"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userService := &services.UserService{}
	userController := controllers.NewUserController(*userService)
	userRouter.GET("/", middlewares.JWTAuthMiddleware(), userController.GetCurrentUserController)
	userRouter.POST("/create/", userController.CreateUserController)
	userRouter.GET("/email-verify/", userController.VerifyEmailController)
	userRouter.POST("/login", userController.LoginController)
	userRouter.GET("/generate-otp/", userController.GenerateOTP)
	userRouter.POST("/verify-otp/", userController.VerifyOTP)
}
