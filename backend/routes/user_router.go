package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go_user_role/backend/controllers"
	"github.com/go_user_role/backend/services"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userService := &services.UserService{}
	userController := controllers.NewUserController(*userService)
	userRouter.GET("/", userController.GetCurrentUserController)
	userRouter.POST("/create/", userController.CreateUserController)
	userRouter.GET("/email-verify/", userController.VerifyEmailController)
	userRouter.POST("/login", userController.LoginController)
	userRouter.GET("/generate-otp/", userController.GenerateOTP)
	userRouter.POST("/verify-otp/", userController.VerifyOTP)
}
