package user_routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/controllers/user_controller"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userHandler := user_controller.UserHandler{}
	userRouter.POST("/create/", userHandler.CreateUserController)
	userRouter.GET("/email-verify/", userHandler.VerifyEmailController)
	userRouter.POST("/login", userHandler.LoginController)
	userRouter.GET("/", middlewares.JWTAuthMiddleware(), userHandler.GetCurrentUserController)
}
