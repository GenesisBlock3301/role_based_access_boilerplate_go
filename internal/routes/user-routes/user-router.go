package user_routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/controller/user_controller"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userRouter.POST("/create/", user_controller.CreateUserController)
	userRouter.GET("/email-verify/", user_controller.VerifyEmailController)
	userRouter.POST("/login", user_controller.LoginController)
	userRouter.GET("/", middleware.JWTAuthMiddleware(), user_controller.GetCurrentUserController)
}
