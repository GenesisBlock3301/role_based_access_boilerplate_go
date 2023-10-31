package routes

import (
	"github.com/gin-gonic/gin"
)

func RootRouter(router *gin.Engine) {
	// static route
	router.Static("/media", ".media")

	apiRouter := router.Group("api/v1")
	roleRouter := apiRouter.Group("/role")
	userRouter := apiRouter.Group("/user")
	// Find all user's router.
	UserRouter(userRouter)
	// Find all role router.
	RoleRouter(roleRouter)
}
