package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/routes/user"
	"github.com/gin-gonic/gin"
)

func RootRouter(router *gin.Engine) {
	// static route
	router.Static("/media", ".media")

	apiRouter := router.Group("api/v1")
	roleRouter := apiRouter.Group("/role")
	userRouter := apiRouter.Group("/user")
	user.UserRouter(userRouter)
	RoleRouter(roleRouter)
}
