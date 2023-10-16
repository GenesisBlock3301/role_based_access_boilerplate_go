package routes

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/routes/user-routes"
	"github.com/gin-gonic/gin"
)

func RootRouter(router *gin.Engine) {
	// static route
	router.Static("/media", ".media")

	apiRouter := router.Group("api/v1")
	//roleRouter := apiRouter.Group("/role_controller-routes-services")
	userRouter := apiRouter.Group("/user")
	user_routes.UserRouter(userRouter)
	//RoleRouter(roleRouter)
}
