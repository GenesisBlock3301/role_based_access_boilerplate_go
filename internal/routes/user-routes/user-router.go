package user_routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/controllers/user_controller"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userRouter.POST("/create/", user_controller.CreateUserController)
	userRouter.GET("/email-verify/", user_controller.VerifyEmailController)
	userRouter.POST("/login", user_controller.LoginController)
	userRouter.GET("/", middlewares.JWTAuthMiddleware(), user_controller.GetCurrentUserController)
}
